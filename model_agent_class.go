/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.5.5
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AgentClass The class of an agent
type AgentClass string

// List of AgentClass
const (
	AGENT_CLASS_IOSYS AgentClass = "iosys"
	AGENT_CLASS_MBUS  AgentClass = "mbus"
)

// All allowed values of AgentClass enum
var AllowedAgentClassEnumValues = []AgentClass{
	"iosys",
	"mbus",
}

func (v *AgentClass) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentClass(value)
	for _, existing := range AllowedAgentClassEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AgentClass", value)
}

// NewAgentClassFromValue returns a pointer to a valid AgentClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentClassFromValue(v string) (*AgentClass, error) {
	ev := AgentClass(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentClass: valid values are %v", v, AllowedAgentClassEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentClass) IsValid() bool {
	for _, existing := range AllowedAgentClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentClass value
func (v AgentClass) Ptr() *AgentClass {
	return &v
}

type NullableAgentClass struct {
	value *AgentClass
	isSet bool
}

func (v NullableAgentClass) Get() *AgentClass {
	return v.value
}

func (v *NullableAgentClass) Set(val *AgentClass) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentClass) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentClass(val *AgentClass) *NullableAgentClass {
	return &NullableAgentClass{value: val, isSet: true}
}

func (v NullableAgentClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
