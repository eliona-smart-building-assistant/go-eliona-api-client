/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// IosysAgentMappingSpecific Specific mapping for IOSYS agents
type IosysAgentMappingSpecific struct {
	IosVar NullableString `json:"iosVar,omitempty"`
	IosType NullableString `json:"iosType,omitempty"`
	Down bool `json:"down"`
	Scale NullableFloat64 `json:"scale,omitempty"`
	Zero NullableFloat64 `json:"zero,omitempty"`
	Mask []int32 `json:"mask,omitempty"`
	MaskAttributes []string `json:"maskAttributes,omitempty"`
	DeadTime NullableInt32 `json:"deadTime,omitempty"`
	DeadBand NullableFloat64 `json:"deadBand,omitempty"`
	Filter NullableString `json:"filter,omitempty"`
	Tau NullableFloat64 `json:"tau,omitempty"`
}

// NewIosysAgentMappingSpecific instantiates a new IosysAgentMappingSpecific object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIosysAgentMappingSpecific(down bool) *IosysAgentMappingSpecific {
	this := IosysAgentMappingSpecific{}
	this.Down = down
	return &this
}

// NewIosysAgentMappingSpecificWithDefaults instantiates a new IosysAgentMappingSpecific object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIosysAgentMappingSpecificWithDefaults() *IosysAgentMappingSpecific {
	this := IosysAgentMappingSpecific{}
	var down bool = false
	this.Down = down
	return &this
}

// GetIosVar returns the IosVar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentMappingSpecific) GetIosVar() string {
	if o == nil || o.IosVar.Get() == nil {
		var ret string
		return ret
	}
	return *o.IosVar.Get()
}

// GetIosVarOk returns a tuple with the IosVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentMappingSpecific) GetIosVarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IosVar.Get(), o.IosVar.IsSet()
}

// HasIosVar returns a boolean if a field has been set.
func (o *IosysAgentMappingSpecific) HasIosVar() bool {
	if o != nil && o.IosVar.IsSet() {
		return true
	}

	return false
}

// SetIosVar gets a reference to the given NullableString and assigns it to the IosVar field.
func (o *IosysAgentMappingSpecific) SetIosVar(v string) {
	o.IosVar.Set(&v)
}
// SetIosVarNil sets the value for IosVar to be an explicit nil
func (o *IosysAgentMappingSpecific) SetIosVarNil() {
	o.IosVar.Set(nil)
}

// UnsetIosVar ensures that no value is present for IosVar, not even an explicit nil
func (o *IosysAgentMappingSpecific) UnsetIosVar() {
	o.IosVar.Unset()
}

// GetIosType returns the IosType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentMappingSpecific) GetIosType() string {
	if o == nil || o.IosType.Get() == nil {
		var ret string
		return ret
	}
	return *o.IosType.Get()
}

// GetIosTypeOk returns a tuple with the IosType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentMappingSpecific) GetIosTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IosType.Get(), o.IosType.IsSet()
}

// HasIosType returns a boolean if a field has been set.
func (o *IosysAgentMappingSpecific) HasIosType() bool {
	if o != nil && o.IosType.IsSet() {
		return true
	}

	return false
}

// SetIosType gets a reference to the given NullableString and assigns it to the IosType field.
func (o *IosysAgentMappingSpecific) SetIosType(v string) {
	o.IosType.Set(&v)
}
// SetIosTypeNil sets the value for IosType to be an explicit nil
func (o *IosysAgentMappingSpecific) SetIosTypeNil() {
	o.IosType.Set(nil)
}

// UnsetIosType ensures that no value is present for IosType, not even an explicit nil
func (o *IosysAgentMappingSpecific) UnsetIosType() {
	o.IosType.Unset()
}

// GetDown returns the Down field value
func (o *IosysAgentMappingSpecific) GetDown() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Down
}

// GetDownOk returns a tuple with the Down field value
// and a boolean to check if the value has been set.
func (o *IosysAgentMappingSpecific) GetDownOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Down, true
}

// SetDown sets field value
func (o *IosysAgentMappingSpecific) SetDown(v bool) {
	o.Down = v
}

