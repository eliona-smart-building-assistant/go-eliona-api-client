/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.6
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
	RuleId NullableInt32 `json:"ruleId,omitempty"`
	// ID of the corresponding asset
	AssetId int32       `json:"assetId"`
	Subtype DataSubtype `json:"subtype"`
	// Name of the attribute of the asset type
	Attribute NullableString `json:"attribute,omitempty"`
	Priority  AlarmPriority  `json:"priority"`
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
	// Message.yaml texts for alarm
	Message   map[string]interface{} `json:"message,omitempty"`
	AssetInfo NullableAsset          `json:"assetInfo,omitempty"`
	RuleInfo  NullableAlarmRule      `json:"ruleInfo,omitempty"`
	// The status code expecting when actually perform the operation. Some values are - 200: updated (ok)  - 201: created - 204: deleted (no content) - 304: unchanged (not modified) - 400: problem (bad request) - 404: not found - 409: duplicated (conflict) - 422: unprocessable
	StatusCode *int32 `json:"statusCode,omitempty"`
}

type _AlarmListen AlarmListen

// NewAlarmListen instantiates a new AlarmListen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlarmListen(assetId int32, subtype DataSubtype, priority AlarmPriority, timestamp time.Time, occurrences int32) *AlarmListen {
	this := AlarmListen{}
	this.AssetId = assetId
	this.Subtype = subtype
	this.Priority = priority
	var requiresAcknowledge bool = false
	this.RequiresAcknowledge = &requiresAcknowledge
	this.Timestamp = timestamp
	this.Occurrences = occurrences
	return &this
}

// NewAlarmListenWithDefaults instantiates a new AlarmListen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlarmListenWithDefaults() *AlarmListen {
	this := AlarmListen{}
	var subtype DataSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	var requiresAcknowledge bool = false
	this.RequiresAcknowledge = &requiresAcknowledge
	return &this
}

// GetRuleId returns the RuleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetRuleId() int32 {
	if o == nil || IsNil(o.RuleId.Get()) {
		var ret int32
		return ret
	}
	return *o.RuleId.Get()
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetRuleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleId.Get(), o.RuleId.IsSet()
}

// HasRuleId returns a boolean if a field has been set.
func (o *AlarmListen) HasRuleId() bool {
	if o != nil && o.RuleId.IsSet() {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given NullableInt32 and assigns it to the RuleId field.
func (o *AlarmListen) SetRuleId(v int32) {
	o.RuleId.Set(&v)
}

// SetRuleIdNil sets the value for RuleId to be an explicit nil
func (o *AlarmListen) SetRuleIdNil() {
	o.RuleId.Set(nil)
}

// UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
func (o *AlarmListen) UnsetRuleId() {
	o.RuleId.Unset()
}

// GetAssetId returns the AssetId field value
func (o *AlarmListen) GetAssetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *AlarmListen) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *AlarmListen) SetAssetId(v int32) {
	o.AssetId = v
}

// GetSubtype returns the Subtype field value
func (o *AlarmListen) GetSubtype() DataSubtype {
	if o == nil {
		var ret DataSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *AlarmListen) GetSubtypeOk() (*DataSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *AlarmListen) SetSubtype(v DataSubtype) {
	o.Subtype = v
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

// GetPriority returns the Priority field value
func (o *AlarmListen) GetPriority() AlarmPriority {
	if o == nil {
		var ret AlarmPriority
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *AlarmListen) GetPriorityOk() (*AlarmPriority, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *AlarmListen) SetPriority(v AlarmPriority) {
	o.Priority = v
}

// GetRequiresAcknowledge returns the RequiresAcknowledge field value if set, zero value otherwise.
func (o *AlarmListen) GetRequiresAcknowledge() bool {
	if o == nil || IsNil(o.RequiresAcknowledge) {
		var ret bool
		return ret
	}
	return *o.RequiresAcknowledge
}

// GetRequiresAcknowledgeOk returns a tuple with the RequiresAcknowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlarmListen) GetRequiresAcknowledgeOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiresAcknowledge) {
		return nil, false
	}
	return o.RequiresAcknowledge, true
}

// HasRequiresAcknowledge returns a boolean if a field has been set.
func (o *AlarmListen) HasRequiresAcknowledge() bool {
	if o != nil && !IsNil(o.RequiresAcknowledge) {
		return true
	}

	return false
}

// SetRequiresAcknowledge gets a reference to the given bool and assigns it to the RequiresAcknowledge field.
func (o *AlarmListen) SetRequiresAcknowledge(v bool) {
	o.RequiresAcknowledge = &v
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

// GetTimestamp returns the Timestamp field value
func (o *AlarmListen) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *AlarmListen) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *AlarmListen) SetTimestamp(v time.Time) {
	o.Timestamp = v
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

// GetOccurrences returns the Occurrences field value
func (o *AlarmListen) GetOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value
// and a boolean to check if the value has been set.
func (o *AlarmListen) GetOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occurrences, true
}

// SetOccurrences sets field value
func (o *AlarmListen) SetOccurrences(v int32) {
	o.Occurrences = v
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

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlarmListen) GetMessage() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlarmListen) GetMessageOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Message) {
		return map[string]interface{}{}, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AlarmListen) HasMessage() bool {
	if o != nil && IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given map[string]interface{} and assigns it to the Message field.
func (o *AlarmListen) SetMessage(v map[string]interface{}) {
	o.Message = v
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
	if o.RuleId.IsSet() {
		toSerialize["ruleId"] = o.RuleId.Get()
	}
	toSerialize["assetId"] = o.AssetId
	toSerialize["subtype"] = o.Subtype
	if o.Attribute.IsSet() {
		toSerialize["attribute"] = o.Attribute.Get()
	}
	toSerialize["priority"] = o.Priority
	if !IsNil(o.RequiresAcknowledge) {
		toSerialize["requiresAcknowledge"] = o.RequiresAcknowledge
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	toSerialize["timestamp"] = o.Timestamp
	if o.GoneTimestamp.IsSet() {
		toSerialize["goneTimestamp"] = o.GoneTimestamp.Get()
	}
	if o.AcknowledgeTimestamp.IsSet() {
		toSerialize["acknowledgeTimestamp"] = o.AcknowledgeTimestamp.Get()
	}
	toSerialize["occurrences"] = o.Occurrences
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
		"assetId",
		"subtype",
		"priority",
		"timestamp",
		"occurrences",
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
