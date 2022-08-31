/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Aggregation Defines the aggregation of data points
type Aggregation struct {
	// ID of the aggregation
	Id NullableInt32 `json:"id,omitempty"`
	// ID of the corresponding asset
	AssetId int32 `json:"assetId"`
	Subtype DataSubtype `json:"subtype"`
	// Name of the attribute which holds the data points
	Attribute *string `json:"attribute,omitempty"`
	// Calculation mode
	Mode NullableString `json:"mode,omitempty"`
	// calculation interval
	Raster string `json:"raster"`
}

// NewAggregation instantiates a new Aggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregation(assetId int32, subtype DataSubtype, raster string) *Aggregation {
	this := Aggregation{}
	this.AssetId = assetId
	this.Subtype = subtype
	this.Raster = raster
	return &this
}

// NewAggregationWithDefaults instantiates a new Aggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationWithDefaults() *Aggregation {
	this := Aggregation{}
	var subtype DataSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Aggregation) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Aggregation) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Aggregation) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Aggregation) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Aggregation) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Aggregation) UnsetId() {
	o.Id.Unset()
}

// GetAssetId returns the AssetId field value
func (o *Aggregation) GetAssetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *Aggregation) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *Aggregation) SetAssetId(v int32) {
	o.AssetId = v
}

// GetSubtype returns the Subtype field value
func (o *Aggregation) GetSubtype() DataSubtype {
	if o == nil {
		var ret DataSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *Aggregation) GetSubtypeOk() (*DataSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *Aggregation) SetSubtype(v DataSubtype) {
	o.Subtype = v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *Aggregation) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregation) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *Aggregation) HasAttribute() bool {
	if o != nil && o.Attribute != nil {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *Aggregation) SetAttribute(v string) {
	o.Attribute = &v
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Aggregation) GetMode() string {
	if o == nil || o.Mode.Get() == nil {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Aggregation) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *Aggregation) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given NullableString and assigns it to the Mode field.
func (o *Aggregation) SetMode(v string) {
	o.Mode.Set(&v)
}
// SetModeNil sets the value for Mode to be an explicit nil
func (o *Aggregation) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *Aggregation) UnsetMode() {
	o.Mode.Unset()
}

// GetRaster returns the Raster field value
func (o *Aggregation) GetRaster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Raster
}

// GetRasterOk returns a tuple with the Raster field value
// and a boolean to check if the value has been set.
func (o *Aggregation) GetRasterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Raster, true
}

// SetRaster sets field value
func (o *Aggregation) SetRaster(v string) {
	o.Raster = v
}

func (o Aggregation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if true {
		toSerialize["assetId"] = o.AssetId
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.Mode.IsSet() {
		toSerialize["mode"] = o.Mode.Get()
	}
	if true {
		toSerialize["raster"] = o.Raster
	}
	return json.Marshal(toSerialize)
}

type NullableAggregation struct {
	value *Aggregation
	isSet bool
}

func (v NullableAggregation) Get() *Aggregation {
	return v.value
}

func (v *NullableAggregation) Set(val *Aggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregation(val *Aggregation) *NullableAggregation {
	return &NullableAggregation{value: val, isSet: true}
}

func (v NullableAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

