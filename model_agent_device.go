/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.12
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AgentDevice struct for AgentDevice
type AgentDevice struct {
	IosysAgentDevice *IosysAgentDevice
	MbusAgentDevice  *MbusAgentDevice
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AgentDevice) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into IosysAgentDevice
	err = json.Unmarshal(data, &dst.IosysAgentDevice)
	if err == nil {
		jsonIosysAgentDevice, _ := json.Marshal(dst.IosysAgentDevice)
		if string(jsonIosysAgentDevice) == "{}" { // empty struct
			dst.IosysAgentDevice = nil
		} else {
			return nil // data stored in dst.IosysAgentDevice, return on the first match
		}
	} else {
		dst.IosysAgentDevice = nil
	}

	// try to unmarshal JSON data into MbusAgentDevice
	err = json.Unmarshal(data, &dst.MbusAgentDevice)
	if err == nil {
		jsonMbusAgentDevice, _ := json.Marshal(dst.MbusAgentDevice)
		if string(jsonMbusAgentDevice) == "{}" { // empty struct
			dst.MbusAgentDevice = nil
		} else {
			return nil // data stored in dst.MbusAgentDevice, return on the first match
		}
	} else {
		dst.MbusAgentDevice = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AgentDevice)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AgentDevice) MarshalJSON() ([]byte, error) {
	if src.IosysAgentDevice != nil {
		return json.Marshal(&src.IosysAgentDevice)
	}

	if src.MbusAgentDevice != nil {
		return json.Marshal(&src.MbusAgentDevice)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAgentDevice struct {
	value *AgentDevice
	isSet bool
}

func (v NullableAgentDevice) Get() *AgentDevice {
	return v.value
}

func (v *NullableAgentDevice) Set(val *AgentDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentDevice(val *AgentDevice) *NullableAgentDevice {
	return &NullableAgentDevice{value: val, isSet: true}
}

func (v NullableAgentDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
