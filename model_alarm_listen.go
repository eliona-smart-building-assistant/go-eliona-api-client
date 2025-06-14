/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.9.4
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the AlarmListen type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlarmListen{}

// AlarmListen struct for AlarmListen
type AlarmListen struct {
	// The id of the corresponding rule
	RuleId int32 `json:"ruleId"`
	// ID of the corresponding asset
	AssetId NullableInt32 `json:"assetId,omitempty"`
	// Type of asset data
	Subtype NullableString `json:"subtype,omitempty"`
	// Name of the attribute of the asset type
	Attribute NullableString `json:"attribute,omitempty"`
	// The priority of the alarm. The lower this value the higher the priority.
	Priority NullableInt32 `json:"priority,omitempty"`
	// Requires the alarm an acknowledgment
	RequiresAcknowledge NullableBool `json:"requiresAcknowledge,omitempty"`
	// The value which triggers the alarm
	Value NullableFloat64 `json:"value,omitempty"`
	// Timestamp of the latest data change
	Timestamp NullableTime `json:"timestamp,omitempty"`
	// Timestamp of the latest data change
	GoneTimestamp NullableTime `json:"goneTimestamp,omitempty"`
	// Timestamp of the latest data change
	AcknowledgeTimestamp NullableTime `json:"acknowledgeTimestamp,omitempty"`
	// How often this alarm is triggered
	Occurrences NullableInt32 `json:"occurrences,omitempty"`
	// Text of acknowledgement
	AcknowledgeText NullableString `json:"acknowledgeText,omitempty"`
	// User who acknowledged the alarm
	AcknowledgeUserId NullableString `json:"acknowledgeUserId,omitempty"`
	// Message.yaml texts for alarm
	Message             map[string]interface{} `json:"message"`
	AcknowledgeUserInfo NullableUser           `json:"acknowledgeUserInfo,omitempty"`
	AssetInfo           NullableAsset          `json:"assetInfo,omitempty"`
	RuleInfo            NullableAlarmRule      `json:"ruleInfo,omitempty"`
	// The status code expecting when actually perform the operation. Some values are - 200: updated (ok)  - 201: created - 204: deleted (no content) - 304: unchanged (not modified) - 400: problem (bad request) - 404: not found - 409: duplicated (conflict) - 422: unprocessable
	StatusCode *int32 `json:"statusCode,omitempty"`
}

type _AlarmListen AlarmListen

// NewAlarmListen instantiates a new AlarmListen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmListen(ruleId int32, message map[string]interface{}) *AlarmListen {
	this := AlarmListen{}
	this.RuleId = ruleId
	var requiresAcknowledge bool = false
	this.RequiresAcknowledge = *NewNullableBool(&requiresAcknowledge)
	this.Message = message
	return &this
}

// NewAlarmListenWithDefaults instantiates a new AlarmListen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmListenWithDefaults() *AlarmListen {
	this := AlarmListen{}
	var requiresAcknowledge bool = false
	this.RequiresAcknowledge = *NewNullableBool(&requiresAcknowledge)
	return &this
}

