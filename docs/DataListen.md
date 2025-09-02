# DataListen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **int32** | ID of the corresponding asset | 
**Subtype** | [**DataSubtype**](DataSubtype.md) |  | [default to INPUT]
**Timestamp** | Pointer to **NullableTime** | Timestamp of the latest data change | [optional] 
**Data** | **map[string]interface{}** | Asset payload | 
**AssetTypeName** | Pointer to **NullableString** | The name of the corresponding asset type | [optional] [readonly] 
**ClientReference** | Pointer to **NullableString** | freely assignable by the client to identify self-created data | [optional] 
**StatusCode** | Pointer to **int32** | The status code expecting when actually perform the operation. Some values are - 200: updated (ok)  - 201: created - 204: deleted (no content) - 304: unchanged (not modified) - 400: problem (bad request) - 404: not found - 409: duplicated (conflict) - 422: unprocessable  | [optional] 

## Methods

### NewDataListen

`func NewDataListen(assetId int32, subtype DataSubtype, data map[string]interface{}, ) *DataListen`

NewDataListen instantiates a new DataListen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataListenWithDefaults

`func NewDataListenWithDefaults() *DataListen`

NewDataListenWithDefaults instantiates a new DataListen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *DataListen) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *DataListen) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *DataListen) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.


### GetSubtype

`func (o *DataListen) GetSubtype() DataSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *DataListen) GetSubtypeOk() (*DataSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *DataListen) SetSubtype(v DataSubtype)`

SetSubtype sets Subtype field to given value.


### GetTimestamp

`func (o *DataListen) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DataListen) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DataListen) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DataListen) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DataListen) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DataListen) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetData

`func (o *DataListen) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DataListen) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DataListen) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetAssetTypeName

`func (o *DataListen) GetAssetTypeName() string`

GetAssetTypeName returns the AssetTypeName field if non-nil, zero value otherwise.

### GetAssetTypeNameOk

`func (o *DataListen) GetAssetTypeNameOk() (*string, bool)`

GetAssetTypeNameOk returns a tuple with the AssetTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypeName

`func (o *DataListen) SetAssetTypeName(v string)`

SetAssetTypeName sets AssetTypeName field to given value.

### HasAssetTypeName

`func (o *DataListen) HasAssetTypeName() bool`

HasAssetTypeName returns a boolean if a field has been set.

### SetAssetTypeNameNil

`func (o *DataListen) SetAssetTypeNameNil(b bool)`

 SetAssetTypeNameNil sets the value for AssetTypeName to be an explicit nil

### UnsetAssetTypeName
`func (o *DataListen) UnsetAssetTypeName()`

UnsetAssetTypeName ensures that no value is present for AssetTypeName, not even an explicit nil
### GetClientReference

`func (o *DataListen) GetClientReference() string`

GetClientReference returns the ClientReference field if non-nil, zero value otherwise.

### GetClientReferenceOk

`func (o *DataListen) GetClientReferenceOk() (*string, bool)`

GetClientReferenceOk returns a tuple with the ClientReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReference

`func (o *DataListen) SetClientReference(v string)`

SetClientReference sets ClientReference field to given value.

### HasClientReference

`func (o *DataListen) HasClientReference() bool`

HasClientReference returns a boolean if a field has been set.

### SetClientReferenceNil

`func (o *DataListen) SetClientReferenceNil(b bool)`

 SetClientReferenceNil sets the value for ClientReference to be an explicit nil

### UnsetClientReference
`func (o *DataListen) UnsetClientReference()`

UnsetClientReference ensures that no value is present for ClientReference, not even an explicit nil
### GetStatusCode

`func (o *DataListen) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *DataListen) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *DataListen) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *DataListen) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


