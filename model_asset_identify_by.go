/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.6
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AssetIdentifyBy the model 'AssetIdentifyBy'
type AssetIdentifyBy string

// List of AssetIdentifyBy
const (
	ASSET_IDENTIFY_BY_RESOURCE_ID AssetIdentifyBy = "resourceId"
	ASSET_IDENTIFY_BY_DEVICE_ID   AssetIdentifyBy = "deviceId"
	ASSET_IDENTIFY_BY_ID          AssetIdentifyBy = "id"
	ASSET_IDENTIFY_BY_GAI_PROJ_ID AssetIdentifyBy = "gai-projId"
)

// All allowed values of AssetIdentifyBy enum
var AllowedAssetIdentifyByEnumValues = []AssetIdentifyBy{
	"resourceId",
	"deviceId",
	"id",
	"gai-projId",
}

func (v *AssetIdentifyBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssetIdentifyBy(value)
	for _, existing := range AllowedAssetIdentifyByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssetIdentifyBy", value)
}

// NewAssetIdentifyByFromValue returns a pointer to a valid AssetIdentifyBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetIdentifyByFromValue(v string) (*AssetIdentifyBy, error) {
	ev := AssetIdentifyBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssetIdentifyBy: valid values are %v", v, AllowedAssetIdentifyByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetIdentifyBy) IsValid() bool {
	for _, existing := range AllowedAssetIdentifyByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssetIdentifyBy value
func (v AssetIdentifyBy) Ptr() *AssetIdentifyBy {
	return &v
}

type NullableAssetIdentifyBy struct {
	value *AssetIdentifyBy
	isSet bool
}

func (v NullableAssetIdentifyBy) Get() *AssetIdentifyBy {
	return v.value
}

func (v *NullableAssetIdentifyBy) Set(val *AssetIdentifyBy) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetIdentifyBy) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetIdentifyBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetIdentifyBy(val *AssetIdentifyBy) *NullableAssetIdentifyBy {
	return &NullableAssetIdentifyBy{value: val, isSet: true}
}

func (v NullableAssetIdentifyBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetIdentifyBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
