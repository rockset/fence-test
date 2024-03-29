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

// checks if the View type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &View{}

// View struct for View
type View struct {
	// ISO-8601 date.
	CreatedAt *string `json:"created_at,omitempty"`
	// Name of the API key that was used to create this object if one was used.
	CreatedByApikeyName *string `json:"created_by_apikey_name,omitempty"`
	// Email of the creator.
	CreatorEmail *string `json:"creator_email,omitempty"`
	// View description.
	Description *string `json:"description,omitempty"`
	// List of entities referenced by view. An entity can be a view, alias or collection.
	Entities []string `json:"entities,omitempty"`
	// ISO-8601 date.
	ModifiedAt *string `json:"modified_at,omitempty"`
	// Name of the view.
	Name *string `json:"name,omitempty"`
	// Email of the owner, note: deprecated and will always be null.
	OwnerEmail *string `json:"owner_email,omitempty"`
	Path *string `json:"path,omitempty"`
	// SQL query of the view.
	QuerySql *string `json:"query_sql,omitempty"`
	// State of the view.
	State *string `json:"state,omitempty"`
	// Name of the workspace.
	Workspace *string `json:"workspace,omitempty"`
}

// NewView instantiates a new View object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewView() *View {
	this := View{}
	return &this
}

// NewViewWithDefaults instantiates a new View object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewWithDefaults() *View {
	this := View{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *View) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *View) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *View) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedByApikeyName returns the CreatedByApikeyName field value if set, zero value otherwise.
func (o *View) GetCreatedByApikeyName() string {
	if o == nil || IsNil(o.CreatedByApikeyName) {
		var ret string
		return ret
	}
	return *o.CreatedByApikeyName
}

// GetCreatedByApikeyNameOk returns a tuple with the CreatedByApikeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetCreatedByApikeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByApikeyName) {
		return nil, false
	}
	return o.CreatedByApikeyName, true
}

// HasCreatedByApikeyName returns a boolean if a field has been set.
func (o *View) HasCreatedByApikeyName() bool {
	if o != nil && !IsNil(o.CreatedByApikeyName) {
		return true
	}

	return false
}

// SetCreatedByApikeyName gets a reference to the given string and assigns it to the CreatedByApikeyName field.
func (o *View) SetCreatedByApikeyName(v string) {
	o.CreatedByApikeyName = &v
}

// GetCreatorEmail returns the CreatorEmail field value if set, zero value otherwise.
func (o *View) GetCreatorEmail() string {
	if o == nil || IsNil(o.CreatorEmail) {
		var ret string
		return ret
	}
	return *o.CreatorEmail
}

// GetCreatorEmailOk returns a tuple with the CreatorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetCreatorEmailOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorEmail) {
		return nil, false
	}
	return o.CreatorEmail, true
}

// HasCreatorEmail returns a boolean if a field has been set.
func (o *View) HasCreatorEmail() bool {
	if o != nil && !IsNil(o.CreatorEmail) {
		return true
	}

	return false
}

// SetCreatorEmail gets a reference to the given string and assigns it to the CreatorEmail field.
func (o *View) SetCreatorEmail(v string) {
	o.CreatorEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *View) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *View) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *View) SetDescription(v string) {
	o.Description = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *View) GetEntities() []string {
	if o == nil || IsNil(o.Entities) {
		var ret []string
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetEntitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *View) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []string and assigns it to the Entities field.
func (o *View) SetEntities(v []string) {
	o.Entities = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *View) GetModifiedAt() string {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetModifiedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *View) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *View) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *View) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *View) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *View) SetName(v string) {
	o.Name = &v
}

// GetOwnerEmail returns the OwnerEmail field value if set, zero value otherwise.
func (o *View) GetOwnerEmail() string {
	if o == nil || IsNil(o.OwnerEmail) {
		var ret string
		return ret
	}
	return *o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetOwnerEmailOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerEmail) {
		return nil, false
	}
	return o.OwnerEmail, true
}

// HasOwnerEmail returns a boolean if a field has been set.
func (o *View) HasOwnerEmail() bool {
	if o != nil && !IsNil(o.OwnerEmail) {
		return true
	}

	return false
}

// SetOwnerEmail gets a reference to the given string and assigns it to the OwnerEmail field.
func (o *View) SetOwnerEmail(v string) {
	o.OwnerEmail = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *View) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *View) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *View) SetPath(v string) {
	o.Path = &v
}

// GetQuerySql returns the QuerySql field value if set, zero value otherwise.
func (o *View) GetQuerySql() string {
	if o == nil || IsNil(o.QuerySql) {
		var ret string
		return ret
	}
	return *o.QuerySql
}

// GetQuerySqlOk returns a tuple with the QuerySql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetQuerySqlOk() (*string, bool) {
	if o == nil || IsNil(o.QuerySql) {
		return nil, false
	}
	return o.QuerySql, true
}

// HasQuerySql returns a boolean if a field has been set.
func (o *View) HasQuerySql() bool {
	if o != nil && !IsNil(o.QuerySql) {
		return true
	}

	return false
}

// SetQuerySql gets a reference to the given string and assigns it to the QuerySql field.
func (o *View) SetQuerySql(v string) {
	o.QuerySql = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *View) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *View) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *View) SetState(v string) {
	o.State = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *View) GetWorkspace() string {
	if o == nil || IsNil(o.Workspace) {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetWorkspaceOk() (*string, bool) {
	if o == nil || IsNil(o.Workspace) {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *View) HasWorkspace() bool {
	if o != nil && !IsNil(o.Workspace) {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *View) SetWorkspace(v string) {
	o.Workspace = &v
}

func (o View) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o View) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedByApikeyName) {
		toSerialize["created_by_apikey_name"] = o.CreatedByApikeyName
	}
	if !IsNil(o.CreatorEmail) {
		toSerialize["creator_email"] = o.CreatorEmail
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OwnerEmail) {
		toSerialize["owner_email"] = o.OwnerEmail
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.QuerySql) {
		toSerialize["query_sql"] = o.QuerySql
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Workspace) {
		toSerialize["workspace"] = o.Workspace
	}
	return toSerialize, nil
}

type NullableView struct {
	value *View
	isSet bool
}

func (v NullableView) Get() *View {
	return v.value
}

func (v *NullableView) Set(val *View) {
	v.value = val
	v.isSet = true
}

func (v NullableView) IsSet() bool {
	return v.isSet
}

func (v *NullableView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableView(val *View) *NullableView {
	return &NullableView{value: val, isSet: true}
}

func (v NullableView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


