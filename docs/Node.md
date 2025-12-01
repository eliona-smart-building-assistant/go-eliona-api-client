# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Unique id for the edge node | [optional] [readonly] 
**SiteId** | Pointer to **NullableString** | ID of the site to which the node belongs | [optional] 
**Vendor** | Pointer to **NullableString** | Vendor name | [optional] 
**Model** | Pointer to **NullableString** | Model name | [optional] 
**Description** | Pointer to **NullableString** | Descriptive text for the edge node | [optional] 
**Enable** | Pointer to **bool** | Is the node enabled or not | [optional] [default to false]
**AssetId** | Pointer to **NullableInt32** | ID of the corresponding asset | [optional] 
**Ident** | Pointer to **NullableString** | UUID to identify the edge node | [optional] [readonly] 
**Password** | Pointer to **NullableString** | Password with which the node identifies itself | [optional] 

## Methods

### NewNode

`func NewNode() *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Node) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Node) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Node) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Node) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Node) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Node) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSiteId

`func (o *Node) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Node) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Node) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Node) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *Node) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *Node) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetVendor

`func (o *Node) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Node) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Node) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Node) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *Node) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *Node) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
### GetModel

`func (o *Node) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Node) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Node) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Node) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *Node) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *Node) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetDescription

`func (o *Node) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Node) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Node) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Node) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Node) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Node) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnable

`func (o *Node) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Node) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Node) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *Node) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetAssetId

`func (o *Node) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Node) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Node) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Node) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *Node) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *Node) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
### GetIdent

`func (o *Node) GetIdent() string`

GetIdent returns the Ident field if non-nil, zero value otherwise.

### GetIdentOk

`func (o *Node) GetIdentOk() (*string, bool)`

GetIdentOk returns a tuple with the Ident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdent

`func (o *Node) SetIdent(v string)`

SetIdent sets Ident field to given value.

### HasIdent

`func (o *Node) HasIdent() bool`

HasIdent returns a boolean if a field has been set.

### SetIdentNil

`func (o *Node) SetIdentNil(b bool)`

 SetIdentNil sets the value for Ident to be an explicit nil

### UnsetIdent
`func (o *Node) UnsetIdent()`

UnsetIdent ensures that no value is present for Ident, not even an explicit nil
### GetPassword

`func (o *Node) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Node) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Node) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Node) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *Node) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *Node) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


