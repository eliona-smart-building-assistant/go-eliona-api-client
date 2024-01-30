/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.1
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AlarmRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlarmRule{}

// AlarmRule Rule for an alarm
type AlarmRule struct {
	// The id of the rule
	Id NullableInt32 `json:"id,omitempty"`
	// ID of the corresponding asset
	AssetId int32       `json:"assetId"`
	Subtype DataSubtype `json:"subtype"`
	// Name of the attribute of the asset type
	Attribute string `json:"attribute"`
	// Rule enabled or not
	Enable   *bool         `json:"enable,omitempty"`
	Priority AlarmPriority `json:"priority"`
	// Requires the alarm an acknowledgment
	RequiresAcknowledge *bool `json:"requiresAcknowledge,omitempty"`
	// Triggers alarm if attribute value equals this value
	Equal NullableFloat64 `json:"equal,omitempty"`
	// Triggers alarm if attribute value is less than value
	Low NullableFloat64 `json:"low,omitempty"`
	// Triggers alarm if attribute value is greater than value
	High NullableFloat64 `json:"high,omitempty"`
	// Texts for alarm
	Message map[string]interface{} `json:"message,omitempty"`
	// List of associated tags
	Tags []string `json:"tags,omitempty"`
	// The subject for the alarm
	Subject NullableString `json:"subject,omitempty"`
	// The url describing the alarm
	Urldoc NullableString `json:"urldoc,omitempty"`
	// Notification
	NotifyOn NullableString `json:"notifyOn,omitempty"`
	// Do not mask
	DontMask NullableBool `json:"dontMask,omitempty"`
	// Check type
	CheckType NullableString `json:"checkType,omitempty"`
	AssetInfo NullableAsset  `json:"assetInfo,omitempty"`
}

type _AlarmRule AlarmRule

// NewAlarmRule instantiates a new AlarmRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmRule(assetId int32, subtype DataSubtype, attribute string, priority AlarmPriority) *AlarmRule {
	this := AlarmRule{}
	this.AssetId = assetId
	this.Subtype = subtype
	this.Attribute = attribute
	var enable bool = true
	this.Enable = &enable
	this.Priority = priority
	var requiresAcknowledge bool = false
	this.RequiresAcknowledge = &requiresAcknowledge
	var dontMask bool = false
	this.DontMask = *NewNullableBool(&dontMask)
	return &this
}

// NewAlarmRuleWithDefaults instantiates a new AlarmRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmRuleWithDefaults() *AlarmRule {
	this := AlarmRule{}
	var subtype DataSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	var enable bool = true
	this.Enable = &enable
	var requiresAcknowledge bool = false
	this.RequiresAcknowledge = &requiresAcknowledge
	var dontMask bool = false
	this.DontMask = *NewNullableBool(&dontMask)
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AlarmRule) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *AlarmRule) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *AlarmRule) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AlarmRule) UnsetId() {
	o.Id.Unset()
}

// GetAssetId returns the AssetId field value
func (o *AlarmRule) GetAssetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *AlarmRule) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *AlarmRule) SetAssetId(v int32) {
	o.AssetId = v
}

// GetSubtype returns the Subtype field value
func (o *AlarmRule) GetSubtype() DataSubtype {
	if o == nil {
		var ret DataSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *AlarmRule) GetSubtypeOk() (*DataSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *AlarmRule) SetSubtype(v DataSubtype) {
	o.Subtype = v
}

// GetAttribute returns the Attribute field value
func (o *AlarmRule) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *AlarmRule) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *AlarmRule) SetAttribute(v string) {
	o.Attribute = v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *AlarmRule) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRule) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *AlarmRule) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *AlarmRule) SetEnable(v bool) {
	o.Enable = &v
}

