# AssetTypeCategoryProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name for this asset type category property | 
**Translation** | Pointer to [**NullableTranslation**](Translation.md) |  | [optional] 
**Unit** | Pointer to **NullableString** | Unit of data | [optional] 

## Methods

### NewAssetTypeCategoryProperty

`func NewAssetTypeCategoryProperty(name string, ) *AssetTypeCategoryProperty`

NewAssetTypeCategoryProperty instantiates a new AssetTypeCategoryProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTypeCategoryPropertyWithDefaults

`func NewAssetTypeCategoryPropertyWithDefaults() *AssetTypeCategoryProperty`

NewAssetTypeCategoryPropertyWithDefaults instantiates a new AssetTypeCategoryProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AssetTypeCategoryProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetTypeCategoryProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetTypeCategoryProperty) SetName(v string)`

SetName sets Name field to given value.


### GetTranslation

`func (o *AssetTypeCategoryProperty) GetTranslation() Translation`

GetTranslation returns the Translation field if non-nil, zero value otherwise.

### GetTranslationOk

`func (o *AssetTypeCategoryProperty) GetTranslationOk() (*Translation, bool)`

GetTranslationOk returns a tuple with the Translation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslation

`func (o *AssetTypeCategoryProperty) SetTranslation(v Translation)`

SetTranslation sets Translation field to given value.

### HasTranslation

`func (o *AssetTypeCategoryProperty) HasTranslation() bool`

HasTranslation returns a boolean if a field has been set.

### SetTranslationNil

`func (o *AssetTypeCategoryProperty) SetTranslationNil(b bool)`

 SetTranslationNil sets the value for Translation to be an explicit nil

### UnsetTranslation
`func (o *AssetTypeCategoryProperty) UnsetTranslation()`

UnsetTranslation ensures that no value is present for Translation, not even an explicit nil
### GetUnit

`func (o *AssetTypeCategoryProperty) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AssetTypeCategoryProperty) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AssetTypeCategoryProperty) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AssetTypeCategoryProperty) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *AssetTypeCategoryProperty) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *AssetTypeCategoryProperty) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


