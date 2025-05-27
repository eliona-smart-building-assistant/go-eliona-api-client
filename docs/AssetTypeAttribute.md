# AssetTypeAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetTypeName** | Pointer to **NullableString** | The unique name for the asset type | [optional] 
**Name** | **string** | Unique key of asset data | 
**Subtype** | [**DataSubtype**](DataSubtype.md) |  | [default to SUBTYPE_INPUT]
**Type** | Pointer to **NullableString** | Name of the type for this attribute: air_quality, battery-voltage, brightness, co2, current, device-info, device-status, energy, flow, frequency, humidity, inputs-and-switches, level, motion, operating-status, people-count, power, presence, pressure, temperature, vehicle-detector, voltage, weather, voc | [optional] 
**Enable** | Pointer to **bool** | Is data active or not | [optional] [default to false]
**Translation** | Pointer to [**NullableTranslation**](Translation.md) |  | [optional] 
**Unit** | Pointer to **NullableString** | Physical unit of numeric data | [optional] 
**Precision** | Pointer to **NullableInt64** | Number of decimal places | [optional] 
**Min** | Pointer to **NullableFloat64** | Lower limit | [optional] 
**Max** | Pointer to **NullableFloat64** | Upper limit | [optional] 
**AggregationMode** | Pointer to **NullableString** | Deprecated: Use the &#39;GET /data-trend-aggregated&#39; endpoint to retrieve aggregated data for periodic rasters without defining aggregations. Aggregation calculation mode.  | [optional] 
**AggregationRasters** | Pointer to **[]string** |  | [optional] 
**Viewer** | Pointer to **NullableBool** | Should the attribute be displayed in viewer | [optional] [default to false]
**Ar** | Pointer to **NullableBool** | Should the attribute be displayed in AR | [optional] [default to false]
**Sequence** | Pointer to **NullableInt64** | Sequence in AR display | [optional] 
**Virtual** | Pointer to **NullableBool** | Is the attribute virtual or not | [optional] 
**Formula** | Pointer to **NullableString** | calculation rule to calculate the value for this attribute | [optional] 
**Scale** | Pointer to **NullableFloat32** | value scale | [optional] 
**Zero** | Pointer to **NullableFloat32** | value scale | [optional] 
**Map** | Pointer to **[]map[string]interface{}** | list of mapping between value and custom text | [optional] 
**SourcePath** | Pointer to **[]string** | source path for attribute value | [optional] 
**IsDigital** | Pointer to **NullableBool** | is attribute digital | [optional] 

## Methods

### NewAssetTypeAttribute

`func NewAssetTypeAttribute(name string, subtype DataSubtype, ) *AssetTypeAttribute`

NewAssetTypeAttribute instantiates a new AssetTypeAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTypeAttributeWithDefaults

`func NewAssetTypeAttributeWithDefaults() *AssetTypeAttribute`

NewAssetTypeAttributeWithDefaults instantiates a new AssetTypeAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetTypeName

`func (o *AssetTypeAttribute) GetAssetTypeName() string`

GetAssetTypeName returns the AssetTypeName field if non-nil, zero value otherwise.

### GetAssetTypeNameOk

`func (o *AssetTypeAttribute) GetAssetTypeNameOk() (*string, bool)`

GetAssetTypeNameOk returns a tuple with the AssetTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypeName

`func (o *AssetTypeAttribute) SetAssetTypeName(v string)`

SetAssetTypeName sets AssetTypeName field to given value.

### HasAssetTypeName

`func (o *AssetTypeAttribute) HasAssetTypeName() bool`

HasAssetTypeName returns a boolean if a field has been set.

### SetAssetTypeNameNil

`func (o *AssetTypeAttribute) SetAssetTypeNameNil(b bool)`

 SetAssetTypeNameNil sets the value for AssetTypeName to be an explicit nil

### UnsetAssetTypeName
`func (o *AssetTypeAttribute) UnsetAssetTypeName()`

UnsetAssetTypeName ensures that no value is present for AssetTypeName, not even an explicit nil
### GetName

