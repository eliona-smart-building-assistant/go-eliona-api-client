/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.0
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentDeviceMappingGeneral type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentDeviceMappingGeneral{}

// AgentDeviceMappingGeneral A general mapping of attributes for an agent
type AgentDeviceMappingGeneral struct {
	Class NullableAgentClass `json:"class,omitempty"`
	// Unique id for the mapping
	Id NullableInt32 `json:"id,omitempty"`
	// The id of the device the mapping belongs to
	DeviceId NullableInt32 `json:"deviceId,omitempty"`
	// Is the mapping enabled or not
	Enable *bool `json:"enable,omitempty"`
	// ID of the corresponding asset
	AssetId NullableInt32 `json:"assetId,omitempty"`
	Subtype DataSubtype   `json:"subtype"`
	// Name of the attribute to map
	Attribute string `json:"attribute"`
}

type _AgentDeviceMappingGeneral AgentDeviceMappingGeneral

// NewAgentDeviceMappingGeneral instantiates a new AgentDeviceMappingGeneral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentDeviceMappingGeneral(subtype DataSubtype, attribute string) *AgentDeviceMappingGeneral {
	this := AgentDeviceMappingGeneral{}
	var enable bool = true
	this.Enable = &enable
	this.Subtype = subtype
	this.Attribute = attribute
	return &this
}

// NewAgentDeviceMappingGeneralWithDefaults instantiates a new AgentDeviceMappingGeneral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentDeviceMappingGeneralWithDefaults() *AgentDeviceMappingGeneral {
	this := AgentDeviceMappingGeneral{}
	var enable bool = true
	this.Enable = &enable
	var subtype DataSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentDeviceMappingGeneral) GetClass() AgentClass {
	if o == nil || IsNil(o.Class.Get()) {
		var ret AgentClass
		return ret
	}
	return *o.Class.Get()
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentDeviceMappingGeneral) GetClassOk() (*AgentClass, bool) {
	if o == nil {
		return nil, false
	}
	return o.Class.Get(), o.Class.IsSet()
}

// HasClass returns a boolean if a field has been set.
func (o *AgentDeviceMappingGeneral) HasClass() bool {
	if o != nil && o.Class.IsSet() {
		return true
	}

	return false
}

// SetClass gets a reference to the given NullableAgentClass and assigns it to the Class field.
func (o *AgentDeviceMappingGeneral) SetClass(v AgentClass) {
	o.Class.Set(&v)
}

// SetClassNil sets the value for Class to be an explicit nil
func (o *AgentDeviceMappingGeneral) SetClassNil() {
	o.Class.Set(nil)
}

// UnsetClass ensures that no value is present for Class, not even an explicit nil
func (o *AgentDeviceMappingGeneral) UnsetClass() {
	o.Class.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentDeviceMappingGeneral) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentDeviceMappingGeneral) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AgentDeviceMappingGeneral) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *AgentDeviceMappingGeneral) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *AgentDeviceMappingGeneral) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AgentDeviceMappingGeneral) UnsetId() {
	o.Id.Unset()
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentDeviceMappingGeneral) GetDeviceId() int32 {
	if o == nil || IsNil(o.DeviceId.Get()) {
		var ret int32
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentDeviceMappingGeneral) GetDeviceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *AgentDeviceMappingGeneral) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableInt32 and assigns it to the DeviceId field.
func (o *AgentDeviceMappingGeneral) SetDeviceId(v int32) {
	o.DeviceId.Set(&v)
}

// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *AgentDeviceMappingGeneral) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *AgentDeviceMappingGeneral) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *AgentDeviceMappingGeneral) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentDeviceMappingGeneral) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *AgentDeviceMappingGeneral) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *AgentDeviceMappingGeneral) SetEnable(v bool) {
	o.Enable = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentDeviceMappingGeneral) GetAssetId() int32 {
	if o == nil || IsNil(o.AssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.AssetId.Get()
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentDeviceMappingGeneral) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetId.Get(), o.AssetId.IsSet()
}

// HasAssetId returns a boolean if a field has been set.
func (o *AgentDeviceMappingGeneral) HasAssetId() bool {
	if o != nil && o.AssetId.IsSet() {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given NullableInt32 and assigns it to the AssetId field.
func (o *AgentDeviceMappingGeneral) SetAssetId(v int32) {
	o.AssetId.Set(&v)
}

// SetAssetIdNil sets the value for AssetId to be an explicit nil
func (o *AgentDeviceMappingGeneral) SetAssetIdNil() {
	o.AssetId.Set(nil)
}

// UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
func (o *AgentDeviceMappingGeneral) UnsetAssetId() {
	o.AssetId.Unset()
}

// GetSubtype returns the Subtype field value
func (o *AgentDeviceMappingGeneral) GetSubtype() DataSubtype {
	if o == nil {
		var ret DataSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *AgentDeviceMappingGeneral) GetSubtypeOk() (*DataSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *AgentDeviceMappingGeneral) SetSubtype(v DataSubtype) {
	o.Subtype = v
}

// GetAttribute returns the Attribute field value
func (o *AgentDeviceMappingGeneral) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *AgentDeviceMappingGeneral) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *AgentDeviceMappingGeneral) SetAttribute(v string) {
	o.Attribute = v
}

func (o AgentDeviceMappingGeneral) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentDeviceMappingGeneral) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Class.IsSet() {
		toSerialize["class"] = o.Class.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.DeviceId.IsSet() {
		toSerialize["deviceId"] = o.DeviceId.Get()
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if o.AssetId.IsSet() {
		toSerialize["assetId"] = o.AssetId.Get()
	}
	toSerialize["subtype"] = o.Subtype
	toSerialize["attribute"] = o.Attribute
	return toSerialize, nil
}

func (o *AgentDeviceMappingGeneral) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subtype",
		"attribute",
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

	varAgentDeviceMappingGeneral := _AgentDeviceMappingGeneral{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentDeviceMappingGeneral)

	if err != nil {
		return err
	}

	*o = AgentDeviceMappingGeneral(varAgentDeviceMappingGeneral)

	return err
}

type NullableAgentDeviceMappingGeneral struct {
	value *AgentDeviceMappingGeneral
	isSet bool
}

func (v NullableAgentDeviceMappingGeneral) Get() *AgentDeviceMappingGeneral {
	return v.value
}

func (v *NullableAgentDeviceMappingGeneral) Set(val *AgentDeviceMappingGeneral) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentDeviceMappingGeneral) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentDeviceMappingGeneral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentDeviceMappingGeneral(val *AgentDeviceMappingGeneral) *NullableAgentDeviceMappingGeneral {
	return &NullableAgentDeviceMappingGeneral{value: val, isSet: true}
}

func (v NullableAgentDeviceMappingGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentDeviceMappingGeneral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
