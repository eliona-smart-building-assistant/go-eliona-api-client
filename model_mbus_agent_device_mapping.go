/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MbusAgentDeviceMapping struct for MbusAgentDeviceMapping
type MbusAgentDeviceMapping struct {
	Class NullableAgentClass `json:"class,omitempty"`
	// Unique id for the mapping
	Id NullableInt32 `json:"id,omitempty"`
	// The id of the device the mapping belongs to
	DeviceId NullableInt32 `json:"deviceId,omitempty"`
	// Is the mapping enabled or not
	Enable *bool `json:"enable,omitempty"`
	// ID of the corresponding asset
	AssetId NullableInt32 `json:"assetId,omitempty"`
	Subtype DataSubtype `json:"subtype"`
	// Name of the attribute to map
	Attribute string `json:"attribute"`
	Field NullableInt32 `json:"field,omitempty"`
	Scale NullableFloat64 `json:"scale,omitempty"`
	Zero NullableFloat64 `json:"zero,omitempty"`
}

// NewMbusAgentDeviceMapping instantiates a new MbusAgentDeviceMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbusAgentDeviceMapping(subtype DataSubtype, attribute string) *MbusAgentDeviceMapping {
	this := MbusAgentDeviceMapping{}
	var enable bool = true
	this.Enable = &enable
	this.Subtype = subtype
	this.Attribute = attribute
	return &this
}

// NewMbusAgentDeviceMappingWithDefaults instantiates a new MbusAgentDeviceMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbusAgentDeviceMappingWithDefaults() *MbusAgentDeviceMapping {
	this := MbusAgentDeviceMapping{}
	var enable bool = true
	this.Enable = &enable
	var subtype DataSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceMapping) GetClass() AgentClass {
	if o == nil || isNil(o.Class.Get()) {
		var ret AgentClass
		return ret
	}
	return *o.Class.Get()
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceMapping) GetClassOk() (*AgentClass, bool) {
	if o == nil {
    return nil, false
	}
	return o.Class.Get(), o.Class.IsSet()
}

// HasClass returns a boolean if a field has been set.
func (o *MbusAgentDeviceMapping) HasClass() bool {
	if o != nil && o.Class.IsSet() {
		return true
	}

	return false
}

// SetClass gets a reference to the given NullableAgentClass and assigns it to the Class field.
func (o *MbusAgentDeviceMapping) SetClass(v AgentClass) {
	o.Class.Set(&v)
}
// SetClassNil sets the value for Class to be an explicit nil
func (o *MbusAgentDeviceMapping) SetClassNil() {
	o.Class.Set(nil)
}

// UnsetClass ensures that no value is present for Class, not even an explicit nil
func (o *MbusAgentDeviceMapping) UnsetClass() {
	o.Class.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceMapping) GetId() int32 {
	if o == nil || isNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceMapping) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MbusAgentDeviceMapping) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *MbusAgentDeviceMapping) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MbusAgentDeviceMapping) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MbusAgentDeviceMapping) UnsetId() {
	o.Id.Unset()
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceMapping) GetDeviceId() int32 {
	if o == nil || isNil(o.DeviceId.Get()) {
		var ret int32
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceMapping) GetDeviceIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *MbusAgentDeviceMapping) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableInt32 and assigns it to the DeviceId field.
func (o *MbusAgentDeviceMapping) SetDeviceId(v int32) {
	o.DeviceId.Set(&v)
}
// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *MbusAgentDeviceMapping) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *MbusAgentDeviceMapping) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *MbusAgentDeviceMapping) GetEnable() bool {
	if o == nil || isNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbusAgentDeviceMapping) GetEnableOk() (*bool, bool) {
	if o == nil || isNil(o.Enable) {
    return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *MbusAgentDeviceMapping) HasEnable() bool {
	if o != nil && !isNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *MbusAgentDeviceMapping) SetEnable(v bool) {
	o.Enable = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceMapping) GetAssetId() int32 {
	if o == nil || isNil(o.AssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.AssetId.Get()
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceMapping) GetAssetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.AssetId.Get(), o.AssetId.IsSet()
}

// HasAssetId returns a boolean if a field has been set.
func (o *MbusAgentDeviceMapping) HasAssetId() bool {
	if o != nil && o.AssetId.IsSet() {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given NullableInt32 and assigns it to the AssetId field.
func (o *MbusAgentDeviceMapping) SetAssetId(v int32) {
	o.AssetId.Set(&v)
}
// SetAssetIdNil sets the value for AssetId to be an explicit nil
func (o *MbusAgentDeviceMapping) SetAssetIdNil() {
	o.AssetId.Set(nil)
}

// UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
func (o *MbusAgentDeviceMapping) UnsetAssetId() {
	o.AssetId.Unset()
}

// GetSubtype returns the Subtype field value
func (o *MbusAgentDeviceMapping) GetSubtype() DataSubtype {
	if o == nil {
		var ret DataSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *MbusAgentDeviceMapping) GetSubtypeOk() (*DataSubtype, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *MbusAgentDeviceMapping) SetSubtype(v DataSubtype) {
	o.Subtype = v
}

// GetAttribute returns the Attribute field value
func (o *MbusAgentDeviceMapping) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *MbusAgentDeviceMapping) GetAttributeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *MbusAgentDeviceMapping) SetAttribute(v string) {
	o.Attribute = v
}

// GetField returns the Field field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceMapping) GetField() int32 {
	if o == nil || isNil(o.Field.Get()) {
		var ret int32
		return ret
	}
	return *o.Field.Get()
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceMapping) GetFieldOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Field.Get(), o.Field.IsSet()
}

