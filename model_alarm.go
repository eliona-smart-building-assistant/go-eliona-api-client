/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// Alarm An alarm
type Alarm struct {
	// The id of the corresponding rule
	RuleId NullableInt32 `json:"ruleId,omitempty"`
	// ID of the corresponding asset
	AssetId int32 `json:"assetId"`
	Subtype HeapSubtype `json:"subtype"`
	// Name of the attribute of the asset type
	Attribute NullableString `json:"attribute,omitempty"`
	Priority AlarmPriority `json:"priority"`
	// Requires the alarm an acknowledgment
	RequiresAcknowledge *bool `json:"requiresAcknowledge,omitempty"`
	// The value which triggers the alarm
	Value NullableFloat64 `json:"value,omitempty"`
	// Timestamp of the latest data change
	Timestamp time.Time `json:"timestamp"`
	// Timestamp of the latest data change
	GoneTimestamp NullableTime `json:"goneTimestamp,omitempty"`
	// Timestamp of the latest data change
	AcknowledgeTimestamp NullableTime `json:"acknowledgeTimestamp,omitempty"`
	// How often this alarm is triggered
	Occurrences int32 `json:"occurrences"`
	// Text of acknowledgement
	AcknowledgeText NullableString `json:"acknowledgeText,omitempty"`
	// User who acknowledged the alarm
	AcknowledgeUserId NullableString `json:"acknowledgeUserId,omitempty"`
	// Message texts for alarm
	Message map[string]interface{} `json:"message,omitempty"`
	AssetInfo NullableAsset `json:"assetInfo,omitempty"`
	RuleInfo NullableAlarmRule `json:"ruleInfo,omitempty"`
}

// NewAlarm instantiates a new Alarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarm(assetId int32, subtype HeapSubtype, priority AlarmPriority, timestamp time.Time, occurrences int32) *Alarm {
	this := Alarm{}
	this.AssetId = assetId
	this.Subtype = subtype
	this.Priority = priority
	var requiresAcknowledge bool = false
	this.RequiresAcknowledge = &requiresAcknowledge
	this.Timestamp = timestamp
	this.Occurrences = occurrences
	return &this
}

// NewAlarmWithDefaults instantiates a new Alarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmWithDefaults() *Alarm {
	this := Alarm{}
	var subtype HeapSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	var requiresAcknowledge bool = false
	this.RequiresAcknowledge = &requiresAcknowledge
	return &this
}

// GetRuleId returns the RuleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alarm) GetRuleId() int32 {
	if o == nil || o.RuleId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RuleId.Get()
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alarm) GetRuleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleId.Get(), o.RuleId.IsSet()
}

