/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.5.9
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the IosysAgentDeviceMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IosysAgentDeviceMapping{}

// IosysAgentDeviceMapping struct for IosysAgentDeviceMapping
type IosysAgentDeviceMapping struct {
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
	Attribute      string          `json:"attribute"`
	IosVar         NullableString  `json:"iosVar,omitempty"`
	IosType        NullableString  `json:"iosType,omitempty"`
	Down           NullableBool    `json:"down,omitempty"`
	Scale          NullableFloat64 `json:"scale,omitempty"`
	Zero           NullableFloat64 `json:"zero,omitempty"`
	Mask           []int64         `json:"mask,omitempty"`
	MaskAttributes []string        `json:"maskAttributes,omitempty"`
	DeadTime       NullableInt32   `json:"deadTime,omitempty"`
	DeadBand       NullableFloat64 `json:"deadBand,omitempty"`
	Filter         NullableString  `json:"filter,omitempty"`
	Tau            NullableFloat64 `json:"tau,omitempty"`
}

type _IosysAgentDeviceMapping IosysAgentDeviceMapping

// NewIosysAgentDeviceMapping instantiates a new IosysAgentDeviceMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIosysAgentDeviceMapping(subtype DataSubtype, attribute string) *IosysAgentDeviceMapping {
	this := IosysAgentDeviceMapping{}
	var enable bool = true
	this.Enable = &enable
	this.Subtype = subtype
	this.Attribute = attribute
	var down bool = false
	this.Down = *NewNullableBool(&down)
	return &this
}

// NewIosysAgentDeviceMappingWithDefaults instantiates a new IosysAgentDeviceMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIosysAgentDeviceMappingWithDefaults() *IosysAgentDeviceMapping {
	this := IosysAgentDeviceMapping{}
	var enable bool = true
	this.Enable = &enable
	var subtype DataSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	var down bool = false
	this.Down = *NewNullableBool(&down)
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetClass() AgentClass {
	if o == nil || IsNil(o.Class.Get()) {
		var ret AgentClass
		return ret
	}
	return *o.Class.Get()
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetClassOk() (*AgentClass, bool) {
	if o == nil {
		return nil, false
	}
	return o.Class.Get(), o.Class.IsSet()
}

// HasClass returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasClass() bool {
	if o != nil && o.Class.IsSet() {
		return true
	}

	return false
}

// SetClass gets a reference to the given NullableAgentClass and assigns it to the Class field.
func (o *IosysAgentDeviceMapping) SetClass(v AgentClass) {
	o.Class.Set(&v)
}

// SetClassNil sets the value for Class to be an explicit nil
func (o *IosysAgentDeviceMapping) SetClassNil() {
	o.Class.Set(nil)
}

// UnsetClass ensures that no value is present for Class, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetClass() {
	o.Class.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *IosysAgentDeviceMapping) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *IosysAgentDeviceMapping) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetId() {
	o.Id.Unset()
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetDeviceId() int32 {
	if o == nil || IsNil(o.DeviceId.Get()) {
		var ret int32
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetDeviceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableInt32 and assigns it to the DeviceId field.
func (o *IosysAgentDeviceMapping) SetDeviceId(v int32) {
	o.DeviceId.Set(&v)
}

// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *IosysAgentDeviceMapping) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *IosysAgentDeviceMapping) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IosysAgentDeviceMapping) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *IosysAgentDeviceMapping) SetEnable(v bool) {
	o.Enable = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetAssetId() int32 {
	if o == nil || IsNil(o.AssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.AssetId.Get()
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetId.Get(), o.AssetId.IsSet()
}

// HasAssetId returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasAssetId() bool {
	if o != nil && o.AssetId.IsSet() {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given NullableInt32 and assigns it to the AssetId field.
func (o *IosysAgentDeviceMapping) SetAssetId(v int32) {
	o.AssetId.Set(&v)
}

// SetAssetIdNil sets the value for AssetId to be an explicit nil
func (o *IosysAgentDeviceMapping) SetAssetIdNil() {
	o.AssetId.Set(nil)
}

// UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetAssetId() {
	o.AssetId.Unset()
}

// GetSubtype returns the Subtype field value
func (o *IosysAgentDeviceMapping) GetSubtype() DataSubtype {
	if o == nil {
		var ret DataSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *IosysAgentDeviceMapping) GetSubtypeOk() (*DataSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *IosysAgentDeviceMapping) SetSubtype(v DataSubtype) {
	o.Subtype = v
}

// GetAttribute returns the Attribute field value
func (o *IosysAgentDeviceMapping) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *IosysAgentDeviceMapping) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *IosysAgentDeviceMapping) SetAttribute(v string) {
	o.Attribute = v
}

// GetIosVar returns the IosVar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetIosVar() string {
	if o == nil || IsNil(o.IosVar.Get()) {
		var ret string
		return ret
	}
	return *o.IosVar.Get()
}

// GetIosVarOk returns a tuple with the IosVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetIosVarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IosVar.Get(), o.IosVar.IsSet()
}

// HasIosVar returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasIosVar() bool {
	if o != nil && o.IosVar.IsSet() {
		return true
	}

	return false
}