// HasField returns a boolean if a field has been set.
func (o *MbusAgentDeviceMapping) HasField() bool {
	if o != nil && o.Field.IsSet() {
		return true
	}

	return false
}

// SetField gets a reference to the given NullableInt32 and assigns it to the Field field.
func (o *MbusAgentDeviceMapping) SetField(v int32) {
	o.Field.Set(&v)
}
// SetFieldNil sets the value for Field to be an explicit nil
func (o *MbusAgentDeviceMapping) SetFieldNil() {
	o.Field.Set(nil)
}

// UnsetField ensures that no value is present for Field, not even an explicit nil
func (o *MbusAgentDeviceMapping) UnsetField() {
	o.Field.Unset()
}

// GetScale returns the Scale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceMapping) GetScale() float64 {
	if o == nil || isNil(o.Scale.Get()) {
		var ret float64
		return ret
	}
	return *o.Scale.Get()
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceMapping) GetScaleOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return o.Scale.Get(), o.Scale.IsSet()
}

// HasScale returns a boolean if a field has been set.
func (o *MbusAgentDeviceMapping) HasScale() bool {
	if o != nil && o.Scale.IsSet() {
		return true
	}

	return false
}

// SetScale gets a reference to the given NullableFloat64 and assigns it to the Scale field.
func (o *MbusAgentDeviceMapping) SetScale(v float64) {
	o.Scale.Set(&v)
}
// SetScaleNil sets the value for Scale to be an explicit nil
func (o *MbusAgentDeviceMapping) SetScaleNil() {
	o.Scale.Set(nil)
}

// UnsetScale ensures that no value is present for Scale, not even an explicit nil
func (o *MbusAgentDeviceMapping) UnsetScale() {
	o.Scale.Unset()
}

// GetZero returns the Zero field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceMapping) GetZero() float64 {
	if o == nil || isNil(o.Zero.Get()) {
		var ret float64
		return ret
	}
	return *o.Zero.Get()
}

// GetZeroOk returns a tuple with the Zero field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceMapping) GetZeroOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return o.Zero.Get(), o.Zero.IsSet()
}

// HasZero returns a boolean if a field has been set.
func (o *MbusAgentDeviceMapping) HasZero() bool {
	if o != nil && o.Zero.IsSet() {
		return true
	}

	return false
}

// SetZero gets a reference to the given NullableFloat64 and assigns it to the Zero field.
func (o *MbusAgentDeviceMapping) SetZero(v float64) {
	o.Zero.Set(&v)
}
// SetZeroNil sets the value for Zero to be an explicit nil
func (o *MbusAgentDeviceMapping) SetZeroNil() {
	o.Zero.Set(nil)
}

// UnsetZero ensures that no value is present for Zero, not even an explicit nil
func (o *MbusAgentDeviceMapping) UnsetZero() {
	o.Zero.Unset()
}

func (o MbusAgentDeviceMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Class.IsSet() {
		toSerialize["class"] = o.Class.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.DeviceId.IsSet() {
		toSerialize["deviceId"] = o.DeviceId.Get()
	}
	if !isNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if o.AssetId.IsSet() {
		toSerialize["assetId"] = o.AssetId.Get()
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if true {
		toSerialize["attribute"] = o.Attribute
	}
	if o.Field.IsSet() {
		toSerialize["field"] = o.Field.Get()
	}
	if o.Scale.IsSet() {
		toSerialize["scale"] = o.Scale.Get()
	}
	if o.Zero.IsSet() {
		toSerialize["zero"] = o.Zero.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMbusAgentDeviceMapping struct {
	value *MbusAgentDeviceMapping
	isSet bool
}

func (v NullableMbusAgentDeviceMapping) Get() *MbusAgentDeviceMapping {
	return v.value
}

func (v *NullableMbusAgentDeviceMapping) Set(val *MbusAgentDeviceMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableMbusAgentDeviceMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableMbusAgentDeviceMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbusAgentDeviceMapping(val *MbusAgentDeviceMapping) *NullableMbusAgentDeviceMapping {
	return &NullableMbusAgentDeviceMapping{value: val, isSet: true}
}

func (v NullableMbusAgentDeviceMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbusAgentDeviceMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


