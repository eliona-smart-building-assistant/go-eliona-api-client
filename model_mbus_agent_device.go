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
)

// checks if the MbusAgentDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbusAgentDevice{}

// MbusAgentDevice struct for MbusAgentDevice
type MbusAgentDevice struct {
	Class NullableAgentClass `json:"class,omitempty"`
	// Unique id for the device
	Id NullableInt32 `json:"id,omitempty"`
	// The id of the agent the device belongs to
	AgentId NullableInt32 `json:"agentId,omitempty"`
	// Is the device enabled or not
	Enable          *bool          `json:"enable,omitempty"`
	Manufacturer    NullableString `json:"manufacturer,omitempty"`
	Model           NullableString `json:"model,omitempty"`
	Address         NullableInt32  `json:"address,omitempty"`
	SecAddress      NullableString `json:"secAddress,omitempty"`
	Raster          NullableString `json:"raster,omitempty"`
	MaxFail         NullableInt32  `json:"maxFail,omitempty"`
	MaxRetry        NullableInt32  `json:"maxRetry,omitempty"`
	SendNke         NullableBool   `json:"sendNke,omitempty"`
	AppResetSubcode NullableInt32  `json:"appResetSubcode,omitempty"`
	MultiFrames     NullableInt32  `json:"multiFrames,omitempty"`
}

// NewMbusAgentDevice instantiates a new MbusAgentDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbusAgentDevice() *MbusAgentDevice {
	this := MbusAgentDevice{}
	var enable bool = false
	this.Enable = &enable
	var maxFail int32 = 4
	this.MaxFail = *NewNullableInt32(&maxFail)
	var maxRetry int32 = 3
	this.MaxRetry = *NewNullableInt32(&maxRetry)
	var sendNke bool = false
	this.SendNke = *NewNullableBool(&sendNke)
	var multiFrames int32 = 0
	this.MultiFrames = *NewNullableInt32(&multiFrames)
	return &this
}

// NewMbusAgentDeviceWithDefaults instantiates a new MbusAgentDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbusAgentDeviceWithDefaults() *MbusAgentDevice {
	this := MbusAgentDevice{}
	var enable bool = false
	this.Enable = &enable
	var maxFail int32 = 4
	this.MaxFail = *NewNullableInt32(&maxFail)
	var maxRetry int32 = 3
	this.MaxRetry = *NewNullableInt32(&maxRetry)
	var sendNke bool = false
	this.SendNke = *NewNullableBool(&sendNke)
	var multiFrames int32 = 0
	this.MultiFrames = *NewNullableInt32(&multiFrames)
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetClass() AgentClass {
	if o == nil || isNil(o.Class.Get()) {
		var ret AgentClass
		return ret
	}
	return *o.Class.Get()
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetClassOk() (*AgentClass, bool) {
	if o == nil {
		return nil, false
	}
	return o.Class.Get(), o.Class.IsSet()
}

// HasClass returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasClass() bool {
	if o != nil && o.Class.IsSet() {
		return true
	}

	return false
}

// SetClass gets a reference to the given NullableAgentClass and assigns it to the Class field.
func (o *MbusAgentDevice) SetClass(v AgentClass) {
	o.Class.Set(&v)
}

// SetClassNil sets the value for Class to be an explicit nil
func (o *MbusAgentDevice) SetClassNil() {
	o.Class.Set(nil)
}

// UnsetClass ensures that no value is present for Class, not even an explicit nil
func (o *MbusAgentDevice) UnsetClass() {
	o.Class.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetId() int32 {
	if o == nil || isNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *MbusAgentDevice) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *MbusAgentDevice) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MbusAgentDevice) UnsetId() {
	o.Id.Unset()
}

// GetAgentId returns the AgentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetAgentId() int32 {
	if o == nil || isNil(o.AgentId.Get()) {
		var ret int32
		return ret
	}
	return *o.AgentId.Get()
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetAgentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgentId.Get(), o.AgentId.IsSet()
}

