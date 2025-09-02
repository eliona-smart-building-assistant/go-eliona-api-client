# DataAggregated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationId** | Pointer to **int32** | ID of the aggregation | [optional] 
**AssetId** | **int32** | ID of the corresponding asset | 
**Subtype** | [**DataSubtype**](DataSubtype.md) |  | [default to INPUT]
**Attribute** | Pointer to **string** | Name of the attribute which holds the data points | [optional] 
**Raster** | **string** | Calculation intervals. | 
**Timestamp** | Pointer to **NullableTime** | Timestamp of this aggregated data set. The timestamp when the timeslot based on raster starts. | [optional] 
**Count** | Pointer to **NullableFloat64** | Count of data points in this aggregated data set | [optional] 
**Average** | Pointer to **NullableFloat64** | Average of all data points for this aggregated data set | [optional] 
**Sum** | Pointer to **NullableFloat64** | Sum of all data points for this aggregated data set | [optional] 
**First** | Pointer to **NullableFloat64** | First data point in this aggregated data set | [optional] 
**Min** | Pointer to **NullableFloat64** | Data point with the most minimal value in this aggregated data set | [optional] 
**Max** | Pointer to **NullableFloat64** | Data point with the most maximal value in this aggregated data set | [optional] 
**Last** | Pointer to **NullableFloat64** | Latest data point in this aggregated data set | [optional] 
**LastTimestamp** | Pointer to **NullableTime** | Timestamp of the latest data point | [optional] 
**AssetTypeName** | Pointer to **NullableString** | The name of the corresponding asset type | [optional] [readonly] 

## Methods

### NewDataAggregated

`func NewDataAggregated(assetId int32, subtype DataSubtype, raster string, ) *DataAggregated`

NewDataAggregated instantiates a new DataAggregated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataAggregatedWithDefaults

`func NewDataAggregatedWithDefaults() *DataAggregated`

NewDataAggregatedWithDefaults instantiates a new DataAggregated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationId

`func (o *DataAggregated) GetAggregationId() int32`

GetAggregationId returns the AggregationId field if non-nil, zero value otherwise.

### GetAggregationIdOk

`func (o *DataAggregated) GetAggregationIdOk() (*int32, bool)`

GetAggregationIdOk returns a tuple with the AggregationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationId

`func (o *DataAggregated) SetAggregationId(v int32)`

SetAggregationId sets AggregationId field to given value.

### HasAggregationId

`func (o *DataAggregated) HasAggregationId() bool`

HasAggregationId returns a boolean if a field has been set.

### GetAssetId

`func (o *DataAggregated) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *DataAggregated) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *DataAggregated) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.


### GetSubtype

`func (o *DataAggregated) GetSubtype() DataSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *DataAggregated) GetSubtypeOk() (*DataSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *DataAggregated) SetSubtype(v DataSubtype)`

SetSubtype sets Subtype field to given value.


### GetAttribute

`func (o *DataAggregated) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *DataAggregated) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *DataAggregated) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *DataAggregated) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetRaster

`func (o *DataAggregated) GetRaster() string`

GetRaster returns the Raster field if non-nil, zero value otherwise.

### GetRasterOk

`func (o *DataAggregated) GetRasterOk() (*string, bool)`

GetRasterOk returns a tuple with the Raster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaster

`func (o *DataAggregated) SetRaster(v string)`

SetRaster sets Raster field to given value.


### GetTimestamp

`func (o *DataAggregated) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DataAggregated) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DataAggregated) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DataAggregated) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DataAggregated) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DataAggregated) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCount

`func (o *DataAggregated) GetCount() float64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DataAggregated) GetCountOk() (*float64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DataAggregated) SetCount(v float64)`

SetCount sets Count field to given value.

### HasCount

`func (o *DataAggregated) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *DataAggregated) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *DataAggregated) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetAverage

