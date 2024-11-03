/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.7.3
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the IosysAgentDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IosysAgentDevice{}

// IosysAgentDevice struct for IosysAgentDevice
type IosysAgentDevice struct {
	Class NullableAgentClass `json:"class,omitempty"`
	// Unique id for the device
	Id NullableInt32 `json:"id,omitempty"`
	// The id of the agent the device belongs to
	AgentId NullableInt32 `json:"agentId,omitempty"`
	// Is the device enabled or not
	Enable *bool `json:"enable,omitempty"`
	// Port of device
	Port NullableInt32 `json:"port,omitempty"`
	// Certificate of device
	Certificate NullableString `json:"certificate,omitempty"`
	// Key for device
	Key NullableString `json:"key,omitempty"`
	// Time in seconds
	Timeout NullableInt32 `json:"timeout,omitempty"`
	// Reconnect
	Reconnect NullableInt32 `json:"reconnect,omitempty"`
}

// NewIosysAgentDevice instantiates a new IosysAgentDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIosysAgentDevice() *IosysAgentDevice {
	this := IosysAgentDevice{}
	var enable bool = false
	this.Enable = &enable
	return &this
}

// NewIosysAgentDeviceWithDefaults instantiates a new IosysAgentDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIosysAgentDeviceWithDefaults() *IosysAgentDevice {
	this := IosysAgentDevice{}
	var enable bool = false
	this.Enable = &enable
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDevice) GetClass() AgentClass {
	if o == nil || IsNil(o.Class.Get()) {
		var ret AgentClass
		return ret
	}
	return *o.Class.Get()
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDevice) GetClassOk() (*AgentClass, bool) {
	if o == nil {
		return nil, false
	}
	return o.Class.Get(), o.Class.IsSet()
}

// HasClass returns a boolean if a field has been set.
func (o *IosysAgentDevice) HasClass() bool {
	if o != nil && o.Class.IsSet() {
		return true
	}

	return false
}

// SetClass gets a reference to the given NullableAgentClass and assigns it to the Class field.
func (o *IosysAgentDevice) SetClass(v AgentClass) {
	o.Class.Set(&v)
}

// SetClassNil sets the value for Class to be an explicit nil
func (o *IosysAgentDevice) SetClassNil() {
	o.Class.Set(nil)
}

// UnsetClass ensures that no value is present for Class, not even an explicit nil
func (o *IosysAgentDevice) UnsetClass() {
	o.Class.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDevice) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDevice) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *IosysAgentDevice) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *IosysAgentDevice) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *IosysAgentDevice) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *IosysAgentDevice) UnsetId() {
	o.Id.Unset()
}

// GetAgentId returns the AgentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDevice) GetAgentId() int32 {
	if o == nil || IsNil(o.AgentId.Get()) {
		var ret int32
		return ret
	}
	return *o.AgentId.Get()
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDevice) GetAgentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentId.Get(), o.AgentId.IsSet()
}

// HasAgentId returns a boolean if a field has been set.
func (o *IosysAgentDevice) HasAgentId() bool {
	if o != nil && o.AgentId.IsSet() {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given NullableInt32 and assigns it to the AgentId field.
func (o *IosysAgentDevice) SetAgentId(v int32) {
	o.AgentId.Set(&v)
}

// SetAgentIdNil sets the value for AgentId to be an explicit nil
func (o *IosysAgentDevice) SetAgentIdNil() {
	o.AgentId.Set(nil)
}

// UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
func (o *IosysAgentDevice) UnsetAgentId() {
	o.AgentId.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *IosysAgentDevice) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosysAgentDevice) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *IosysAgentDevice) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *IosysAgentDevice) SetEnable(v bool) {
	o.Enable = &v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDevice) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDevice) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *IosysAgentDevice) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *IosysAgentDevice) SetPort(v int32) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil
func (o *IosysAgentDevice) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *IosysAgentDevice) UnsetPort() {
	o.Port.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDevice) GetCertificate() string {
	if o == nil || IsNil(o.Certificate.Get()) {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDevice) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *IosysAgentDevice) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *IosysAgentDevice) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *IosysAgentDevice) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *IosysAgentDevice) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDevice) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDevice) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *IosysAgentDevice) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *IosysAgentDevice) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *IosysAgentDevice) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *IosysAgentDevice) UnsetKey() {
	o.Key.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDevice) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout.Get()) {
		var ret int32
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDevice) GetTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *IosysAgentDevice) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableInt32 and assigns it to the Timeout field.
func (o *IosysAgentDevice) SetTimeout(v int32) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *IosysAgentDevice) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *IosysAgentDevice) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetReconnect returns the Reconnect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDevice) GetReconnect() int32 {
	if o == nil || IsNil(o.Reconnect.Get()) {
		var ret int32
		return ret
	}
	return *o.Reconnect.Get()
}

// GetReconnectOk returns a tuple with the Reconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDevice) GetReconnectOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reconnect.Get(), o.Reconnect.IsSet()
}

// HasReconnect returns a boolean if a field has been set.
func (o *IosysAgentDevice) HasReconnect() bool {
	if o != nil && o.Reconnect.IsSet() {
		return true
	}

	return false
}

// SetReconnect gets a reference to the given NullableInt32 and assigns it to the Reconnect field.
func (o *IosysAgentDevice) SetReconnect(v int32) {
	o.Reconnect.Set(&v)
}

// SetReconnectNil sets the value for Reconnect to be an explicit nil
func (o *IosysAgentDevice) SetReconnectNil() {
	o.Reconnect.Set(nil)
}

// UnsetReconnect ensures that no value is present for Reconnect, not even an explicit nil
func (o *IosysAgentDevice) UnsetReconnect() {
	o.Reconnect.Unset()
}

func (o IosysAgentDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IosysAgentDevice) ToMap() (map[string]interface{}, error) {
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
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	if o.Reconnect.IsSet() {
		toSerialize["reconnect"] = o.Reconnect.Get()
	}
	return toSerialize, nil
}

type NullableIosysAgentDevice struct {
	value *IosysAgentDevice
	isSet bool
}

func (v NullableIosysAgentDevice) Get() *IosysAgentDevice {
	return v.value
}

func (v *NullableIosysAgentDevice) Set(val *IosysAgentDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableIosysAgentDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableIosysAgentDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIosysAgentDevice(val *IosysAgentDevice) *NullableIosysAgentDevice {
	return &NullableIosysAgentDevice{value: val, isSet: true}
}

func (v NullableIosysAgentDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIosysAgentDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
