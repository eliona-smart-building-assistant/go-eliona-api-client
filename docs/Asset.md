# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **NullableString** | A custom identifier for the resource, which can be utilized to identify it for various operations such as updates, deletions, and other related tasks. If this identifier is not provided, the application will generate a new value for it automatically | [optional] 
**Id** | Pointer to **NullableInt32** | The internal Id of asset | [optional] [readonly] 
**DeviceIds** | Pointer to **[]string** | A list of unique device ids | [optional] 
**ProjectId** | **string** | ID of the project to which the asset belongs | 
**GlobalAssetIdentifier** | **string** | Unique identifier for the asset | 
**Name** | Pointer to **NullableString** | Alternate text for the asset to display in frontend | [optional] 
**AssetType** | **string** | Reference to asset type by name | 
**Latitude** | Pointer to **NullableFloat64** | Latitude coordinate (GPS) of the asset | [optional] 
**Longitude** | Pointer to **NullableFloat64** | Longitude coordinate (GPS) of the asset | [optional] 
**IsTracker** | Pointer to **NullableBool** | Does the asset function as a tracker and capture its position by itself | [optional] [default to false]
**TrackerId** | Pointer to **NullableInt32** | Tracker Asset Id | [optional] 
**Description** | Pointer to **NullableString** | Textual description for this asset | [optional] 
**ParentFunctionalAssetId** | Pointer to **NullableInt32** | The id of an asset which groups this asset as a functional child | [optional] 
**FunctionalAssetIdPath** | Pointer to **[]int32** | The hierarchical path of functional ids of the asset | [optional] [readonly] 
**ParentLocationalAssetId** | Pointer to **NullableInt32** | The id of an asset which groups this asset as a locational child | [optional] 
**LocationalAssetIdPath** | Pointer to **[]int32** | The hierarchical path of locational ids of the asset | [optional] [readonly] 
**ParentFunctionalIdentifier** | Pointer to **NullableString** | The identifier specified by the identifiedBy parameter classifies this asset as a functional child. | [optional] 
**ParentLocationalIdentifier** | Pointer to **NullableString** | The identifier specified by the identifiedBy parameter classifies this asset as a locational child | [optional] 
**Tags** | Pointer to **[]string** | List of associated tags | [optional] 
**AllowedInactivity** | Pointer to **NullableString** |  | [optional] 
**ChildrenInfo** | Pointer to [**[]Asset**](Asset.md) | List of children for this asset. | [optional] [readonly] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | A list of files attached to the asset | [optional] 

## Methods

### NewAsset

`func NewAsset(projectId string, globalAssetIdentifier string, assetType string, ) *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *Asset) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Asset) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Asset) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Asset) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *Asset) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *Asset) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetId

`func (o *Asset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Asset) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Asset) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Asset) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Asset) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDeviceIds

`func (o *Asset) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *Asset) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *Asset) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *Asset) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### SetDeviceIdsNil

`func (o *Asset) SetDeviceIdsNil(b bool)`

 SetDeviceIdsNil sets the value for DeviceIds to be an explicit nil

### UnsetDeviceIds
`func (o *Asset) UnsetDeviceIds()`

UnsetDeviceIds ensures that no value is present for DeviceIds, not even an explicit nil
### GetProjectId

`func (o *Asset) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Asset) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Asset) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetGlobalAssetIdentifier

`func (o *Asset) GetGlobalAssetIdentifier() string`

GetGlobalAssetIdentifier returns the GlobalAssetIdentifier field if non-nil, zero value otherwise.

### GetGlobalAssetIdentifierOk

`func (o *Asset) GetGlobalAssetIdentifierOk() (*string, bool)`

GetGlobalAssetIdentifierOk returns a tuple with the GlobalAssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAssetIdentifier

`func (o *Asset) SetGlobalAssetIdentifier(v string)`

SetGlobalAssetIdentifier sets GlobalAssetIdentifier field to given value.


### GetName

