/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.7.2
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AssetType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetType{}

// AssetType A type of assets
type AssetType struct {
	// The unique name for this asset type
	Name string `json:"name"`
	// Is this a customer created type or not
	Custom *bool `json:"custom,omitempty"`
	// The vendor providing assets of this type
	Vendor NullableString `json:"vendor,omitempty"`
	// The specific model of assets of this type
	Model       NullableString      `json:"model,omitempty"`
	Translation NullableTranslation `json:"translation,omitempty"`
	// The url describing assets of this type
	Urldoc NullableString `json:"urldoc,omitempty"`
	// Icon name corresponding to assets of this type: blind, building, button, closable, elevator, environment, fallback, filling, gateway, light, mailbox, parking, people, power, rack, storey, trash, ventilation, vibration, water, weather
	Icon NullableString `json:"icon,omitempty"`
	// Asset types payload function
	PayloadFunction   NullableString `json:"payloadFunction,omitempty"`
	AllowedInactivity NullableString `json:"allowedInactivity,omitempty"`
	// Function as a tracker
	IsTracker NullableBool `json:"isTracker,omitempty"`
	// List of named attributes
	Attributes []AssetTypeAttribute `json:"attributes,omitempty"`
}

type _AssetType AssetType

// NewAssetType instantiates a new AssetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetType(name string) *AssetType {
	this := AssetType{}
	this.Name = name
	var custom bool = true
	this.Custom = &custom
	var isTracker bool = false
	this.IsTracker = *NewNullableBool(&isTracker)
	return &this
}

// NewAssetTypeWithDefaults instantiates a new AssetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTypeWithDefaults() *AssetType {
	this := AssetType{}
	var custom bool = true
	this.Custom = &custom
	var isTracker bool = false
	this.IsTracker = *NewNullableBool(&isTracker)
	return &this
}

// GetName returns the Name field value
func (o *AssetType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AssetType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AssetType) SetName(v string) {
	o.Name = v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *AssetType) GetCustom() bool {
	if o == nil || IsNil(o.Custom) {
		var ret bool
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetType) GetCustomOk() (*bool, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *AssetType) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given bool and assigns it to the Custom field.
func (o *AssetType) SetCustom(v bool) {
	o.Custom = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetType) GetVendor() string {
	if o == nil || IsNil(o.Vendor.Get()) {
		var ret string
		return ret
	}
	return *o.Vendor.Get()
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetType) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vendor.Get(), o.Vendor.IsSet()
}

// HasVendor returns a boolean if a field has been set.
func (o *AssetType) HasVendor() bool {
	if o != nil && o.Vendor.IsSet() {
		return true
	}

	return false
}

// SetVendor gets a reference to the given NullableString and assigns it to the Vendor field.
func (o *AssetType) SetVendor(v string) {
	o.Vendor.Set(&v)
}

// SetVendorNil sets the value for Vendor to be an explicit nil
func (o *AssetType) SetVendorNil() {
	o.Vendor.Set(nil)
}

// UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
func (o *AssetType) UnsetVendor() {
	o.Vendor.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetType) GetModel() string {
	if o == nil || IsNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *AssetType) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *AssetType) SetModel(v string) {
	o.Model.Set(&v)
}

// SetModelNil sets the value for Model to be an explicit nil
func (o *AssetType) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *AssetType) UnsetModel() {
	o.Model.Unset()
}

// GetTranslation returns the Translation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetType) GetTranslation() Translation {
	if o == nil || IsNil(o.Translation.Get()) {
		var ret Translation
		return ret
	}
	return *o.Translation.Get()
}

// GetTranslationOk returns a tuple with the Translation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetType) GetTranslationOk() (*Translation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Translation.Get(), o.Translation.IsSet()
}

// HasTranslation returns a boolean if a field has been set.
func (o *AssetType) HasTranslation() bool {
	if o != nil && o.Translation.IsSet() {
		return true
	}

	return false
}

// SetTranslation gets a reference to the given NullableTranslation and assigns it to the Translation field.
func (o *AssetType) SetTranslation(v Translation) {
	o.Translation.Set(&v)
}

// SetTranslationNil sets the value for Translation to be an explicit nil
func (o *AssetType) SetTranslationNil() {
	o.Translation.Set(nil)
}

// UnsetTranslation ensures that no value is present for Translation, not even an explicit nil
func (o *AssetType) UnsetTranslation() {
	o.Translation.Unset()
}

// GetUrldoc returns the Urldoc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetType) GetUrldoc() string {
	if o == nil || IsNil(o.Urldoc.Get()) {
		var ret string
		return ret
	}
	return *o.Urldoc.Get()
}

// GetUrldocOk returns a tuple with the Urldoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetType) GetUrldocOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Urldoc.Get(), o.Urldoc.IsSet()
}

// HasUrldoc returns a boolean if a field has been set.
func (o *AssetType) HasUrldoc() bool {
	if o != nil && o.Urldoc.IsSet() {
		return true
	}

	return false
}

// SetUrldoc gets a reference to the given NullableString and assigns it to the Urldoc field.
func (o *AssetType) SetUrldoc(v string) {
	o.Urldoc.Set(&v)
}

// SetUrldocNil sets the value for Urldoc to be an explicit nil
func (o *AssetType) SetUrldocNil() {
	o.Urldoc.Set(nil)
}

