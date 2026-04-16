# CalculationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | The id of the rule | [optional] [readonly] 
**AssetId** | **int32** | ID of the corresponding asset | 
**Subtype** | [**DataSubtype**](DataSubtype.md) |  | [default to INPUT]
**Attribute** | **string** | Name of the attribute of the asset type to be calculated | 
**Virtual** | Pointer to **NullableBool** | Is the calculation attribute virtual or not | [optional] 
**Formula** | Pointer to **NullableString** | calculation rule to calculate the value for the attribute | [optional] 
**UnitId** | Pointer to **NullableInt32** | ID of measurement unit | [optional] 
**Filter** | Pointer to **map[string]interface{}** | Filter definition for calculation rule | [optional] 

## Methods

### NewCalculationRule

`func NewCalculationRule(assetId int32, subtype DataSubtype, attribute string, ) *CalculationRule`

NewCalculationRule instantiates a new CalculationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculationRuleWithDefaults

`func NewCalculationRuleWithDefaults() *CalculationRule`

NewCalculationRuleWithDefaults instantiates a new CalculationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CalculationRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalculationRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalculationRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CalculationRule) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CalculationRule) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CalculationRule) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAssetId

`func (o *CalculationRule) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *CalculationRule) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *CalculationRule) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.


### GetSubtype

`func (o *CalculationRule) GetSubtype() DataSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CalculationRule) GetSubtypeOk() (*DataSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CalculationRule) SetSubtype(v DataSubtype)`

SetSubtype sets Subtype field to given value.


### GetAttribute

`func (o *CalculationRule) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *CalculationRule) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *CalculationRule) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetVirtual

`func (o *CalculationRule) GetVirtual() bool`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *CalculationRule) GetVirtualOk() (*bool, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *CalculationRule) SetVirtual(v bool)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *CalculationRule) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.

### SetVirtualNil

`func (o *CalculationRule) SetVirtualNil(b bool)`

 SetVirtualNil sets the value for Virtual to be an explicit nil

### UnsetVirtual
`func (o *CalculationRule) UnsetVirtual()`

UnsetVirtual ensures that no value is present for Virtual, not even an explicit nil
### GetFormula

`func (o *CalculationRule) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *CalculationRule) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *CalculationRule) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *CalculationRule) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### SetFormulaNil

`func (o *CalculationRule) SetFormulaNil(b bool)`

 SetFormulaNil sets the value for Formula to be an explicit nil

### UnsetFormula
`func (o *CalculationRule) UnsetFormula()`

UnsetFormula ensures that no value is present for Formula, not even an explicit nil
### GetUnitId

`func (o *CalculationRule) GetUnitId() int32`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *CalculationRule) GetUnitIdOk() (*int32, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *CalculationRule) SetUnitId(v int32)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *CalculationRule) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *CalculationRule) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *CalculationRule) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetFilter

`func (o *CalculationRule) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CalculationRule) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CalculationRule) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CalculationRule) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *CalculationRule) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *CalculationRule) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


