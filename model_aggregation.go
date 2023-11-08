/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.5.5
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the Aggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Aggregation{}

// Aggregation Defines the aggregation of data points
type Aggregation struct {
	// ID of the aggregation
	Id NullableInt32 `json:"id,omitempty"`
	// ID of the corresponding asset
	AssetId int32       `json:"assetId"`
	Subtype DataSubtype `json:"subtype"`
	// Name of the attribute which holds the data points
	Attribute *string `json:"attribute,omitempty"`
	// Calculation mode
	Mode string `json:"mode"`
	// calculation interval
	Raster NullableString `json:"raster,omitempty"`
}

// NewAggregation instantiates a new Aggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregation(assetId int32, subtype DataSubtype, mode string) *Aggregation {
	this := Aggregation{}
	this.AssetId = assetId
	this.Subtype = subtype
	this.Mode = mode
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
	if o == nil || IsNil(o.Id.Get()) {
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
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregation) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *Aggregation) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *Aggregation) SetAttribute(v string) {
	o.Attribute = &v
}

// GetMode returns the Mode field value
func (o *Aggregation) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *Aggregation) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *Aggregation) SetMode(v string) {
	o.Mode = v
}

// GetRaster returns the Raster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Aggregation) GetRaster() string {
	if o == nil || IsNil(o.Raster.Get()) {
		var ret string
		return ret
	}
	return *o.Raster.Get()
}

// GetRasterOk returns a tuple with the Raster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Aggregation) GetRasterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Raster.Get(), o.Raster.IsSet()
}

// HasRaster returns a boolean if a field has been set.
func (o *Aggregation) HasRaster() bool {
	if o != nil && o.Raster.IsSet() {
		return true
	}

	return false
}

// SetRaster gets a reference to the given NullableString and assigns it to the Raster field.
func (o *Aggregation) SetRaster(v string) {
	o.Raster.Set(&v)
}

// SetRasterNil sets the value for Raster to be an explicit nil
func (o *Aggregation) SetRasterNil() {
	o.Raster.Set(nil)
}

// UnsetRaster ensures that no value is present for Raster, not even an explicit nil
func (o *Aggregation) UnsetRaster() {
	o.Raster.Unset()
}

func (o Aggregation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Aggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["assetId"] = o.AssetId
	toSerialize["subtype"] = o.Subtype
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	toSerialize["mode"] = o.Mode
	if o.Raster.IsSet() {
		toSerialize["raster"] = o.Raster.Get()
	}
	return toSerialize, nil
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
