/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.4.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MbusAgentDeviceMappingSpecific Specific mapping for MBUS agents
type MbusAgentDeviceMappingSpecific struct {
	Field NullableInt32 `json:"field,omitempty"`
	Scale NullableFloat64 `json:"scale,omitempty"`
	Zero NullableFloat64 `json:"zero,omitempty"`
}

// NewMbusAgentDeviceMappingSpecific instantiates a new MbusAgentDeviceMappingSpecific object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbusAgentDeviceMappingSpecific() *MbusAgentDeviceMappingSpecific {
	this := MbusAgentDeviceMappingSpecific{}
	return &this
}

// NewMbusAgentDeviceMappingSpecificWithDefaults instantiates a new MbusAgentDeviceMappingSpecific object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbusAgentDeviceMappingSpecificWithDefaults() *MbusAgentDeviceMappingSpecific {
	this := MbusAgentDeviceMappingSpecific{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceMappingSpecific) GetField() int32 {
	if o == nil || isNil(o.Field.Get()) {
		var ret int32
		return ret
	}
	return *o.Field.Get()
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceMappingSpecific) GetFieldOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Field.Get(), o.Field.IsSet()
}

// HasField returns a boolean if a field has been set.
func (o *MbusAgentDeviceMappingSpecific) HasField() bool {
	if o != nil && o.Field.IsSet() {
		return true
	}

	return false
}

// SetField gets a reference to the given NullableInt32 and assigns it to the Field field.
func (o *MbusAgentDeviceMappingSpecific) SetField(v int32) {
	o.Field.Set(&v)
}
// SetFieldNil sets the value for Field to be an explicit nil
func (o *MbusAgentDeviceMappingSpecific) SetFieldNil() {
	o.Field.Set(nil)
}

// UnsetField ensures that no value is present for Field, not even an explicit nil
func (o *MbusAgentDeviceMappingSpecific) UnsetField() {
	o.Field.Unset()
}

// GetScale returns the Scale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceMappingSpecific) GetScale() float64 {
	if o == nil || isNil(o.Scale.Get()) {
		var ret float64
		return ret
	}
	return *o.Scale.Get()
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceMappingSpecific) GetScaleOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return o.Scale.Get(), o.Scale.IsSet()
}

// HasScale returns a boolean if a field has been set.
func (o *MbusAgentDeviceMappingSpecific) HasScale() bool {
	if o != nil && o.Scale.IsSet() {
		return true
	}

	return false
}

// SetScale gets a reference to the given NullableFloat64 and assigns it to the Scale field.
func (o *MbusAgentDeviceMappingSpecific) SetScale(v float64) {
	o.Scale.Set(&v)
}
// SetScaleNil sets the value for Scale to be an explicit nil
func (o *MbusAgentDeviceMappingSpecific) SetScaleNil() {
	o.Scale.Set(nil)
}

// UnsetScale ensures that no value is present for Scale, not even an explicit nil
func (o *MbusAgentDeviceMappingSpecific) UnsetScale() {
	o.Scale.Unset()
}

// GetZero returns the Zero field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceMappingSpecific) GetZero() float64 {
	if o == nil || isNil(o.Zero.Get()) {
		var ret float64
		return ret
	}
	return *o.Zero.Get()
}

// GetZeroOk returns a tuple with the Zero field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceMappingSpecific) GetZeroOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return o.Zero.Get(), o.Zero.IsSet()
}

// HasZero returns a boolean if a field has been set.
func (o *MbusAgentDeviceMappingSpecific) HasZero() bool {
	if o != nil && o.Zero.IsSet() {
		return true
	}

	return false
}

// SetZero gets a reference to the given NullableFloat64 and assigns it to the Zero field.
func (o *MbusAgentDeviceMappingSpecific) SetZero(v float64) {
	o.Zero.Set(&v)
}
// SetZeroNil sets the value for Zero to be an explicit nil
func (o *MbusAgentDeviceMappingSpecific) SetZeroNil() {
	o.Zero.Set(nil)
}

// UnsetZero ensures that no value is present for Zero, not even an explicit nil
func (o *MbusAgentDeviceMappingSpecific) UnsetZero() {
	o.Zero.Unset()
}

func (o MbusAgentDeviceMappingSpecific) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableMbusAgentDeviceMappingSpecific struct {
	value *MbusAgentDeviceMappingSpecific
	isSet bool
}

func (v NullableMbusAgentDeviceMappingSpecific) Get() *MbusAgentDeviceMappingSpecific {
	return v.value
}

func (v *NullableMbusAgentDeviceMappingSpecific) Set(val *MbusAgentDeviceMappingSpecific) {
	v.value = val
	v.isSet = true
}

func (v NullableMbusAgentDeviceMappingSpecific) IsSet() bool {
	return v.isSet
}

func (v *NullableMbusAgentDeviceMappingSpecific) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbusAgentDeviceMappingSpecific(val *MbusAgentDeviceMappingSpecific) *NullableMbusAgentDeviceMappingSpecific {
	return &NullableMbusAgentDeviceMappingSpecific{value: val, isSet: true}
}

func (v NullableMbusAgentDeviceMappingSpecific) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbusAgentDeviceMappingSpecific) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


