# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | The internal Id of dashboard | [optional] [readonly] 
**Name** | **string** | The name for this dashboard | 
**UserId** | **string** | ID of the user who owns the dashboard | 
**Sequence** | Pointer to **NullableInt32** | The sequence of the dashboard | [optional] [default to 0]
**Widgets** | Pointer to [**[]Widget**](Widget.md) | List of widgets on this dashboard (order matches the order of widgets on the dashboard) | [optional] 
**Public** | Pointer to **NullableBool** | Is the dashboard public and not bound to a dedicated user | [optional] [default to false]
**ProjectId** | Pointer to **NullableString** | ID of the project to which the dashboard belongs. The project ID is deprecated due to the removal of projects and is broadly replaced by sites. This filter now uses siteId, which is internally represented as a UUID. | [optional] 

## Methods

### NewDashboard

`func NewDashboard(name string, userId string, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dashboard) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dashboard) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dashboard) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Dashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Dashboard) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Dashboard) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Dashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dashboard) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *Dashboard) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Dashboard) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Dashboard) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSequence

`func (o *Dashboard) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *Dashboard) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *Dashboard) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *Dashboard) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequenceNil

`func (o *Dashboard) SetSequenceNil(b bool)`

 SetSequenceNil sets the value for Sequence to be an explicit nil

### UnsetSequence
`func (o *Dashboard) UnsetSequence()`

UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
### GetWidgets

`func (o *Dashboard) GetWidgets() []Widget`

GetWidgets returns the Widgets field if non-nil, zero value otherwise.

### GetWidgetsOk

`func (o *Dashboard) GetWidgetsOk() (*[]Widget, bool)`

GetWidgetsOk returns a tuple with the Widgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgets

`func (o *Dashboard) SetWidgets(v []Widget)`

SetWidgets sets Widgets field to given value.

### HasWidgets

`func (o *Dashboard) HasWidgets() bool`

HasWidgets returns a boolean if a field has been set.

### SetWidgetsNil

`func (o *Dashboard) SetWidgetsNil(b bool)`

 SetWidgetsNil sets the value for Widgets to be an explicit nil

### UnsetWidgets
`func (o *Dashboard) UnsetWidgets()`

UnsetWidgets ensures that no value is present for Widgets, not even an explicit nil
### GetPublic

`func (o *Dashboard) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Dashboard) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Dashboard) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *Dashboard) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### SetPublicNil

`func (o *Dashboard) SetPublicNil(b bool)`

 SetPublicNil sets the value for Public to be an explicit nil

### UnsetPublic
`func (o *Dashboard) UnsetPublic()`

UnsetPublic ensures that no value is present for Public, not even an explicit nil
### GetProjectId

`func (o *Dashboard) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Dashboard) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Dashboard) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Dashboard) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Dashboard) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Dashboard) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


