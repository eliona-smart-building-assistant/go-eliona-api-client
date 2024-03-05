/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.6
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

// checks if the DataListen type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataListen{}

// DataListen struct for DataListen
type DataListen struct {
	// ID of the corresponding asset
	AssetId int32       `json:"assetId"`
	Subtype DataSubtype `json:"subtype"`
	// Timestamp of the latest data change
	Timestamp NullableTime `json:"timestamp,omitempty"`
	// Asset payload
	Data map[string]interface{} `json:"data"`
	// The name of the corresponding asset type
	AssetTypeName NullableString `json:"assetTypeName,omitempty"`
	// freely assignable by the client to identify self-created data
	ClientReference NullableString `json:"clientReference,omitempty"`
	// The status code expecting when actually perform the operation. Some values are - 200: updated (ok)  - 201: created - 204: deleted (no content) - 304: unchanged (not modified) - 400: problem (bad request) - 404: not found - 409: duplicated (conflict) - 422: unprocessable
	StatusCode *int32 `json:"statusCode,omitempty"`
}

type _DataListen DataListen

// NewDataListen instantiates a new DataListen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataListen(assetId int32, subtype DataSubtype, data map[string]interface{}) *DataListen {
	this := DataListen{}
	this.AssetId = assetId
	this.Subtype = subtype
	this.Data = data
	return &this
}

// NewDataListenWithDefaults instantiates a new DataListen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataListenWithDefaults() *DataListen {
	this := DataListen{}
	var subtype DataSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	return &this
}

// GetAssetId returns the AssetId field value
func (o *DataListen) GetAssetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *DataListen) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *DataListen) SetAssetId(v int32) {
	o.AssetId = v
}

// GetSubtype returns the Subtype field value
func (o *DataListen) GetSubtype() DataSubtype {
	if o == nil {
		var ret DataSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *DataListen) GetSubtypeOk() (*DataSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *DataListen) SetSubtype(v DataSubtype) {
	o.Subtype = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataListen) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataListen) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DataListen) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *DataListen) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *DataListen) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *DataListen) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetData returns the Data field value
func (o *DataListen) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DataListen) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DataListen) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetAssetTypeName returns the AssetTypeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataListen) GetAssetTypeName() string {
	if o == nil || IsNil(o.AssetTypeName.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTypeName.Get()
}

// GetAssetTypeNameOk returns a tuple with the AssetTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataListen) GetAssetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTypeName.Get(), o.AssetTypeName.IsSet()
}

// HasAssetTypeName returns a boolean if a field has been set.
func (o *DataListen) HasAssetTypeName() bool {
	if o != nil && o.AssetTypeName.IsSet() {
		return true
	}

	return false
}

// SetAssetTypeName gets a reference to the given NullableString and assigns it to the AssetTypeName field.
func (o *DataListen) SetAssetTypeName(v string) {
	o.AssetTypeName.Set(&v)
}

// SetAssetTypeNameNil sets the value for AssetTypeName to be an explicit nil
func (o *DataListen) SetAssetTypeNameNil() {
	o.AssetTypeName.Set(nil)
}

// UnsetAssetTypeName ensures that no value is present for AssetTypeName, not even an explicit nil
func (o *DataListen) UnsetAssetTypeName() {
	o.AssetTypeName.Unset()
}

// GetClientReference returns the ClientReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataListen) GetClientReference() string {
	if o == nil || IsNil(o.ClientReference.Get()) {
		var ret string
		return ret
	}
	return *o.ClientReference.Get()
}

// GetClientReferenceOk returns a tuple with the ClientReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataListen) GetClientReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientReference.Get(), o.ClientReference.IsSet()
}

// HasClientReference returns a boolean if a field has been set.
func (o *DataListen) HasClientReference() bool {
	if o != nil && o.ClientReference.IsSet() {
		return true
	}

	return false
}

// SetClientReference gets a reference to the given NullableString and assigns it to the ClientReference field.
func (o *DataListen) SetClientReference(v string) {
	o.ClientReference.Set(&v)
}

// SetClientReferenceNil sets the value for ClientReference to be an explicit nil
func (o *DataListen) SetClientReferenceNil() {
	o.ClientReference.Set(nil)
}

// UnsetClientReference ensures that no value is present for ClientReference, not even an explicit nil
func (o *DataListen) UnsetClientReference() {
	o.ClientReference.Unset()
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *DataListen) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataListen) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *DataListen) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *DataListen) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o DataListen) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataListen) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	return toSerialize, nil
}

func (o *DataListen) UnmarshalJSON(data []byte) (err error) {
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

	varDataListen := _DataListen{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataListen)

	if err != nil {
		return err
	}

	*o = DataListen(varDataListen)

	return err
}

type NullableDataListen struct {
	value *DataListen
	isSet bool
}

func (v NullableDataListen) Get() *DataListen {
	return v.value
}

func (v *NullableDataListen) Set(val *DataListen) {
	v.value = val
	v.isSet = true
}

func (v NullableDataListen) IsSet() bool {
	return v.isSet
}

func (v *NullableDataListen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataListen(val *DataListen) *NullableDataListen {
	return &NullableDataListen{value: val, isSet: true}
}

func (v NullableDataListen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataListen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
