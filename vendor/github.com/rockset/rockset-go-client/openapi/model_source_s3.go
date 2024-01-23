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

// checks if the SourceS3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceS3{}

// SourceS3 struct for SourceS3
type SourceS3 struct {
	// Address of S3 bucket containing data.
	Bucket string `json:"bucket"`
	ObjectBytesDownloaded *int64 `json:"object_bytes_downloaded,omitempty"`
	ObjectBytesTotal *int64 `json:"object_bytes_total,omitempty"`
	ObjectCountDownloaded *int64 `json:"object_count_downloaded,omitempty"`
	ObjectCountTotal *int64 `json:"object_count_total,omitempty"`
	// Glob-style pattern that selects keys to ingest. Only either prefix or pattern can be specified.
	Pattern *string `json:"pattern,omitempty"`
	// Prefix that selects keys to ingest.
	Prefix *string `json:"prefix,omitempty"`
	// List of prefixes to paths from which data should be ingested.
	Prefixes []string `json:"prefixes,omitempty"`
	// AWS region containing source bucket.
	Region *string `json:"region,omitempty"`
	Settings *SourceS3Settings `json:"settings,omitempty"`
}

// NewSourceS3 instantiates a new SourceS3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceS3(bucket string) *SourceS3 {
	this := SourceS3{}
	this.Bucket = bucket
	return &this
}

// NewSourceS3WithDefaults instantiates a new SourceS3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceS3WithDefaults() *SourceS3 {
	this := SourceS3{}
	return &this
}

// GetBucket returns the Bucket field value
func (o *SourceS3) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *SourceS3) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *SourceS3) SetBucket(v string) {
	o.Bucket = v
}

// GetObjectBytesDownloaded returns the ObjectBytesDownloaded field value if set, zero value otherwise.
func (o *SourceS3) GetObjectBytesDownloaded() int64 {
	if o == nil || IsNil(o.ObjectBytesDownloaded) {
		var ret int64
		return ret
	}
	return *o.ObjectBytesDownloaded
}

// GetObjectBytesDownloadedOk returns a tuple with the ObjectBytesDownloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceS3) GetObjectBytesDownloadedOk() (*int64, bool) {
	if o == nil || IsNil(o.ObjectBytesDownloaded) {
		return nil, false
	}
	return o.ObjectBytesDownloaded, true
}

// HasObjectBytesDownloaded returns a boolean if a field has been set.
func (o *SourceS3) HasObjectBytesDownloaded() bool {
	if o != nil && !IsNil(o.ObjectBytesDownloaded) {
		return true
	}

	return false
}

// SetObjectBytesDownloaded gets a reference to the given int64 and assigns it to the ObjectBytesDownloaded field.
func (o *SourceS3) SetObjectBytesDownloaded(v int64) {
	o.ObjectBytesDownloaded = &v
}

// GetObjectBytesTotal returns the ObjectBytesTotal field value if set, zero value otherwise.
func (o *SourceS3) GetObjectBytesTotal() int64 {
	if o == nil || IsNil(o.ObjectBytesTotal) {
		var ret int64
		return ret
	}
	return *o.ObjectBytesTotal
}

// GetObjectBytesTotalOk returns a tuple with the ObjectBytesTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceS3) GetObjectBytesTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.ObjectBytesTotal) {
		return nil, false
	}
	return o.ObjectBytesTotal, true
}

// HasObjectBytesTotal returns a boolean if a field has been set.
func (o *SourceS3) HasObjectBytesTotal() bool {
	if o != nil && !IsNil(o.ObjectBytesTotal) {
		return true
	}

	return false
}

// SetObjectBytesTotal gets a reference to the given int64 and assigns it to the ObjectBytesTotal field.
func (o *SourceS3) SetObjectBytesTotal(v int64) {
	o.ObjectBytesTotal = &v
}

// GetObjectCountDownloaded returns the ObjectCountDownloaded field value if set, zero value otherwise.
func (o *SourceS3) GetObjectCountDownloaded() int64 {
	if o == nil || IsNil(o.ObjectCountDownloaded) {
		var ret int64
		return ret
	}
	return *o.ObjectCountDownloaded
}

