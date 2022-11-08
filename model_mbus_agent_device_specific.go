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

// MbusAgentDeviceSpecific Specific device for MBUS agents
type MbusAgentDeviceSpecific struct {
	Manufacturer NullableString `json:"manufacturer,omitempty"`
	Model NullableString `json:"model,omitempty"`
	Address NullableInt32 `json:"address,omitempty"`
	SecAddress NullableString `json:"secAddress,omitempty"`
	Raster NullableString `json:"raster,omitempty"`
	MaxFail NullableInt32 `json:"maxFail,omitempty"`
	MaxRetry NullableInt32 `json:"maxRetry,omitempty"`
	SendNke NullableBool `json:"sendNke,omitempty"`
	AppResetSubcode NullableInt32 `json:"appResetSubcode,omitempty"`
	MultiFrames NullableInt32 `json:"multiFrames,omitempty"`
}

// NewMbusAgentDeviceSpecific instantiates a new MbusAgentDeviceSpecific object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbusAgentDeviceSpecific() *MbusAgentDeviceSpecific {
	this := MbusAgentDeviceSpecific{}
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

// NewMbusAgentDeviceSpecificWithDefaults instantiates a new MbusAgentDeviceSpecific object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbusAgentDeviceSpecificWithDefaults() *MbusAgentDeviceSpecific {
	this := MbusAgentDeviceSpecific{}
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

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceSpecific) GetManufacturer() string {
	if o == nil || isNil(o.Manufacturer.Get()) {
		var ret string
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceSpecific) GetManufacturerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *MbusAgentDeviceSpecific) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableString and assigns it to the Manufacturer field.
func (o *MbusAgentDeviceSpecific) SetManufacturer(v string) {
	o.Manufacturer.Set(&v)
}
// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *MbusAgentDeviceSpecific) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *MbusAgentDeviceSpecific) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceSpecific) GetModel() string {
	if o == nil || isNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceSpecific) GetModelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *MbusAgentDeviceSpecific) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *MbusAgentDeviceSpecific) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *MbusAgentDeviceSpecific) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *MbusAgentDeviceSpecific) UnsetModel() {
	o.Model.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceSpecific) GetAddress() int32 {
	if o == nil || isNil(o.Address.Get()) {
		var ret int32
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceSpecific) GetAddressOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *MbusAgentDeviceSpecific) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableInt32 and assigns it to the Address field.
func (o *MbusAgentDeviceSpecific) SetAddress(v int32) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *MbusAgentDeviceSpecific) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *MbusAgentDeviceSpecific) UnsetAddress() {
	o.Address.Unset()
}

// GetSecAddress returns the SecAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceSpecific) GetSecAddress() string {
	if o == nil || isNil(o.SecAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SecAddress.Get()
}

// GetSecAddressOk returns a tuple with the SecAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceSpecific) GetSecAddressOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SecAddress.Get(), o.SecAddress.IsSet()
}

// HasSecAddress returns a boolean if a field has been set.
func (o *MbusAgentDeviceSpecific) HasSecAddress() bool {
	if o != nil && o.SecAddress.IsSet() {
		return true
	}

	return false
}

// SetSecAddress gets a reference to the given NullableString and assigns it to the SecAddress field.
func (o *MbusAgentDeviceSpecific) SetSecAddress(v string) {
	o.SecAddress.Set(&v)
}
// SetSecAddressNil sets the value for SecAddress to be an explicit nil
func (o *MbusAgentDeviceSpecific) SetSecAddressNil() {
	o.SecAddress.Set(nil)
}

// UnsetSecAddress ensures that no value is present for SecAddress, not even an explicit nil
func (o *MbusAgentDeviceSpecific) UnsetSecAddress() {
	o.SecAddress.Unset()
}

// GetRaster returns the Raster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceSpecific) GetRaster() string {
	if o == nil || isNil(o.Raster.Get()) {
		var ret string
		return ret
	}
	return *o.Raster.Get()
}

// GetRasterOk returns a tuple with the Raster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceSpecific) GetRasterOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Raster.Get(), o.Raster.IsSet()
}

