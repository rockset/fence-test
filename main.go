package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rockset/rockset-go-client/retry"
	"github.com/rockset/rockset-go-client/wait"
)

type Worker struct {
	rc         *rockset.RockClient
	wg         *sync.WaitGroup
	Workspace  string
	Collection string
	l          *slog.Logger
	truncate   bool
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	l := slog.New(slog.NewTextHandler(os.Stderr, nil))

	r := retry.NewExponential()
	r.MaxBackoff = 2 * time.Second
	rc, err := rockset.NewClient(rockset.WithRetry(r))
	if err != nil {
		panic(err)
	}

	pfx := "w"
	ws := RandomString(pfx, 6)
	collection := "test"

	if _, err = rc.CreateWorkspace(ctx, ws); err != nil {
		panic(err)
	}
	l.Info("created workspace", "ws", ws)
	defer RemoveWorkspace(l, rc, ws)

	w := wait.New(rc)
	if err = w.UntilWorkspaceAvailable(ctx, ws); err != nil {
		panic(err)
	}
	l.Info("workspace available", "ws", ws)

	if err := createAndWait(ctx, l, rc, ws, collection); err != nil {
		panic(err)
	}

	var count = 10
	var offset string
	for idx := 0; idx < count; idx++ {
		t0 := time.Now()
		slog.Info("loop", "idx", idx)

		// query and see how many documents there are (should be 0)
		res, err := rc.Query(ctx, fmt.Sprintf("SELECT _id FROM %s.%s", ws, collection))
		if err != nil {
			panic(err)
		}
		if len(res.Results) != 0 {
			panic(fmt.Sprintf("expected 0 docs, got %d", len(res.Results)))
		}

		// write docs, get offset
		if offset, err = addDocs(ctx, rc, ws, collection, idx, count); err != nil {
			panic(err)
		}
		l.Info("documents written", "ws", ws, "c", collection, "count", count, "Δ", time.Since(t0).String())

		// fence (loop and get collection commit)
		if err = waitForOffset(ctx, rc, ws, collection, offset); err != nil {
			panic(err)
		}
		slog.Info("write fence passed", "Δ", time.Since(t0).String())

		res, err = rc.Query(ctx, fmt.Sprintf("SELECT _id FROM %s.%s", ws, collection))
		if err != nil {
			panic(err)
		}
		if len(res.Results) != count {
			var ids []string
			for _, r := range res.Results {
				ids = append(ids, r["_id"].(string))
			}
			panic(fmt.Sprintf("expected %d docs, got %d: %v", count, len(res.Results), ids))
		}

		// delete all docs in the collection
		if offset, err = deleteDocs(ctx, rc, ws, collection); err != nil {
			panic(err)
		}

		// fence (loop and get collection commit)
		if err = waitForOffset(ctx, rc, ws, collection, offset); err != nil {
			panic(err)
		}

		slog.Info("delete fence passed", "Δ", time.Since(t0).String())
	}
}

func addDocs(ctx context.Context, rc *rockset.RockClient, ws, collection string, idx, count int) (string, error) {
	var offset string
	var ids []string
	for i := 0; i < count; i++ {
		doc := map[string]interface{}{}
		doc["idx"] = idx
		doc["i"] = i
		res, err := rc.AddDocumentsWithOffset(ctx, ws, collection, []interface{}{doc})
		if err != nil {
			return "", err
		}

		if len(res.Data) != 1 {
			return "", fmt.Errorf("expected 1 result, got %d", len(res.Data))
		}
		first := res.GetData()[0]
		if first.GetStatus() != "ADDED" {
			return "", fmt.Errorf("expected status ADDED, got %s", first.GetStatus())
		}

		ids = append(ids, first.GetId())
		offset = res.GetLastOffset()
	}
	slog.Info("documents written", "ws", ws, "c", collection, "count", count, "ids", ids)

	return offset, nil
}

func waitForOffset(ctx context.Context, rc *rockset.RockClient, ws, collection, offset string) error {
	for {
		qr, err := rc.GetCollectionCommit(ctx, ws, collection, []string{offset})
		if err != nil {
			return err
		}

		if qr.GetPassed() {
			slog.Info("fence passed", "ws", ws, "c", collection, "offset", offset)
			break
		}

		time.Sleep(time.Second)
	}

	return nil
}

func deleteDocs(ctx context.Context, rc *rockset.RockClient, ws, collection string) (string, error) {
	res, err := rc.Query(ctx, fmt.Sprintf("SELECT _id FROM %s.%s", ws, collection))
	if err != nil {
		return "", err
	}

	ids := make([]string, len(res.Results))
	for i, r := range res.Results {
		if i > 10000 {
			break
		}
		ids[i] = r["_id"].(string)
	}

	docs, err := rc.DeleteDocumentsWithOffset(ctx, ws, collection, ids)
	if err != nil {
		return "", err
	}

	for _, d := range docs.GetData() {
		if d.GetStatus() != "DELETED" {
			slog.Error("failed to delete", "ws", ws, "c", collection, "_id", d.GetId(), "staus", d.GetStatus())
		}
	}
	slog.Info("deleted documents", "count", len(docs.GetData()), "offset", docs.GetLastOffset(), "ids", ids)

	return docs.GetLastOffset(), nil
}

func createAndWait(ctx context.Context, l *slog.Logger, rc *rockset.RockClient, ws, collection string) error {
	c, err := rc.CreateCollection(ctx, ws, collection)
	if err != nil {
		return err
	}
	l.Debug("created", "ws", c.GetWorkspace(), "c", c.GetName())

	w := wait.New(rc)
	if err := w.UntilCollectionReady(ctx, ws, collection); err != nil {
		return err
	}
	l.Info("collection ready", "ws", ws, "c", collection)

	return nil
}

func RemoveWorkspace(l *slog.Logger, rc *rockset.RockClient, ws string) {
	l.Info("removing workspace", "ws", ws)
	for {
		cols, err := rc.ListCollections(context.Background(), option.WithWorkspace(ws))
		if err != nil {
			l.Error("failed to list collections", "ws", ws, "err", err)
			break
		}
		if len(cols) == 0 {
			l.Debug("all collections deleted", "ws", ws)
			break
		}

		for _, c := range cols {
			switch s := c.GetStatus(); s {
			case "DELETING":
				l.Debug("waiting for collection to be deleted", "ws", ws, "c", c.GetName())
				continue
			case "DELETED":
				continue
			case "READY":
			// nothing
			default:
				slog.Warn("unexpected status", "ws", ws, "c", c.GetName(), "status", s)
				continue
			}

			if err = rc.DeleteCollection(context.Background(), ws, c.GetName()); err != nil {
				l.Error("failed to delete collection", "ws", ws, "c", c.GetName(), "err", err)
				continue
			}
			slog.Info("deleted", "ws", ws, "c", c.GetName())
		}
		time.Sleep(time.Second)
	}

	if err := rc.DeleteWorkspace(context.TODO(), ws); err != nil {
		l.Error("failed to delete", "ws", ws, "err", err)
		return
	}
	l.Info("deleted", "ws", ws)
}

func RandomString(pfx string, length int) string {
	return pfx + "_" + stringWithCharset(length, charset)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
