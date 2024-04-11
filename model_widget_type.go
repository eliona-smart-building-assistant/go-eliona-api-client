/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.10
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the WidgetType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetType{}

// WidgetType A frontend widget
type WidgetType struct {
	// The internal Id of widget type
	Id NullableInt32 `json:"id,omitempty"`
	// The unique name for this widget type
	Name string `json:"name"`
	// Is this a customer created type or not
	Custom      *bool               `json:"custom,omitempty"`
	Translation NullableTranslation `json:"translation"`
	// Icon name corresponding to assets used in this widget
	Icon NullableString `json:"icon,omitempty"`
	// Show alarms in widget
	WithAlarm NullableBool `json:"withAlarm,omitempty"`
	// Show selection for timespan in widget
	WithTimespan NullableBool `json:"withTimespan,omitempty"`
	// A list of elements for this widget (order matches the order of elements for this type)
	Elements []WidgetTypeElement `json:"elements"`
}

type _WidgetType WidgetType

// NewWidgetType instantiates a new WidgetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetType(name string, translation NullableTranslation, elements []WidgetTypeElement) *WidgetType {
	this := WidgetType{}
	this.Name = name
	var custom bool = true
	this.Custom = &custom
	this.Translation = translation
	var withAlarm bool = false
	this.WithAlarm = *NewNullableBool(&withAlarm)
	var withTimespan bool = false
	this.WithTimespan = *NewNullableBool(&withTimespan)
	this.Elements = elements
	return &this
}

// NewWidgetTypeWithDefaults instantiates a new WidgetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetTypeWithDefaults() *WidgetType {
	this := WidgetType{}
	var custom bool = true
	this.Custom = &custom
	var withAlarm bool = false
	this.WithAlarm = *NewNullableBool(&withAlarm)
	var withTimespan bool = false
	this.WithTimespan = *NewNullableBool(&withTimespan)
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetType) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *WidgetType) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *WidgetType) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *WidgetType) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *WidgetType) UnsetId() {
	o.Id.Unset()
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

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *WidgetType) GetCustom() bool {
	if o == nil || IsNil(o.Custom) {
		var ret bool
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetType) GetCustomOk() (*bool, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *WidgetType) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given bool and assigns it to the Custom field.
func (o *WidgetType) SetCustom(v bool) {
	o.Custom = &v
}

// GetTranslation returns the Translation field value
// If the value is explicit nil, the zero value for Translation will be returned
func (o *WidgetType) GetTranslation() Translation {
	if o == nil || o.Translation.Get() == nil {
		var ret Translation
		return ret
	}

	return *o.Translation.Get()
}

// GetTranslationOk returns a tuple with the Translation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetType) GetTranslationOk() (*Translation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Translation.Get(), o.Translation.IsSet()
}

// SetTranslation sets field value
func (o *WidgetType) SetTranslation(v Translation) {
	o.Translation.Set(&v)
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetType) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetType) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *WidgetType) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *WidgetType) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *WidgetType) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *WidgetType) UnsetIcon() {
	o.Icon.Unset()
}

// GetWithAlarm returns the WithAlarm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetType) GetWithAlarm() bool {
	if o == nil || IsNil(o.WithAlarm.Get()) {
		var ret bool
		return ret
	}
	return *o.WithAlarm.Get()
}

// GetWithAlarmOk returns a tuple with the WithAlarm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetType) GetWithAlarmOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithAlarm.Get(), o.WithAlarm.IsSet()
}

// HasWithAlarm returns a boolean if a field has been set.
func (o *WidgetType) HasWithAlarm() bool {
	if o != nil && o.WithAlarm.IsSet() {
		return true
	}

	return false
}

// SetWithAlarm gets a reference to the given NullableBool and assigns it to the WithAlarm field.
func (o *WidgetType) SetWithAlarm(v bool) {
	o.WithAlarm.Set(&v)
}

// SetWithAlarmNil sets the value for WithAlarm to be an explicit nil
func (o *WidgetType) SetWithAlarmNil() {
	o.WithAlarm.Set(nil)
}

// UnsetWithAlarm ensures that no value is present for WithAlarm, not even an explicit nil
func (o *WidgetType) UnsetWithAlarm() {
	o.WithAlarm.Unset()
}

// GetWithTimespan returns the WithTimespan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetType) GetWithTimespan() bool {
	if o == nil || IsNil(o.WithTimespan.Get()) {
		var ret bool
		return ret
	}
	return *o.WithTimespan.Get()
}

// GetWithTimespanOk returns a tuple with the WithTimespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetType) GetWithTimespanOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithTimespan.Get(), o.WithTimespan.IsSet()
}

// HasWithTimespan returns a boolean if a field has been set.
func (o *WidgetType) HasWithTimespan() bool {
	if o != nil && o.WithTimespan.IsSet() {
		return true
	}

	return false
}

// SetWithTimespan gets a reference to the given NullableBool and assigns it to the WithTimespan field.
func (o *WidgetType) SetWithTimespan(v bool) {
	o.WithTimespan.Set(&v)
}

// SetWithTimespanNil sets the value for WithTimespan to be an explicit nil
func (o *WidgetType) SetWithTimespanNil() {
	o.WithTimespan.Set(nil)
}

// UnsetWithTimespan ensures that no value is present for WithTimespan, not even an explicit nil
func (o *WidgetType) UnsetWithTimespan() {
	o.WithTimespan.Unset()
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	toSerialize["translation"] = o.Translation.Get()
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.WithAlarm.IsSet() {
		toSerialize["withAlarm"] = o.WithAlarm.Get()
	}
	if o.WithTimespan.IsSet() {
		toSerialize["withTimespan"] = o.WithTimespan.Get()
	}
	toSerialize["elements"] = o.Elements
	return toSerialize, nil
}

func (o *WidgetType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"translation",
		"elements",
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

	varWidgetType := _WidgetType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWidgetType)

	if err != nil {
		return err
	}

	*o = WidgetType(varWidgetType)

	return err
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
