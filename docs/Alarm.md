# Alarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | **int32** | The id of the corresponding rule | [readonly] 
**AssetId** | Pointer to **NullableInt32** | ID of the corresponding asset | [optional] [readonly] 
**Subtype** | Pointer to **NullableString** | Type of asset data | [optional] [readonly] 
**Attribute** | Pointer to **NullableString** | Name of the attribute of the asset type | [optional] [readonly] 
**Priority** | Pointer to **NullableInt32** | The priority of the alarm. The lower this value the higher the priority. | [optional] 
**RequiresAcknowledge** | Pointer to **NullableBool** | Requires the alarm an acknowledgment | [optional] [default to false]
**Value** | Pointer to **NullableFloat64** | The value which triggers the alarm | [optional] 
**Timestamp** | Pointer to **NullableTime** | Timestamp of the latest data change | [optional] 
**GoneTimestamp** | Pointer to **NullableTime** | Timestamp of the latest data change | [optional] [readonly] 
**AcknowledgeTimestamp** | Pointer to **NullableTime** | Timestamp of the latest data change | [optional] [readonly] 
**Occurrences** | Pointer to **NullableInt32** | How often this alarm is triggered | [optional] [readonly] 
**AcknowledgeText** | Pointer to **NullableString** | Text of acknowledgement | [optional] [readonly] 
**AcknowledgeUserId** | Pointer to **NullableString** | User who acknowledged the alarm | [optional] [readonly] 
**Message** | **map[string]interface{}** | Message.yaml texts for alarm | 
**AcknowledgeUserInfo** | Pointer to [**NullableUser**](User.md) |  | [optional] 
**AssetInfo** | Pointer to [**NullableAsset**](Asset.md) |  | [optional] 
**RuleInfo** | Pointer to [**NullableAlarmRule**](AlarmRule.md) |  | [optional] 

## Methods

### NewAlarm

`func NewAlarm(ruleId int32, message map[string]interface{}, ) *Alarm`

NewAlarm instantiates a new Alarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmWithDefaults

`func NewAlarmWithDefaults() *Alarm`

NewAlarmWithDefaults instantiates a new Alarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *Alarm) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *Alarm) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *Alarm) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.


### GetAssetId

`func (o *Alarm) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Alarm) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Alarm) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Alarm) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *Alarm) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *Alarm) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
### GetSubtype

`func (o *Alarm) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Alarm) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Alarm) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *Alarm) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *Alarm) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *Alarm) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetAttribute

`func (o *Alarm) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Alarm) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Alarm) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *Alarm) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### SetAttributeNil

`func (o *Alarm) SetAttributeNil(b bool)`

 SetAttributeNil sets the value for Attribute to be an explicit nil

### UnsetAttribute
`func (o *Alarm) UnsetAttribute()`

UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
### GetPriority

`func (o *Alarm) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Alarm) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Alarm) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Alarm) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *Alarm) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *Alarm) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetRequiresAcknowledge

`func (o *Alarm) GetRequiresAcknowledge() bool`

GetRequiresAcknowledge returns the RequiresAcknowledge field if non-nil, zero value otherwise.

### GetRequiresAcknowledgeOk

`func (o *Alarm) GetRequiresAcknowledgeOk() (*bool, bool)`

GetRequiresAcknowledgeOk returns a tuple with the RequiresAcknowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAcknowledge

`func (o *Alarm) SetRequiresAcknowledge(v bool)`

SetRequiresAcknowledge sets RequiresAcknowledge field to given value.

### HasRequiresAcknowledge

`func (o *Alarm) HasRequiresAcknowledge() bool`

HasRequiresAcknowledge returns a boolean if a field has been set.

### SetRequiresAcknowledgeNil

`func (o *Alarm) SetRequiresAcknowledgeNil(b bool)`

 SetRequiresAcknowledgeNil sets the value for RequiresAcknowledge to be an explicit nil

### UnsetRequiresAcknowledge
`func (o *Alarm) UnsetRequiresAcknowledge()`

UnsetRequiresAcknowledge ensures that no value is present for RequiresAcknowledge, not even an explicit nil
### GetValue

`func (o *Alarm) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Alarm) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Alarm) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *Alarm) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Alarm) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Alarm) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTimestamp

`func (o *Alarm) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Alarm) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Alarm) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Alarm) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *Alarm) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Alarm) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetGoneTimestamp

`func (o *Alarm) GetGoneTimestamp() time.Time`