// HasAgentId returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasAgentId() bool {
	if o != nil && o.AgentId.IsSet() {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given NullableInt32 and assigns it to the AgentId field.
func (o *MbusAgentDevice) SetAgentId(v int32) {
	o.AgentId.Set(&v)
}

// SetAgentIdNil sets the value for AgentId to be an explicit nil
func (o *MbusAgentDevice) SetAgentIdNil() {
	o.AgentId.Set(nil)
}

// UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
func (o *MbusAgentDevice) UnsetAgentId() {
	o.AgentId.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *MbusAgentDevice) GetEnable() bool {
	if o == nil || isNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbusAgentDevice) GetEnableOk() (*bool, bool) {
	if o == nil || isNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasEnable() bool {
	if o != nil && !isNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *MbusAgentDevice) SetEnable(v bool) {
	o.Enable = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetManufacturer() string {
	if o == nil || isNil(o.Manufacturer.Get()) {
		var ret string
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetManufacturerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableString and assigns it to the Manufacturer field.
func (o *MbusAgentDevice) SetManufacturer(v string) {
	o.Manufacturer.Set(&v)
}

// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *MbusAgentDevice) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *MbusAgentDevice) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetModel() string {
	if o == nil || isNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *MbusAgentDevice) SetModel(v string) {
	o.Model.Set(&v)
}

// SetModelNil sets the value for Model to be an explicit nil
func (o *MbusAgentDevice) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *MbusAgentDevice) UnsetModel() {
	o.Model.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetAddress() int32 {
	if o == nil || isNil(o.Address.Get()) {
		var ret int32
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetAddressOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableInt32 and assigns it to the Address field.
func (o *MbusAgentDevice) SetAddress(v int32) {
	o.Address.Set(&v)
}

// SetAddressNil sets the value for Address to be an explicit nil
func (o *MbusAgentDevice) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *MbusAgentDevice) UnsetAddress() {
	o.Address.Unset()
}

// GetSecAddress returns the SecAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetSecAddress() string {
	if o == nil || isNil(o.SecAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SecAddress.Get()
}

// GetSecAddressOk returns a tuple with the SecAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetSecAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecAddress.Get(), o.SecAddress.IsSet()
}

// HasSecAddress returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasSecAddress() bool {
	if o != nil && o.SecAddress.IsSet() {
		return true
	}

	return false
}

// SetSecAddress gets a reference to the given NullableString and assigns it to the SecAddress field.
func (o *MbusAgentDevice) SetSecAddress(v string) {
	o.SecAddress.Set(&v)
}

// SetSecAddressNil sets the value for SecAddress to be an explicit nil
func (o *MbusAgentDevice) SetSecAddressNil() {
	o.SecAddress.Set(nil)
}

// UnsetSecAddress ensures that no value is present for SecAddress, not even an explicit nil
func (o *MbusAgentDevice) UnsetSecAddress() {
	o.SecAddress.Unset()
}

// GetRaster returns the Raster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetRaster() string {
	if o == nil || isNil(o.Raster.Get()) {
		var ret string
		return ret
	}
	return *o.Raster.Get()
}

// GetRasterOk returns a tuple with the Raster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetRasterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Raster.Get(), o.Raster.IsSet()
}

// HasRaster returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasRaster() bool {
	if o != nil && o.Raster.IsSet() {
		return true
	}

	return false
}

// SetRaster gets a reference to the given NullableString and assigns it to the Raster field.
func (o *MbusAgentDevice) SetRaster(v string) {
	o.Raster.Set(&v)
}

// SetRasterNil sets the value for Raster to be an explicit nil
func (o *MbusAgentDevice) SetRasterNil() {
	o.Raster.Set(nil)
}

// UnsetRaster ensures that no value is present for Raster, not even an explicit nil
func (o *MbusAgentDevice) UnsetRaster() {
	o.Raster.Unset()
}

// GetMaxFail returns the MaxFail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetMaxFail() int32 {
	if o == nil || isNil(o.MaxFail.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxFail.Get()
}

// GetMaxFailOk returns a tuple with the MaxFail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetMaxFailOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxFail.Get(), o.MaxFail.IsSet()
}

// HasMaxFail returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasMaxFail() bool {
	if o != nil && o.MaxFail.IsSet() {
		return true
	}

	return false
}

// SetMaxFail gets a reference to the given NullableInt32 and assigns it to the MaxFail field.
func (o *MbusAgentDevice) SetMaxFail(v int32) {
	o.MaxFail.Set(&v)
}

// SetMaxFailNil sets the value for MaxFail to be an explicit nil
func (o *MbusAgentDevice) SetMaxFailNil() {
	o.MaxFail.Set(nil)
}

// UnsetMaxFail ensures that no value is present for MaxFail, not even an explicit nil
func (o *MbusAgentDevice) UnsetMaxFail() {
	o.MaxFail.Unset()
}

// GetMaxRetry returns the MaxRetry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetMaxRetry() int32 {
	if o == nil || isNil(o.MaxRetry.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxRetry.Get()
}

// GetMaxRetryOk returns a tuple with the MaxRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetMaxRetryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRetry.Get(), o.MaxRetry.IsSet()
}

// HasMaxRetry returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasMaxRetry() bool {
	if o != nil && o.MaxRetry.IsSet() {
		return true
	}

	return false
}

// SetMaxRetry gets a reference to the given NullableInt32 and assigns it to the MaxRetry field.
func (o *MbusAgentDevice) SetMaxRetry(v int32) {
	o.MaxRetry.Set(&v)
}

// SetMaxRetryNil sets the value for MaxRetry to be an explicit nil
func (o *MbusAgentDevice) SetMaxRetryNil() {
	o.MaxRetry.Set(nil)
}

// UnsetMaxRetry ensures that no value is present for MaxRetry, not even an explicit nil
func (o *MbusAgentDevice) UnsetMaxRetry() {
	o.MaxRetry.Unset()
}