// GetScale returns the Scale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentMappingSpecific) GetScale() float64 {
	if o == nil || o.Scale.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Scale.Get()
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentMappingSpecific) GetScaleOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scale.Get(), o.Scale.IsSet()
}

// HasScale returns a boolean if a field has been set.
func (o *IosysAgentMappingSpecific) HasScale() bool {
	if o != nil && o.Scale.IsSet() {
		return true
	}

	return false
}

// SetScale gets a reference to the given NullableFloat64 and assigns it to the Scale field.
func (o *IosysAgentMappingSpecific) SetScale(v float64) {
	o.Scale.Set(&v)
}
// SetScaleNil sets the value for Scale to be an explicit nil
func (o *IosysAgentMappingSpecific) SetScaleNil() {
	o.Scale.Set(nil)
}

// UnsetScale ensures that no value is present for Scale, not even an explicit nil
func (o *IosysAgentMappingSpecific) UnsetScale() {
	o.Scale.Unset()
}

// GetZero returns the Zero field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentMappingSpecific) GetZero() float64 {
	if o == nil || o.Zero.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Zero.Get()
}

// GetZeroOk returns a tuple with the Zero field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentMappingSpecific) GetZeroOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zero.Get(), o.Zero.IsSet()
}

// HasZero returns a boolean if a field has been set.
func (o *IosysAgentMappingSpecific) HasZero() bool {
	if o != nil && o.Zero.IsSet() {
		return true
	}

	return false
}

// SetZero gets a reference to the given NullableFloat64 and assigns it to the Zero field.
func (o *IosysAgentMappingSpecific) SetZero(v float64) {
	o.Zero.Set(&v)
}
// SetZeroNil sets the value for Zero to be an explicit nil
func (o *IosysAgentMappingSpecific) SetZeroNil() {
	o.Zero.Set(nil)
}

// UnsetZero ensures that no value is present for Zero, not even an explicit nil
func (o *IosysAgentMappingSpecific) UnsetZero() {
	o.Zero.Unset()
}

// GetMask returns the Mask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentMappingSpecific) GetMask() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentMappingSpecific) GetMaskOk() ([]int32, bool) {
	if o == nil || o.Mask == nil {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *IosysAgentMappingSpecific) HasMask() bool {
	if o != nil && o.Mask != nil {
		return true
	}

	return false
}

// SetMask gets a reference to the given []int32 and assigns it to the Mask field.
func (o *IosysAgentMappingSpecific) SetMask(v []int32) {
	o.Mask = v
}

// GetMaskAttributes returns the MaskAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentMappingSpecific) GetMaskAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MaskAttributes
}

// GetMaskAttributesOk returns a tuple with the MaskAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentMappingSpecific) GetMaskAttributesOk() ([]string, bool) {
	if o == nil || o.MaskAttributes == nil {
		return nil, false
	}
	return o.MaskAttributes, true
}

// HasMaskAttributes returns a boolean if a field has been set.
func (o *IosysAgentMappingSpecific) HasMaskAttributes() bool {
	if o != nil && o.MaskAttributes != nil {
		return true
	}

	return false
}

// SetMaskAttributes gets a reference to the given []string and assigns it to the MaskAttributes field.
func (o *IosysAgentMappingSpecific) SetMaskAttributes(v []string) {
	o.MaskAttributes = v
}

// GetDeadTime returns the DeadTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentMappingSpecific) GetDeadTime() int32 {
	if o == nil || o.DeadTime.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DeadTime.Get()
}

// GetDeadTimeOk returns a tuple with the DeadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentMappingSpecific) GetDeadTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeadTime.Get(), o.DeadTime.IsSet()
}

// HasDeadTime returns a boolean if a field has been set.
func (o *IosysAgentMappingSpecific) HasDeadTime() bool {
	if o != nil && o.DeadTime.IsSet() {
		return true
	}

	return false
}

// SetDeadTime gets a reference to the given NullableInt32 and assigns it to the DeadTime field.
func (o *IosysAgentMappingSpecific) SetDeadTime(v int32) {
	o.DeadTime.Set(&v)
}
// SetDeadTimeNil sets the value for DeadTime to be an explicit nil
func (o *IosysAgentMappingSpecific) SetDeadTimeNil() {
	o.DeadTime.Set(nil)
}