GetGoneTimestamp returns the GoneTimestamp field if non-nil, zero value otherwise.

### GetGoneTimestampOk

`func (o *Alarm) GetGoneTimestampOk() (*time.Time, bool)`

GetGoneTimestampOk returns a tuple with the GoneTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoneTimestamp

`func (o *Alarm) SetGoneTimestamp(v time.Time)`

SetGoneTimestamp sets GoneTimestamp field to given value.

### HasGoneTimestamp

`func (o *Alarm) HasGoneTimestamp() bool`

HasGoneTimestamp returns a boolean if a field has been set.

### SetGoneTimestampNil

`func (o *Alarm) SetGoneTimestampNil(b bool)`

 SetGoneTimestampNil sets the value for GoneTimestamp to be an explicit nil

### UnsetGoneTimestamp
`func (o *Alarm) UnsetGoneTimestamp()`

UnsetGoneTimestamp ensures that no value is present for GoneTimestamp, not even an explicit nil
### GetAcknowledgeTimestamp

`func (o *Alarm) GetAcknowledgeTimestamp() time.Time`

GetAcknowledgeTimestamp returns the AcknowledgeTimestamp field if non-nil, zero value otherwise.

### GetAcknowledgeTimestampOk

`func (o *Alarm) GetAcknowledgeTimestampOk() (*time.Time, bool)`

GetAcknowledgeTimestampOk returns a tuple with the AcknowledgeTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeTimestamp

`func (o *Alarm) SetAcknowledgeTimestamp(v time.Time)`

SetAcknowledgeTimestamp sets AcknowledgeTimestamp field to given value.

### HasAcknowledgeTimestamp

`func (o *Alarm) HasAcknowledgeTimestamp() bool`

HasAcknowledgeTimestamp returns a boolean if a field has been set.

### SetAcknowledgeTimestampNil

`func (o *Alarm) SetAcknowledgeTimestampNil(b bool)`

 SetAcknowledgeTimestampNil sets the value for AcknowledgeTimestamp to be an explicit nil

### UnsetAcknowledgeTimestamp
`func (o *Alarm) UnsetAcknowledgeTimestamp()`

UnsetAcknowledgeTimestamp ensures that no value is present for AcknowledgeTimestamp, not even an explicit nil
### GetOccurrences

`func (o *Alarm) GetOccurrences() int32`

GetOccurrences returns the Occurrences field if non-nil, zero value otherwise.

### GetOccurrencesOk

`func (o *Alarm) GetOccurrencesOk() (*int32, bool)`

GetOccurrencesOk returns a tuple with the Occurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrences

`func (o *Alarm) SetOccurrences(v int32)`

SetOccurrences sets Occurrences field to given value.

### HasOccurrences

`func (o *Alarm) HasOccurrences() bool`

HasOccurrences returns a boolean if a field has been set.

### SetOccurrencesNil

`func (o *Alarm) SetOccurrencesNil(b bool)`

 SetOccurrencesNil sets the value for Occurrences to be an explicit nil

### UnsetOccurrences
`func (o *Alarm) UnsetOccurrences()`

UnsetOccurrences ensures that no value is present for Occurrences, not even an explicit nil
### GetAcknowledgeText

`func (o *Alarm) GetAcknowledgeText() string`

GetAcknowledgeText returns the AcknowledgeText field if non-nil, zero value otherwise.

### GetAcknowledgeTextOk

`func (o *Alarm) GetAcknowledgeTextOk() (*string, bool)`

GetAcknowledgeTextOk returns a tuple with the AcknowledgeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeText

`func (o *Alarm) SetAcknowledgeText(v string)`

SetAcknowledgeText sets AcknowledgeText field to given value.

### HasAcknowledgeText

`func (o *Alarm) HasAcknowledgeText() bool`

HasAcknowledgeText returns a boolean if a field has been set.

### SetAcknowledgeTextNil

`func (o *Alarm) SetAcknowledgeTextNil(b bool)`

 SetAcknowledgeTextNil sets the value for AcknowledgeText to be an explicit nil

### UnsetAcknowledgeText
`func (o *Alarm) UnsetAcknowledgeText()`

UnsetAcknowledgeText ensures that no value is present for AcknowledgeText, not even an explicit nil
### GetAcknowledgeUserId

`func (o *Alarm) GetAcknowledgeUserId() string`

GetAcknowledgeUserId returns the AcknowledgeUserId field if non-nil, zero value otherwise.