// GetObjectCountDownloadedOk returns a tuple with the ObjectCountDownloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceS3) GetObjectCountDownloadedOk() (*int64, bool) {
	if o == nil || IsNil(o.ObjectCountDownloaded) {
		return nil, false
	}
	return o.ObjectCountDownloaded, true
}

// HasObjectCountDownloaded returns a boolean if a field has been set.
func (o *SourceS3) HasObjectCountDownloaded() bool {
	if o != nil && !IsNil(o.ObjectCountDownloaded) {
		return true
	}

	return false
}

// SetObjectCountDownloaded gets a reference to the given int64 and assigns it to the ObjectCountDownloaded field.
func (o *SourceS3) SetObjectCountDownloaded(v int64) {
	o.ObjectCountDownloaded = &v
}

// GetObjectCountTotal returns the ObjectCountTotal field value if set, zero value otherwise.
func (o *SourceS3) GetObjectCountTotal() int64 {
	if o == nil || IsNil(o.ObjectCountTotal) {
		var ret int64
		return ret
	}
	return *o.ObjectCountTotal
}

// GetObjectCountTotalOk returns a tuple with the ObjectCountTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceS3) GetObjectCountTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.ObjectCountTotal) {
		return nil, false
	}
	return o.ObjectCountTotal, true
}

// HasObjectCountTotal returns a boolean if a field has been set.
func (o *SourceS3) HasObjectCountTotal() bool {
	if o != nil && !IsNil(o.ObjectCountTotal) {
		return true
	}

	return false
}

// SetObjectCountTotal gets a reference to the given int64 and assigns it to the ObjectCountTotal field.
func (o *SourceS3) SetObjectCountTotal(v int64) {
	o.ObjectCountTotal = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *SourceS3) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceS3) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SourceS3) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *SourceS3) SetPattern(v string) {
	o.Pattern = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *SourceS3) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceS3) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *SourceS3) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *SourceS3) SetPrefix(v string) {
	o.Prefix = &v
}

// GetPrefixes returns the Prefixes field value if set, zero value otherwise.
func (o *SourceS3) GetPrefixes() []string {
	if o == nil || IsNil(o.Prefixes) {
		var ret []string
		return ret
	}
	return o.Prefixes
}

// GetPrefixesOk returns a tuple with the Prefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceS3) GetPrefixesOk() ([]string, bool) {
	if o == nil || IsNil(o.Prefixes) {
		return nil, false
	}
	return o.Prefixes, true
}

// HasPrefixes returns a boolean if a field has been set.
func (o *SourceS3) HasPrefixes() bool {
	if o != nil && !IsNil(o.Prefixes) {
		return true
	}

	return false
}

// SetPrefixes gets a reference to the given []string and assigns it to the Prefixes field.
func (o *SourceS3) SetPrefixes(v []string) {
	o.Prefixes = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *SourceS3) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceS3) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *SourceS3) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *SourceS3) SetRegion(v string) {
	o.Region = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SourceS3) GetSettings() SourceS3Settings {
	if o == nil || IsNil(o.Settings) {
		var ret SourceS3Settings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceS3) GetSettingsOk() (*SourceS3Settings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SourceS3) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SourceS3Settings and assigns it to the Settings field.
func (o *SourceS3) SetSettings(v SourceS3Settings) {
	o.Settings = &v
}

func (o SourceS3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceS3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket"] = o.Bucket
	// skip: object_bytes_downloaded is readOnly
	// skip: object_bytes_total is readOnly
	// skip: object_count_downloaded is readOnly
	// skip: object_count_total is readOnly
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	// skip: prefixes is readOnly
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableSourceS3 struct {
	value *SourceS3
	isSet bool
}

func (v NullableSourceS3) Get() *SourceS3 {
	return v.value
}

func (v *NullableSourceS3) Set(val *SourceS3) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceS3) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceS3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceS3(val *SourceS3) *NullableSourceS3 {
	return &NullableSourceS3{value: val, isSet: true}
}

func (v NullableSourceS3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceS3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

