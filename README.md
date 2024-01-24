# Fence Test

This repo contains a test for the Rockset fence API.

To run it, you need to set the following environment variables:

* `ROCKSET_APIKEY` - your Rockset API key
* `ROCKSET_APISERVER` - the Rockset API server to use, e.g. `api.usw2a1.rockset.com`

Then run `go run -mode=vendor main.go` to run the test.
