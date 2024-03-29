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

// checks if the Organization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organization{}

// Organization An organization in Rockset is a container for users and collections.
type Organization struct {
	// List of clusters associated with this org.
	Clusters []Cluster `json:"clusters,omitempty"`
	// ISO-8601 date.
	CreatedAt *string `json:"created_at,omitempty"`
	// Name of the organization.
	DisplayName *string `json:"display_name,omitempty"`
	// Organization's unique external ID within Rockset.
	ExternalId *string `json:"external_id,omitempty"`
	// Unique identifier for the organization.
	Id *string `json:"id,omitempty"`
	// Rockset's global AWS user.
	RocksetUser *string `json:"rockset_user,omitempty"`
	// Connection name of SSO connection.
	SsoConnection *string `json:"sso_connection,omitempty"`
	// Whether or not SSO is the only permitted form of auth.
	SsoOnly *bool `json:"sso_only,omitempty"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization() *Organization {
	this := Organization{}
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *Organization) GetClusters() []Cluster {
	if o == nil || IsNil(o.Clusters) {
		var ret []Cluster
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetClustersOk() ([]Cluster, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *Organization) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []Cluster and assigns it to the Clusters field.
func (o *Organization) SetClusters(v []Cluster) {
	o.Clusters = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Organization) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Organization) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Organization) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Organization) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Organization) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Organization) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Organization) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Organization) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Organization) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Organization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Organization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Organization) SetId(v string) {
	o.Id = &v
}

// GetRocksetUser returns the RocksetUser field value if set, zero value otherwise.
func (o *Organization) GetRocksetUser() string {
	if o == nil || IsNil(o.RocksetUser) {
		var ret string
		return ret
	}
	return *o.RocksetUser
}

// GetRocksetUserOk returns a tuple with the RocksetUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetRocksetUserOk() (*string, bool) {
	if o == nil || IsNil(o.RocksetUser) {
		return nil, false
	}
	return o.RocksetUser, true
}

// HasRocksetUser returns a boolean if a field has been set.
func (o *Organization) HasRocksetUser() bool {
	if o != nil && !IsNil(o.RocksetUser) {
		return true
	}

	return false
}

// SetRocksetUser gets a reference to the given string and assigns it to the RocksetUser field.
func (o *Organization) SetRocksetUser(v string) {
	o.RocksetUser = &v
}

// GetSsoConnection returns the SsoConnection field value if set, zero value otherwise.
func (o *Organization) GetSsoConnection() string {
	if o == nil || IsNil(o.SsoConnection) {
		var ret string
		return ret
	}
	return *o.SsoConnection
}

// GetSsoConnectionOk returns a tuple with the SsoConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetSsoConnectionOk() (*string, bool) {
	if o == nil || IsNil(o.SsoConnection) {
		return nil, false
	}
	return o.SsoConnection, true
}

// HasSsoConnection returns a boolean if a field has been set.
func (o *Organization) HasSsoConnection() bool {
	if o != nil && !IsNil(o.SsoConnection) {
		return true
	}

	return false
}

// SetSsoConnection gets a reference to the given string and assigns it to the SsoConnection field.
func (o *Organization) SetSsoConnection(v string) {
	o.SsoConnection = &v
}

// GetSsoOnly returns the SsoOnly field value if set, zero value otherwise.
func (o *Organization) GetSsoOnly() bool {
	if o == nil || IsNil(o.SsoOnly) {
		var ret bool
		return ret
	}
	return *o.SsoOnly
}

// GetSsoOnlyOk returns a tuple with the SsoOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetSsoOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.SsoOnly) {
		return nil, false
	}
	return o.SsoOnly, true
}

// HasSsoOnly returns a boolean if a field has been set.
func (o *Organization) HasSsoOnly() bool {
	if o != nil && !IsNil(o.SsoOnly) {
		return true
	}

	return false
}

// SetSsoOnly gets a reference to the given bool and assigns it to the SsoOnly field.
func (o *Organization) SetSsoOnly(v bool) {
	o.SsoOnly = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RocksetUser) {
		toSerialize["rockset_user"] = o.RocksetUser
	}
	if !IsNil(o.SsoConnection) {
		toSerialize["sso_connection"] = o.SsoConnection
	}
	if !IsNil(o.SsoOnly) {
		toSerialize["sso_only"] = o.SsoOnly
	}
	return toSerialize, nil
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