// GetPriority returns the Priority field value
func (o *AlarmRule) GetPriority() AlarmPriority {
	if o == nil {
		var ret AlarmPriority
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *AlarmRule) GetPriorityOk() (*AlarmPriority, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *AlarmRule) SetPriority(v AlarmPriority) {
	o.Priority = v
}

// GetRequiresAcknowledge returns the RequiresAcknowledge field value if set, zero value otherwise.
func (o *AlarmRule) GetRequiresAcknowledge() bool {
	if o == nil || IsNil(o.RequiresAcknowledge) {
		var ret bool
		return ret
	}
	return *o.RequiresAcknowledge
}

// GetRequiresAcknowledgeOk returns a tuple with the RequiresAcknowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmRule) GetRequiresAcknowledgeOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiresAcknowledge) {
		return nil, false
	}
	return o.RequiresAcknowledge, true
}

// HasRequiresAcknowledge returns a boolean if a field has been set.
func (o *AlarmRule) HasRequiresAcknowledge() bool {
	if o != nil && !IsNil(o.RequiresAcknowledge) {
		return true
	}

	return false
}

// SetRequiresAcknowledge gets a reference to the given bool and assigns it to the RequiresAcknowledge field.
func (o *AlarmRule) SetRequiresAcknowledge(v bool) {
	o.RequiresAcknowledge = &v
}

// GetEqual returns the Equal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetEqual() float64 {
	if o == nil || IsNil(o.Equal.Get()) {
		var ret float64
		return ret
	}
	return *o.Equal.Get()
}

// GetEqualOk returns a tuple with the Equal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetEqualOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Equal.Get(), o.Equal.IsSet()
}

// HasEqual returns a boolean if a field has been set.
func (o *AlarmRule) HasEqual() bool {
	if o != nil && o.Equal.IsSet() {
		return true
	}

	return false
}

// SetEqual gets a reference to the given NullableFloat64 and assigns it to the Equal field.
func (o *AlarmRule) SetEqual(v float64) {
	o.Equal.Set(&v)
}

// SetEqualNil sets the value for Equal to be an explicit nil
func (o *AlarmRule) SetEqualNil() {
	o.Equal.Set(nil)
}

// UnsetEqual ensures that no value is present for Equal, not even an explicit nil
func (o *AlarmRule) UnsetEqual() {
	o.Equal.Unset()
}

// GetLow returns the Low field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetLow() float64 {
	if o == nil || IsNil(o.Low.Get()) {
		var ret float64
		return ret
	}
	return *o.Low.Get()
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetLowOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Low.Get(), o.Low.IsSet()
}

// HasLow returns a boolean if a field has been set.
func (o *AlarmRule) HasLow() bool {
	if o != nil && o.Low.IsSet() {
		return true
	}

	return false
}

// SetLow gets a reference to the given NullableFloat64 and assigns it to the Low field.
func (o *AlarmRule) SetLow(v float64) {
	o.Low.Set(&v)
}

// SetLowNil sets the value for Low to be an explicit nil
func (o *AlarmRule) SetLowNil() {
	o.Low.Set(nil)
}

// UnsetLow ensures that no value is present for Low, not even an explicit nil
func (o *AlarmRule) UnsetLow() {
	o.Low.Unset()
}

// GetHigh returns the High field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetHigh() float64 {
	if o == nil || IsNil(o.High.Get()) {
		var ret float64
		return ret
	}
	return *o.High.Get()
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetHighOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.High.Get(), o.High.IsSet()
}

// HasHigh returns a boolean if a field has been set.
func (o *AlarmRule) HasHigh() bool {
	if o != nil && o.High.IsSet() {
		return true
	}

	return false
}

// SetHigh gets a reference to the given NullableFloat64 and assigns it to the High field.
func (o *AlarmRule) SetHigh(v float64) {
	o.High.Set(&v)
}

// SetHighNil sets the value for High to be an explicit nil
func (o *AlarmRule) SetHighNil() {
	o.High.Set(nil)
}

