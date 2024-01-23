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

// checks if the AsyncQueryOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncQueryOptions{}

// AsyncQueryOptions struct for AsyncQueryOptions
type AsyncQueryOptions struct {
	// If the query completes before the client timeout, the results are returned. Otherwise if the client timeout is exceeded, the query id will be returned, and the query will continue to run in the background for up to 30 minutes. (The 30 minute timeout can be configured lower with timeout_ms.) `async_options.client_timeout_ms` only applies when `async` is true. The default value of `client_timeout_ms` is 0, so async query requests will immediately return with a query id by default. 
	ClientTimeoutMs *int64 `json:"client_timeout_ms,omitempty"`
	// [DEPRECATED] Use the query request `max_initial_results` instead. The maximum number of results you will receive as a client. If the query exceeds this limit, the remaining results can be requested using a returned pagination cursor. In addition, there is a maximum response size of 100MiB so fewer than `max_results` may be returned.
	MaxInitialResults *int64 `json:"max_initial_results,omitempty"`
	// [DEPRECATED] Use the query request `timeout_ms` instead. The maximum amount of time that the system will attempt to complete query execution before aborting the query and returning an error. This must be set to a value that is greater than or equal to the client timeout, and the maximum value of this timeout is 30 minutes.
	TimeoutMs *int64 `json:"timeout_ms,omitempty"`
}

// NewAsyncQueryOptions instantiates a new AsyncQueryOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncQueryOptions() *AsyncQueryOptions {
	this := AsyncQueryOptions{}
	return &this
}

// NewAsyncQueryOptionsWithDefaults instantiates a new AsyncQueryOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncQueryOptionsWithDefaults() *AsyncQueryOptions {
	this := AsyncQueryOptions{}
	return &this
}

// GetClientTimeoutMs returns the ClientTimeoutMs field value if set, zero value otherwise.
func (o *AsyncQueryOptions) GetClientTimeoutMs() int64 {
	if o == nil || IsNil(o.ClientTimeoutMs) {
		var ret int64
		return ret
	}
	return *o.ClientTimeoutMs
}

// GetClientTimeoutMsOk returns a tuple with the ClientTimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncQueryOptions) GetClientTimeoutMsOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientTimeoutMs) {
		return nil, false
	}
	return o.ClientTimeoutMs, true
}

// HasClientTimeoutMs returns a boolean if a field has been set.
func (o *AsyncQueryOptions) HasClientTimeoutMs() bool {
	if o != nil && !IsNil(o.ClientTimeoutMs) {
		return true
	}

	return false
}

// SetClientTimeoutMs gets a reference to the given int64 and assigns it to the ClientTimeoutMs field.
func (o *AsyncQueryOptions) SetClientTimeoutMs(v int64) {
	o.ClientTimeoutMs = &v
}

// GetMaxInitialResults returns the MaxInitialResults field value if set, zero value otherwise.
func (o *AsyncQueryOptions) GetMaxInitialResults() int64 {
	if o == nil || IsNil(o.MaxInitialResults) {
		var ret int64
		return ret
	}
	return *o.MaxInitialResults
}

// GetMaxInitialResultsOk returns a tuple with the MaxInitialResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncQueryOptions) GetMaxInitialResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxInitialResults) {
		return nil, false
	}
	return o.MaxInitialResults, true
}

// HasMaxInitialResults returns a boolean if a field has been set.
func (o *AsyncQueryOptions) HasMaxInitialResults() bool {
	if o != nil && !IsNil(o.MaxInitialResults) {
		return true
	}

	return false
}

// SetMaxInitialResults gets a reference to the given int64 and assigns it to the MaxInitialResults field.
func (o *AsyncQueryOptions) SetMaxInitialResults(v int64) {
	o.MaxInitialResults = &v
}

// GetTimeoutMs returns the TimeoutMs field value if set, zero value otherwise.
func (o *AsyncQueryOptions) GetTimeoutMs() int64 {
	if o == nil || IsNil(o.TimeoutMs) {
		var ret int64
		return ret
	}
	return *o.TimeoutMs
}

// GetTimeoutMsOk returns a tuple with the TimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncQueryOptions) GetTimeoutMsOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeoutMs) {
		return nil, false
	}
	return o.TimeoutMs, true
}

// HasTimeoutMs returns a boolean if a field has been set.
func (o *AsyncQueryOptions) HasTimeoutMs() bool {
	if o != nil && !IsNil(o.TimeoutMs) {
		return true
	}

	return false
}

// SetTimeoutMs gets a reference to the given int64 and assigns it to the TimeoutMs field.
func (o *AsyncQueryOptions) SetTimeoutMs(v int64) {
	o.TimeoutMs = &v
}

func (o AsyncQueryOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncQueryOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientTimeoutMs) {
		toSerialize["client_timeout_ms"] = o.ClientTimeoutMs
	}
	if !IsNil(o.MaxInitialResults) {
		toSerialize["max_initial_results"] = o.MaxInitialResults
	}
	if !IsNil(o.TimeoutMs) {
		toSerialize["timeout_ms"] = o.TimeoutMs
	}
	return toSerialize, nil
}

type NullableAsyncQueryOptions struct {
	value *AsyncQueryOptions
	isSet bool
}

func (v NullableAsyncQueryOptions) Get() *AsyncQueryOptions {
	return v.value
}

func (v *NullableAsyncQueryOptions) Set(val *AsyncQueryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncQueryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncQueryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncQueryOptions(val *AsyncQueryOptions) *NullableAsyncQueryOptions {
	return &NullableAsyncQueryOptions{value: val, isSet: true}
}

func (v NullableAsyncQueryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncQueryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


