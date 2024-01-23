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

// checks if the FormatParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormatParams{}

// FormatParams struct for FormatParams
type FormatParams struct {
	Avro map[string]interface{} `json:"avro,omitempty"`
	Csv *CsvParams `json:"csv,omitempty"`
	// Source data is in json format.
	Json *bool `json:"json,omitempty"`
	MssqlDms *bool `json:"mssql_dms,omitempty"`
	MysqlDms *bool `json:"mysql_dms,omitempty"`
	OracleDms *bool `json:"oracle_dms,omitempty"`
	PostgresDms *bool `json:"postgres_dms,omitempty"`
	Xml *XmlParams `json:"xml,omitempty"`
}

// NewFormatParams instantiates a new FormatParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormatParams() *FormatParams {
	this := FormatParams{}
	return &this
}

// NewFormatParamsWithDefaults instantiates a new FormatParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatParamsWithDefaults() *FormatParams {
	this := FormatParams{}
	return &this
}

// GetAvro returns the Avro field value if set, zero value otherwise.
func (o *FormatParams) GetAvro() map[string]interface{} {
	if o == nil || IsNil(o.Avro) {
		var ret map[string]interface{}
		return ret
	}
	return o.Avro
}

// GetAvroOk returns a tuple with the Avro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatParams) GetAvroOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Avro) {
		return map[string]interface{}{}, false
	}
	return o.Avro, true
}

// HasAvro returns a boolean if a field has been set.
func (o *FormatParams) HasAvro() bool {
	if o != nil && !IsNil(o.Avro) {
		return true
	}

	return false
}

// SetAvro gets a reference to the given map[string]interface{} and assigns it to the Avro field.
func (o *FormatParams) SetAvro(v map[string]interface{}) {
	o.Avro = v
}

// GetCsv returns the Csv field value if set, zero value otherwise.
func (o *FormatParams) GetCsv() CsvParams {
	if o == nil || IsNil(o.Csv) {
		var ret CsvParams
		return ret
	}
	return *o.Csv
}

// GetCsvOk returns a tuple with the Csv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatParams) GetCsvOk() (*CsvParams, bool) {
	if o == nil || IsNil(o.Csv) {
		return nil, false
	}
	return o.Csv, true
}

// HasCsv returns a boolean if a field has been set.
func (o *FormatParams) HasCsv() bool {
	if o != nil && !IsNil(o.Csv) {
		return true
	}

	return false
}

// SetCsv gets a reference to the given CsvParams and assigns it to the Csv field.
func (o *FormatParams) SetCsv(v CsvParams) {
	o.Csv = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *FormatParams) GetJson() bool {
	if o == nil || IsNil(o.Json) {
		var ret bool
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatParams) GetJsonOk() (*bool, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *FormatParams) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given bool and assigns it to the Json field.
func (o *FormatParams) SetJson(v bool) {
	o.Json = &v
}

// GetMssqlDms returns the MssqlDms field value if set, zero value otherwise.
func (o *FormatParams) GetMssqlDms() bool {
	if o == nil || IsNil(o.MssqlDms) {
		var ret bool
		return ret
	}
	return *o.MssqlDms
}

// GetMssqlDmsOk returns a tuple with the MssqlDms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatParams) GetMssqlDmsOk() (*bool, bool) {
	if o == nil || IsNil(o.MssqlDms) {
		return nil, false
	}
	return o.MssqlDms, true
}

// HasMssqlDms returns a boolean if a field has been set.
func (o *FormatParams) HasMssqlDms() bool {
	if o != nil && !IsNil(o.MssqlDms) {
		return true
	}

	return false
}

// SetMssqlDms gets a reference to the given bool and assigns it to the MssqlDms field.
func (o *FormatParams) SetMssqlDms(v bool) {
	o.MssqlDms = &v
}

// GetMysqlDms returns the MysqlDms field value if set, zero value otherwise.
func (o *FormatParams) GetMysqlDms() bool {
	if o == nil || IsNil(o.MysqlDms) {
		var ret bool
		return ret
	}
	return *o.MysqlDms
}

// GetMysqlDmsOk returns a tuple with the MysqlDms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatParams) GetMysqlDmsOk() (*bool, bool) {
	if o == nil || IsNil(o.MysqlDms) {
		return nil, false
	}
	return o.MysqlDms, true
}