`func (o *AssetTypeAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetTypeAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetTypeAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetSubtype

`func (o *AssetTypeAttribute) GetSubtype() DataSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AssetTypeAttribute) GetSubtypeOk() (*DataSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AssetTypeAttribute) SetSubtype(v DataSubtype)`

SetSubtype sets Subtype field to given value.


### GetType

`func (o *AssetTypeAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssetTypeAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssetTypeAttribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssetTypeAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AssetTypeAttribute) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AssetTypeAttribute) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetEnable

`func (o *AssetTypeAttribute) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AssetTypeAttribute) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AssetTypeAttribute) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AssetTypeAttribute) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetTranslation

`func (o *AssetTypeAttribute) GetTranslation() Translation`

GetTranslation returns the Translation field if non-nil, zero value otherwise.

### GetTranslationOk

`func (o *AssetTypeAttribute) GetTranslationOk() (*Translation, bool)`

GetTranslationOk returns a tuple with the Translation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslation

`func (o *AssetTypeAttribute) SetTranslation(v Translation)`

SetTranslation sets Translation field to given value.

### HasTranslation

`func (o *AssetTypeAttribute) HasTranslation() bool`

HasTranslation returns a boolean if a field has been set.

### SetTranslationNil

`func (o *AssetTypeAttribute) SetTranslationNil(b bool)`

 SetTranslationNil sets the value for Translation to be an explicit nil

### UnsetTranslation
`func (o *AssetTypeAttribute) UnsetTranslation()`

UnsetTranslation ensures that no value is present for Translation, not even an explicit nil
### GetUnit

`func (o *AssetTypeAttribute) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AssetTypeAttribute) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AssetTypeAttribute) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AssetTypeAttribute) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *AssetTypeAttribute) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *AssetTypeAttribute) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetPrecision

`func (o *AssetTypeAttribute) GetPrecision() int64`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *AssetTypeAttribute) GetPrecisionOk() (*int64, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *AssetTypeAttribute) SetPrecision(v int64)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *AssetTypeAttribute) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### SetPrecisionNil

`func (o *AssetTypeAttribute) SetPrecisionNil(b bool)`

 SetPrecisionNil sets the value for Precision to be an explicit nil

### UnsetPrecision
`func (o *AssetTypeAttribute) UnsetPrecision()`

UnsetPrecision ensures that no value is present for Precision, not even an explicit nil
### GetMin

