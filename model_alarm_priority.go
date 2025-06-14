/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.9.4
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AlarmPriority The priority of the alarm. The lower this value the higher the priority.
type AlarmPriority int32

// List of AlarmPriority
const (
	ALARM_PRIORITY_HEIGHT AlarmPriority = 1
	ALARM_PRIORITY_MEDIUM AlarmPriority = 2
	ALARM_PRIORITY_LOW    AlarmPriority = 3
	ALARM_PRIORITY_INFO   AlarmPriority = 10
)

// All allowed values of AlarmPriority enum
var AllowedAlarmPriorityEnumValues = []AlarmPriority{
	1,
	2,
	3,
	10,
}

func (v *AlarmPriority) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlarmPriority(value)
	for _, existing := range AllowedAlarmPriorityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlarmPriority", value)
}

// NewAlarmPriorityFromValue returns a pointer to a valid AlarmPriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlarmPriorityFromValue(v int32) (*AlarmPriority, error) {
	ev := AlarmPriority(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlarmPriority: valid values are %v", v, AllowedAlarmPriorityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlarmPriority) IsValid() bool {
	for _, existing := range AllowedAlarmPriorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlarmPriority value
func (v AlarmPriority) Ptr() *AlarmPriority {
	return &v
}

type NullableAlarmPriority struct {
	value *AlarmPriority
	isSet bool
}

func (v NullableAlarmPriority) Get() *AlarmPriority {
	return v.value
}

func (v *NullableAlarmPriority) Set(val *AlarmPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmPriority(val *AlarmPriority) *NullableAlarmPriority {
	return &NullableAlarmPriority{value: val, isSet: true}
}

func (v NullableAlarmPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