// HasMysqlDms returns a boolean if a field has been set.
func (o *FormatParams) HasMysqlDms() bool {
	if o != nil && !IsNil(o.MysqlDms) {
		return true
	}

	return false
}

// SetMysqlDms gets a reference to the given bool and assigns it to the MysqlDms field.
func (o *FormatParams) SetMysqlDms(v bool) {
	o.MysqlDms = &v
}

// GetOracleDms returns the OracleDms field value if set, zero value otherwise.
func (o *FormatParams) GetOracleDms() bool {
	if o == nil || IsNil(o.OracleDms) {
		var ret bool
		return ret
	}
	return *o.OracleDms
}

// GetOracleDmsOk returns a tuple with the OracleDms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatParams) GetOracleDmsOk() (*bool, bool) {
	if o == nil || IsNil(o.OracleDms) {
		return nil, false
	}
	return o.OracleDms, true
}

// HasOracleDms returns a boolean if a field has been set.
func (o *FormatParams) HasOracleDms() bool {
	if o != nil && !IsNil(o.OracleDms) {
		return true
	}

	return false
}

// SetOracleDms gets a reference to the given bool and assigns it to the OracleDms field.
func (o *FormatParams) SetOracleDms(v bool) {
	o.OracleDms = &v
}

// GetPostgresDms returns the PostgresDms field value if set, zero value otherwise.
func (o *FormatParams) GetPostgresDms() bool {
	if o == nil || IsNil(o.PostgresDms) {
		var ret bool
		return ret
	}
	return *o.PostgresDms
}

// GetPostgresDmsOk returns a tuple with the PostgresDms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatParams) GetPostgresDmsOk() (*bool, bool) {
	if o == nil || IsNil(o.PostgresDms) {
		return nil, false
	}
	return o.PostgresDms, true
}

// HasPostgresDms returns a boolean if a field has been set.
func (o *FormatParams) HasPostgresDms() bool {
	if o != nil && !IsNil(o.PostgresDms) {
		return true
	}

	return false
}

// SetPostgresDms gets a reference to the given bool and assigns it to the PostgresDms field.
func (o *FormatParams) SetPostgresDms(v bool) {
	o.PostgresDms = &v
}

// GetXml returns the Xml field value if set, zero value otherwise.
func (o *FormatParams) GetXml() XmlParams {
	if o == nil || IsNil(o.Xml) {
		var ret XmlParams
		return ret
	}
	return *o.Xml
}

// GetXmlOk returns a tuple with the Xml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatParams) GetXmlOk() (*XmlParams, bool) {
	if o == nil || IsNil(o.Xml) {
		return nil, false
	}
	return o.Xml, true
}

// HasXml returns a boolean if a field has been set.
func (o *FormatParams) HasXml() bool {
	if o != nil && !IsNil(o.Xml) {
		return true
	}

	return false
}

// SetXml gets a reference to the given XmlParams and assigns it to the Xml field.
func (o *FormatParams) SetXml(v XmlParams) {
	o.Xml = &v
}

func (o FormatParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormatParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Avro) {
		toSerialize["avro"] = o.Avro
	}
	if !IsNil(o.Csv) {
		toSerialize["csv"] = o.Csv
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}
	if !IsNil(o.MssqlDms) {
		toSerialize["mssql_dms"] = o.MssqlDms
	}
	if !IsNil(o.MysqlDms) {
		toSerialize["mysql_dms"] = o.MysqlDms
	}
	if !IsNil(o.OracleDms) {
		toSerialize["oracle_dms"] = o.OracleDms
	}
	if !IsNil(o.PostgresDms) {
		toSerialize["postgres_dms"] = o.PostgresDms
	}
	if !IsNil(o.Xml) {
		toSerialize["xml"] = o.Xml
	}
	return toSerialize, nil
}

type NullableFormatParams struct {
	value *FormatParams
	isSet bool
}

func (v NullableFormatParams) Get() *FormatParams {
	return v.value
}

func (v *NullableFormatParams) Set(val *FormatParams) {
	v.value = val
	v.isSet = true
}

func (v NullableFormatParams) IsSet() bool {
	return v.isSet
}

func (v *NullableFormatParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormatParams(val *FormatParams) *NullableFormatParams {
	return &NullableFormatParams{value: val, isSet: true}
}

func (v NullableFormatParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormatParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