`func (o *Asset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Asset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Asset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Asset) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Asset) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Asset) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAssetType

`func (o *Asset) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *Asset) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *Asset) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetLatitude

`func (o *Asset) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Asset) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Asset) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Asset) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *Asset) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *Asset) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *Asset) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Asset) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Asset) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Asset) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *Asset) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *Asset) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetIsTracker

`func (o *Asset) GetIsTracker() bool`

GetIsTracker returns the IsTracker field if non-nil, zero value otherwise.

### GetIsTrackerOk

`func (o *Asset) GetIsTrackerOk() (*bool, bool)`

GetIsTrackerOk returns a tuple with the IsTracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTracker

`func (o *Asset) SetIsTracker(v bool)`

SetIsTracker sets IsTracker field to given value.

### HasIsTracker

`func (o *Asset) HasIsTracker() bool`

HasIsTracker returns a boolean if a field has been set.

### SetIsTrackerNil

`func (o *Asset) SetIsTrackerNil(b bool)`

 SetIsTrackerNil sets the value for IsTracker to be an explicit nil

### UnsetIsTracker
`func (o *Asset) UnsetIsTracker()`

UnsetIsTracker ensures that no value is present for IsTracker, not even an explicit nil
### GetTrackerId

`func (o *Asset) GetTrackerId() int32`

GetTrackerId returns the TrackerId field if non-nil, zero value otherwise.

### GetTrackerIdOk

`func (o *Asset) GetTrackerIdOk() (*int32, bool)`

GetTrackerIdOk returns a tuple with the TrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackerId

`func (o *Asset) SetTrackerId(v int32)`

SetTrackerId sets TrackerId field to given value.

### HasTrackerId

`func (o *Asset) HasTrackerId() bool`

HasTrackerId returns a boolean if a field has been set.

### SetTrackerIdNil

`func (o *Asset) SetTrackerIdNil(b bool)`

 SetTrackerIdNil sets the value for TrackerId to be an explicit nil

### UnsetTrackerId
`func (o *Asset) UnsetTrackerId()`

UnsetTrackerId ensures that no value is present for TrackerId, not even an explicit nil
### GetDescription

`func (o *Asset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Asset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Asset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Asset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Asset) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Asset) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentFunctionalAssetId

`func (o *Asset) GetParentFunctionalAssetId() int32`

GetParentFunctionalAssetId returns the ParentFunctionalAssetId field if non-nil, zero value otherwise.

### GetParentFunctionalAssetIdOk

`func (o *Asset) GetParentFunctionalAssetIdOk() (*int32, bool)`

GetParentFunctionalAssetIdOk returns a tuple with the ParentFunctionalAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFunctionalAssetId

`func (o *Asset) SetParentFunctionalAssetId(v int32)`

SetParentFunctionalAssetId sets ParentFunctionalAssetId field to given value.

### HasParentFunctionalAssetId

`func (o *Asset) HasParentFunctionalAssetId() bool`

HasParentFunctionalAssetId returns a boolean if a field has been set.

### SetParentFunctionalAssetIdNil

`func (o *Asset) SetParentFunctionalAssetIdNil(b bool)`

 SetParentFunctionalAssetIdNil sets the value for ParentFunctionalAssetId to be an explicit nil

### UnsetParentFunctionalAssetId
`func (o *Asset) UnsetParentFunctionalAssetId()`

UnsetParentFunctionalAssetId ensures that no value is present for ParentFunctionalAssetId, not even an explicit nil
### GetFunctionalAssetIdPath

`func (o *Asset) GetFunctionalAssetIdPath() []int32`

GetFunctionalAssetIdPath returns the FunctionalAssetIdPath field if non-nil, zero value otherwise.

### GetFunctionalAssetIdPathOk

`func (o *Asset) GetFunctionalAssetIdPathOk() (*[]int32, bool)`

GetFunctionalAssetIdPathOk returns a tuple with the FunctionalAssetIdPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalAssetIdPath

`func (o *Asset) SetFunctionalAssetIdPath(v []int32)`