### GetAcknowledgeUserIdOk

`func (o *Alarm) GetAcknowledgeUserIdOk() (*string, bool)`

GetAcknowledgeUserIdOk returns a tuple with the AcknowledgeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeUserId

`func (o *Alarm) SetAcknowledgeUserId(v string)`

SetAcknowledgeUserId sets AcknowledgeUserId field to given value.

### HasAcknowledgeUserId

`func (o *Alarm) HasAcknowledgeUserId() bool`

HasAcknowledgeUserId returns a boolean if a field has been set.

### SetAcknowledgeUserIdNil

`func (o *Alarm) SetAcknowledgeUserIdNil(b bool)`

 SetAcknowledgeUserIdNil sets the value for AcknowledgeUserId to be an explicit nil

### UnsetAcknowledgeUserId
`func (o *Alarm) UnsetAcknowledgeUserId()`

UnsetAcknowledgeUserId ensures that no value is present for AcknowledgeUserId, not even an explicit nil
### GetMessage

`func (o *Alarm) GetMessage() map[string]interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Alarm) GetMessageOk() (*map[string]interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Alarm) SetMessage(v map[string]interface{})`

SetMessage sets Message field to given value.


### GetAcknowledgeUserInfo

`func (o *Alarm) GetAcknowledgeUserInfo() User`

GetAcknowledgeUserInfo returns the AcknowledgeUserInfo field if non-nil, zero value otherwise.

### GetAcknowledgeUserInfoOk

`func (o *Alarm) GetAcknowledgeUserInfoOk() (*User, bool)`

GetAcknowledgeUserInfoOk returns a tuple with the AcknowledgeUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeUserInfo

`func (o *Alarm) SetAcknowledgeUserInfo(v User)`

SetAcknowledgeUserInfo sets AcknowledgeUserInfo field to given value.

### HasAcknowledgeUserInfo

`func (o *Alarm) HasAcknowledgeUserInfo() bool`

HasAcknowledgeUserInfo returns a boolean if a field has been set.

### SetAcknowledgeUserInfoNil

`func (o *Alarm) SetAcknowledgeUserInfoNil(b bool)`

 SetAcknowledgeUserInfoNil sets the value for AcknowledgeUserInfo to be an explicit nil

### UnsetAcknowledgeUserInfo
`func (o *Alarm) UnsetAcknowledgeUserInfo()`

UnsetAcknowledgeUserInfo ensures that no value is present for AcknowledgeUserInfo, not even an explicit nil
### GetAssetInfo

`func (o *Alarm) GetAssetInfo() Asset`

GetAssetInfo returns the AssetInfo field if non-nil, zero value otherwise.

### GetAssetInfoOk

`func (o *Alarm) GetAssetInfoOk() (*Asset, bool)`

GetAssetInfoOk returns a tuple with the AssetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetInfo

`func (o *Alarm) SetAssetInfo(v Asset)`

SetAssetInfo sets AssetInfo field to given value.

### HasAssetInfo

`func (o *Alarm) HasAssetInfo() bool`

HasAssetInfo returns a boolean if a field has been set.

### SetAssetInfoNil

`func (o *Alarm) SetAssetInfoNil(b bool)`

 SetAssetInfoNil sets the value for AssetInfo to be an explicit nil

### UnsetAssetInfo
`func (o *Alarm) UnsetAssetInfo()`

UnsetAssetInfo ensures that no value is present for AssetInfo, not even an explicit nil
### GetRuleInfo

`func (o *Alarm) GetRuleInfo() AlarmRule`

GetRuleInfo returns the RuleInfo field if non-nil, zero value otherwise.

### GetRuleInfoOk

`func (o *Alarm) GetRuleInfoOk() (*AlarmRule, bool)`

GetRuleInfoOk returns a tuple with the RuleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleInfo

`func (o *Alarm) SetRuleInfo(v AlarmRule)`

SetRuleInfo sets RuleInfo field to given value.

### HasRuleInfo

`func (o *Alarm) HasRuleInfo() bool`

HasRuleInfo returns a boolean if a field has been set.

### SetRuleInfoNil

`func (o *Alarm) SetRuleInfoNil(b bool)`

 SetRuleInfoNil sets the value for RuleInfo to be an explicit nil

### UnsetRuleInfo
`func (o *Alarm) UnsetRuleInfo()`

UnsetRuleInfo ensures that no value is present for RuleInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