`func (o *DataAggregated) GetAverage() float64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *DataAggregated) GetAverageOk() (*float64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *DataAggregated) SetAverage(v float64)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *DataAggregated) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### SetAverageNil

`func (o *DataAggregated) SetAverageNil(b bool)`

 SetAverageNil sets the value for Average to be an explicit nil

### UnsetAverage
`func (o *DataAggregated) UnsetAverage()`

UnsetAverage ensures that no value is present for Average, not even an explicit nil
### GetSum

`func (o *DataAggregated) GetSum() float64`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *DataAggregated) GetSumOk() (*float64, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *DataAggregated) SetSum(v float64)`

SetSum sets Sum field to given value.

### HasSum

`func (o *DataAggregated) HasSum() bool`

HasSum returns a boolean if a field has been set.

### SetSumNil

`func (o *DataAggregated) SetSumNil(b bool)`

 SetSumNil sets the value for Sum to be an explicit nil

### UnsetSum
`func (o *DataAggregated) UnsetSum()`

UnsetSum ensures that no value is present for Sum, not even an explicit nil
### GetFirst

`func (o *DataAggregated) GetFirst() float64`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *DataAggregated) GetFirstOk() (*float64, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *DataAggregated) SetFirst(v float64)`

SetFirst sets First field to given value.

### HasFirst

`func (o *DataAggregated) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### SetFirstNil

`func (o *DataAggregated) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *DataAggregated) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil
### GetMin

`func (o *DataAggregated) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *DataAggregated) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *DataAggregated) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *DataAggregated) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *DataAggregated) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *DataAggregated) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *DataAggregated) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *DataAggregated) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *DataAggregated) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *DataAggregated) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *DataAggregated) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *DataAggregated) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetLast

`func (o *DataAggregated) GetLast() float64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *DataAggregated) GetLastOk() (*float64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *DataAggregated) SetLast(v float64)`

SetLast sets Last field to given value.

### HasLast

`func (o *DataAggregated) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLastNil

`func (o *DataAggregated) SetLastNil(b bool)`

 SetLastNil sets the value for Last to be an explicit nil

### UnsetLast
`func (o *DataAggregated) UnsetLast()`

UnsetLast ensures that no value is present for Last, not even an explicit nil
### GetLastTimestamp

`func (o *DataAggregated) GetLastTimestamp() time.Time`

GetLastTimestamp returns the LastTimestamp field if non-nil, zero value otherwise.

### GetLastTimestampOk

`func (o *DataAggregated) GetLastTimestampOk() (*time.Time, bool)`

GetLastTimestampOk returns a tuple with the LastTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimestamp

`func (o *DataAggregated) SetLastTimestamp(v time.Time)`

SetLastTimestamp sets LastTimestamp field to given value.

### HasLastTimestamp

`func (o *DataAggregated) HasLastTimestamp() bool`

HasLastTimestamp returns a boolean if a field has been set.

### SetLastTimestampNil

`func (o *DataAggregated) SetLastTimestampNil(b bool)`

 SetLastTimestampNil sets the value for LastTimestamp to be an explicit nil

### UnsetLastTimestamp
`func (o *DataAggregated) UnsetLastTimestamp()`

UnsetLastTimestamp ensures that no value is present for LastTimestamp, not even an explicit nil
### GetAssetTypeName

`func (o *DataAggregated) GetAssetTypeName() string`

GetAssetTypeName returns the AssetTypeName field if non-nil, zero value otherwise.

### GetAssetTypeNameOk

`func (o *DataAggregated) GetAssetTypeNameOk() (*string, bool)`

GetAssetTypeNameOk returns a tuple with the AssetTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypeName

`func (o *DataAggregated) SetAssetTypeName(v string)`

SetAssetTypeName sets AssetTypeName field to given value.

### HasAssetTypeName

`func (o *DataAggregated) HasAssetTypeName() bool`

HasAssetTypeName returns a boolean if a field has been set.

### SetAssetTypeNameNil

`func (o *DataAggregated) SetAssetTypeNameNil(b bool)`

 SetAssetTypeNameNil sets the value for AssetTypeName to be an explicit nil

### UnsetAssetTypeName
`func (o *DataAggregated) UnsetAssetTypeName()`

UnsetAssetTypeName ensures that no value is present for AssetTypeName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