// GetRuleId returns the RuleId field value
func (o *AlarmListen) GetRuleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *AlarmListen) GetRuleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *AlarmListen) SetRuleId(v int32) {
	o.RuleId = v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetAssetId() int32 {
	if o == nil || IsNil(o.AssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.AssetId.Get()
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetId.Get(), o.AssetId.IsSet()
}

// HasAssetId returns a boolean if a field has been set.
func (o *AlarmListen) HasAssetId() bool {
	if o != nil && o.AssetId.IsSet() {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given NullableInt32 and assigns it to the AssetId field.
func (o *AlarmListen) SetAssetId(v int32) {
	o.AssetId.Set(&v)
}

// SetAssetIdNil sets the value for AssetId to be an explicit nil
func (o *AlarmListen) SetAssetIdNil() {
	o.AssetId.Set(nil)
}

// UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
func (o *AlarmListen) UnsetAssetId() {
	o.AssetId.Unset()
}

// GetSubtype returns the Subtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetSubtype() string {
	if o == nil || IsNil(o.Subtype.Get()) {
		var ret string
		return ret
	}
	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// HasSubtype returns a boolean if a field has been set.
func (o *AlarmListen) HasSubtype() bool {
	if o != nil && o.Subtype.IsSet() {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given NullableString and assigns it to the Subtype field.
func (o *AlarmListen) SetSubtype(v string) {
	o.Subtype.Set(&v)
}

// SetSubtypeNil sets the value for Subtype to be an explicit nil
func (o *AlarmListen) SetSubtypeNil() {
	o.Subtype.Set(nil)
}

// UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
func (o *AlarmListen) UnsetSubtype() {
	o.Subtype.Unset()
}

// GetAttribute returns the Attribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetAttribute() string {
	if o == nil || IsNil(o.Attribute.Get()) {
		var ret string
		return ret
	}
	return *o.Attribute.Get()
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attribute.Get(), o.Attribute.IsSet()
}

// HasAttribute returns a boolean if a field has been set.
func (o *AlarmListen) HasAttribute() bool {
	if o != nil && o.Attribute.IsSet() {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given NullableString and assigns it to the Attribute field.
func (o *AlarmListen) SetAttribute(v string) {
	o.Attribute.Set(&v)
}

// SetAttributeNil sets the value for Attribute to be an explicit nil
func (o *AlarmListen) SetAttributeNil() {
	o.Attribute.Set(nil)
}

// UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
func (o *AlarmListen) UnsetAttribute() {
	o.Attribute.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *AlarmListen) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *AlarmListen) SetPriority(v int32) {
	o.Priority.Set(&v)
}

// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *AlarmListen) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *AlarmListen) UnsetPriority() {
	o.Priority.Unset()
}

// GetRequiresAcknowledge returns the RequiresAcknowledge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetRequiresAcknowledge() bool {
	if o == nil || IsNil(o.RequiresAcknowledge.Get()) {
		var ret bool
		return ret
	}
	return *o.RequiresAcknowledge.Get()
}

// GetRequiresAcknowledgeOk returns a tuple with the RequiresAcknowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetRequiresAcknowledgeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiresAcknowledge.Get(), o.RequiresAcknowledge.IsSet()
}

// HasRequiresAcknowledge returns a boolean if a field has been set.
func (o *AlarmListen) HasRequiresAcknowledge() bool {
	if o != nil && o.RequiresAcknowledge.IsSet() {
		return true
	}

	return false
}

// SetRequiresAcknowledge gets a reference to the given NullableBool and assigns it to the RequiresAcknowledge field.
func (o *AlarmListen) SetRequiresAcknowledge(v bool) {
	o.RequiresAcknowledge.Set(&v)
}

// SetRequiresAcknowledgeNil sets the value for RequiresAcknowledge to be an explicit nil
func (o *AlarmListen) SetRequiresAcknowledgeNil() {
	o.RequiresAcknowledge.Set(nil)
}

// UnsetRequiresAcknowledge ensures that no value is present for RequiresAcknowledge, not even an explicit nil
func (o *AlarmListen) UnsetRequiresAcknowledge() {
	o.RequiresAcknowledge.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetValue() float64 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret float64
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *AlarmListen) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableFloat64 and assigns it to the Value field.
func (o *AlarmListen) SetValue(v float64) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *AlarmListen) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *AlarmListen) UnsetValue() {
	o.Value.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AlarmListen) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *AlarmListen) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *AlarmListen) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *AlarmListen) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetGoneTimestamp returns the GoneTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetGoneTimestamp() time.Time {
	if o == nil || IsNil(o.GoneTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.GoneTimestamp.Get()
}

// GetGoneTimestampOk returns a tuple with the GoneTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetGoneTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.GoneTimestamp.Get(), o.GoneTimestamp.IsSet()
}

// HasGoneTimestamp returns a boolean if a field has been set.
func (o *AlarmListen) HasGoneTimestamp() bool {
	if o != nil && o.GoneTimestamp.IsSet() {
		return true
	}

	return false
}