// UnsetUrldoc ensures that no value is present for Urldoc, not even an explicit nil
func (o *AssetType) UnsetUrldoc() {
	o.Urldoc.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetType) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetType) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *AssetType) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *AssetType) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *AssetType) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *AssetType) UnsetIcon() {
	o.Icon.Unset()
}

// GetPayloadFunction returns the PayloadFunction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetType) GetPayloadFunction() string {
	if o == nil || IsNil(o.PayloadFunction.Get()) {
		var ret string
		return ret
	}
	return *o.PayloadFunction.Get()
}

// GetPayloadFunctionOk returns a tuple with the PayloadFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetType) GetPayloadFunctionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayloadFunction.Get(), o.PayloadFunction.IsSet()
}

// HasPayloadFunction returns a boolean if a field has been set.
func (o *AssetType) HasPayloadFunction() bool {
	if o != nil && o.PayloadFunction.IsSet() {
		return true
	}

	return false
}

// SetPayloadFunction gets a reference to the given NullableString and assigns it to the PayloadFunction field.
func (o *AssetType) SetPayloadFunction(v string) {
	o.PayloadFunction.Set(&v)
}

// SetPayloadFunctionNil sets the value for PayloadFunction to be an explicit nil
func (o *AssetType) SetPayloadFunctionNil() {
	o.PayloadFunction.Set(nil)
}

// UnsetPayloadFunction ensures that no value is present for PayloadFunction, not even an explicit nil
func (o *AssetType) UnsetPayloadFunction() {
	o.PayloadFunction.Unset()
}

// GetAllowedInactivity returns the AllowedInactivity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetType) GetAllowedInactivity() string {
	if o == nil || IsNil(o.AllowedInactivity.Get()) {
		var ret string
		return ret
	}
	return *o.AllowedInactivity.Get()
}

// GetAllowedInactivityOk returns a tuple with the AllowedInactivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetType) GetAllowedInactivityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedInactivity.Get(), o.AllowedInactivity.IsSet()
}

// HasAllowedInactivity returns a boolean if a field has been set.
func (o *AssetType) HasAllowedInactivity() bool {
	if o != nil && o.AllowedInactivity.IsSet() {
		return true
	}

	return false
}

// SetAllowedInactivity gets a reference to the given NullableString and assigns it to the AllowedInactivity field.
func (o *AssetType) SetAllowedInactivity(v string) {
	o.AllowedInactivity.Set(&v)
}

// SetAllowedInactivityNil sets the value for AllowedInactivity to be an explicit nil
func (o *AssetType) SetAllowedInactivityNil() {
	o.AllowedInactivity.Set(nil)
}

// UnsetAllowedInactivity ensures that no value is present for AllowedInactivity, not even an explicit nil
func (o *AssetType) UnsetAllowedInactivity() {
	o.AllowedInactivity.Unset()
}

// GetIsTracker returns the IsTracker field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetType) GetIsTracker() bool {
	if o == nil || IsNil(o.IsTracker.Get()) {
		var ret bool
		return ret
	}
	return *o.IsTracker.Get()
}

// GetIsTrackerOk returns a tuple with the IsTracker field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetType) GetIsTrackerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsTracker.Get(), o.IsTracker.IsSet()
}

// HasIsTracker returns a boolean if a field has been set.
func (o *AssetType) HasIsTracker() bool {
	if o != nil && o.IsTracker.IsSet() {
		return true
	}

	return false
}

// SetIsTracker gets a reference to the given NullableBool and assigns it to the IsTracker field.
func (o *AssetType) SetIsTracker(v bool) {
	o.IsTracker.Set(&v)
}

// SetIsTrackerNil sets the value for IsTracker to be an explicit nil
func (o *AssetType) SetIsTrackerNil() {
	o.IsTracker.Set(nil)
}

// UnsetIsTracker ensures that no value is present for IsTracker, not even an explicit nil
func (o *AssetType) UnsetIsTracker() {
	o.IsTracker.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetType) GetAttributes() []AssetTypeAttribute {
	if o == nil {
		var ret []AssetTypeAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetType) GetAttributesOk() ([]AssetTypeAttribute, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AssetType) HasAttributes() bool {
	if o != nil && IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []AssetTypeAttribute and assigns it to the Attributes field.
func (o *AssetType) SetAttributes(v []AssetTypeAttribute) {
	o.Attributes = v
}

func (o AssetType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if o.Vendor.IsSet() {
		toSerialize["vendor"] = o.Vendor.Get()
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if o.Translation.IsSet() {
		toSerialize["translation"] = o.Translation.Get()
	}
	if o.Urldoc.IsSet() {
		toSerialize["urldoc"] = o.Urldoc.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.PayloadFunction.IsSet() {
		toSerialize["payloadFunction"] = o.PayloadFunction.Get()
	}
	if o.AllowedInactivity.IsSet() {
		toSerialize["allowedInactivity"] = o.AllowedInactivity.Get()
	}
	if o.IsTracker.IsSet() {
		toSerialize["isTracker"] = o.IsTracker.Get()
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *AssetType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAssetType := _AssetType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssetType)

	if err != nil {
		return err
	}

	*o = AssetType(varAssetType)

	return err
}

type NullableAssetType struct {
	value *AssetType
	isSet bool
}

func (v NullableAssetType) Get() *AssetType {
	return v.value
}

func (v *NullableAssetType) Set(val *AssetType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetType(val *AssetType) *NullableAssetType {
	return &NullableAssetType{value: val, isSet: true}
}

func (v NullableAssetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
