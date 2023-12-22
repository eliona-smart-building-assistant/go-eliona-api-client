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

// checks if the Tag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tag{}

// Tag A tag
type Tag struct {
	// The name of the tag
	Name string `json:"name"`
	// Id of the color
	ColorId NullableInt32 `json:"colorId,omitempty"`
	// Is this a tag custom or not
	Custom NullableBool `json:"custom,omitempty"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag(name string) *Tag {
	this := Tag{}
	this.Name = name
	var custom bool = true
	this.Custom = *NewNullableBool(&custom)
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	var custom bool = true
	this.Custom = *NewNullableBool(&custom)
	return &this
}

// GetName returns the Name field value
func (o *Tag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tag) SetName(v string) {
	o.Name = v
}

// GetColorId returns the ColorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetColorId() int32 {
	if o == nil || IsNil(o.ColorId.Get()) {
		var ret int32
		return ret
	}
	return *o.ColorId.Get()
}

// GetColorIdOk returns a tuple with the ColorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetColorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ColorId.Get(), o.ColorId.IsSet()
}

// HasColorId returns a boolean if a field has been set.
func (o *Tag) HasColorId() bool {
	if o != nil && o.ColorId.IsSet() {
		return true
	}

	return false
}

// SetColorId gets a reference to the given NullableInt32 and assigns it to the ColorId field.
func (o *Tag) SetColorId(v int32) {
	o.ColorId.Set(&v)
}

// SetColorIdNil sets the value for ColorId to be an explicit nil
func (o *Tag) SetColorIdNil() {
	o.ColorId.Set(nil)
}

// UnsetColorId ensures that no value is present for ColorId, not even an explicit nil
func (o *Tag) UnsetColorId() {
	o.ColorId.Unset()
}

// GetCustom returns the Custom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetCustom() bool {
	if o == nil || IsNil(o.Custom.Get()) {
		var ret bool
		return ret
	}
	return *o.Custom.Get()
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetCustomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Custom.Get(), o.Custom.IsSet()
}

// HasCustom returns a boolean if a field has been set.
func (o *Tag) HasCustom() bool {
	if o != nil && o.Custom.IsSet() {
		return true
	}

	return false
}

// SetCustom gets a reference to the given NullableBool and assigns it to the Custom field.
func (o *Tag) SetCustom(v bool) {
	o.Custom.Set(&v)
}

// SetCustomNil sets the value for Custom to be an explicit nil
func (o *Tag) SetCustomNil() {
	o.Custom.Set(nil)
}

// UnsetCustom ensures that no value is present for Custom, not even an explicit nil
func (o *Tag) UnsetCustom() {
	o.Custom.Unset()
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.ColorId.IsSet() {
		toSerialize["colorId"] = o.ColorId.Get()
	}
	if o.Custom.IsSet() {
		toSerialize["custom"] = o.Custom.Get()
	}
	return toSerialize, nil
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