// SetIosVar gets a reference to the given NullableString and assigns it to the IosVar field.
func (o *IosysAgentDeviceMapping) SetIosVar(v string) {
	o.IosVar.Set(&v)
}

// SetIosVarNil sets the value for IosVar to be an explicit nil
func (o *IosysAgentDeviceMapping) SetIosVarNil() {
	o.IosVar.Set(nil)
}

// UnsetIosVar ensures that no value is present for IosVar, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetIosVar() {
	o.IosVar.Unset()
}

// GetIosType returns the IosType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetIosType() string {
	if o == nil || IsNil(o.IosType.Get()) {
		var ret string
		return ret
	}
	return *o.IosType.Get()
}

// GetIosTypeOk returns a tuple with the IosType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetIosTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IosType.Get(), o.IosType.IsSet()
}

// HasIosType returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasIosType() bool {
	if o != nil && o.IosType.IsSet() {
		return true
	}

	return false
}

// SetIosType gets a reference to the given NullableString and assigns it to the IosType field.
func (o *IosysAgentDeviceMapping) SetIosType(v string) {
	o.IosType.Set(&v)
}

// SetIosTypeNil sets the value for IosType to be an explicit nil
func (o *IosysAgentDeviceMapping) SetIosTypeNil() {
	o.IosType.Set(nil)
}

// UnsetIosType ensures that no value is present for IosType, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetIosType() {
	o.IosType.Unset()
}

// GetDown returns the Down field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetDown() bool {
	if o == nil || IsNil(o.Down.Get()) {
		var ret bool
		return ret
	}
	return *o.Down.Get()
}

// GetDownOk returns a tuple with the Down field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetDownOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Down.Get(), o.Down.IsSet()
}

// HasDown returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasDown() bool {
	if o != nil && o.Down.IsSet() {
		return true
	}

	return false
}

// SetDown gets a reference to the given NullableBool and assigns it to the Down field.
func (o *IosysAgentDeviceMapping) SetDown(v bool) {
	o.Down.Set(&v)
}

// SetDownNil sets the value for Down to be an explicit nil
func (o *IosysAgentDeviceMapping) SetDownNil() {
	o.Down.Set(nil)
}

// UnsetDown ensures that no value is present for Down, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetDown() {
	o.Down.Unset()
}

// GetScale returns the Scale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetScale() float64 {
	if o == nil || IsNil(o.Scale.Get()) {
		var ret float64
		return ret
	}
	return *o.Scale.Get()
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetScaleOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scale.Get(), o.Scale.IsSet()
}

// HasScale returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasScale() bool {
	if o != nil && o.Scale.IsSet() {
		return true
	}

	return false
}

// SetScale gets a reference to the given NullableFloat64 and assigns it to the Scale field.
func (o *IosysAgentDeviceMapping) SetScale(v float64) {
	o.Scale.Set(&v)
}

