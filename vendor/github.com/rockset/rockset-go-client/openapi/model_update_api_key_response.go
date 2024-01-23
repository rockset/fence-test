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

// checks if the UpdateApiKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateApiKeyResponse{}

// UpdateApiKeyResponse struct for UpdateApiKeyResponse
type UpdateApiKeyResponse struct {
	Data *ApiKey `json:"data,omitempty"`
}

// NewUpdateApiKeyResponse instantiates a new UpdateApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApiKeyResponse() *UpdateApiKeyResponse {
	this := UpdateApiKeyResponse{}
	return &this
}

// NewUpdateApiKeyResponseWithDefaults instantiates a new UpdateApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApiKeyResponseWithDefaults() *UpdateApiKeyResponse {
	this := UpdateApiKeyResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateApiKeyResponse) GetData() ApiKey {
	if o == nil || IsNil(o.Data) {
		var ret ApiKey
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiKeyResponse) GetDataOk() (*ApiKey, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateApiKeyResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ApiKey and assigns it to the Data field.
func (o *UpdateApiKeyResponse) SetData(v ApiKey) {
	o.Data = &v
}

func (o UpdateApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateApiKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUpdateApiKeyResponse struct {
	value *UpdateApiKeyResponse
	isSet bool
}

func (v NullableUpdateApiKeyResponse) Get() *UpdateApiKeyResponse {
	return v.value
}

func (v *NullableUpdateApiKeyResponse) Set(val *UpdateApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApiKeyResponse(val *UpdateApiKeyResponse) *NullableUpdateApiKeyResponse {
	return &NullableUpdateApiKeyResponse{value: val, isSet: true}
}

func (v NullableUpdateApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