// HasRuleId returns a boolean if a field has been set.
func (o *Alarm) HasRuleId() bool {
	if o != nil && o.RuleId.IsSet() {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given NullableInt32 and assigns it to the RuleId field.
func (o *Alarm) SetRuleId(v int32) {
	o.RuleId.Set(&v)
}
// SetRuleIdNil sets the value for RuleId to be an explicit nil
func (o *Alarm) SetRuleIdNil() {
	o.RuleId.Set(nil)
}

// UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
func (o *Alarm) UnsetRuleId() {
	o.RuleId.Unset()
}

// GetAssetId returns the AssetId field value
func (o *Alarm) GetAssetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *Alarm) SetAssetId(v int32) {
	o.AssetId = v
}

// GetSubtype returns the Subtype field value
func (o *Alarm) GetSubtype() HeapSubtype {
	if o == nil {
		var ret HeapSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetSubtypeOk() (*HeapSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *Alarm) SetSubtype(v HeapSubtype) {
	o.Subtype = v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alarm) GetAttribute() string {
	if o == nil || o.Attribute.Get() == nil {
		var ret string
		return ret
	}
	return *o.Attribute.Get()
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alarm) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attribute.Get(), o.Attribute.IsSet()
}

// HasAttribute returns a boolean if a field has been set.
func (o *Alarm) HasAttribute() bool {
	if o != nil && o.Attribute.IsSet() {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given NullableString and assigns it to the Attribute field.
func (o *Alarm) SetAttribute(v string) {
	o.Attribute.Set(&v)
}
// SetAttributeNil sets the value for Attribute to be an explicit nil
func (o *Alarm) SetAttributeNil() {
	o.Attribute.Set(nil)
}

// UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
func (o *Alarm) UnsetAttribute() {
	o.Attribute.Unset()
}

// GetPriority returns the Priority field value
func (o *Alarm) GetPriority() AlarmPriority {
	if o == nil {
		var ret AlarmPriority
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetPriorityOk() (*AlarmPriority, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *Alarm) SetPriority(v AlarmPriority) {
	o.Priority = v
}

// GetRequiresAcknowledge returns the RequiresAcknowledge field value if set, zero value otherwise.
func (o *Alarm) GetRequiresAcknowledge() bool {
	if o == nil || o.RequiresAcknowledge == nil {
		var ret bool
		return ret
	}
	return *o.RequiresAcknowledge
}

// GetRequiresAcknowledgeOk returns a tuple with the RequiresAcknowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alarm) GetRequiresAcknowledgeOk() (*bool, bool) {
	if o == nil || o.RequiresAcknowledge == nil {
		return nil, false
	}
	return o.RequiresAcknowledge, true
}

// HasRequiresAcknowledge returns a boolean if a field has been set.
func (o *Alarm) HasRequiresAcknowledge() bool {
	if o != nil && o.RequiresAcknowledge != nil {
		return true
	}

	return false
}

// SetRequiresAcknowledge gets a reference to the given bool and assigns it to the RequiresAcknowledge field.
func (o *Alarm) SetRequiresAcknowledge(v bool) {
	o.RequiresAcknowledge = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alarm) GetValue() float64 {
	if o == nil || o.Value.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alarm) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *Alarm) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableFloat64 and assigns it to the Value field.
func (o *Alarm) SetValue(v float64) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *Alarm) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *Alarm) UnsetValue() {
	o.Value.Unset()
}

// GetTimestamp returns the Timestamp field value
func (o *Alarm) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Alarm) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetGoneTimestamp returns the GoneTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alarm) GetGoneTimestamp() time.Time {
	if o == nil || o.GoneTimestamp.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.GoneTimestamp.Get()
}

// GetGoneTimestampOk returns a tuple with the GoneTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alarm) GetGoneTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.GoneTimestamp.Get(), o.GoneTimestamp.IsSet()
}

// HasGoneTimestamp returns a boolean if a field has been set.
func (o *Alarm) HasGoneTimestamp() bool {
	if o != nil && o.GoneTimestamp.IsSet() {
		return true
	}

	return false
}

// SetGoneTimestamp gets a reference to the given NullableTime and assigns it to the GoneTimestamp field.
func (o *Alarm) SetGoneTimestamp(v time.Time) {
	o.GoneTimestamp.Set(&v)
}
// SetGoneTimestampNil sets the value for GoneTimestamp to be an explicit nil
func (o *Alarm) SetGoneTimestampNil() {
	o.GoneTimestamp.Set(nil)
}

// UnsetGoneTimestamp ensures that no value is present for GoneTimestamp, not even an explicit nil
func (o *Alarm) UnsetGoneTimestamp() {
	o.GoneTimestamp.Unset()
}

// GetAcknowledgeTimestamp returns the AcknowledgeTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alarm) GetAcknowledgeTimestamp() time.Time {
	if o == nil || o.AcknowledgeTimestamp.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AcknowledgeTimestamp.Get()
}

// GetAcknowledgeTimestampOk returns a tuple with the AcknowledgeTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alarm) GetAcknowledgeTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcknowledgeTimestamp.Get(), o.AcknowledgeTimestamp.IsSet()
}

// HasAcknowledgeTimestamp returns a boolean if a field has been set.
func (o *Alarm) HasAcknowledgeTimestamp() bool {
	if o != nil && o.AcknowledgeTimestamp.IsSet() {
		return true
	}

	return false
}