// UnsetDeadTime ensures that no value is present for DeadTime, not even an explicit nil
func (o *IosysAgentMappingSpecific) UnsetDeadTime() {
	o.DeadTime.Unset()
}

// GetDeadBand returns the DeadBand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentMappingSpecific) GetDeadBand() float64 {
	if o == nil || o.DeadBand.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DeadBand.Get()
}

// GetDeadBandOk returns a tuple with the DeadBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentMappingSpecific) GetDeadBandOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeadBand.Get(), o.DeadBand.IsSet()
}

// HasDeadBand returns a boolean if a field has been set.
func (o *IosysAgentMappingSpecific) HasDeadBand() bool {
	if o != nil && o.DeadBand.IsSet() {
		return true
	}

	return false
}

// SetDeadBand gets a reference to the given NullableFloat64 and assigns it to the DeadBand field.
func (o *IosysAgentMappingSpecific) SetDeadBand(v float64) {
	o.DeadBand.Set(&v)
}
// SetDeadBandNil sets the value for DeadBand to be an explicit nil
func (o *IosysAgentMappingSpecific) SetDeadBandNil() {
	o.DeadBand.Set(nil)
}

// UnsetDeadBand ensures that no value is present for DeadBand, not even an explicit nil
func (o *IosysAgentMappingSpecific) UnsetDeadBand() {
	o.DeadBand.Unset()
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentMappingSpecific) GetFilter() string {
	if o == nil || o.Filter.Get() == nil {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentMappingSpecific) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *IosysAgentMappingSpecific) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableString and assigns it to the Filter field.
func (o *IosysAgentMappingSpecific) SetFilter(v string) {
	o.Filter.Set(&v)
}
// SetFilterNil sets the value for Filter to be an explicit nil
func (o *IosysAgentMappingSpecific) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *IosysAgentMappingSpecific) UnsetFilter() {
	o.Filter.Unset()
}

// GetTau returns the Tau field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IosysAgentMappingSpecific) GetTau() float64 {
	if o == nil || o.Tau.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Tau.Get()
}

// GetTauOk returns a tuple with the Tau field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IosysAgentMappingSpecific) GetTauOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tau.Get(), o.Tau.IsSet()
}

// HasTau returns a boolean if a field has been set.
func (o *IosysAgentMappingSpecific) HasTau() bool {
	if o != nil && o.Tau.IsSet() {
		return true
	}

	return false
}

// SetTau gets a reference to the given NullableFloat64 and assigns it to the Tau field.
func (o *IosysAgentMappingSpecific) SetTau(v float64) {
	o.Tau.Set(&v)
}
// SetTauNil sets the value for Tau to be an explicit nil
func (o *IosysAgentMappingSpecific) SetTauNil() {
	o.Tau.Set(nil)
}

// UnsetTau ensures that no value is present for Tau, not even an explicit nil
func (o *IosysAgentMappingSpecific) UnsetTau() {
	o.Tau.Unset()
}

func (o IosysAgentMappingSpecific) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IosVar.IsSet() {
		toSerialize["iosVar"] = o.IosVar.Get()
	}
	if o.IosType.IsSet() {
		toSerialize["iosType"] = o.IosType.Get()
	}
	if true {
		toSerialize["down"] = o.Down
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
	return json.Marshal(toSerialize)
}

type NullableIosysAgentMappingSpecific struct {
	value *IosysAgentMappingSpecific
	isSet bool
}

func (v NullableIosysAgentMappingSpecific) Get() *IosysAgentMappingSpecific {
	return v.value
}

func (v *NullableIosysAgentMappingSpecific) Set(val *IosysAgentMappingSpecific) {
	v.value = val
	v.isSet = true
}

func (v NullableIosysAgentMappingSpecific) IsSet() bool {
	return v.isSet
}

func (v *NullableIosysAgentMappingSpecific) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIosysAgentMappingSpecific(val *IosysAgentMappingSpecific) *NullableIosysAgentMappingSpecific {
	return &NullableIosysAgentMappingSpecific{value: val, isSet: true}
}

func (v NullableIosysAgentMappingSpecific) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIosysAgentMappingSpecific) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