`func (o *AssetTypeAttribute) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *AssetTypeAttribute) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *AssetTypeAttribute) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *AssetTypeAttribute) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *AssetTypeAttribute) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *AssetTypeAttribute) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *AssetTypeAttribute) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *AssetTypeAttribute) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *AssetTypeAttribute) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *AssetTypeAttribute) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *AssetTypeAttribute) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *AssetTypeAttribute) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetAggregationMode

`func (o *AssetTypeAttribute) GetAggregationMode() string`

GetAggregationMode returns the AggregationMode field if non-nil, zero value otherwise.

### GetAggregationModeOk

`func (o *AssetTypeAttribute) GetAggregationModeOk() (*string, bool)`

GetAggregationModeOk returns a tuple with the AggregationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationMode

`func (o *AssetTypeAttribute) SetAggregationMode(v string)`

SetAggregationMode sets AggregationMode field to given value.

### HasAggregationMode

`func (o *AssetTypeAttribute) HasAggregationMode() bool`

HasAggregationMode returns a boolean if a field has been set.

### SetAggregationModeNil

`func (o *AssetTypeAttribute) SetAggregationModeNil(b bool)`

 SetAggregationModeNil sets the value for AggregationMode to be an explicit nil

### UnsetAggregationMode
`func (o *AssetTypeAttribute) UnsetAggregationMode()`

UnsetAggregationMode ensures that no value is present for AggregationMode, not even an explicit nil
### GetAggregationRasters

`func (o *AssetTypeAttribute) GetAggregationRasters() []string`

GetAggregationRasters returns the AggregationRasters field if non-nil, zero value otherwise.

### GetAggregationRastersOk

`func (o *AssetTypeAttribute) GetAggregationRastersOk() (*[]string, bool)`

GetAggregationRastersOk returns a tuple with the AggregationRasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationRasters

`func (o *AssetTypeAttribute) SetAggregationRasters(v []string)`

SetAggregationRasters sets AggregationRasters field to given value.

### HasAggregationRasters

`func (o *AssetTypeAttribute) HasAggregationRasters() bool`

HasAggregationRasters returns a boolean if a field has been set.

### GetViewer

`func (o *AssetTypeAttribute) GetViewer() bool`

GetViewer returns the Viewer field if non-nil, zero value otherwise.

### GetViewerOk

`func (o *AssetTypeAttribute) GetViewerOk() (*bool, bool)`

GetViewerOk returns a tuple with the Viewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewer

`func (o *AssetTypeAttribute) SetViewer(v bool)`

SetViewer sets Viewer field to given value.

### HasViewer

`func (o *AssetTypeAttribute) HasViewer() bool`

HasViewer returns a boolean if a field has been set.

### SetViewerNil

`func (o *AssetTypeAttribute) SetViewerNil(b bool)`

 SetViewerNil sets the value for Viewer to be an explicit nil

### UnsetViewer
`func (o *AssetTypeAttribute) UnsetViewer()`

UnsetViewer ensures that no value is present for Viewer, not even an explicit nil
### GetAr

`func (o *AssetTypeAttribute) GetAr() bool`

GetAr returns the Ar field if non-nil, zero value otherwise.

### GetArOk

`func (o *AssetTypeAttribute) GetArOk() (*bool, bool)`

GetArOk returns a tuple with the Ar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAr

`func (o *AssetTypeAttribute) SetAr(v bool)`

SetAr sets Ar field to given value.

### HasAr

`func (o *AssetTypeAttribute) HasAr() bool`

HasAr returns a boolean if a field has been set.

### SetArNil

`func (o *AssetTypeAttribute) SetArNil(b bool)`

 SetArNil sets the value for Ar to be an explicit nil

### UnsetAr
`func (o *AssetTypeAttribute) UnsetAr()`

UnsetAr ensures that no value is present for Ar, not even an explicit nil
### GetSequence

`func (o *AssetTypeAttribute) GetSequence() int64`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *AssetTypeAttribute) GetSequenceOk() (*int64, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *AssetTypeAttribute) SetSequence(v int64)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *AssetTypeAttribute) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequenceNil

`func (o *AssetTypeAttribute) SetSequenceNil(b bool)`

 SetSequenceNil sets the value for Sequence to be an explicit nil

### UnsetSequence
`func (o *AssetTypeAttribute) UnsetSequence()`

UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
### GetVirtual

`func (o *AssetTypeAttribute) GetVirtual() bool`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *AssetTypeAttribute) GetVirtualOk() (*bool, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *AssetTypeAttribute) SetVirtual(v bool)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *AssetTypeAttribute) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.

### SetVirtualNil

`func (o *AssetTypeAttribute) SetVirtualNil(b bool)`

 SetVirtualNil sets the value for Virtual to be an explicit nil

### UnsetVirtual
`func (o *AssetTypeAttribute) UnsetVirtual()`

UnsetVirtual ensures that no value is present for Virtual, not even an explicit nil
### GetFormula

`func (o *AssetTypeAttribute) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *AssetTypeAttribute) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *AssetTypeAttribute) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *AssetTypeAttribute) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### SetFormulaNil

`func (o *AssetTypeAttribute) SetFormulaNil(b bool)`

 SetFormulaNil sets the value for Formula to be an explicit nil

### UnsetFormula
`func (o *AssetTypeAttribute) UnsetFormula()`

UnsetFormula ensures that no value is present for Formula, not even an explicit nil
### GetScale

`func (o *AssetTypeAttribute) GetScale() float32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *AssetTypeAttribute) GetScaleOk() (*float32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *AssetTypeAttribute) SetScale(v float32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *AssetTypeAttribute) HasScale() bool`

