/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.5.7
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ValueMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueMapping{}

// ValueMapping Mapping between value and custom text
type ValueMapping struct {
	// Value that should be mapped
	Value *int32 `json:"value,omitempty"`
	// Custom text to display instead of value
	Text *string `json:"text,omitempty"`
}

// NewValueMapping instantiates a new ValueMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueMapping() *ValueMapping {
	this := ValueMapping{}
	return &this
}

// NewValueMappingWithDefaults instantiates a new ValueMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueMappingWithDefaults() *ValueMapping {
	this := ValueMapping{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ValueMapping) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueMapping) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ValueMapping) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *ValueMapping) SetValue(v int32) {
	o.Value = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ValueMapping) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueMapping) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ValueMapping) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *ValueMapping) SetText(v string) {
	o.Text = &v
}

func (o ValueMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableValueMapping struct {
	value *ValueMapping
	isSet bool
}

func (v NullableValueMapping) Get() *ValueMapping {
	return v.value
}

func (v *NullableValueMapping) Set(val *ValueMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableValueMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableValueMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueMapping(val *ValueMapping) *NullableValueMapping {
	return &NullableValueMapping{value: val, isSet: true}
}

func (v NullableValueMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
