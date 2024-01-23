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

// checks if the QueryLambdaTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryLambdaTag{}

// QueryLambdaTag struct for QueryLambdaTag
type QueryLambdaTag struct {
	// Name of Query Lambda tag.
	TagName *string `json:"tag_name,omitempty"`
	Version *QueryLambdaVersion `json:"version,omitempty"`
}

// NewQueryLambdaTag instantiates a new QueryLambdaTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLambdaTag() *QueryLambdaTag {
	this := QueryLambdaTag{}
	return &this
}

// NewQueryLambdaTagWithDefaults instantiates a new QueryLambdaTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLambdaTagWithDefaults() *QueryLambdaTag {
	this := QueryLambdaTag{}
	return &this
}

// GetTagName returns the TagName field value if set, zero value otherwise.
func (o *QueryLambdaTag) GetTagName() string {
	if o == nil || IsNil(o.TagName) {
		var ret string
		return ret
	}
	return *o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaTag) GetTagNameOk() (*string, bool) {
	if o == nil || IsNil(o.TagName) {
		return nil, false
	}
	return o.TagName, true
}

// HasTagName returns a boolean if a field has been set.
func (o *QueryLambdaTag) HasTagName() bool {
	if o != nil && !IsNil(o.TagName) {
		return true
	}

	return false
}

// SetTagName gets a reference to the given string and assigns it to the TagName field.
func (o *QueryLambdaTag) SetTagName(v string) {
	o.TagName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *QueryLambdaTag) GetVersion() QueryLambdaVersion {
	if o == nil || IsNil(o.Version) {
		var ret QueryLambdaVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLambdaTag) GetVersionOk() (*QueryLambdaVersion, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *QueryLambdaTag) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given QueryLambdaVersion and assigns it to the Version field.
func (o *QueryLambdaTag) SetVersion(v QueryLambdaVersion) {
	o.Version = &v
}

func (o QueryLambdaTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLambdaTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TagName) {
		toSerialize["tag_name"] = o.TagName
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableQueryLambdaTag struct {
	value *QueryLambdaTag
	isSet bool
}

func (v NullableQueryLambdaTag) Get() *QueryLambdaTag {
	return v.value
}

func (v *NullableQueryLambdaTag) Set(val *QueryLambdaTag) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLambdaTag) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLambdaTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLambdaTag(val *QueryLambdaTag) *NullableQueryLambdaTag {
	return &NullableQueryLambdaTag{value: val, isSet: true}
}

func (v NullableQueryLambdaTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLambdaTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