// UnsetHigh ensures that no value is present for High, not even an explicit nil
func (o *AlarmRule) UnsetHigh() {
	o.High.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetMessage() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetMessageOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Message) {
		return map[string]interface{}{}, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AlarmRule) HasMessage() bool {
	if o != nil && IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given map[string]interface{} and assigns it to the Message field.
func (o *AlarmRule) SetMessage(v map[string]interface{}) {
	o.Message = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AlarmRule) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AlarmRule) SetTags(v []string) {
	o.Tags = v
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetSubject() string {
	if o == nil || IsNil(o.Subject.Get()) {
		var ret string
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *AlarmRule) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullableString and assigns it to the Subject field.
func (o *AlarmRule) SetSubject(v string) {
	o.Subject.Set(&v)
}

// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *AlarmRule) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *AlarmRule) UnsetSubject() {
	o.Subject.Unset()
}

// GetUrldoc returns the Urldoc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetUrldoc() string {
	if o == nil || IsNil(o.Urldoc.Get()) {
		var ret string
		return ret
	}
	return *o.Urldoc.Get()
}

// GetUrldocOk returns a tuple with the Urldoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetUrldocOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Urldoc.Get(), o.Urldoc.IsSet()
}

// HasUrldoc returns a boolean if a field has been set.
func (o *AlarmRule) HasUrldoc() bool {
	if o != nil && o.Urldoc.IsSet() {
		return true
	}

	return false
}

// SetUrldoc gets a reference to the given NullableString and assigns it to the Urldoc field.
func (o *AlarmRule) SetUrldoc(v string) {
	o.Urldoc.Set(&v)
}

// SetUrldocNil sets the value for Urldoc to be an explicit nil
func (o *AlarmRule) SetUrldocNil() {
	o.Urldoc.Set(nil)
}

// UnsetUrldoc ensures that no value is present for Urldoc, not even an explicit nil
func (o *AlarmRule) UnsetUrldoc() {
	o.Urldoc.Unset()
}

// GetNotifyOn returns the NotifyOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetNotifyOn() string {
	if o == nil || IsNil(o.NotifyOn.Get()) {
		var ret string
		return ret
	}
	return *o.NotifyOn.Get()
}

// GetNotifyOnOk returns a tuple with the NotifyOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetNotifyOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotifyOn.Get(), o.NotifyOn.IsSet()
}

// HasNotifyOn returns a boolean if a field has been set.
func (o *AlarmRule) HasNotifyOn() bool {
	if o != nil && o.NotifyOn.IsSet() {
		return true
	}

	return false
}

// SetNotifyOn gets a reference to the given NullableString and assigns it to the NotifyOn field.
func (o *AlarmRule) SetNotifyOn(v string) {
	o.NotifyOn.Set(&v)
}

// SetNotifyOnNil sets the value for NotifyOn to be an explicit nil
func (o *AlarmRule) SetNotifyOnNil() {
	o.NotifyOn.Set(nil)
}

// UnsetNotifyOn ensures that no value is present for NotifyOn, not even an explicit nil
func (o *AlarmRule) UnsetNotifyOn() {
	o.NotifyOn.Unset()
}

// GetDontMask returns the DontMask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetDontMask() bool {
	if o == nil || IsNil(o.DontMask.Get()) {
		var ret bool
		return ret
	}
	return *o.DontMask.Get()
}

// GetDontMaskOk returns a tuple with the DontMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetDontMaskOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DontMask.Get(), o.DontMask.IsSet()
}

// HasDontMask returns a boolean if a field has been set.
func (o *AlarmRule) HasDontMask() bool {
	if o != nil && o.DontMask.IsSet() {
		return true
	}

	return false
}

// SetDontMask gets a reference to the given NullableBool and assigns it to the DontMask field.
func (o *AlarmRule) SetDontMask(v bool) {
	o.DontMask.Set(&v)
}

// SetDontMaskNil sets the value for DontMask to be an explicit nil
func (o *AlarmRule) SetDontMaskNil() {
	o.DontMask.Set(nil)
}

// UnsetDontMask ensures that no value is present for DontMask, not even an explicit nil
func (o *AlarmRule) UnsetDontMask() {
	o.DontMask.Unset()
}

// GetCheckType returns the CheckType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetCheckType() string {
	if o == nil || IsNil(o.CheckType.Get()) {
		var ret string
		return ret
	}
	return *o.CheckType.Get()
}

