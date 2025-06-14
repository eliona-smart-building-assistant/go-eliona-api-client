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

// checks if the App type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &App{}

// App An app
type App struct {
	// Name of the app
	Name string `json:"name"`
	// Is the app active or inactive
	Active NullableBool `json:"active,omitempty"`
	// Is the app already registered or not
	Registered NullableBool `json:"registered,omitempty"`
	// Delivers the apps metadata to handle it in the app store
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// the apps version
	Version NullableString `json:"version,omitempty"`
}

type _App App

// NewApp instantiates a new App object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApp(name string) *App {
	this := App{}
	this.Name = name
	return &this
}

// NewAppWithDefaults instantiates a new App object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppWithDefaults() *App {
	this := App{}
	return &this
}

// GetName returns the Name field value
func (o *App) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *App) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *App) SetName(v string) {
	o.Name = v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *App) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *App) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *App) SetActive(v bool) {
	o.Active.Set(&v)
}

// SetActiveNil sets the value for Active to be an explicit nil
func (o *App) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *App) UnsetActive() {
	o.Active.Unset()
}

// GetRegistered returns the Registered field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *App) GetRegistered() bool {
	if o == nil || IsNil(o.Registered.Get()) {
		var ret bool
		return ret
	}
	return *o.Registered.Get()
}

// GetRegisteredOk returns a tuple with the Registered field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetRegisteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Registered.Get(), o.Registered.IsSet()
}

// HasRegistered returns a boolean if a field has been set.
func (o *App) HasRegistered() bool {
	if o != nil && o.Registered.IsSet() {
		return true
	}

	return false
}

// SetRegistered gets a reference to the given NullableBool and assigns it to the Registered field.
func (o *App) SetRegistered(v bool) {
	o.Registered.Set(&v)
}

// SetRegisteredNil sets the value for Registered to be an explicit nil
func (o *App) SetRegisteredNil() {
	o.Registered.Set(nil)
}

// UnsetRegistered ensures that no value is present for Registered, not even an explicit nil
func (o *App) UnsetRegistered() {
	o.Registered.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *App) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *App) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *App) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *App) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *App) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *App) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *App) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *App) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *App) UnsetVersion() {
	o.Version.Unset()
}

func (o App) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o App) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	if o.Registered.IsSet() {
		toSerialize["registered"] = o.Registered.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	return toSerialize, nil
}

func (o *App) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varApp := _App{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApp)

	if err != nil {
		return err
	}

	*o = App(varApp)

	return err
}

type NullableApp struct {
	value *App
	isSet bool
}

func (v NullableApp) Get() *App {
	return v.value
}

func (v *NullableApp) Set(val *App) {
	v.value = val
	v.isSet = true
}

func (v NullableApp) IsSet() bool {
	return v.isSet
}

func (v *NullableApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApp(val *App) *NullableApp {
	return &NullableApp{value: val, isSet: true}
}

func (v NullableApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