// SetAcknowledgeTimestamp gets a reference to the given NullableTime and assigns it to the AcknowledgeTimestamp field.
func (o *Alarm) SetAcknowledgeTimestamp(v time.Time) {
	o.AcknowledgeTimestamp.Set(&v)
}
// SetAcknowledgeTimestampNil sets the value for AcknowledgeTimestamp to be an explicit nil
func (o *Alarm) SetAcknowledgeTimestampNil() {
	o.AcknowledgeTimestamp.Set(nil)
}

// UnsetAcknowledgeTimestamp ensures that no value is present for AcknowledgeTimestamp, not even an explicit nil
func (o *Alarm) UnsetAcknowledgeTimestamp() {
	o.AcknowledgeTimestamp.Unset()
}

// GetOccurrences returns the Occurrences field value
func (o *Alarm) GetOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value
// and a boolean to check if the value has been set.
func (o *Alarm) GetOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occurrences, true
}

// SetOccurrences sets field value
func (o *Alarm) SetOccurrences(v int32) {
	o.Occurrences = v
}

// GetAcknowledgeText returns the AcknowledgeText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alarm) GetAcknowledgeText() string {
	if o == nil || o.AcknowledgeText.Get() == nil {
		var ret string
		return ret
	}
	return *o.AcknowledgeText.Get()
}

// GetAcknowledgeTextOk returns a tuple with the AcknowledgeText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alarm) GetAcknowledgeTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcknowledgeText.Get(), o.AcknowledgeText.IsSet()
}

// HasAcknowledgeText returns a boolean if a field has been set.
func (o *Alarm) HasAcknowledgeText() bool {
	if o != nil && o.AcknowledgeText.IsSet() {
		return true
	}

	return false
}

// SetAcknowledgeText gets a reference to the given NullableString and assigns it to the AcknowledgeText field.
func (o *Alarm) SetAcknowledgeText(v string) {
	o.AcknowledgeText.Set(&v)
}
// SetAcknowledgeTextNil sets the value for AcknowledgeText to be an explicit nil
func (o *Alarm) SetAcknowledgeTextNil() {
	o.AcknowledgeText.Set(nil)
}

// UnsetAcknowledgeText ensures that no value is present for AcknowledgeText, not even an explicit nil
func (o *Alarm) UnsetAcknowledgeText() {
	o.AcknowledgeText.Unset()
}

// GetAcknowledgeUserId returns the AcknowledgeUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alarm) GetAcknowledgeUserId() string {
	if o == nil || o.AcknowledgeUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AcknowledgeUserId.Get()
}

// GetAcknowledgeUserIdOk returns a tuple with the AcknowledgeUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alarm) GetAcknowledgeUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcknowledgeUserId.Get(), o.AcknowledgeUserId.IsSet()
}

// HasAcknowledgeUserId returns a boolean if a field has been set.
func (o *Alarm) HasAcknowledgeUserId() bool {
	if o != nil && o.AcknowledgeUserId.IsSet() {
		return true
	}

	return false
}

// SetAcknowledgeUserId gets a reference to the given NullableString and assigns it to the AcknowledgeUserId field.
func (o *Alarm) SetAcknowledgeUserId(v string) {
	o.AcknowledgeUserId.Set(&v)
}
// SetAcknowledgeUserIdNil sets the value for AcknowledgeUserId to be an explicit nil
func (o *Alarm) SetAcknowledgeUserIdNil() {
	o.AcknowledgeUserId.Set(nil)
}

// UnsetAcknowledgeUserId ensures that no value is present for AcknowledgeUserId, not even an explicit nil
func (o *Alarm) UnsetAcknowledgeUserId() {
	o.AcknowledgeUserId.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alarm) GetMessage() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alarm) GetMessageOk() (map[string]interface{}, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Alarm) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given map[string]interface{} and assigns it to the Message field.
func (o *Alarm) SetMessage(v map[string]interface{}) {
	o.Message = v
}

