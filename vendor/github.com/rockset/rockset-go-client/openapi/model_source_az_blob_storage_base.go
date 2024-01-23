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

// checks if the SourceAzBlobStorageBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceAzBlobStorageBase{}

// SourceAzBlobStorageBase struct for SourceAzBlobStorageBase
type SourceAzBlobStorageBase struct {
	Settings *SourceAzBlobStorageSettings `json:"settings,omitempty"`
}

// NewSourceAzBlobStorageBase instantiates a new SourceAzBlobStorageBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAzBlobStorageBase() *SourceAzBlobStorageBase {
	this := SourceAzBlobStorageBase{}
	return &this
}

// NewSourceAzBlobStorageBaseWithDefaults instantiates a new SourceAzBlobStorageBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAzBlobStorageBaseWithDefaults() *SourceAzBlobStorageBase {
	this := SourceAzBlobStorageBase{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SourceAzBlobStorageBase) GetSettings() SourceAzBlobStorageSettings {
	if o == nil || IsNil(o.Settings) {
		var ret SourceAzBlobStorageSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAzBlobStorageBase) GetSettingsOk() (*SourceAzBlobStorageSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SourceAzBlobStorageBase) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SourceAzBlobStorageSettings and assigns it to the Settings field.
func (o *SourceAzBlobStorageBase) SetSettings(v SourceAzBlobStorageSettings) {
	o.Settings = &v
}

func (o SourceAzBlobStorageBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceAzBlobStorageBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableSourceAzBlobStorageBase struct {
	value *SourceAzBlobStorageBase
	isSet bool
}

func (v NullableSourceAzBlobStorageBase) Get() *SourceAzBlobStorageBase {
	return v.value
}

func (v *NullableSourceAzBlobStorageBase) Set(val *SourceAzBlobStorageBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAzBlobStorageBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAzBlobStorageBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAzBlobStorageBase(val *SourceAzBlobStorageBase) *NullableSourceAzBlobStorageBase {
	return &NullableSourceAzBlobStorageBase{value: val, isSet: true}
}

func (v NullableSourceAzBlobStorageBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAzBlobStorageBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