// GetCheckTypeOk returns a tuple with the CheckType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetCheckTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckType.Get(), o.CheckType.IsSet()
}

// HasCheckType returns a boolean if a field has been set.
func (o *AlarmRule) HasCheckType() bool {
	if o != nil && o.CheckType.IsSet() {
		return true
	}

	return false
}

// SetCheckType gets a reference to the given NullableString and assigns it to the CheckType field.
func (o *AlarmRule) SetCheckType(v string) {
	o.CheckType.Set(&v)
}

// SetCheckTypeNil sets the value for CheckType to be an explicit nil
func (o *AlarmRule) SetCheckTypeNil() {
	o.CheckType.Set(nil)
}

// UnsetCheckType ensures that no value is present for CheckType, not even an explicit nil
func (o *AlarmRule) UnsetCheckType() {
	o.CheckType.Unset()
}

// GetAssetInfo returns the AssetInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmRule) GetAssetInfo() Asset {
	if o == nil || IsNil(o.AssetInfo.Get()) {
		var ret Asset
		return ret
	}
	return *o.AssetInfo.Get()
}

// GetAssetInfoOk returns a tuple with the AssetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmRule) GetAssetInfoOk() (*Asset, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetInfo.Get(), o.AssetInfo.IsSet()
}

// HasAssetInfo returns a boolean if a field has been set.
func (o *AlarmRule) HasAssetInfo() bool {
	if o != nil && o.AssetInfo.IsSet() {
		return true
	}

	return false
}

// SetAssetInfo gets a reference to the given NullableAsset and assigns it to the AssetInfo field.
func (o *AlarmRule) SetAssetInfo(v Asset) {
	o.AssetInfo.Set(&v)
}

// SetAssetInfoNil sets the value for AssetInfo to be an explicit nil
func (o *AlarmRule) SetAssetInfoNil() {
	o.AssetInfo.Set(nil)
}

// UnsetAssetInfo ensures that no value is present for AssetInfo, not even an explicit nil
func (o *AlarmRule) UnsetAssetInfo() {
	o.AssetInfo.Unset()
}

func (o AlarmRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlarmRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["assetId"] = o.AssetId
	toSerialize["subtype"] = o.Subtype
	toSerialize["attribute"] = o.Attribute
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	toSerialize["priority"] = o.Priority
	if !IsNil(o.RequiresAcknowledge) {
		toSerialize["requiresAcknowledge"] = o.RequiresAcknowledge
	}
	if o.Equal.IsSet() {
		toSerialize["equal"] = o.Equal.Get()
	}
	if o.Low.IsSet() {
		toSerialize["low"] = o.Low.Get()
	}
	if o.High.IsSet() {
		toSerialize["high"] = o.High.Get()
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Subject.IsSet() {
		toSerialize["subject"] = o.Subject.Get()
	}
	if o.Urldoc.IsSet() {
		toSerialize["urldoc"] = o.Urldoc.Get()
	}
	if o.NotifyOn.IsSet() {
		toSerialize["notifyOn"] = o.NotifyOn.Get()
	}
	if o.DontMask.IsSet() {
		toSerialize["dontMask"] = o.DontMask.Get()
	}
	if o.CheckType.IsSet() {
		toSerialize["checkType"] = o.CheckType.Get()
	}
	if o.AssetInfo.IsSet() {
		toSerialize["assetInfo"] = o.AssetInfo.Get()
	}
	return toSerialize, nil
}

func (o *AlarmRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assetId",
		"subtype",
		"attribute",
		"priority",
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

	varAlarmRule := _AlarmRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlarmRule)

	if err != nil {
		return err
	}

	*o = AlarmRule(varAlarmRule)

	return err
}

type NullableAlarmRule struct {
	value *AlarmRule
	isSet bool
}

func (v NullableAlarmRule) Get() *AlarmRule {
	return v.value
}

func (v *NullableAlarmRule) Set(val *AlarmRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmRule(val *AlarmRule) *NullableAlarmRule {
	return &NullableAlarmRule{value: val, isSet: true}
}

func (v NullableAlarmRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
