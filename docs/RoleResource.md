# RoleResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Path name of the role resource | [optional] 
**ViewPermission** | Pointer to **NullableBool** | Role has view permission for this resource | [optional] [default to false]
**EditPermission** | Pointer to **NullableBool** | Role has edit permission for this resource | [optional] [default to false]
**ManagePermission** | Pointer to **NullableBool** | Role has manage permission for this resource | [optional] [default to false]

## Methods

### NewRoleResource

`func NewRoleResource() *RoleResource`

NewRoleResource instantiates a new RoleResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleResourceWithDefaults

`func NewRoleResourceWithDefaults() *RoleResource`

NewRoleResourceWithDefaults instantiates a new RoleResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RoleResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RoleResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RoleResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RoleResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetViewPermission

`func (o *RoleResource) GetViewPermission() bool`

GetViewPermission returns the ViewPermission field if non-nil, zero value otherwise.

### GetViewPermissionOk

`func (o *RoleResource) GetViewPermissionOk() (*bool, bool)`

GetViewPermissionOk returns a tuple with the ViewPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPermission

`func (o *RoleResource) SetViewPermission(v bool)`

SetViewPermission sets ViewPermission field to given value.

### HasViewPermission

`func (o *RoleResource) HasViewPermission() bool`

HasViewPermission returns a boolean if a field has been set.

### SetViewPermissionNil

`func (o *RoleResource) SetViewPermissionNil(b bool)`

 SetViewPermissionNil sets the value for ViewPermission to be an explicit nil

### UnsetViewPermission
`func (o *RoleResource) UnsetViewPermission()`

UnsetViewPermission ensures that no value is present for ViewPermission, not even an explicit nil
### GetEditPermission

`func (o *RoleResource) GetEditPermission() bool`

GetEditPermission returns the EditPermission field if non-nil, zero value otherwise.

### GetEditPermissionOk

`func (o *RoleResource) GetEditPermissionOk() (*bool, bool)`

GetEditPermissionOk returns a tuple with the EditPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditPermission

`func (o *RoleResource) SetEditPermission(v bool)`

SetEditPermission sets EditPermission field to given value.

### HasEditPermission

`func (o *RoleResource) HasEditPermission() bool`

HasEditPermission returns a boolean if a field has been set.

### SetEditPermissionNil

`func (o *RoleResource) SetEditPermissionNil(b bool)`

 SetEditPermissionNil sets the value for EditPermission to be an explicit nil

### UnsetEditPermission
`func (o *RoleResource) UnsetEditPermission()`

UnsetEditPermission ensures that no value is present for EditPermission, not even an explicit nil
### GetManagePermission

`func (o *RoleResource) GetManagePermission() bool`

GetManagePermission returns the ManagePermission field if non-nil, zero value otherwise.

### GetManagePermissionOk

`func (o *RoleResource) GetManagePermissionOk() (*bool, bool)`

GetManagePermissionOk returns a tuple with the ManagePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagePermission

`func (o *RoleResource) SetManagePermission(v bool)`

SetManagePermission sets ManagePermission field to given value.

### HasManagePermission

`func (o *RoleResource) HasManagePermission() bool`

HasManagePermission returns a boolean if a field has been set.

### SetManagePermissionNil

`func (o *RoleResource) SetManagePermissionNil(b bool)`

 SetManagePermissionNil sets the value for ManagePermission to be an explicit nil

### UnsetManagePermission
`func (o *RoleResource) UnsetManagePermission()`

UnsetManagePermission ensures that no value is present for ManagePermission, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


