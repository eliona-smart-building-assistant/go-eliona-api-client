/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.9.4
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Patch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Patch{}

// Patch A patch for an app
type Patch struct {
	// Name of the app
	AppName string `json:"appName"`
	// Name of the patch
	Name string `json:"name"`
	// Is the app active or inactive
	Active NullableBool `json:"active,omitempty"`
	// Is the app already applied or not
	Applied NullableBool `json:"applied,omitempty"`
}

type _Patch Patch

// NewPatch instantiates a new Patch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatch(appName string, name string) *Patch {
	this := Patch{}
	this.AppName = appName
	this.Name = name
	return &this
}

// NewPatchWithDefaults instantiates a new Patch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchWithDefaults() *Patch {
	this := Patch{}
	return &this
}

// GetAppName returns the AppName field value
func (o *Patch) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *Patch) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *Patch) SetAppName(v string) {
	o.AppName = v
}

// GetName returns the Name field value
func (o *Patch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Patch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Patch) SetName(v string) {
	o.Name = v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Patch) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Patch) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *Patch) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *Patch) SetActive(v bool) {
	o.Active.Set(&v)
}

// SetActiveNil sets the value for Active to be an explicit nil
func (o *Patch) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *Patch) UnsetActive() {
	o.Active.Unset()
}

// GetApplied returns the Applied field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Patch) GetApplied() bool {
	if o == nil || IsNil(o.Applied.Get()) {
		var ret bool
		return ret
	}
	return *o.Applied.Get()
}

// GetAppliedOk returns a tuple with the Applied field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Patch) GetAppliedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Applied.Get(), o.Applied.IsSet()
}

// HasApplied returns a boolean if a field has been set.
func (o *Patch) HasApplied() bool {
	if o != nil && o.Applied.IsSet() {
		return true
	}

	return false
}

// SetApplied gets a reference to the given NullableBool and assigns it to the Applied field.
func (o *Patch) SetApplied(v bool) {
	o.Applied.Set(&v)
}

// SetAppliedNil sets the value for Applied to be an explicit nil
func (o *Patch) SetAppliedNil() {
	o.Applied.Set(nil)
}

// UnsetApplied ensures that no value is present for Applied, not even an explicit nil
func (o *Patch) UnsetApplied() {
	o.Applied.Unset()
}

func (o Patch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Patch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appName"] = o.AppName
	toSerialize["name"] = o.Name
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	if o.Applied.IsSet() {
		toSerialize["applied"] = o.Applied.Get()
	}
	return toSerialize, nil
}

func (o *Patch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appName",
		"name",
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

	varPatch := _Patch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatch)

	if err != nil {
		return err
	}

	*o = Patch(varPatch)

	return err
}

type NullablePatch struct {
	value *Patch
	isSet bool
}

func (v NullablePatch) Get() *Patch {
	return v.value
}

func (v *NullablePatch) Set(val *Patch) {
	v.value = val
	v.isSet = true
}

func (v NullablePatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatch(val *Patch) *NullablePatch {
	return &NullablePatch{value: val, isSet: true}
}

func (v NullablePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