// GetAssetInfo returns the AssetInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alarm) GetAssetInfo() Asset {
	if o == nil || o.AssetInfo.Get() == nil {
		var ret Asset
		return ret
	}
	return *o.AssetInfo.Get()
}

// GetAssetInfoOk returns a tuple with the AssetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alarm) GetAssetInfoOk() (*Asset, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetInfo.Get(), o.AssetInfo.IsSet()
}

// HasAssetInfo returns a boolean if a field has been set.
func (o *Alarm) HasAssetInfo() bool {
	if o != nil && o.AssetInfo.IsSet() {
		return true
	}

	return false
}

// SetAssetInfo gets a reference to the given NullableAsset and assigns it to the AssetInfo field.
func (o *Alarm) SetAssetInfo(v Asset) {
	o.AssetInfo.Set(&v)
}
// SetAssetInfoNil sets the value for AssetInfo to be an explicit nil
func (o *Alarm) SetAssetInfoNil() {
	o.AssetInfo.Set(nil)
}

// UnsetAssetInfo ensures that no value is present for AssetInfo, not even an explicit nil
func (o *Alarm) UnsetAssetInfo() {
	o.AssetInfo.Unset()
}

// GetRuleInfo returns the RuleInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Alarm) GetRuleInfo() AlarmRule {
	if o == nil || o.RuleInfo.Get() == nil {
		var ret AlarmRule
		return ret
	}
	return *o.RuleInfo.Get()
}

// GetRuleInfoOk returns a tuple with the RuleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Alarm) GetRuleInfoOk() (*AlarmRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleInfo.Get(), o.RuleInfo.IsSet()
}

// HasRuleInfo returns a boolean if a field has been set.
func (o *Alarm) HasRuleInfo() bool {
	if o != nil && o.RuleInfo.IsSet() {
		return true
	}

	return false
}

// SetRuleInfo gets a reference to the given NullableAlarmRule and assigns it to the RuleInfo field.
func (o *Alarm) SetRuleInfo(v AlarmRule) {
	o.RuleInfo.Set(&v)
}
// SetRuleInfoNil sets the value for RuleInfo to be an explicit nil
func (o *Alarm) SetRuleInfoNil() {
	o.RuleInfo.Set(nil)
}

// UnsetRuleInfo ensures that no value is present for RuleInfo, not even an explicit nil
func (o *Alarm) UnsetRuleInfo() {
	o.RuleInfo.Unset()
}

func (o Alarm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RuleId.IsSet() {
		toSerialize["ruleId"] = o.RuleId.Get()
	}
	if true {
		toSerialize["assetId"] = o.AssetId
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Attribute.IsSet() {
		toSerialize["attribute"] = o.Attribute.Get()
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if o.RequiresAcknowledge != nil {
		toSerialize["requiresAcknowledge"] = o.RequiresAcknowledge
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.GoneTimestamp.IsSet() {
		toSerialize["goneTimestamp"] = o.GoneTimestamp.Get()
	}
	if o.AcknowledgeTimestamp.IsSet() {
		toSerialize["acknowledgeTimestamp"] = o.AcknowledgeTimestamp.Get()
	}
	if true {
		toSerialize["occurrences"] = o.Occurrences
	}
	if o.AcknowledgeText.IsSet() {
		toSerialize["acknowledgeText"] = o.AcknowledgeText.Get()
	}
	if o.AcknowledgeUserId.IsSet() {
		toSerialize["acknowledgeUserId"] = o.AcknowledgeUserId.Get()
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.AssetInfo.IsSet() {
		toSerialize["assetInfo"] = o.AssetInfo.Get()
	}
	if o.RuleInfo.IsSet() {
		toSerialize["ruleInfo"] = o.RuleInfo.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAlarm struct {
	value *Alarm
	isSet bool
}

func (v NullableAlarm) Get() *Alarm {
	return v.value
}

func (v *NullableAlarm) Set(val *Alarm) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarm(val *Alarm) *NullableAlarm {
	return &NullableAlarm{value: val, isSet: true}
}

func (v NullableAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