// HasRaster returns a boolean if a field has been set.
func (o *MbusAgentDeviceSpecific) HasRaster() bool {
	if o != nil && o.Raster.IsSet() {
		return true
	}

	return false
}

// SetRaster gets a reference to the given NullableString and assigns it to the Raster field.
func (o *MbusAgentDeviceSpecific) SetRaster(v string) {
	o.Raster.Set(&v)
}
// SetRasterNil sets the value for Raster to be an explicit nil
func (o *MbusAgentDeviceSpecific) SetRasterNil() {
	o.Raster.Set(nil)
}

// UnsetRaster ensures that no value is present for Raster, not even an explicit nil
func (o *MbusAgentDeviceSpecific) UnsetRaster() {
	o.Raster.Unset()
}

// GetMaxFail returns the MaxFail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceSpecific) GetMaxFail() int32 {
	if o == nil || isNil(o.MaxFail.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxFail.Get()
}

// GetMaxFailOk returns a tuple with the MaxFail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceSpecific) GetMaxFailOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxFail.Get(), o.MaxFail.IsSet()
}

// HasMaxFail returns a boolean if a field has been set.
func (o *MbusAgentDeviceSpecific) HasMaxFail() bool {
	if o != nil && o.MaxFail.IsSet() {
		return true
	}

	return false
}

// SetMaxFail gets a reference to the given NullableInt32 and assigns it to the MaxFail field.
func (o *MbusAgentDeviceSpecific) SetMaxFail(v int32) {
	o.MaxFail.Set(&v)
}
// SetMaxFailNil sets the value for MaxFail to be an explicit nil
func (o *MbusAgentDeviceSpecific) SetMaxFailNil() {
	o.MaxFail.Set(nil)
}

// UnsetMaxFail ensures that no value is present for MaxFail, not even an explicit nil
func (o *MbusAgentDeviceSpecific) UnsetMaxFail() {
	o.MaxFail.Unset()
}

// GetMaxRetry returns the MaxRetry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceSpecific) GetMaxRetry() int32 {
	if o == nil || isNil(o.MaxRetry.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxRetry.Get()
}

// GetMaxRetryOk returns a tuple with the MaxRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceSpecific) GetMaxRetryOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxRetry.Get(), o.MaxRetry.IsSet()
}

// HasMaxRetry returns a boolean if a field has been set.
func (o *MbusAgentDeviceSpecific) HasMaxRetry() bool {
	if o != nil && o.MaxRetry.IsSet() {
		return true
	}

	return false
}

// SetMaxRetry gets a reference to the given NullableInt32 and assigns it to the MaxRetry field.
func (o *MbusAgentDeviceSpecific) SetMaxRetry(v int32) {
	o.MaxRetry.Set(&v)
}
// SetMaxRetryNil sets the value for MaxRetry to be an explicit nil
func (o *MbusAgentDeviceSpecific) SetMaxRetryNil() {
	o.MaxRetry.Set(nil)
}

// UnsetMaxRetry ensures that no value is present for MaxRetry, not even an explicit nil
func (o *MbusAgentDeviceSpecific) UnsetMaxRetry() {
	o.MaxRetry.Unset()
}

// GetSendNke returns the SendNke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceSpecific) GetSendNke() bool {
	if o == nil || isNil(o.SendNke.Get()) {
		var ret bool
		return ret
	}
	return *o.SendNke.Get()
}

// GetSendNkeOk returns a tuple with the SendNke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceSpecific) GetSendNkeOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.SendNke.Get(), o.SendNke.IsSet()
}

// HasSendNke returns a boolean if a field has been set.
func (o *MbusAgentDeviceSpecific) HasSendNke() bool {
	if o != nil && o.SendNke.IsSet() {
		return true
	}

	return false
}

// SetSendNke gets a reference to the given NullableBool and assigns it to the SendNke field.
func (o *MbusAgentDeviceSpecific) SetSendNke(v bool) {
	o.SendNke.Set(&v)
}
// SetSendNkeNil sets the value for SendNke to be an explicit nil
func (o *MbusAgentDeviceSpecific) SetSendNkeNil() {
	o.SendNke.Set(nil)
}

