# AlarmRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | The id of the rule | [optional] [readonly] 
**AssetId** | **int32** | ID of the corresponding asset | 
**Subtype** | [**DataSubtype**](DataSubtype.md) |  | [default to INPUT]
**Attribute** | **string** | Name of the attribute of the asset type | 
**Enable** | Pointer to **bool** | Rule enabled or not | [optional] [default to true]
**Priority** | [**AlarmPriority**](AlarmPriority.md) |  | 
**RequiresAcknowledge** | Pointer to **bool** | Requires the alarm an acknowledgment | [optional] [default to false]
**Equal** | Pointer to **NullableFloat64** | Triggers alarm if attribute value equals this value | [optional] 
**Low** | Pointer to **NullableFloat64** | Triggers alarm if attribute value is less than value | [optional] 
**High** | Pointer to **NullableFloat64** | Triggers alarm if attribute value is greater than value | [optional] 
**Message** | Pointer to **map[string]interface{}** | Texts for alarm | [optional] 
**Tags** | Pointer to **[]string** | List of associated tags | [optional] 
**Subject** | Pointer to **NullableString** | The subject for the alarm | [optional] 
**Urldoc** | Pointer to **NullableString** | The url describing the alarm | [optional] 
**Params** | Pointer to **map[string]interface{}** | Parameter for aggregated alarms | [optional] 
**NotifyOn** | Pointer to **NullableString** | Notification | [optional] 
**DontMask** | Pointer to **NullableBool** | Do not mask | [optional] [default to false]
**CheckType** | Pointer to **NullableString** | Check type | [optional] 
**AssetInfo** | Pointer to [**NullableAsset**](Asset.md) |  | [optional] 

## Methods

### NewAlarmRule

`func NewAlarmRule(assetId int32, subtype DataSubtype, attribute string, priority AlarmPriority, ) *AlarmRule`

NewAlarmRule instantiates a new AlarmRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmRuleWithDefaults

`func NewAlarmRuleWithDefaults() *AlarmRule`

NewAlarmRuleWithDefaults instantiates a new AlarmRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlarmRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlarmRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlarmRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlarmRule) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AlarmRule) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AlarmRule) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAssetId

`func (o *AlarmRule) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AlarmRule) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AlarmRule) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.


### GetSubtype

`func (o *AlarmRule) GetSubtype() DataSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AlarmRule) GetSubtypeOk() (*DataSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AlarmRule) SetSubtype(v DataSubtype)`

SetSubtype sets Subtype field to given value.


### GetAttribute

`func (o *AlarmRule) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AlarmRule) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AlarmRule) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetEnable

`func (o *AlarmRule) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AlarmRule) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AlarmRule) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AlarmRule) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPriority

`func (o *AlarmRule) GetPriority() AlarmPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AlarmRule) GetPriorityOk() (*AlarmPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AlarmRule) SetPriority(v AlarmPriority)`

SetPriority sets Priority field to given value.


### GetRequiresAcknowledge

`func (o *AlarmRule) GetRequiresAcknowledge() bool`

GetRequiresAcknowledge returns the RequiresAcknowledge field if non-nil, zero value otherwise.

### GetRequiresAcknowledgeOk

`func (o *AlarmRule) GetRequiresAcknowledgeOk() (*bool, bool)`

GetRequiresAcknowledgeOk returns a tuple with the RequiresAcknowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresAcknowledge

`func (o *AlarmRule) SetRequiresAcknowledge(v bool)`

SetRequiresAcknowledge sets RequiresAcknowledge field to given value.

### HasRequiresAcknowledge

`func (o *AlarmRule) HasRequiresAcknowledge() bool`

HasRequiresAcknowledge returns a boolean if a field has been set.

### GetEqual

`func (o *AlarmRule) GetEqual() float64`

GetEqual returns the Equal field if non-nil, zero value otherwise.

### GetEqualOk

`func (o *AlarmRule) GetEqualOk() (*float64, bool)`

GetEqualOk returns a tuple with the Equal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqual

`func (o *AlarmRule) SetEqual(v float64)`

SetEqual sets Equal field to given value.

### HasEqual

`func (o *AlarmRule) HasEqual() bool`

HasEqual returns a boolean if a field has been set.

### SetEqualNil

`func (o *AlarmRule) SetEqualNil(b bool)`

 SetEqualNil sets the value for Equal to be an explicit nil

### UnsetEqual
`func (o *AlarmRule) UnsetEqual()`

UnsetEqual ensures that no value is present for Equal, not even an explicit nil
### GetLow

`func (o *AlarmRule) GetLow() float64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *AlarmRule) GetLowOk() (*float64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *AlarmRule) SetLow(v float64)`

SetLow sets Low field to given value.

### HasLow

`func (o *AlarmRule) HasLow() bool`

HasLow returns a boolean if a field has been set.

### SetLowNil

`func (o *AlarmRule) SetLowNil(b bool)`

 SetLowNil sets the value for Low to be an explicit nil

### UnsetLow
`func (o *AlarmRule) UnsetLow()`

UnsetLow ensures that no value is present for Low, not even an explicit nil
### GetHigh

`func (o *AlarmRule) GetHigh() float64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *AlarmRule) GetHighOk() (*float64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *AlarmRule) SetHigh(v float64)`

SetHigh sets High field to given value.

### HasHigh

`func (o *AlarmRule) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### SetHighNil

