# AssetType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name for this asset type | 
**Custom** | Pointer to **bool** | Is this a customer created type or not | [optional] [default to true]
**Vendor** | Pointer to **NullableString** | The vendor providing assets of this type | [optional] 
**Model** | Pointer to **NullableString** | The specific model of assets of this type | [optional] 
**Translation** | Pointer to [**NullableTranslation**](Translation.md) |  | [optional] 
**Urldoc** | Pointer to **NullableString** | The url describing assets of this type | [optional] 
**Icon** | Pointer to **NullableString** | Icon name corresponding to assets of this type: blind, building, button, closable, elevator, environment, fallback, filling, gateway, light, mailbox, parking, people, power, rack, storey, trash, ventilation, vibration, water, weather | [optional] 
**PayloadFunction** | Pointer to **NullableString** | Asset types payload function | [optional] 
**AllowedInactivity** | Pointer to **NullableString** |  | [optional] 
**Attributes** | Pointer to [**[]AssetTypeAttribute**](AssetTypeAttribute.md) | List of named attributes | [optional] 

## Methods

### NewAssetType

`func NewAssetType(name string, ) *AssetType`

NewAssetType instantiates a new AssetType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTypeWithDefaults

`func NewAssetTypeWithDefaults() *AssetType`

NewAssetTypeWithDefaults instantiates a new AssetType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AssetType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetType) SetName(v string)`

SetName sets Name field to given value.


### GetCustom

`func (o *AssetType) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *AssetType) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *AssetType) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *AssetType) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetVendor

`func (o *AssetType) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AssetType) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AssetType) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AssetType) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *AssetType) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *AssetType) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
### GetModel

`func (o *AssetType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AssetType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AssetType) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AssetType) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *AssetType) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *AssetType) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetTranslation

`func (o *AssetType) GetTranslation() Translation`

GetTranslation returns the Translation field if non-nil, zero value otherwise.

### GetTranslationOk

`func (o *AssetType) GetTranslationOk() (*Translation, bool)`

GetTranslationOk returns a tuple with the Translation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslation

`func (o *AssetType) SetTranslation(v Translation)`

SetTranslation sets Translation field to given value.

### HasTranslation

`func (o *AssetType) HasTranslation() bool`

HasTranslation returns a boolean if a field has been set.

### SetTranslationNil

`func (o *AssetType) SetTranslationNil(b bool)`

 SetTranslationNil sets the value for Translation to be an explicit nil

### UnsetTranslation
`func (o *AssetType) UnsetTranslation()`

UnsetTranslation ensures that no value is present for Translation, not even an explicit nil
### GetUrldoc

`func (o *AssetType) GetUrldoc() string`

GetUrldoc returns the Urldoc field if non-nil, zero value otherwise.

### GetUrldocOk

`func (o *AssetType) GetUrldocOk() (*string, bool)`

GetUrldocOk returns a tuple with the Urldoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrldoc

`func (o *AssetType) SetUrldoc(v string)`

SetUrldoc sets Urldoc field to given value.

### HasUrldoc

`func (o *AssetType) HasUrldoc() bool`

HasUrldoc returns a boolean if a field has been set.

### SetUrldocNil

`func (o *AssetType) SetUrldocNil(b bool)`

 SetUrldocNil sets the value for Urldoc to be an explicit nil

### UnsetUrldoc
`func (o *AssetType) UnsetUrldoc()`

UnsetUrldoc ensures that no value is present for Urldoc, not even an explicit nil
### GetIcon

`func (o *AssetType) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AssetType) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AssetType) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AssetType) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *AssetType) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *AssetType) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetPayloadFunction

`func (o *AssetType) GetPayloadFunction() string`

GetPayloadFunction returns the PayloadFunction field if non-nil, zero value otherwise.

### GetPayloadFunctionOk

`func (o *AssetType) GetPayloadFunctionOk() (*string, bool)`

GetPayloadFunctionOk returns a tuple with the PayloadFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadFunction

`func (o *AssetType) SetPayloadFunction(v string)`

SetPayloadFunction sets PayloadFunction field to given value.

### HasPayloadFunction

`func (o *AssetType) HasPayloadFunction() bool`

HasPayloadFunction returns a boolean if a field has been set.

### SetPayloadFunctionNil

`func (o *AssetType) SetPayloadFunctionNil(b bool)`

 SetPayloadFunctionNil sets the value for PayloadFunction to be an explicit nil

### UnsetPayloadFunction
`func (o *AssetType) UnsetPayloadFunction()`

UnsetPayloadFunction ensures that no value is present for PayloadFunction, not even an explicit nil
### GetAllowedInactivity

`func (o *AssetType) GetAllowedInactivity() string`

GetAllowedInactivity returns the AllowedInactivity field if non-nil, zero value otherwise.

### GetAllowedInactivityOk

`func (o *AssetType) GetAllowedInactivityOk() (*string, bool)`

GetAllowedInactivityOk returns a tuple with the AllowedInactivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInactivity

`func (o *AssetType) SetAllowedInactivity(v string)`

SetAllowedInactivity sets AllowedInactivity field to given value.

### HasAllowedInactivity

`func (o *AssetType) HasAllowedInactivity() bool`

HasAllowedInactivity returns a boolean if a field has been set.

### SetAllowedInactivityNil

`func (o *AssetType) SetAllowedInactivityNil(b bool)`

 SetAllowedInactivityNil sets the value for AllowedInactivity to be an explicit nil

### UnsetAllowedInactivity
`func (o *AssetType) UnsetAllowedInactivity()`

UnsetAllowedInactivity ensures that no value is present for AllowedInactivity, not even an explicit nil
### GetAttributes

`func (o *AssetType) GetAttributes() []AssetTypeAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AssetType) GetAttributesOk() (*[]AssetTypeAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AssetType) SetAttributes(v []AssetTypeAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AssetType) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *AssetType) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *AssetType) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


