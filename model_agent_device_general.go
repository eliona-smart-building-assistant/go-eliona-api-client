/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.8.1
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AgentDeviceGeneral type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentDeviceGeneral{}

// AgentDeviceGeneral A general device for an agent
type AgentDeviceGeneral struct {
	Class NullableAgentClass `json:"class,omitempty"`
	// Unique id for the device
	Id NullableInt32 `json:"id,omitempty"`
	// The id of the agent the device belongs to
	AgentId NullableInt32 `json:"agentId,omitempty"`
	// Is the device enabled or not
	Enable *bool `json:"enable,omitempty"`
}

// NewAgentDeviceGeneral instantiates a new AgentDeviceGeneral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentDeviceGeneral() *AgentDeviceGeneral {
	this := AgentDeviceGeneral{}
	var enable bool = false
	this.Enable = &enable
	return &this
}

// NewAgentDeviceGeneralWithDefaults instantiates a new AgentDeviceGeneral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentDeviceGeneralWithDefaults() *AgentDeviceGeneral {
	this := AgentDeviceGeneral{}
	var enable bool = false
	this.Enable = &enable
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentDeviceGeneral) GetClass() AgentClass {
	if o == nil || IsNil(o.Class.Get()) {
		var ret AgentClass
		return ret
	}
	return *o.Class.Get()
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentDeviceGeneral) GetClassOk() (*AgentClass, bool) {
	if o == nil {
		return nil, false
	}
	return o.Class.Get(), o.Class.IsSet()
}

// HasClass returns a boolean if a field has been set.
func (o *AgentDeviceGeneral) HasClass() bool {
	if o != nil && o.Class.IsSet() {
		return true
	}

	return false
}

// SetClass gets a reference to the given NullableAgentClass and assigns it to the Class field.
func (o *AgentDeviceGeneral) SetClass(v AgentClass) {
	o.Class.Set(&v)
}

// SetClassNil sets the value for Class to be an explicit nil
func (o *AgentDeviceGeneral) SetClassNil() {
	o.Class.Set(nil)
}

// UnsetClass ensures that no value is present for Class, not even an explicit nil
func (o *AgentDeviceGeneral) UnsetClass() {
	o.Class.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentDeviceGeneral) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentDeviceGeneral) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AgentDeviceGeneral) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *AgentDeviceGeneral) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *AgentDeviceGeneral) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AgentDeviceGeneral) UnsetId() {
	o.Id.Unset()
}

// GetAgentId returns the AgentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentDeviceGeneral) GetAgentId() int32 {
	if o == nil || IsNil(o.AgentId.Get()) {
		var ret int32
		return ret
	}
	return *o.AgentId.Get()
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentDeviceGeneral) GetAgentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentId.Get(), o.AgentId.IsSet()
}

// HasAgentId returns a boolean if a field has been set.
func (o *AgentDeviceGeneral) HasAgentId() bool {
	if o != nil && o.AgentId.IsSet() {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given NullableInt32 and assigns it to the AgentId field.
func (o *AgentDeviceGeneral) SetAgentId(v int32) {
	o.AgentId.Set(&v)
}

// SetAgentIdNil sets the value for AgentId to be an explicit nil
func (o *AgentDeviceGeneral) SetAgentIdNil() {
	o.AgentId.Set(nil)
}

// UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
func (o *AgentDeviceGeneral) UnsetAgentId() {
	o.AgentId.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *AgentDeviceGeneral) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentDeviceGeneral) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *AgentDeviceGeneral) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *AgentDeviceGeneral) SetEnable(v bool) {
	o.Enable = &v
}

func (o AgentDeviceGeneral) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentDeviceGeneral) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Class.IsSet() {
		toSerialize["class"] = o.Class.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.AgentId.IsSet() {
		toSerialize["agentId"] = o.AgentId.Get()
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	return toSerialize, nil
}

type NullableAgentDeviceGeneral struct {
	value *AgentDeviceGeneral
	isSet bool
}

func (v NullableAgentDeviceGeneral) Get() *AgentDeviceGeneral {
	return v.value
}

func (v *NullableAgentDeviceGeneral) Set(val *AgentDeviceGeneral) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentDeviceGeneral) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentDeviceGeneral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentDeviceGeneral(val *AgentDeviceGeneral) *NullableAgentDeviceGeneral {
	return &NullableAgentDeviceGeneral{value: val, isSet: true}
}

func (v NullableAgentDeviceGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentDeviceGeneral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
