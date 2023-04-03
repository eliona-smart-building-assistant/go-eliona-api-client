/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.4.12
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AgentDeviceMapping struct for AgentDeviceMapping
type AgentDeviceMapping struct {
	IosysAgentDeviceMapping *IosysAgentDeviceMapping
	MbusAgentDeviceMapping  *MbusAgentDeviceMapping
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AgentDeviceMapping) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into IosysAgentDeviceMapping
	err = json.Unmarshal(data, &dst.IosysAgentDeviceMapping)
	if err == nil {
		jsonIosysAgentDeviceMapping, _ := json.Marshal(dst.IosysAgentDeviceMapping)
		if string(jsonIosysAgentDeviceMapping) == "{}" { // empty struct
			dst.IosysAgentDeviceMapping = nil
		} else {
			return nil // data stored in dst.IosysAgentDeviceMapping, return on the first match
		}
	} else {
		dst.IosysAgentDeviceMapping = nil
	}

	// try to unmarshal JSON data into MbusAgentDeviceMapping
	err = json.Unmarshal(data, &dst.MbusAgentDeviceMapping)
	if err == nil {
		jsonMbusAgentDeviceMapping, _ := json.Marshal(dst.MbusAgentDeviceMapping)
		if string(jsonMbusAgentDeviceMapping) == "{}" { // empty struct
			dst.MbusAgentDeviceMapping = nil
		} else {
			return nil // data stored in dst.MbusAgentDeviceMapping, return on the first match
		}
	} else {
		dst.MbusAgentDeviceMapping = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AgentDeviceMapping)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AgentDeviceMapping) MarshalJSON() ([]byte, error) {
	if src.IosysAgentDeviceMapping != nil {
		return json.Marshal(&src.IosysAgentDeviceMapping)
	}

	if src.MbusAgentDeviceMapping != nil {
		return json.Marshal(&src.MbusAgentDeviceMapping)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAgentDeviceMapping struct {
	value *AgentDeviceMapping
	isSet bool
}

func (v NullableAgentDeviceMapping) Get() *AgentDeviceMapping {
	return v.value
}

func (v *NullableAgentDeviceMapping) Set(val *AgentDeviceMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentDeviceMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentDeviceMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentDeviceMapping(val *AgentDeviceMapping) *NullableAgentDeviceMapping {
	return &NullableAgentDeviceMapping{value: val, isSet: true}
}

func (v NullableAgentDeviceMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentDeviceMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