// GetSendNke returns the SendNke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetSendNke() bool {
	if o == nil || isNil(o.SendNke.Get()) {
		var ret bool
		return ret
	}
	return *o.SendNke.Get()
}

// GetSendNkeOk returns a tuple with the SendNke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetSendNkeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendNke.Get(), o.SendNke.IsSet()
}

// HasSendNke returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasSendNke() bool {
	if o != nil && o.SendNke.IsSet() {
		return true
	}

	return false
}

// SetSendNke gets a reference to the given NullableBool and assigns it to the SendNke field.
func (o *MbusAgentDevice) SetSendNke(v bool) {
	o.SendNke.Set(&v)
}

// SetSendNkeNil sets the value for SendNke to be an explicit nil
func (o *MbusAgentDevice) SetSendNkeNil() {
	o.SendNke.Set(nil)
}

// UnsetSendNke ensures that no value is present for SendNke, not even an explicit nil
func (o *MbusAgentDevice) UnsetSendNke() {
	o.SendNke.Unset()
}

// GetAppResetSubcode returns the AppResetSubcode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetAppResetSubcode() int32 {
	if o == nil || isNil(o.AppResetSubcode.Get()) {
		var ret int32
		return ret
	}
	return *o.AppResetSubcode.Get()
}

// GetAppResetSubcodeOk returns a tuple with the AppResetSubcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetAppResetSubcodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppResetSubcode.Get(), o.AppResetSubcode.IsSet()
}

// HasAppResetSubcode returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasAppResetSubcode() bool {
	if o != nil && o.AppResetSubcode.IsSet() {
		return true
	}

	return false
}

// SetAppResetSubcode gets a reference to the given NullableInt32 and assigns it to the AppResetSubcode field.
func (o *MbusAgentDevice) SetAppResetSubcode(v int32) {
	o.AppResetSubcode.Set(&v)
}

// SetAppResetSubcodeNil sets the value for AppResetSubcode to be an explicit nil
func (o *MbusAgentDevice) SetAppResetSubcodeNil() {
	o.AppResetSubcode.Set(nil)
}

// UnsetAppResetSubcode ensures that no value is present for AppResetSubcode, not even an explicit nil
func (o *MbusAgentDevice) UnsetAppResetSubcode() {
	o.AppResetSubcode.Unset()
}

// GetMultiFrames returns the MultiFrames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDevice) GetMultiFrames() int32 {
	if o == nil || isNil(o.MultiFrames.Get()) {
		var ret int32
		return ret
	}
	return *o.MultiFrames.Get()
}

// GetMultiFramesOk returns a tuple with the MultiFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDevice) GetMultiFramesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MultiFrames.Get(), o.MultiFrames.IsSet()
}

// HasMultiFrames returns a boolean if a field has been set.
func (o *MbusAgentDevice) HasMultiFrames() bool {
	if o != nil && o.MultiFrames.IsSet() {
		return true
	}

	return false
}

// SetMultiFrames gets a reference to the given NullableInt32 and assigns it to the MultiFrames field.
func (o *MbusAgentDevice) SetMultiFrames(v int32) {
	o.MultiFrames.Set(&v)
}

// SetMultiFramesNil sets the value for MultiFrames to be an explicit nil
func (o *MbusAgentDevice) SetMultiFramesNil() {
	o.MultiFrames.Set(nil)
}

// UnsetMultiFrames ensures that no value is present for MultiFrames, not even an explicit nil
func (o *MbusAgentDevice) UnsetMultiFrames() {
	o.MultiFrames.Unset()
}

func (o MbusAgentDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbusAgentDevice) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.SecAddress.IsSet() {
		toSerialize["secAddress"] = o.SecAddress.Get()
	}
	if o.Raster.IsSet() {
		toSerialize["raster"] = o.Raster.Get()
	}
	if o.MaxFail.IsSet() {
		toSerialize["maxFail"] = o.MaxFail.Get()
	}
	if o.MaxRetry.IsSet() {
		toSerialize["maxRetry"] = o.MaxRetry.Get()
	}
	if o.SendNke.IsSet() {
		toSerialize["sendNke"] = o.SendNke.Get()
	}
	if o.AppResetSubcode.IsSet() {
		toSerialize["appResetSubcode"] = o.AppResetSubcode.Get()
	}
	if o.MultiFrames.IsSet() {
		toSerialize["multiFrames"] = o.MultiFrames.Get()
	}
	return toSerialize, nil
}

type NullableMbusAgentDevice struct {
	value *MbusAgentDevice
	isSet bool
}

func (v NullableMbusAgentDevice) Get() *MbusAgentDevice {
	return v.value
}

func (v *NullableMbusAgentDevice) Set(val *MbusAgentDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableMbusAgentDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableMbusAgentDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbusAgentDevice(val *MbusAgentDevice) *NullableMbusAgentDevice {
	return &NullableMbusAgentDevice{value: val, isSet: true}
}

func (v NullableMbusAgentDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbusAgentDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
