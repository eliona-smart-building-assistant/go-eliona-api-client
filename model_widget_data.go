/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.8
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the WidgetData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetData{}

// WidgetData Data for a widget
type WidgetData struct {
	// The internal Id of widget data
	Id NullableInt32 `json:"id,omitempty"`
	// Position of the element in widget type
	ElementSequence NullableInt32 `json:"elementSequence,omitempty"`
	// The master asset id of this widget
	AssetId NullableInt32 `json:"assetId,omitempty"`
	// individual config parameters depending on category
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewWidgetData instantiates a new WidgetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetData() *WidgetData {
	this := WidgetData{}
	return &this
}

// NewWidgetDataWithDefaults instantiates a new WidgetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetDataWithDefaults() *WidgetData {
	this := WidgetData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetData) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetData) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *WidgetData) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *WidgetData) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *WidgetData) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *WidgetData) UnsetId() {
	o.Id.Unset()
}

// GetElementSequence returns the ElementSequence field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetData) GetElementSequence() int32 {
	if o == nil || IsNil(o.ElementSequence.Get()) {
		var ret int32
		return ret
	}
	return *o.ElementSequence.Get()
}

// GetElementSequenceOk returns a tuple with the ElementSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetData) GetElementSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ElementSequence.Get(), o.ElementSequence.IsSet()
}

// HasElementSequence returns a boolean if a field has been set.
func (o *WidgetData) HasElementSequence() bool {
	if o != nil && o.ElementSequence.IsSet() {
		return true
	}

	return false
}

// SetElementSequence gets a reference to the given NullableInt32 and assigns it to the ElementSequence field.
func (o *WidgetData) SetElementSequence(v int32) {
	o.ElementSequence.Set(&v)
}

// SetElementSequenceNil sets the value for ElementSequence to be an explicit nil
func (o *WidgetData) SetElementSequenceNil() {
	o.ElementSequence.Set(nil)
}

// UnsetElementSequence ensures that no value is present for ElementSequence, not even an explicit nil
func (o *WidgetData) UnsetElementSequence() {
	o.ElementSequence.Unset()
}

// GetAssetId returns the AssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetData) GetAssetId() int32 {
	if o == nil || IsNil(o.AssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.AssetId.Get()
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetData) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetId.Get(), o.AssetId.IsSet()
}

// HasAssetId returns a boolean if a field has been set.
func (o *WidgetData) HasAssetId() bool {
	if o != nil && o.AssetId.IsSet() {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given NullableInt32 and assigns it to the AssetId field.
func (o *WidgetData) SetAssetId(v int32) {
	o.AssetId.Set(&v)
}

// SetAssetIdNil sets the value for AssetId to be an explicit nil
func (o *WidgetData) SetAssetIdNil() {
	o.AssetId.Set(nil)
}

// UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
func (o *WidgetData) UnsetAssetId() {
	o.AssetId.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetData) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetData) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WidgetData) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *WidgetData) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o WidgetData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ElementSequence.IsSet() {
		toSerialize["elementSequence"] = o.ElementSequence.Get()
	}
	if o.AssetId.IsSet() {
		toSerialize["assetId"] = o.AssetId.Get()
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableWidgetData struct {
	value *WidgetData
	isSet bool
}

func (v NullableWidgetData) Get() *WidgetData {
	return v.value
}

func (v *NullableWidgetData) Set(val *WidgetData) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetData) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetData(val *WidgetData) *NullableWidgetData {
	return &NullableWidgetData{value: val, isSet: true}
}

func (v NullableWidgetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