// SetGoneTimestamp gets a reference to the given NullableTime and assigns it to the GoneTimestamp field.
func (o *AlarmListen) SetGoneTimestamp(v time.Time) {
	o.GoneTimestamp.Set(&v)
}

// SetGoneTimestampNil sets the value for GoneTimestamp to be an explicit nil
func (o *AlarmListen) SetGoneTimestampNil() {
	o.GoneTimestamp.Set(nil)
}

// UnsetGoneTimestamp ensures that no value is present for GoneTimestamp, not even an explicit nil
func (o *AlarmListen) UnsetGoneTimestamp() {
	o.GoneTimestamp.Unset()
}

// GetAcknowledgeTimestamp returns the AcknowledgeTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetAcknowledgeTimestamp() time.Time {
	if o == nil || IsNil(o.AcknowledgeTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AcknowledgeTimestamp.Get()
}

// GetAcknowledgeTimestampOk returns a tuple with the AcknowledgeTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetAcknowledgeTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcknowledgeTimestamp.Get(), o.AcknowledgeTimestamp.IsSet()
}

// HasAcknowledgeTimestamp returns a boolean if a field has been set.
func (o *AlarmListen) HasAcknowledgeTimestamp() bool {
	if o != nil && o.AcknowledgeTimestamp.IsSet() {
		return true
	}

	return false
}

// SetAcknowledgeTimestamp gets a reference to the given NullableTime and assigns it to the AcknowledgeTimestamp field.
func (o *AlarmListen) SetAcknowledgeTimestamp(v time.Time) {
	o.AcknowledgeTimestamp.Set(&v)
}

// SetAcknowledgeTimestampNil sets the value for AcknowledgeTimestamp to be an explicit nil
func (o *AlarmListen) SetAcknowledgeTimestampNil() {
	o.AcknowledgeTimestamp.Set(nil)
}

// UnsetAcknowledgeTimestamp ensures that no value is present for AcknowledgeTimestamp, not even an explicit nil
func (o *AlarmListen) UnsetAcknowledgeTimestamp() {
	o.AcknowledgeTimestamp.Unset()
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetOccurrences() int32 {
	if o == nil || IsNil(o.Occurrences.Get()) {
		var ret int32
		return ret
	}
	return *o.Occurrences.Get()
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Occurrences.Get(), o.Occurrences.IsSet()
}

// HasOccurrences returns a boolean if a field has been set.
func (o *AlarmListen) HasOccurrences() bool {
	if o != nil && o.Occurrences.IsSet() {
		return true
	}

	return false
}

// SetOccurrences gets a reference to the given NullableInt32 and assigns it to the Occurrences field.
func (o *AlarmListen) SetOccurrences(v int32) {
	o.Occurrences.Set(&v)
}

// SetOccurrencesNil sets the value for Occurrences to be an explicit nil
func (o *AlarmListen) SetOccurrencesNil() {
	o.Occurrences.Set(nil)
}

// UnsetOccurrences ensures that no value is present for Occurrences, not even an explicit nil
func (o *AlarmListen) UnsetOccurrences() {
	o.Occurrences.Unset()
}

// GetAcknowledgeText returns the AcknowledgeText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetAcknowledgeText() string {
	if o == nil || IsNil(o.AcknowledgeText.Get()) {
		var ret string
		return ret
	}
	return *o.AcknowledgeText.Get()
}

// GetAcknowledgeTextOk returns a tuple with the AcknowledgeText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetAcknowledgeTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcknowledgeText.Get(), o.AcknowledgeText.IsSet()
}

// HasAcknowledgeText returns a boolean if a field has been set.
func (o *AlarmListen) HasAcknowledgeText() bool {
	if o != nil && o.AcknowledgeText.IsSet() {
		return true
	}

	return false
}

// SetAcknowledgeText gets a reference to the given NullableString and assigns it to the AcknowledgeText field.
func (o *AlarmListen) SetAcknowledgeText(v string) {
	o.AcknowledgeText.Set(&v)
}

// SetAcknowledgeTextNil sets the value for AcknowledgeText to be an explicit nil
func (o *AlarmListen) SetAcknowledgeTextNil() {
	o.AcknowledgeText.Set(nil)
}

