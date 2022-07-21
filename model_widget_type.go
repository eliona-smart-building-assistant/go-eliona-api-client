/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// WidgetType A frontend widget
type WidgetType struct {
	// The unique name for this widget type
	Name string `json:"name"`
	// Is this a customer created type or not
	Custom bool `json:"custom"`
	Translation Translation `json:"translation"`
	// Icon name corresponding to assets used in this widget
	Icon *string `json:"icon,omitempty"`
	// Show alarms in widget
	WithAlarm *bool `json:"withAlarm,omitempty"`
	// Show selection for timespan in widget
	WithTimespan *bool `json:"withTimespan,omitempty"`
	// A list of elements for this widget
	Elements []WidgetTypeElement `json:"elements"`
}

// NewWidgetType instantiates a new WidgetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetType(name string, custom bool, translation Translation, elements []WidgetTypeElement) *WidgetType {
	this := WidgetType{}
	this.Name = name
	this.Custom = custom
	this.Translation = translation
	var withAlarm bool = false
	this.WithAlarm = &withAlarm
	var withTimespan bool = false
	this.WithTimespan = &withTimespan
	this.Elements = elements
	return &this
}

// NewWidgetTypeWithDefaults instantiates a new WidgetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetTypeWithDefaults() *WidgetType {
	this := WidgetType{}
	var custom bool = true
	this.Custom = custom
	var withAlarm bool = false
	this.WithAlarm = &withAlarm
	var withTimespan bool = false
	this.WithTimespan = &withTimespan
	return &this
}

// GetName returns the Name field value
func (o *WidgetType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WidgetType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WidgetType) SetName(v string) {
	o.Name = v
}

// GetCustom returns the Custom field value
func (o *WidgetType) GetCustom() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *WidgetType) GetCustomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Custom, true
}

// SetCustom sets field value
func (o *WidgetType) SetCustom(v bool) {
	o.Custom = v
}

// GetTranslation returns the Translation field value
func (o *WidgetType) GetTranslation() Translation {
	if o == nil {
		var ret Translation
		return ret
	}

	return o.Translation
}

// GetTranslationOk returns a tuple with the Translation field value
// and a boolean to check if the value has been set.
func (o *WidgetType) GetTranslationOk() (*Translation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Translation, true
}

// SetTranslation sets field value
func (o *WidgetType) SetTranslation(v Translation) {
	o.Translation = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *WidgetType) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetType) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *WidgetType) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *WidgetType) SetIcon(v string) {
	o.Icon = &v
}

// GetWithAlarm returns the WithAlarm field value if set, zero value otherwise.
func (o *WidgetType) GetWithAlarm() bool {
	if o == nil || o.WithAlarm == nil {
		var ret bool
		return ret
	}
	return *o.WithAlarm
}

// GetWithAlarmOk returns a tuple with the WithAlarm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetType) GetWithAlarmOk() (*bool, bool) {
	if o == nil || o.WithAlarm == nil {
		return nil, false
	}
	return o.WithAlarm, true
}

// HasWithAlarm returns a boolean if a field has been set.
func (o *WidgetType) HasWithAlarm() bool {
	if o != nil && o.WithAlarm != nil {
		return true
	}

	return false
}

// SetWithAlarm gets a reference to the given bool and assigns it to the WithAlarm field.
func (o *WidgetType) SetWithAlarm(v bool) {
	o.WithAlarm = &v
}

// GetWithTimespan returns the WithTimespan field value if set, zero value otherwise.
func (o *WidgetType) GetWithTimespan() bool {
	if o == nil || o.WithTimespan == nil {
		var ret bool
		return ret
	}
	return *o.WithTimespan
}

// GetWithTimespanOk returns a tuple with the WithTimespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetType) GetWithTimespanOk() (*bool, bool) {
	if o == nil || o.WithTimespan == nil {
		return nil, false
	}
	return o.WithTimespan, true
}

// HasWithTimespan returns a boolean if a field has been set.
func (o *WidgetType) HasWithTimespan() bool {
	if o != nil && o.WithTimespan != nil {
		return true
	}

	return false
}

// SetWithTimespan gets a reference to the given bool and assigns it to the WithTimespan field.
func (o *WidgetType) SetWithTimespan(v bool) {
	o.WithTimespan = &v
}

// GetElements returns the Elements field value
func (o *WidgetType) GetElements() []WidgetTypeElement {
	if o == nil {
		var ret []WidgetTypeElement
		return ret
	}

	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value
// and a boolean to check if the value has been set.
func (o *WidgetType) GetElementsOk() ([]WidgetTypeElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Elements, true
}

// SetElements sets field value
func (o *WidgetType) SetElements(v []WidgetTypeElement) {
	o.Elements = v
}

func (o WidgetType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["custom"] = o.Custom
	}
	if true {
		toSerialize["translation"] = o.Translation
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.WithAlarm != nil {
		toSerialize["withAlarm"] = o.WithAlarm
	}
	if o.WithTimespan != nil {
		toSerialize["withTimespan"] = o.WithTimespan
	}
	if true {
		toSerialize["elements"] = o.Elements
	}
	return json.Marshal(toSerialize)
}

type NullableWidgetType struct {
	value *WidgetType
	isSet bool
}

func (v NullableWidgetType) Get() *WidgetType {
	return v.value
}

func (v *NullableWidgetType) Set(val *WidgetType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetType(val *WidgetType) *NullableWidgetType {
	return &NullableWidgetType{value: val, isSet: true}
}

func (v NullableWidgetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


