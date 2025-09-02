# AttributeDisplay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **int32** | ID of the corresponding asset | 
**Subtype** | [**DataSubtype**](DataSubtype.md) |  | [default to INPUT]
**Attribute** | **string** | Name of the attribute of the asset type | 
**Unit** | Pointer to **NullableString** | Physical unit of numeric data | [optional] 
**Precision** | Pointer to **NullableInt64** | Number of decimal places | [optional] 
**Min** | Pointer to **NullableFloat64** | Lower limit | [optional] 
**Max** | Pointer to **NullableFloat64** | Upper limit | [optional] 
**Viewer** | Pointer to **NullableBool** | Should the attribute be displayed in viewer | [optional] [default to false]
**Ar** | Pointer to **NullableBool** | Should the attribute be displayed in AR | [optional] [default to false]
**Sequence** | Pointer to **NullableInt64** | Sequence in AR display | [optional] 
**Map** | Pointer to **[]map[string]interface{}** | list of mapping between value and custom text | [optional] 

## Methods

### NewAttributeDisplay

`func NewAttributeDisplay(assetId int32, subtype DataSubtype, attribute string, ) *AttributeDisplay`

NewAttributeDisplay instantiates a new AttributeDisplay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeDisplayWithDefaults

`func NewAttributeDisplayWithDefaults() *AttributeDisplay`

NewAttributeDisplayWithDefaults instantiates a new AttributeDisplay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *AttributeDisplay) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AttributeDisplay) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AttributeDisplay) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.


### GetSubtype

`func (o *AttributeDisplay) GetSubtype() DataSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AttributeDisplay) GetSubtypeOk() (*DataSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AttributeDisplay) SetSubtype(v DataSubtype)`

SetSubtype sets Subtype field to given value.


### GetAttribute

`func (o *AttributeDisplay) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AttributeDisplay) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AttributeDisplay) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetUnit

`func (o *AttributeDisplay) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AttributeDisplay) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AttributeDisplay) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AttributeDisplay) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *AttributeDisplay) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *AttributeDisplay) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetPrecision

`func (o *AttributeDisplay) GetPrecision() int64`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *AttributeDisplay) GetPrecisionOk() (*int64, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *AttributeDisplay) SetPrecision(v int64)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *AttributeDisplay) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### SetPrecisionNil

`func (o *AttributeDisplay) SetPrecisionNil(b bool)`

 SetPrecisionNil sets the value for Precision to be an explicit nil

### UnsetPrecision
`func (o *AttributeDisplay) UnsetPrecision()`

UnsetPrecision ensures that no value is present for Precision, not even an explicit nil
### GetMin

`func (o *AttributeDisplay) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *AttributeDisplay) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *AttributeDisplay) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *AttributeDisplay) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *AttributeDisplay) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *AttributeDisplay) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *AttributeDisplay) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *AttributeDisplay) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *AttributeDisplay) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *AttributeDisplay) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *AttributeDisplay) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *AttributeDisplay) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetViewer

`func (o *AttributeDisplay) GetViewer() bool`

GetViewer returns the Viewer field if non-nil, zero value otherwise.

### GetViewerOk

`func (o *AttributeDisplay) GetViewerOk() (*bool, bool)`

GetViewerOk returns a tuple with the Viewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewer

`func (o *AttributeDisplay) SetViewer(v bool)`

SetViewer sets Viewer field to given value.

### HasViewer

`func (o *AttributeDisplay) HasViewer() bool`

HasViewer returns a boolean if a field has been set.

### SetViewerNil

`func (o *AttributeDisplay) SetViewerNil(b bool)`

 SetViewerNil sets the value for Viewer to be an explicit nil

### UnsetViewer
`func (o *AttributeDisplay) UnsetViewer()`

UnsetViewer ensures that no value is present for Viewer, not even an explicit nil
### GetAr

`func (o *AttributeDisplay) GetAr() bool`

GetAr returns the Ar field if non-nil, zero value otherwise.

### GetArOk

`func (o *AttributeDisplay) GetArOk() (*bool, bool)`

GetArOk returns a tuple with the Ar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAr

`func (o *AttributeDisplay) SetAr(v bool)`

SetAr sets Ar field to given value.

### HasAr

`func (o *AttributeDisplay) HasAr() bool`

HasAr returns a boolean if a field has been set.

### SetArNil

`func (o *AttributeDisplay) SetArNil(b bool)`

 SetArNil sets the value for Ar to be an explicit nil

### UnsetAr
`func (o *AttributeDisplay) UnsetAr()`

UnsetAr ensures that no value is present for Ar, not even an explicit nil
### GetSequence

`func (o *AttributeDisplay) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *AttributeDisplay) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *AttributeDisplay) SetSequence(v int64)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *AttributeDisplay) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequenceNil

`func (o *AttributeDisplay) SetSequenceNil(b bool)`

 SetSequenceNil sets the value for Sequence to be an explicit nil

### UnsetSequence
`func (o *AttributeDisplay) UnsetSequence()`

UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
### GetMap

`func (o *AttributeDisplay) GetMap() []map[string]interface{}`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *AttributeDisplay) GetMapOk() (*[]map[string]interface{}, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *AttributeDisplay) SetMap(v []map[string]interface{})`

SetMap sets Map field to given value.

### HasMap

`func (o *AttributeDisplay) HasMap() bool`

HasMap returns a boolean if a field has been set.

### SetMapNil

`func (o *AttributeDisplay) SetMapNil(b bool)`

 SetMapNil sets the value for Map to be an explicit nil

### UnsetMap
`func (o *AttributeDisplay) UnsetMap()`

UnsetMap ensures that no value is present for Map, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