// UnsetAcknowledgeText ensures that no value is present for AcknowledgeText, not even an explicit nil
func (o *AlarmListen) UnsetAcknowledgeText() {
	o.AcknowledgeText.Unset()
}

// GetAcknowledgeUserId returns the AcknowledgeUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetAcknowledgeUserId() string {
	if o == nil || IsNil(o.AcknowledgeUserId.Get()) {
		var ret string
		return ret
	}
	return *o.AcknowledgeUserId.Get()
}

// GetAcknowledgeUserIdOk returns a tuple with the AcknowledgeUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetAcknowledgeUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcknowledgeUserId.Get(), o.AcknowledgeUserId.IsSet()
}

// HasAcknowledgeUserId returns a boolean if a field has been set.
func (o *AlarmListen) HasAcknowledgeUserId() bool {
	if o != nil && o.AcknowledgeUserId.IsSet() {
		return true
	}

	return false
}

// SetAcknowledgeUserId gets a reference to the given NullableString and assigns it to the AcknowledgeUserId field.
func (o *AlarmListen) SetAcknowledgeUserId(v string) {
	o.AcknowledgeUserId.Set(&v)
}

// SetAcknowledgeUserIdNil sets the value for AcknowledgeUserId to be an explicit nil
func (o *AlarmListen) SetAcknowledgeUserIdNil() {
	o.AcknowledgeUserId.Set(nil)
}

// UnsetAcknowledgeUserId ensures that no value is present for AcknowledgeUserId, not even an explicit nil
func (o *AlarmListen) UnsetAcknowledgeUserId() {
	o.AcknowledgeUserId.Unset()
}

// GetMessage returns the Message field value
func (o *AlarmListen) GetMessage() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlarmListen) GetMessageOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *AlarmListen) SetMessage(v map[string]interface{}) {
	o.Message = v
}

// GetAcknowledgeUserInfo returns the AcknowledgeUserInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetAcknowledgeUserInfo() User {
	if o == nil || IsNil(o.AcknowledgeUserInfo.Get()) {
		var ret User
		return ret
	}
	return *o.AcknowledgeUserInfo.Get()
}

// GetAcknowledgeUserInfoOk returns a tuple with the AcknowledgeUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetAcknowledgeUserInfoOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcknowledgeUserInfo.Get(), o.AcknowledgeUserInfo.IsSet()
}

// HasAcknowledgeUserInfo returns a boolean if a field has been set.
func (o *AlarmListen) HasAcknowledgeUserInfo() bool {
	if o != nil && o.AcknowledgeUserInfo.IsSet() {
		return true
	}

	return false
}

// SetAcknowledgeUserInfo gets a reference to the given NullableUser and assigns it to the AcknowledgeUserInfo field.
func (o *AlarmListen) SetAcknowledgeUserInfo(v User) {
	o.AcknowledgeUserInfo.Set(&v)
}

// SetAcknowledgeUserInfoNil sets the value for AcknowledgeUserInfo to be an explicit nil
func (o *AlarmListen) SetAcknowledgeUserInfoNil() {
	o.AcknowledgeUserInfo.Set(nil)
}

// UnsetAcknowledgeUserInfo ensures that no value is present for AcknowledgeUserInfo, not even an explicit nil
func (o *AlarmListen) UnsetAcknowledgeUserInfo() {
	o.AcknowledgeUserInfo.Unset()
}

// GetAssetInfo returns the AssetInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetAssetInfo() Asset {
	if o == nil || IsNil(o.AssetInfo.Get()) {
		var ret Asset
		return ret
	}
	return *o.AssetInfo.Get()
}

// GetAssetInfoOk returns a tuple with the AssetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetAssetInfoOk() (*Asset, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetInfo.Get(), o.AssetInfo.IsSet()
}

// HasAssetInfo returns a boolean if a field has been set.
func (o *AlarmListen) HasAssetInfo() bool {
	if o != nil && o.AssetInfo.IsSet() {
		return true
	}

	return false
}

// SetAssetInfo gets a reference to the given NullableAsset and assigns it to the AssetInfo field.
func (o *AlarmListen) SetAssetInfo(v Asset) {
	o.AssetInfo.Set(&v)
}

