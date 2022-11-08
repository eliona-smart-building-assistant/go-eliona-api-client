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

// WidgetTypeElement An element for widget types
type WidgetTypeElement struct {
	// The internal Id of widget element
	Id NullableInt32 `json:"id,omitempty"`
	// The category for this element
	Category string `json:"category"`
	// sequence of element in widget; if not defined the index in array is taken
	Sequence NullableInt32 `json:"sequence,omitempty"`
	// individual config parameters depending on category
	Config map[string]interface{} `json:"config,omitempty"`
}

// NewWidgetTypeElement instantiates a new WidgetTypeElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetTypeElement(category string) *WidgetTypeElement {
	this := WidgetTypeElement{}
	this.Category = category
	return &this
}

// NewWidgetTypeElementWithDefaults instantiates a new WidgetTypeElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetTypeElementWithDefaults() *WidgetTypeElement {
	this := WidgetTypeElement{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetTypeElement) GetId() int32 {
	if o == nil || isNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetTypeElement) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *WidgetTypeElement) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *WidgetTypeElement) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *WidgetTypeElement) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *WidgetTypeElement) UnsetId() {
	o.Id.Unset()
}

// GetCategory returns the Category field value
func (o *WidgetTypeElement) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *WidgetTypeElement) GetCategoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *WidgetTypeElement) SetCategory(v string) {
	o.Category = v
}

// GetSequence returns the Sequence field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetTypeElement) GetSequence() int32 {
	if o == nil || isNil(o.Sequence.Get()) {
		var ret int32
		return ret
	}
	return *o.Sequence.Get()
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetTypeElement) GetSequenceOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Sequence.Get(), o.Sequence.IsSet()
}

// HasSequence returns a boolean if a field has been set.
func (o *WidgetTypeElement) HasSequence() bool {
	if o != nil && o.Sequence.IsSet() {
		return true
	}

	return false
}

// SetSequence gets a reference to the given NullableInt32 and assigns it to the Sequence field.
func (o *WidgetTypeElement) SetSequence(v int32) {
	o.Sequence.Set(&v)
}
// SetSequenceNil sets the value for Sequence to be an explicit nil
func (o *WidgetTypeElement) SetSequenceNil() {
	o.Sequence.Set(nil)
}

// UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
func (o *WidgetTypeElement) UnsetSequence() {
	o.Sequence.Unset()
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetTypeElement) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetTypeElement) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Config) {
    return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *WidgetTypeElement) HasConfig() bool {
	if o != nil && isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *WidgetTypeElement) SetConfig(v map[string]interface{}) {
	o.Config = v
}

func (o WidgetTypeElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if o.Sequence.IsSet() {
		toSerialize["sequence"] = o.Sequence.Get()
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableWidgetTypeElement struct {
	value *WidgetTypeElement
	isSet bool
}

func (v NullableWidgetTypeElement) Get() *WidgetTypeElement {
	return v.value
}

func (v *NullableWidgetTypeElement) Set(val *WidgetTypeElement) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetTypeElement) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetTypeElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetTypeElement(val *WidgetTypeElement) *NullableWidgetTypeElement {
	return &NullableWidgetTypeElement{value: val, isSet: true}
}

func (v NullableWidgetTypeElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetTypeElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


