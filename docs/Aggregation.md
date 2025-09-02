# Aggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | ID of the aggregation | [optional] [readonly] 
**AssetId** | **int32** | ID of the corresponding asset | 
**Subtype** | [**DataSubtype**](DataSubtype.md) |  | [default to INPUT]
**Attribute** | Pointer to **string** | Name of the attribute which holds the data points | [optional] 
**Mode** | **string** | Calculation mode | 
**Raster** | Pointer to **NullableString** | calculation interval | [optional] 

## Methods

### NewAggregation

`func NewAggregation(assetId int32, subtype DataSubtype, mode string, ) *Aggregation`

NewAggregation instantiates a new Aggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationWithDefaults

`func NewAggregationWithDefaults() *Aggregation`

NewAggregationWithDefaults instantiates a new Aggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Aggregation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Aggregation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Aggregation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Aggregation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Aggregation) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Aggregation) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAssetId

`func (o *Aggregation) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Aggregation) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Aggregation) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.


### GetSubtype

`func (o *Aggregation) GetSubtype() DataSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Aggregation) GetSubtypeOk() (*DataSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Aggregation) SetSubtype(v DataSubtype)`

SetSubtype sets Subtype field to given value.


### GetAttribute

`func (o *Aggregation) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Aggregation) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Aggregation) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *Aggregation) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetMode

`func (o *Aggregation) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Aggregation) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Aggregation) SetMode(v string)`

SetMode sets Mode field to given value.


### GetRaster

`func (o *Aggregation) GetRaster() string`

GetRaster returns the Raster field if non-nil, zero value otherwise.

### GetRasterOk

`func (o *Aggregation) GetRasterOk() (*string, bool)`

GetRasterOk returns a tuple with the Raster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaster

`func (o *Aggregation) SetRaster(v string)`

SetRaster sets Raster field to given value.

### HasRaster

`func (o *Aggregation) HasRaster() bool`

HasRaster returns a boolean if a field has been set.

### SetRasterNil

`func (o *Aggregation) SetRasterNil(b bool)`

 SetRasterNil sets the value for Raster to be an explicit nil

### UnsetRaster
`func (o *Aggregation) UnsetRaster()`

UnsetRaster ensures that no value is present for Raster, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