SetFunctionalAssetIdPath sets FunctionalAssetIdPath field to given value.

### HasFunctionalAssetIdPath

`func (o *Asset) HasFunctionalAssetIdPath() bool`

HasFunctionalAssetIdPath returns a boolean if a field has been set.

### SetFunctionalAssetIdPathNil

`func (o *Asset) SetFunctionalAssetIdPathNil(b bool)`

 SetFunctionalAssetIdPathNil sets the value for FunctionalAssetIdPath to be an explicit nil

### UnsetFunctionalAssetIdPath
`func (o *Asset) UnsetFunctionalAssetIdPath()`

UnsetFunctionalAssetIdPath ensures that no value is present for FunctionalAssetIdPath, not even an explicit nil
### GetParentLocationalAssetId

`func (o *Asset) GetParentLocationalAssetId() int32`

GetParentLocationalAssetId returns the ParentLocationalAssetId field if non-nil, zero value otherwise.

### GetParentLocationalAssetIdOk

`func (o *Asset) GetParentLocationalAssetIdOk() (*int32, bool)`

GetParentLocationalAssetIdOk returns a tuple with the ParentLocationalAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLocationalAssetId

`func (o *Asset) SetParentLocationalAssetId(v int32)`

SetParentLocationalAssetId sets ParentLocationalAssetId field to given value.

### HasParentLocationalAssetId

`func (o *Asset) HasParentLocationalAssetId() bool`

HasParentLocationalAssetId returns a boolean if a field has been set.

### SetParentLocationalAssetIdNil

`func (o *Asset) SetParentLocationalAssetIdNil(b bool)`

 SetParentLocationalAssetIdNil sets the value for ParentLocationalAssetId to be an explicit nil

### UnsetParentLocationalAssetId
`func (o *Asset) UnsetParentLocationalAssetId()`

UnsetParentLocationalAssetId ensures that no value is present for ParentLocationalAssetId, not even an explicit nil
### GetLocationalAssetIdPath

`func (o *Asset) GetLocationalAssetIdPath() []int32`

GetLocationalAssetIdPath returns the LocationalAssetIdPath field if non-nil, zero value otherwise.

### GetLocationalAssetIdPathOk

`func (o *Asset) GetLocationalAssetIdPathOk() (*[]int32, bool)`

GetLocationalAssetIdPathOk returns a tuple with the LocationalAssetIdPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationalAssetIdPath

`func (o *Asset) SetLocationalAssetIdPath(v []int32)`

SetLocationalAssetIdPath sets LocationalAssetIdPath field to given value.

### HasLocationalAssetIdPath

`func (o *Asset) HasLocationalAssetIdPath() bool`

HasLocationalAssetIdPath returns a boolean if a field has been set.

### SetLocationalAssetIdPathNil

`func (o *Asset) SetLocationalAssetIdPathNil(b bool)`

 SetLocationalAssetIdPathNil sets the value for LocationalAssetIdPath to be an explicit nil

### UnsetLocationalAssetIdPath
`func (o *Asset) UnsetLocationalAssetIdPath()`

UnsetLocationalAssetIdPath ensures that no value is present for LocationalAssetIdPath, not even an explicit nil
### GetParentFunctionalIdentifier

`func (o *Asset) GetParentFunctionalIdentifier() string`

GetParentFunctionalIdentifier returns the ParentFunctionalIdentifier field if non-nil, zero value otherwise.

### GetParentFunctionalIdentifierOk

`func (o *Asset) GetParentFunctionalIdentifierOk() (*string, bool)`

GetParentFunctionalIdentifierOk returns a tuple with the ParentFunctionalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFunctionalIdentifier

`func (o *Asset) SetParentFunctionalIdentifier(v string)`

SetParentFunctionalIdentifier sets ParentFunctionalIdentifier field to given value.

### HasParentFunctionalIdentifier

`func (o *Asset) HasParentFunctionalIdentifier() bool`

HasParentFunctionalIdentifier returns a boolean if a field has been set.