// SetScaleNil sets the value for Scale to be an explicit nil
func (o *IosysAgentDeviceMapping) SetScaleNil() {
	o.Scale.Set(nil)
}

// UnsetScale ensures that no value is present for Scale, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetScale() {
	o.Scale.Unset()
}

// GetZero returns the Zero field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetZero() float64 {
	if o == nil || IsNil(o.Zero.Get()) {
		var ret float64
		return ret
	}
	return *o.Zero.Get()
}

// GetZeroOk returns a tuple with the Zero field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetZeroOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zero.Get(), o.Zero.IsSet()
}

// HasZero returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasZero() bool {
	if o != nil && o.Zero.IsSet() {
		return true
	}

	return false
}

// SetZero gets a reference to the given NullableFloat64 and assigns it to the Zero field.
func (o *IosysAgentDeviceMapping) SetZero(v float64) {
	o.Zero.Set(&v)
}

// SetZeroNil sets the value for Zero to be an explicit nil
func (o *IosysAgentDeviceMapping) SetZeroNil() {
	o.Zero.Set(nil)
}

// UnsetZero ensures that no value is present for Zero, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetZero() {
	o.Zero.Unset()
}

// GetMask returns the Mask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetMask() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetMaskOk() ([]int64, bool) {
	if o == nil || IsNil(o.Mask) {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasMask() bool {
	if o != nil && IsNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given []int64 and assigns it to the Mask field.
func (o *IosysAgentDeviceMapping) SetMask(v []int64) {
	o.Mask = v
}

// GetMaskAttributes returns the MaskAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetMaskAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MaskAttributes
}

// GetMaskAttributesOk returns a tuple with the MaskAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetMaskAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.MaskAttributes) {
		return nil, false
	}
	return o.MaskAttributes, true
}

// HasMaskAttributes returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasMaskAttributes() bool {
	if o != nil && IsNil(o.MaskAttributes) {
		return true
	}

	return false
}

// SetMaskAttributes gets a reference to the given []string and assigns it to the MaskAttributes field.
func (o *IosysAgentDeviceMapping) SetMaskAttributes(v []string) {
	o.MaskAttributes = v
}

// GetDeadTime returns the DeadTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetDeadTime() int32 {
	if o == nil || IsNil(o.DeadTime.Get()) {
		var ret int32
		return ret
	}
	return *o.DeadTime.Get()
}

// GetDeadTimeOk returns a tuple with the DeadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetDeadTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeadTime.Get(), o.DeadTime.IsSet()
}

// HasDeadTime returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasDeadTime() bool {
	if o != nil && o.DeadTime.IsSet() {
		return true
	}

	return false
}

// SetDeadTime gets a reference to the given NullableInt32 and assigns it to the DeadTime field.
func (o *IosysAgentDeviceMapping) SetDeadTime(v int32) {
	o.DeadTime.Set(&v)
}

// SetDeadTimeNil sets the value for DeadTime to be an explicit nil
func (o *IosysAgentDeviceMapping) SetDeadTimeNil() {
	o.DeadTime.Set(nil)
}

// UnsetDeadTime ensures that no value is present for DeadTime, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetDeadTime() {
	o.DeadTime.Unset()
}

// GetDeadBand returns the DeadBand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetDeadBand() float64 {
	if o == nil || IsNil(o.DeadBand.Get()) {
		var ret float64
		return ret
	}
	return *o.DeadBand.Get()
}

// GetDeadBandOk returns a tuple with the DeadBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetDeadBandOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeadBand.Get(), o.DeadBand.IsSet()
}

// HasDeadBand returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasDeadBand() bool {
	if o != nil && o.DeadBand.IsSet() {
		return true
	}

	return false
}

// SetDeadBand gets a reference to the given NullableFloat64 and assigns it to the DeadBand field.
func (o *IosysAgentDeviceMapping) SetDeadBand(v float64) {
	o.DeadBand.Set(&v)
}

