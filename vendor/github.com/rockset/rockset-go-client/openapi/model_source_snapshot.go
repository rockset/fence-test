/*
REST API

Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SourceSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceSnapshot{}

// SourceSnapshot struct for SourceSnapshot
type SourceSnapshot struct {
	// RRN of the snapshot that the new collection will be created from.
	SourceSnapshotRrn *string `json:"source_snapshot_rrn,omitempty"`
}

// NewSourceSnapshot instantiates a new SourceSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceSnapshot() *SourceSnapshot {
	this := SourceSnapshot{}
	return &this
}

// NewSourceSnapshotWithDefaults instantiates a new SourceSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceSnapshotWithDefaults() *SourceSnapshot {
	this := SourceSnapshot{}
	return &this
}

// GetSourceSnapshotRrn returns the SourceSnapshotRrn field value if set, zero value otherwise.
func (o *SourceSnapshot) GetSourceSnapshotRrn() string {
	if o == nil || IsNil(o.SourceSnapshotRrn) {
		var ret string
		return ret
	}
	return *o.SourceSnapshotRrn
}

// GetSourceSnapshotRrnOk returns a tuple with the SourceSnapshotRrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSnapshot) GetSourceSnapshotRrnOk() (*string, bool) {
	if o == nil || IsNil(o.SourceSnapshotRrn) {
		return nil, false
	}
	return o.SourceSnapshotRrn, true
}

// HasSourceSnapshotRrn returns a boolean if a field has been set.
func (o *SourceSnapshot) HasSourceSnapshotRrn() bool {
	if o != nil && !IsNil(o.SourceSnapshotRrn) {
		return true
	}

	return false
}

// SetSourceSnapshotRrn gets a reference to the given string and assigns it to the SourceSnapshotRrn field.
func (o *SourceSnapshot) SetSourceSnapshotRrn(v string) {
	o.SourceSnapshotRrn = &v
}

func (o SourceSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceSnapshotRrn) {
		toSerialize["source_snapshot_rrn"] = o.SourceSnapshotRrn
	}
	return toSerialize, nil
}

type NullableSourceSnapshot struct {
	value *SourceSnapshot
	isSet bool
}

func (v NullableSourceSnapshot) Get() *SourceSnapshot {
	return v.value
}

func (v *NullableSourceSnapshot) Set(val *SourceSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceSnapshot(val *SourceSnapshot) *NullableSourceSnapshot {
	return &NullableSourceSnapshot{value: val, isSet: true}
}

func (v NullableSourceSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


