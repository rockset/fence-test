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

// checks if the SourceMongoDb type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceMongoDb{}

// SourceMongoDb struct for SourceMongoDb
type SourceMongoDb struct {
	// MongoDB collection name.
	CollectionName string `json:"collection_name"`
	// MongoDB database name containing this collection.
	DatabaseName string `json:"database_name"`
	// Whether to get the full document from the MongoDB change stream to enable multi-field expression transformations. Selecting this option will increase load on your upstream MongoDB database.
	RetrieveFullDocument *bool `json:"retrieve_full_document,omitempty"`
	Status *StatusMongoDb `json:"status,omitempty"`
}

// NewSourceMongoDb instantiates a new SourceMongoDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceMongoDb(collectionName string, databaseName string) *SourceMongoDb {
	this := SourceMongoDb{}
	this.CollectionName = collectionName
	this.DatabaseName = databaseName
	return &this
}

// NewSourceMongoDbWithDefaults instantiates a new SourceMongoDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceMongoDbWithDefaults() *SourceMongoDb {
	this := SourceMongoDb{}
	return &this
}

// GetCollectionName returns the CollectionName field value
func (o *SourceMongoDb) GetCollectionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value
// and a boolean to check if the value has been set.
func (o *SourceMongoDb) GetCollectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionName, true
}

// SetCollectionName sets field value
func (o *SourceMongoDb) SetCollectionName(v string) {
	o.CollectionName = v
}

// GetDatabaseName returns the DatabaseName field value
func (o *SourceMongoDb) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *SourceMongoDb) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value
func (o *SourceMongoDb) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetRetrieveFullDocument returns the RetrieveFullDocument field value if set, zero value otherwise.
func (o *SourceMongoDb) GetRetrieveFullDocument() bool {
	if o == nil || IsNil(o.RetrieveFullDocument) {
		var ret bool
		return ret
	}
	return *o.RetrieveFullDocument
}

// GetRetrieveFullDocumentOk returns a tuple with the RetrieveFullDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceMongoDb) GetRetrieveFullDocumentOk() (*bool, bool) {
	if o == nil || IsNil(o.RetrieveFullDocument) {
		return nil, false
	}
	return o.RetrieveFullDocument, true
}

// HasRetrieveFullDocument returns a boolean if a field has been set.
func (o *SourceMongoDb) HasRetrieveFullDocument() bool {
	if o != nil && !IsNil(o.RetrieveFullDocument) {
		return true
	}

	return false
}

// SetRetrieveFullDocument gets a reference to the given bool and assigns it to the RetrieveFullDocument field.
func (o *SourceMongoDb) SetRetrieveFullDocument(v bool) {
	o.RetrieveFullDocument = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SourceMongoDb) GetStatus() StatusMongoDb {
	if o == nil || IsNil(o.Status) {
		var ret StatusMongoDb
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceMongoDb) GetStatusOk() (*StatusMongoDb, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SourceMongoDb) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StatusMongoDb and assigns it to the Status field.
func (o *SourceMongoDb) SetStatus(v StatusMongoDb) {
	o.Status = &v
}

func (o SourceMongoDb) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceMongoDb) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection_name"] = o.CollectionName
	toSerialize["database_name"] = o.DatabaseName
	if !IsNil(o.RetrieveFullDocument) {
		toSerialize["retrieve_full_document"] = o.RetrieveFullDocument
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableSourceMongoDb struct {
	value *SourceMongoDb
	isSet bool
}

func (v NullableSourceMongoDb) Get() *SourceMongoDb {
	return v.value
}

func (v *NullableSourceMongoDb) Set(val *SourceMongoDb) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceMongoDb) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceMongoDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceMongoDb(val *SourceMongoDb) *NullableSourceMongoDb {
	return &NullableSourceMongoDb{value: val, isSet: true}
}

func (v NullableSourceMongoDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceMongoDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