`func (o *AlarmRule) SetHighNil(b bool)`

 SetHighNil sets the value for High to be an explicit nil

### UnsetHigh
`func (o *AlarmRule) UnsetHigh()`

UnsetHigh ensures that no value is present for High, not even an explicit nil
### GetMessage

`func (o *AlarmRule) GetMessage() map[string]interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlarmRule) GetMessageOk() (*map[string]interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlarmRule) SetMessage(v map[string]interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlarmRule) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AlarmRule) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AlarmRule) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTags

`func (o *AlarmRule) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AlarmRule) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AlarmRule) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AlarmRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AlarmRule) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AlarmRule) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetSubject

`func (o *AlarmRule) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlarmRule) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlarmRule) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AlarmRule) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *AlarmRule) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *AlarmRule) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetUrldoc

`func (o *AlarmRule) GetUrldoc() string`

GetUrldoc returns the Urldoc field if non-nil, zero value otherwise.

### GetUrldocOk

`func (o *AlarmRule) GetUrldocOk() (*string, bool)`

GetUrldocOk returns a tuple with the Urldoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrldoc

`func (o *AlarmRule) SetUrldoc(v string)`

SetUrldoc sets Urldoc field to given value.

### HasUrldoc

`func (o *AlarmRule) HasUrldoc() bool`

HasUrldoc returns a boolean if a field has been set.

### SetUrldocNil

`func (o *AlarmRule) SetUrldocNil(b bool)`

 SetUrldocNil sets the value for Urldoc to be an explicit nil

### UnsetUrldoc
`func (o *AlarmRule) UnsetUrldoc()`

UnsetUrldoc ensures that no value is present for Urldoc, not even an explicit nil
### GetParams

`func (o *AlarmRule) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *AlarmRule) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *AlarmRule) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *AlarmRule) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *AlarmRule) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *AlarmRule) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetNotifyOn

`func (o *AlarmRule) GetNotifyOn() string`

GetNotifyOn returns the NotifyOn field if non-nil, zero value otherwise.

### GetNotifyOnOk

`func (o *AlarmRule) GetNotifyOnOk() (*string, bool)`

GetNotifyOnOk returns a tuple with the NotifyOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOn

`func (o *AlarmRule) SetNotifyOn(v string)`

SetNotifyOn sets NotifyOn field to given value.

### HasNotifyOn

`func (o *AlarmRule) HasNotifyOn() bool`

HasNotifyOn returns a boolean if a field has been set.

### SetNotifyOnNil

`func (o *AlarmRule) SetNotifyOnNil(b bool)`

 SetNotifyOnNil sets the value for NotifyOn to be an explicit nil

### UnsetNotifyOn
`func (o *AlarmRule) UnsetNotifyOn()`

UnsetNotifyOn ensures that no value is present for NotifyOn, not even an explicit nil
### GetDontMask

`func (o *AlarmRule) GetDontMask() bool`

GetDontMask returns the DontMask field if non-nil, zero value otherwise.

### GetDontMaskOk

`func (o *AlarmRule) GetDontMaskOk() (*bool, bool)`

GetDontMaskOk returns a tuple with the DontMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDontMask

`func (o *AlarmRule) SetDontMask(v bool)`

SetDontMask sets DontMask field to given value.

### HasDontMask

`func (o *AlarmRule) HasDontMask() bool`

HasDontMask returns a boolean if a field has been set.

### SetDontMaskNil

`func (o *AlarmRule) SetDontMaskNil(b bool)`

 SetDontMaskNil sets the value for DontMask to be an explicit nil

### UnsetDontMask
`func (o *AlarmRule) UnsetDontMask()`

UnsetDontMask ensures that no value is present for DontMask, not even an explicit nil
### GetCheckType

`func (o *AlarmRule) GetCheckType() string`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *AlarmRule) GetCheckTypeOk() (*string, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *AlarmRule) SetCheckType(v string)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *AlarmRule) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### SetCheckTypeNil

`func (o *AlarmRule) SetCheckTypeNil(b bool)`

 SetCheckTypeNil sets the value for CheckType to be an explicit nil

### UnsetCheckType
`func (o *AlarmRule) UnsetCheckType()`

UnsetCheckType ensures that no value is present for CheckType, not even an explicit nil
### GetAssetInfo

`func (o *AlarmRule) GetAssetInfo() Asset`

GetAssetInfo returns the AssetInfo field if non-nil, zero value otherwise.

### GetAssetInfoOk

`func (o *AlarmRule) GetAssetInfoOk() (*Asset, bool)`

GetAssetInfoOk returns a tuple with the AssetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetInfo

`func (o *AlarmRule) SetAssetInfo(v Asset)`

SetAssetInfo sets AssetInfo field to given value.

### HasAssetInfo

`func (o *AlarmRule) HasAssetInfo() bool`

HasAssetInfo returns a boolean if a field has been set.

### SetAssetInfoNil

`func (o *AlarmRule) SetAssetInfoNil(b bool)`

 SetAssetInfoNil sets the value for AssetInfo to be an explicit nil

### UnsetAssetInfo
`func (o *AlarmRule) UnsetAssetInfo()`

UnsetAssetInfo ensures that no value is present for AssetInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


