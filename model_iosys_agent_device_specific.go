/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.4.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// IosysAgentDeviceSpecific Specific device for IOSYS agents
type IosysAgentDeviceSpecific struct {
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

// NewIosysAgentDeviceSpecific instantiates a new IosysAgentDeviceSpecific object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIosysAgentDeviceSpecific() *IosysAgentDeviceSpecific {
	this := IosysAgentDeviceSpecific{}
	return &this
}

// NewIosysAgentDeviceSpecificWithDefaults instantiates a new IosysAgentDeviceSpecific object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIosysAgentDeviceSpecificWithDefaults() *IosysAgentDeviceSpecific {
	this := IosysAgentDeviceSpecific{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceSpecific) GetPort() int32 {
	if o == nil || isNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceSpecific) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *IosysAgentDeviceSpecific) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *IosysAgentDeviceSpecific) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *IosysAgentDeviceSpecific) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *IosysAgentDeviceSpecific) UnsetPort() {
	o.Port.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceSpecific) GetCertificate() string {
	if o == nil || isNil(o.Certificate.Get()) {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceSpecific) GetCertificateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *IosysAgentDeviceSpecific) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *IosysAgentDeviceSpecific) SetCertificate(v string) {
	o.Certificate.Set(&v)
}
// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *IosysAgentDeviceSpecific) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *IosysAgentDeviceSpecific) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceSpecific) GetKey() string {
	if o == nil || isNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceSpecific) GetKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *IosysAgentDeviceSpecific) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *IosysAgentDeviceSpecific) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *IosysAgentDeviceSpecific) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *IosysAgentDeviceSpecific) UnsetKey() {
	o.Key.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceSpecific) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout.Get()) {
		var ret int32
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceSpecific) GetTimeoutOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *IosysAgentDeviceSpecific) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableInt32 and assigns it to the Timeout field.
func (o *IosysAgentDeviceSpecific) SetTimeout(v int32) {
	o.Timeout.Set(&v)
}
// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *IosysAgentDeviceSpecific) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *IosysAgentDeviceSpecific) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetReconnect returns the Reconnect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceSpecific) GetReconnect() int32 {
	if o == nil || isNil(o.Reconnect.Get()) {
		var ret int32
		return ret
	}
	return *o.Reconnect.Get()
}

// GetReconnectOk returns a tuple with the Reconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceSpecific) GetReconnectOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Reconnect.Get(), o.Reconnect.IsSet()
}

// HasReconnect returns a boolean if a field has been set.
func (o *IosysAgentDeviceSpecific) HasReconnect() bool {
	if o != nil && o.Reconnect.IsSet() {
		return true
	}

	return false
}

// SetReconnect gets a reference to the given NullableInt32 and assigns it to the Reconnect field.
func (o *IosysAgentDeviceSpecific) SetReconnect(v int32) {
	o.Reconnect.Set(&v)
}
// SetReconnectNil sets the value for Reconnect to be an explicit nil
func (o *IosysAgentDeviceSpecific) SetReconnectNil() {
	o.Reconnect.Set(nil)
}

// UnsetReconnect ensures that no value is present for Reconnect, not even an explicit nil
func (o *IosysAgentDeviceSpecific) UnsetReconnect() {
	o.Reconnect.Unset()
}

func (o IosysAgentDeviceSpecific) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullableIosysAgentDeviceSpecific struct {
	value *IosysAgentDeviceSpecific
	isSet bool
}

func (v NullableIosysAgentDeviceSpecific) Get() *IosysAgentDeviceSpecific {
	return v.value
}

func (v *NullableIosysAgentDeviceSpecific) Set(val *IosysAgentDeviceSpecific) {
	v.value = val
	v.isSet = true
}

func (v NullableIosysAgentDeviceSpecific) IsSet() bool {
	return v.isSet
}

func (v *NullableIosysAgentDeviceSpecific) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIosysAgentDeviceSpecific(val *IosysAgentDeviceSpecific) *NullableIosysAgentDeviceSpecific {
	return &NullableIosysAgentDeviceSpecific{value: val, isSet: true}
}

func (v NullableIosysAgentDeviceSpecific) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIosysAgentDeviceSpecific) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