// UnsetSendNke ensures that no value is present for SendNke, not even an explicit nil
func (o *MbusAgentDeviceSpecific) UnsetSendNke() {
	o.SendNke.Unset()
}

// GetAppResetSubcode returns the AppResetSubcode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceSpecific) GetAppResetSubcode() int32 {
	if o == nil || isNil(o.AppResetSubcode.Get()) {
		var ret int32
		return ret
	}
	return *o.AppResetSubcode.Get()
}

// GetAppResetSubcodeOk returns a tuple with the AppResetSubcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceSpecific) GetAppResetSubcodeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppResetSubcode.Get(), o.AppResetSubcode.IsSet()
}

// HasAppResetSubcode returns a boolean if a field has been set.
func (o *MbusAgentDeviceSpecific) HasAppResetSubcode() bool {
	if o != nil && o.AppResetSubcode.IsSet() {
		return true
	}

	return false
}

// SetAppResetSubcode gets a reference to the given NullableInt32 and assigns it to the AppResetSubcode field.
func (o *MbusAgentDeviceSpecific) SetAppResetSubcode(v int32) {
	o.AppResetSubcode.Set(&v)
}
// SetAppResetSubcodeNil sets the value for AppResetSubcode to be an explicit nil
func (o *MbusAgentDeviceSpecific) SetAppResetSubcodeNil() {
	o.AppResetSubcode.Set(nil)
}

// UnsetAppResetSubcode ensures that no value is present for AppResetSubcode, not even an explicit nil
func (o *MbusAgentDeviceSpecific) UnsetAppResetSubcode() {
	o.AppResetSubcode.Unset()
}

// GetMultiFrames returns the MultiFrames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MbusAgentDeviceSpecific) GetMultiFrames() int32 {
	if o == nil || isNil(o.MultiFrames.Get()) {
		var ret int32
		return ret
	}
	return *o.MultiFrames.Get()
}

// GetMultiFramesOk returns a tuple with the MultiFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MbusAgentDeviceSpecific) GetMultiFramesOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MultiFrames.Get(), o.MultiFrames.IsSet()
}

// HasMultiFrames returns a boolean if a field has been set.
func (o *MbusAgentDeviceSpecific) HasMultiFrames() bool {
	if o != nil && o.MultiFrames.IsSet() {
		return true
	}

	return false
}

// SetMultiFrames gets a reference to the given NullableInt32 and assigns it to the MultiFrames field.
func (o *MbusAgentDeviceSpecific) SetMultiFrames(v int32) {
	o.MultiFrames.Set(&v)
}
// SetMultiFramesNil sets the value for MultiFrames to be an explicit nil
func (o *MbusAgentDeviceSpecific) SetMultiFramesNil() {
	o.MultiFrames.Set(nil)
}

// UnsetMultiFrames ensures that no value is present for MultiFrames, not even an explicit nil
func (o *MbusAgentDeviceSpecific) UnsetMultiFrames() {
	o.MultiFrames.Unset()
}

func (o MbusAgentDeviceSpecific) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullableMbusAgentDeviceSpecific struct {
	value *MbusAgentDeviceSpecific
	isSet bool
}

func (v NullableMbusAgentDeviceSpecific) Get() *MbusAgentDeviceSpecific {
	return v.value
}

func (v *NullableMbusAgentDeviceSpecific) Set(val *MbusAgentDeviceSpecific) {
	v.value = val
	v.isSet = true
}

func (v NullableMbusAgentDeviceSpecific) IsSet() bool {
	return v.isSet
}

func (v *NullableMbusAgentDeviceSpecific) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbusAgentDeviceSpecific(val *MbusAgentDeviceSpecific) *NullableMbusAgentDeviceSpecific {
	return &NullableMbusAgentDeviceSpecific{value: val, isSet: true}
}

func (v NullableMbusAgentDeviceSpecific) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbusAgentDeviceSpecific) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


