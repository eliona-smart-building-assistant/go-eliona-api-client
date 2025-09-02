# DataTrendAggregated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **int32** | ID of the corresponding asset | 
**Subtype** | Pointer to [**DataSubtype**](DataSubtype.md) |  | [optional] [default to INPUT]
**Attribute** | Pointer to **string** | Name of the attribute which holds the data points | [optional] 
**Timestamp** | Pointer to **NullableTime** | Timestamp of this aggregated data set. The timestamp when the timeslot based on raster starts. | [optional] 
**Count** | Pointer to **NullableFloat64** | Count of data points in this aggregated data set | [optional] 
**Average** | Pointer to **NullableFloat64** | Average of all data points for this aggregated data set | [optional] 
**Sum** | Pointer to **NullableFloat64** | Sum of all data points for this aggregated data set | [optional] 
**First** | Pointer to **NullableFloat64** | First data point in this aggregated data set | [optional] 
**Min** | Pointer to **NullableFloat64** | Data point with the most minimal value in this aggregated data set | [optional] 
**Max** | Pointer to **NullableFloat64** | Data point with the most maximal value in this aggregated data set | [optional] 
**Last** | Pointer to **NullableFloat64** | Latest data point in this aggregated data set | [optional] 

## Methods

### NewDataTrendAggregated

`func NewDataTrendAggregated(assetId int32, ) *DataTrendAggregated`

NewDataTrendAggregated instantiates a new DataTrendAggregated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTrendAggregatedWithDefaults

`func NewDataTrendAggregatedWithDefaults() *DataTrendAggregated`

NewDataTrendAggregatedWithDefaults instantiates a new DataTrendAggregated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *DataTrendAggregated) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *DataTrendAggregated) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *DataTrendAggregated) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.


### GetSubtype

`func (o *DataTrendAggregated) GetSubtype() DataSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *DataTrendAggregated) GetSubtypeOk() (*DataSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *DataTrendAggregated) SetSubtype(v DataSubtype)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *DataTrendAggregated) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetAttribute

`func (o *DataTrendAggregated) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *DataTrendAggregated) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *DataTrendAggregated) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *DataTrendAggregated) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetTimestamp

`func (o *DataTrendAggregated) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DataTrendAggregated) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DataTrendAggregated) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DataTrendAggregated) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DataTrendAggregated) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DataTrendAggregated) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCount

`func (o *DataTrendAggregated) GetCount() float64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DataTrendAggregated) GetCountOk() (*float64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DataTrendAggregated) SetCount(v float64)`

SetCount sets Count field to given value.

### HasCount

`func (o *DataTrendAggregated) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *DataTrendAggregated) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *DataTrendAggregated) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetAverage

`func (o *DataTrendAggregated) GetAverage() float64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *DataTrendAggregated) GetAverageOk() (*float64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *DataTrendAggregated) SetAverage(v float64)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *DataTrendAggregated) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### SetAverageNil

`func (o *DataTrendAggregated) SetAverageNil(b bool)`

 SetAverageNil sets the value for Average to be an explicit nil

### UnsetAverage
`func (o *DataTrendAggregated) UnsetAverage()`

UnsetAverage ensures that no value is present for Average, not even an explicit nil
### GetSum

`func (o *DataTrendAggregated) GetSum() float64`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *DataTrendAggregated) GetSumOk() (*float64, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *DataTrendAggregated) SetSum(v float64)`

SetSum sets Sum field to given value.

### HasSum

`func (o *DataTrendAggregated) HasSum() bool`

HasSum returns a boolean if a field has been set.

### SetSumNil

`func (o *DataTrendAggregated) SetSumNil(b bool)`

 SetSumNil sets the value for Sum to be an explicit nil

### UnsetSum
`func (o *DataTrendAggregated) UnsetSum()`

UnsetSum ensures that no value is present for Sum, not even an explicit nil
### GetFirst

`func (o *DataTrendAggregated) GetFirst() float64`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *DataTrendAggregated) GetFirstOk() (*float64, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *DataTrendAggregated) SetFirst(v float64)`

SetFirst sets First field to given value.

### HasFirst

`func (o *DataTrendAggregated) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### SetFirstNil

`func (o *DataTrendAggregated) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *DataTrendAggregated) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil
### GetMin

`func (o *DataTrendAggregated) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *DataTrendAggregated) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *DataTrendAggregated) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *DataTrendAggregated) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *DataTrendAggregated) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *DataTrendAggregated) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *DataTrendAggregated) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *DataTrendAggregated) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *DataTrendAggregated) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *DataTrendAggregated) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *DataTrendAggregated) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *DataTrendAggregated) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetLast

`func (o *DataTrendAggregated) GetLast() float64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *DataTrendAggregated) GetLastOk() (*float64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *DataTrendAggregated) SetLast(v float64)`

SetLast sets Last field to given value.

### HasLast

`func (o *DataTrendAggregated) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLastNil

`func (o *DataTrendAggregated) SetLastNil(b bool)`

 SetLastNil sets the value for Last to be an explicit nil

### UnsetLast
`func (o *DataTrendAggregated) UnsetLast()`

UnsetLast ensures that no value is present for Last, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


