# Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **int32** | ID of the corresponding asset | 
**Subtype** | [**DataSubtype**](DataSubtype.md) |  | [default to INPUT]
**Timestamp** | Pointer to **NullableTime** | Timestamp of the latest data change | [optional] 
**Data** | **map[string]interface{}** | Asset payload | 
**AssetTypeName** | Pointer to **NullableString** | The name of the corresponding asset type | [optional] [readonly] 
**ClientReference** | Pointer to **NullableString** | freely assignable by the client to identify self-created data | [optional] 

## Methods

### NewData

`func NewData(assetId int32, subtype DataSubtype, data map[string]interface{}, ) *Data`

NewData instantiates a new Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWithDefaults

`func NewDataWithDefaults() *Data`

NewDataWithDefaults instantiates a new Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *Data) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Data) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Data) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.


### GetSubtype

`func (o *Data) GetSubtype() DataSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Data) GetSubtypeOk() (*DataSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Data) SetSubtype(v DataSubtype)`

SetSubtype sets Subtype field to given value.


### GetTimestamp

`func (o *Data) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Data) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Data) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Data) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *Data) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Data) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetData

`func (o *Data) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Data) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Data) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetAssetTypeName

`func (o *Data) GetAssetTypeName() string`

GetAssetTypeName returns the AssetTypeName field if non-nil, zero value otherwise.

### GetAssetTypeNameOk

`func (o *Data) GetAssetTypeNameOk() (*string, bool)`

GetAssetTypeNameOk returns a tuple with the AssetTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypeName

`func (o *Data) SetAssetTypeName(v string)`

SetAssetTypeName sets AssetTypeName field to given value.

### HasAssetTypeName

`func (o *Data) HasAssetTypeName() bool`

HasAssetTypeName returns a boolean if a field has been set.

### SetAssetTypeNameNil

`func (o *Data) SetAssetTypeNameNil(b bool)`

 SetAssetTypeNameNil sets the value for AssetTypeName to be an explicit nil

### UnsetAssetTypeName
`func (o *Data) UnsetAssetTypeName()`

UnsetAssetTypeName ensures that no value is present for AssetTypeName, not even an explicit nil
### GetClientReference

`func (o *Data) GetClientReference() string`

GetClientReference returns the ClientReference field if non-nil, zero value otherwise.

### GetClientReferenceOk

`func (o *Data) GetClientReferenceOk() (*string, bool)`

GetClientReferenceOk returns a tuple with the ClientReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReference

`func (o *Data) SetClientReference(v string)`

SetClientReference sets ClientReference field to given value.

### HasClientReference

`func (o *Data) HasClientReference() bool`

HasClientReference returns a boolean if a field has been set.

### SetClientReferenceNil

`func (o *Data) SetClientReferenceNil(b bool)`

 SetClientReferenceNil sets the value for ClientReference to be an explicit nil

### UnsetClientReference
`func (o *Data) UnsetClientReference()`

UnsetClientReference ensures that no value is present for ClientReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


