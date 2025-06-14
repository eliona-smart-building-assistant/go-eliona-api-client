/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.9.4
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Data type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Data{}

// Data Data for assets
type Data struct {
	// ID of the corresponding asset
	AssetId int32       `json:"assetId"`
	Subtype DataSubtype `json:"subtype"`
	// Timestamp of the latest data change
	Timestamp NullableTime `json:"timestamp,omitempty"`
	// Asset payload
	Data map[string]interface{} `json:"data"`
	// The name of the corresponding asset type
	// Deprecated
	AssetTypeName NullableString `json:"assetTypeName,omitempty"`
	// freely assignable by the client to identify self-created data
	ClientReference NullableString `json:"clientReference,omitempty"`
}

type _Data Data

// NewData instantiates a new Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewData(assetId int32, subtype DataSubtype, data map[string]interface{}) *Data {
	this := Data{}
	this.AssetId = assetId
	this.Subtype = subtype
	this.Data = data
	return &this
}

// NewDataWithDefaults instantiates a new Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataWithDefaults() *Data {
	this := Data{}
	var subtype DataSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	return &this
}

// GetAssetId returns the AssetId field value
func (o *Data) GetAssetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *Data) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *Data) SetAssetId(v int32) {
	o.AssetId = v
}

// GetSubtype returns the Subtype field value
func (o *Data) GetSubtype() DataSubtype {
	if o == nil {
		var ret DataSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *Data) GetSubtypeOk() (*DataSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *Data) SetSubtype(v DataSubtype) {
	o.Subtype = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Data) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Data) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Data) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *Data) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *Data) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *Data) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetData returns the Data field value
func (o *Data) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Data) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *Data) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetAssetTypeName returns the AssetTypeName field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *Data) GetAssetTypeName() string {
	if o == nil || IsNil(o.AssetTypeName.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTypeName.Get()
}

// GetAssetTypeNameOk returns a tuple with the AssetTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *Data) GetAssetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTypeName.Get(), o.AssetTypeName.IsSet()
}

// HasAssetTypeName returns a boolean if a field has been set.
func (o *Data) HasAssetTypeName() bool {
	if o != nil && o.AssetTypeName.IsSet() {
		return true
	}

	return false
}

// SetAssetTypeName gets a reference to the given NullableString and assigns it to the AssetTypeName field.
// Deprecated
func (o *Data) SetAssetTypeName(v string) {
	o.AssetTypeName.Set(&v)
}

// SetAssetTypeNameNil sets the value for AssetTypeName to be an explicit nil
func (o *Data) SetAssetTypeNameNil() {
	o.AssetTypeName.Set(nil)
}

// UnsetAssetTypeName ensures that no value is present for AssetTypeName, not even an explicit nil
func (o *Data) UnsetAssetTypeName() {
	o.AssetTypeName.Unset()
}

// GetClientReference returns the ClientReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Data) GetClientReference() string {
	if o == nil || IsNil(o.ClientReference.Get()) {
		var ret string
		return ret
	}
	return *o.ClientReference.Get()
}

// GetClientReferenceOk returns a tuple with the ClientReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Data) GetClientReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientReference.Get(), o.ClientReference.IsSet()
}

// HasClientReference returns a boolean if a field has been set.
func (o *Data) HasClientReference() bool {
	if o != nil && o.ClientReference.IsSet() {
		return true
	}

	return false
}

// SetClientReference gets a reference to the given NullableString and assigns it to the ClientReference field.
func (o *Data) SetClientReference(v string) {
	o.ClientReference.Set(&v)
}

// SetClientReferenceNil sets the value for ClientReference to be an explicit nil
func (o *Data) SetClientReferenceNil() {
	o.ClientReference.Set(nil)
}

// UnsetClientReference ensures that no value is present for ClientReference, not even an explicit nil
func (o *Data) UnsetClientReference() {
	o.ClientReference.Unset()
}

func (o Data) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assetId"] = o.AssetId
	toSerialize["subtype"] = o.Subtype
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	toSerialize["data"] = o.Data
	if o.AssetTypeName.IsSet() {
		toSerialize["assetTypeName"] = o.AssetTypeName.Get()
	}
	if o.ClientReference.IsSet() {
		toSerialize["clientReference"] = o.ClientReference.Get()
	}
	return toSerialize, nil
}

func (o *Data) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assetId",
		"subtype",
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varData := _Data{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varData)

	if err != nil {
		return err
	}

	*o = Data(varData)

	return err
}

type NullableData struct {
	value *Data
	isSet bool
}

func (v NullableData) Get() *Data {
	return v.value
}

func (v *NullableData) Set(val *Data) {
	v.value = val
	v.isSet = true
}

func (v NullableData) IsSet() bool {
	return v.isSet
}

func (v *NullableData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableData(val *Data) *NullableData {
	return &NullableData{value: val, isSet: true}
}

func (v NullableData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
