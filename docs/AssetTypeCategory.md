# AssetTypeCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name for this asset type category | 
**Translation** | Pointer to [**NullableTranslation**](Translation.md) |  | [optional] 
**Properties** | Pointer to [**[]AssetTypeCategoryProperty**](AssetTypeCategoryProperty.md) | List of asset type category properties | [optional] 

## Methods

### NewAssetTypeCategory

`func NewAssetTypeCategory(name string, ) *AssetTypeCategory`

NewAssetTypeCategory instantiates a new AssetTypeCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTypeCategoryWithDefaults

`func NewAssetTypeCategoryWithDefaults() *AssetTypeCategory`

NewAssetTypeCategoryWithDefaults instantiates a new AssetTypeCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AssetTypeCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetTypeCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetTypeCategory) SetName(v string)`

SetName sets Name field to given value.


### GetTranslation

`func (o *AssetTypeCategory) GetTranslation() Translation`

GetTranslation returns the Translation field if non-nil, zero value otherwise.

### GetTranslationOk

`func (o *AssetTypeCategory) GetTranslationOk() (*Translation, bool)`

GetTranslationOk returns a tuple with the Translation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslation

`func (o *AssetTypeCategory) SetTranslation(v Translation)`

SetTranslation sets Translation field to given value.

### HasTranslation

`func (o *AssetTypeCategory) HasTranslation() bool`

HasTranslation returns a boolean if a field has been set.

### SetTranslationNil

`func (o *AssetTypeCategory) SetTranslationNil(b bool)`

 SetTranslationNil sets the value for Translation to be an explicit nil

### UnsetTranslation
`func (o *AssetTypeCategory) UnsetTranslation()`

UnsetTranslation ensures that no value is present for Translation, not even an explicit nil
### GetProperties

`func (o *AssetTypeCategory) GetProperties() []AssetTypeCategoryProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AssetTypeCategory) GetPropertiesOk() (*[]AssetTypeCategoryProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AssetTypeCategory) SetProperties(v []AssetTypeCategoryProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AssetTypeCategory) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *AssetTypeCategory) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *AssetTypeCategory) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


