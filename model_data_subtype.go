/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// DataSubtype Type of asset data
type DataSubtype string

// List of DataSubtype
const (
	SUBTYPE_INPUT DataSubtype = "input"
	SUBTYPE_INFO DataSubtype = "info"
	SUBTYPE_STATUS DataSubtype = "status"
	SUBTYPE_OUTPUT DataSubtype = "output"
)

// All allowed values of DataSubtype enum
var AllowedDataSubtypeEnumValues = []DataSubtype{
	"input",
	"info",
	"status",
	"output",
}

func (v *DataSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataSubtype(value)
	for _, existing := range AllowedDataSubtypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataSubtype", value)
}

// NewDataSubtypeFromValue returns a pointer to a valid DataSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSubtypeFromValue(v string) (*DataSubtype, error) {
	ev := DataSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataSubtype: valid values are %v", v, AllowedDataSubtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSubtype) IsValid() bool {
	for _, existing := range AllowedDataSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSubtype value
func (v DataSubtype) Ptr() *DataSubtype {
	return &v
}

type NullableDataSubtype struct {
	value *DataSubtype
	isSet bool
}

func (v NullableDataSubtype) Get() *DataSubtype {
	return v.value
}

func (v *NullableDataSubtype) Set(val *DataSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSubtype(val *DataSubtype) *NullableDataSubtype {
	return &NullableDataSubtype{value: val, isSet: true}
}

func (v NullableDataSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