### SetParentFunctionalIdentifierNil

`func (o *Asset) SetParentFunctionalIdentifierNil(b bool)`

 SetParentFunctionalIdentifierNil sets the value for ParentFunctionalIdentifier to be an explicit nil

### UnsetParentFunctionalIdentifier
`func (o *Asset) UnsetParentFunctionalIdentifier()`

UnsetParentFunctionalIdentifier ensures that no value is present for ParentFunctionalIdentifier, not even an explicit nil
### GetParentLocationalIdentifier

`func (o *Asset) GetParentLocationalIdentifier() string`

GetParentLocationalIdentifier returns the ParentLocationalIdentifier field if non-nil, zero value otherwise.

### GetParentLocationalIdentifierOk

`func (o *Asset) GetParentLocationalIdentifierOk() (*string, bool)`

GetParentLocationalIdentifierOk returns a tuple with the ParentLocationalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLocationalIdentifier

`func (o *Asset) SetParentLocationalIdentifier(v string)`

SetParentLocationalIdentifier sets ParentLocationalIdentifier field to given value.

### HasParentLocationalIdentifier

`func (o *Asset) HasParentLocationalIdentifier() bool`

HasParentLocationalIdentifier returns a boolean if a field has been set.

### SetParentLocationalIdentifierNil

`func (o *Asset) SetParentLocationalIdentifierNil(b bool)`

 SetParentLocationalIdentifierNil sets the value for ParentLocationalIdentifier to be an explicit nil

### UnsetParentLocationalIdentifier
`func (o *Asset) UnsetParentLocationalIdentifier()`

UnsetParentLocationalIdentifier ensures that no value is present for ParentLocationalIdentifier, not even an explicit nil
### GetTags

`func (o *Asset) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Asset) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Asset) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Asset) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Asset) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Asset) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAllowedInactivity

`func (o *Asset) GetAllowedInactivity() string`

GetAllowedInactivity returns the AllowedInactivity field if non-nil, zero value otherwise.

### GetAllowedInactivityOk

`func (o *Asset) GetAllowedInactivityOk() (*string, bool)`

GetAllowedInactivityOk returns a tuple with the AllowedInactivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInactivity

`func (o *Asset) SetAllowedInactivity(v string)`

SetAllowedInactivity sets AllowedInactivity field to given value.

### HasAllowedInactivity

`func (o *Asset) HasAllowedInactivity() bool`

HasAllowedInactivity returns a boolean if a field has been set.

### SetAllowedInactivityNil

`func (o *Asset) SetAllowedInactivityNil(b bool)`

 SetAllowedInactivityNil sets the value for AllowedInactivity to be an explicit nil

### UnsetAllowedInactivity
`func (o *Asset) UnsetAllowedInactivity()`

UnsetAllowedInactivity ensures that no value is present for AllowedInactivity, not even an explicit nil
### GetChildrenInfo

`func (o *Asset) GetChildrenInfo() []Asset`

GetChildrenInfo returns the ChildrenInfo field if non-nil, zero value otherwise.

### GetChildrenInfoOk

`func (o *Asset) GetChildrenInfoOk() (*[]Asset, bool)`

GetChildrenInfoOk returns a tuple with the ChildrenInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenInfo

`func (o *Asset) SetChildrenInfo(v []Asset)`

SetChildrenInfo sets ChildrenInfo field to given value.

### HasChildrenInfo

`func (o *Asset) HasChildrenInfo() bool`

HasChildrenInfo returns a boolean if a field has been set.

### SetChildrenInfoNil

`func (o *Asset) SetChildrenInfoNil(b bool)`

 SetChildrenInfoNil sets the value for ChildrenInfo to be an explicit nil

### UnsetChildrenInfo
`func (o *Asset) UnsetChildrenInfo()`

UnsetChildrenInfo ensures that no value is present for ChildrenInfo, not even an explicit nil
### GetAttachments

`func (o *Asset) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Asset) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Asset) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Asset) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *Asset) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *Asset) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