// SetDeadBandNil sets the value for DeadBand to be an explicit nil
func (o *IosysAgentDeviceMapping) SetDeadBandNil() {
	o.DeadBand.Set(nil)
}

// UnsetDeadBand ensures that no value is present for DeadBand, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetDeadBand() {
	o.DeadBand.Unset()
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetFilter() string {
	if o == nil || IsNil(o.Filter.Get()) {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableString and assigns it to the Filter field.
func (o *IosysAgentDeviceMapping) SetFilter(v string) {
	o.Filter.Set(&v)
}

// SetFilterNil sets the value for Filter to be an explicit nil
func (o *IosysAgentDeviceMapping) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetFilter() {
	o.Filter.Unset()
}

// GetTau returns the Tau field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentDeviceMapping) GetTau() float64 {
	if o == nil || IsNil(o.Tau.Get()) {
		var ret float64
		return ret
	}
	return *o.Tau.Get()
}

// GetTauOk returns a tuple with the Tau field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentDeviceMapping) GetTauOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tau.Get(), o.Tau.IsSet()
}

// HasTau returns a boolean if a field has been set.
func (o *IosysAgentDeviceMapping) HasTau() bool {
	if o != nil && o.Tau.IsSet() {
		return true
	}

	return false
}

// SetTau gets a reference to the given NullableFloat64 and assigns it to the Tau field.
func (o *IosysAgentDeviceMapping) SetTau(v float64) {
	o.Tau.Set(&v)
}

// SetTauNil sets the value for Tau to be an explicit nil
func (o *IosysAgentDeviceMapping) SetTauNil() {
	o.Tau.Set(nil)
}

// UnsetTau ensures that no value is present for Tau, not even an explicit nil
func (o *IosysAgentDeviceMapping) UnsetTau() {
	o.Tau.Unset()
}

func (o IosysAgentDeviceMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IosysAgentDeviceMapping) ToMap() (map[string]interface{}, error) {
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
	if o.IosVar.IsSet() {
		toSerialize["iosVar"] = o.IosVar.Get()
	}
	if o.IosType.IsSet() {
		toSerialize["iosType"] = o.IosType.Get()
	}
	if o.Down.IsSet() {
		toSerialize["down"] = o.Down.Get()
	}
	if o.Scale.IsSet() {
		toSerialize["scale"] = o.Scale.Get()
	}
	if o.Zero.IsSet() {
		toSerialize["zero"] = o.Zero.Get()
	}
	if o.Mask != nil {
		toSerialize["mask"] = o.Mask
	}
	if o.MaskAttributes != nil {
		toSerialize["maskAttributes"] = o.MaskAttributes
	}
	if o.DeadTime.IsSet() {
		toSerialize["deadTime"] = o.DeadTime.Get()
	}
	if o.DeadBand.IsSet() {
		toSerialize["deadBand"] = o.DeadBand.Get()
	}
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if o.Tau.IsSet() {
		toSerialize["tau"] = o.Tau.Get()
	}
	return toSerialize, nil
}

func (o *IosysAgentDeviceMapping) UnmarshalJSON(data []byte) (err error) {
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

	varIosysAgentDeviceMapping := _IosysAgentDeviceMapping{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIosysAgentDeviceMapping)

	if err != nil {
		return err
	}

	*o = IosysAgentDeviceMapping(varIosysAgentDeviceMapping)

	return err
}

type NullableIosysAgentDeviceMapping struct {
	value *IosysAgentDeviceMapping
	isSet bool
}

func (v NullableIosysAgentDeviceMapping) Get() *IosysAgentDeviceMapping {
	return v.value
}

func (v *NullableIosysAgentDeviceMapping) Set(val *IosysAgentDeviceMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableIosysAgentDeviceMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableIosysAgentDeviceMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIosysAgentDeviceMapping(val *IosysAgentDeviceMapping) *NullableIosysAgentDeviceMapping {
	return &NullableIosysAgentDeviceMapping{value: val, isSet: true}
}

func (v NullableIosysAgentDeviceMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIosysAgentDeviceMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