HasScale returns a boolean if a field has been set.

### SetScaleNil

`func (o *AssetTypeAttribute) SetScaleNil(b bool)`

 SetScaleNil sets the value for Scale to be an explicit nil

### UnsetScale
`func (o *AssetTypeAttribute) UnsetScale()`

UnsetScale ensures that no value is present for Scale, not even an explicit nil
### GetZero

`func (o *AssetTypeAttribute) GetZero() float32`

GetZero returns the Zero field if non-nil, zero value otherwise.

### GetZeroOk

`func (o *AssetTypeAttribute) GetZeroOk() (*float32, bool)`

GetZeroOk returns a tuple with the Zero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZero

`func (o *AssetTypeAttribute) SetZero(v float32)`

SetZero sets Zero field to given value.

### HasZero

`func (o *AssetTypeAttribute) HasZero() bool`

HasZero returns a boolean if a field has been set.

### SetZeroNil

`func (o *AssetTypeAttribute) SetZeroNil(b bool)`

 SetZeroNil sets the value for Zero to be an explicit nil

### UnsetZero
`func (o *AssetTypeAttribute) UnsetZero()`

UnsetZero ensures that no value is present for Zero, not even an explicit nil
### GetMap

`func (o *AssetTypeAttribute) GetMap() []map[string]interface{}`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *AssetTypeAttribute) GetMapOk() (*[]map[string]interface{}, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *AssetTypeAttribute) SetMap(v []map[string]interface{})`

SetMap sets Map field to given value.

### HasMap

`func (o *AssetTypeAttribute) HasMap() bool`

HasMap returns a boolean if a field has been set.

### SetMapNil

`func (o *AssetTypeAttribute) SetMapNil(b bool)`

 SetMapNil sets the value for Map to be an explicit nil

### UnsetMap
`func (o *AssetTypeAttribute) UnsetMap()`

UnsetMap ensures that no value is present for Map, not even an explicit nil
### GetSourcePath

`func (o *AssetTypeAttribute) GetSourcePath() []string`

GetSourcePath returns the SourcePath field if non-nil, zero value otherwise.

### GetSourcePathOk

`func (o *AssetTypeAttribute) GetSourcePathOk() (*[]string, bool)`

GetSourcePathOk returns a tuple with the SourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePath

`func (o *AssetTypeAttribute) SetSourcePath(v []string)`

SetSourcePath sets SourcePath field to given value.

### HasSourcePath

`func (o *AssetTypeAttribute) HasSourcePath() bool`

HasSourcePath returns a boolean if a field has been set.

### SetSourcePathNil

`func (o *AssetTypeAttribute) SetSourcePathNil(b bool)`

 SetSourcePathNil sets the value for SourcePath to be an explicit nil

### UnsetSourcePath
`func (o *AssetTypeAttribute) UnsetSourcePath()`

UnsetSourcePath ensures that no value is present for SourcePath, not even an explicit nil
### GetIsDigital

`func (o *AssetTypeAttribute) GetIsDigital() bool`

GetIsDigital returns the IsDigital field if non-nil, zero value otherwise.

### GetIsDigitalOk

`func (o *AssetTypeAttribute) GetIsDigitalOk() (*bool, bool)`

GetIsDigitalOk returns a tuple with the IsDigital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDigital

`func (o *AssetTypeAttribute) SetIsDigital(v bool)`

SetIsDigital sets IsDigital field to given value.

### HasIsDigital

`func (o *AssetTypeAttribute) HasIsDigital() bool`

HasIsDigital returns a boolean if a field has been set.

### SetIsDigitalNil

`func (o *AssetTypeAttribute) SetIsDigitalNil(b bool)`

 SetIsDigitalNil sets the value for IsDigital to be an explicit nil

### UnsetIsDigital
`func (o *AssetTypeAttribute) UnsetIsDigital()`

UnsetIsDigital ensures that no value is present for IsDigital, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


