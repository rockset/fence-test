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

// checks if the AwsAccessKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsAccessKey{}

// AwsAccessKey struct for AwsAccessKey
type AwsAccessKey struct {
	// AWS access key ID.
	AwsAccessKeyId string `json:"aws_access_key_id"`
	// AWS secret access key.
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
}

// NewAwsAccessKey instantiates a new AwsAccessKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsAccessKey(awsAccessKeyId string, awsSecretAccessKey string) *AwsAccessKey {
	this := AwsAccessKey{}
	this.AwsAccessKeyId = awsAccessKeyId
	this.AwsSecretAccessKey = awsSecretAccessKey
	return &this
}

// NewAwsAccessKeyWithDefaults instantiates a new AwsAccessKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsAccessKeyWithDefaults() *AwsAccessKey {
	this := AwsAccessKey{}
	return &this
}

// GetAwsAccessKeyId returns the AwsAccessKeyId field value
func (o *AwsAccessKey) GetAwsAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccessKeyId
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value
// and a boolean to check if the value has been set.
func (o *AwsAccessKey) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccessKeyId, true
}

// SetAwsAccessKeyId sets field value
func (o *AwsAccessKey) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId = v
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value
func (o *AwsAccessKey) GetAwsSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsSecretAccessKey
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *AwsAccessKey) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsSecretAccessKey, true
}

// SetAwsSecretAccessKey sets field value
func (o *AwsAccessKey) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey = v
}

func (o AwsAccessKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsAccessKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aws_access_key_id"] = o.AwsAccessKeyId
	toSerialize["aws_secret_access_key"] = o.AwsSecretAccessKey
	return toSerialize, nil
}

type NullableAwsAccessKey struct {
	value *AwsAccessKey
	isSet bool
}

func (v NullableAwsAccessKey) Get() *AwsAccessKey {
	return v.value
}

func (v *NullableAwsAccessKey) Set(val *AwsAccessKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsAccessKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsAccessKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsAccessKey(val *AwsAccessKey) *NullableAwsAccessKey {
	return &NullableAwsAccessKey{value: val, isSet: true}
}

func (v NullableAwsAccessKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsAccessKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