// SetAssetInfoNil sets the value for AssetInfo to be an explicit nil
func (o *AlarmListen) SetAssetInfoNil() {
	o.AssetInfo.Set(nil)
}

// UnsetAssetInfo ensures that no value is present for AssetInfo, not even an explicit nil
func (o *AlarmListen) UnsetAssetInfo() {
	o.AssetInfo.Unset()
}

// GetRuleInfo returns the RuleInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetRuleInfo() AlarmRule {
	if o == nil || IsNil(o.RuleInfo.Get()) {
		var ret AlarmRule
		return ret
	}
	return *o.RuleInfo.Get()
}

// GetRuleInfoOk returns a tuple with the RuleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetRuleInfoOk() (*AlarmRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleInfo.Get(), o.RuleInfo.IsSet()
}

// HasRuleInfo returns a boolean if a field has been set.
func (o *AlarmListen) HasRuleInfo() bool {
	if o != nil && o.RuleInfo.IsSet() {
		return true
	}

	return false
}

// SetRuleInfo gets a reference to the given NullableAlarmRule and assigns it to the RuleInfo field.
func (o *AlarmListen) SetRuleInfo(v AlarmRule) {
	o.RuleInfo.Set(&v)
}

// SetRuleInfoNil sets the value for RuleInfo to be an explicit nil
func (o *AlarmListen) SetRuleInfoNil() {
	o.RuleInfo.Set(nil)
}

// UnsetRuleInfo ensures that no value is present for RuleInfo, not even an explicit nil
func (o *AlarmListen) UnsetRuleInfo() {
	o.RuleInfo.Unset()
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *AlarmListen) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmListen) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *AlarmListen) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *AlarmListen) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o AlarmListen) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlarmListen) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ruleId"] = o.RuleId
	if o.AssetId.IsSet() {
		toSerialize["assetId"] = o.AssetId.Get()
	}
	if o.Subtype.IsSet() {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	if o.Attribute.IsSet() {
		toSerialize["attribute"] = o.Attribute.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.RequiresAcknowledge.IsSet() {
		toSerialize["requiresAcknowledge"] = o.RequiresAcknowledge.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.GoneTimestamp.IsSet() {
		toSerialize["goneTimestamp"] = o.GoneTimestamp.Get()
	}
	if o.AcknowledgeTimestamp.IsSet() {
		toSerialize["acknowledgeTimestamp"] = o.AcknowledgeTimestamp.Get()
	}
	if o.Occurrences.IsSet() {
		toSerialize["occurrences"] = o.Occurrences.Get()
	}
	if o.AcknowledgeText.IsSet() {
		toSerialize["acknowledgeText"] = o.AcknowledgeText.Get()
	}
	if o.AcknowledgeUserId.IsSet() {
		toSerialize["acknowledgeUserId"] = o.AcknowledgeUserId.Get()
	}
	toSerialize["message"] = o.Message
	if o.AcknowledgeUserInfo.IsSet() {
		toSerialize["acknowledgeUserInfo"] = o.AcknowledgeUserInfo.Get()
	}
	if o.AssetInfo.IsSet() {
		toSerialize["assetInfo"] = o.AssetInfo.Get()
	}
	if o.RuleInfo.IsSet() {
		toSerialize["ruleInfo"] = o.RuleInfo.Get()
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	return toSerialize, nil
}

func (o *AlarmListen) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ruleId",
		"message",
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

	varAlarmListen := _AlarmListen{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlarmListen)

	if err != nil {
		return err
	}

	*o = AlarmListen(varAlarmListen)

	return err
}

type NullableAlarmListen struct {
	value *AlarmListen
	isSet bool
}

func (v NullableAlarmListen) Get() *AlarmListen {
	return v.value
}

func (v *NullableAlarmListen) Set(val *AlarmListen) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmListen) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmListen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmListen(val *AlarmListen) *NullableAlarmListen {
	return &NullableAlarmListen{value: val, isSet: true}
}

func (v NullableAlarmListen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmListen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
