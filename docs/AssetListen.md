# AssetListen

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
**StatusCode** | Pointer to **int32** | The status code expecting when actually perform the operation. Some values are - 200: updated (ok)  - 201: created - 204: deleted (no content) - 304: unchanged (not modified) - 400: problem (bad request) - 404: not found - 409: duplicated (conflict) - 422: unprocessable  | [optional] 

## Methods

### NewAssetListen

`func NewAssetListen(projectId string, globalAssetIdentifier string, assetType string, ) *AssetListen`

NewAssetListen instantiates a new AssetListen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetListenWithDefaults

`func NewAssetListenWithDefaults() *AssetListen`

NewAssetListenWithDefaults instantiates a new AssetListen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *AssetListen) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AssetListen) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AssetListen) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AssetListen) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *AssetListen) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *AssetListen) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetId

`func (o *AssetListen) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetListen) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetListen) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AssetListen) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AssetListen) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AssetListen) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDeviceIds

`func (o *AssetListen) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *AssetListen) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *AssetListen) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *AssetListen) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### SetDeviceIdsNil

`func (o *AssetListen) SetDeviceIdsNil(b bool)`

 SetDeviceIdsNil sets the value for DeviceIds to be an explicit nil

### UnsetDeviceIds
`func (o *AssetListen) UnsetDeviceIds()`

UnsetDeviceIds ensures that no value is present for DeviceIds, not even an explicit nil
### GetProjectId

`func (o *AssetListen) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AssetListen) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AssetListen) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetGlobalAssetIdentifier

`func (o *AssetListen) GetGlobalAssetIdentifier() string`

GetGlobalAssetIdentifier returns the GlobalAssetIdentifier field if non-nil, zero value otherwise.

### GetGlobalAssetIdentifierOk

`func (o *AssetListen) GetGlobalAssetIdentifierOk() (*string, bool)`

GetGlobalAssetIdentifierOk returns a tuple with the GlobalAssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAssetIdentifier

`func (o *AssetListen) SetGlobalAssetIdentifier(v string)`

SetGlobalAssetIdentifier sets GlobalAssetIdentifier field to given value.


### GetName

`func (o *AssetListen) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetListen) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetListen) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetListen) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AssetListen) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AssetListen) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAssetType

`func (o *AssetListen) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AssetListen) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AssetListen) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetLatitude

`func (o *AssetListen) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *AssetListen) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *AssetListen) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *AssetListen) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *AssetListen) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *AssetListen) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *AssetListen) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *AssetListen) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *AssetListen) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *AssetListen) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *AssetListen) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *AssetListen) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetIsTracker

`func (o *AssetListen) GetIsTracker() bool`

GetIsTracker returns the IsTracker field if non-nil, zero value otherwise.

### GetIsTrackerOk

`func (o *AssetListen) GetIsTrackerOk() (*bool, bool)`

GetIsTrackerOk returns a tuple with the IsTracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTracker

`func (o *AssetListen) SetIsTracker(v bool)`

SetIsTracker sets IsTracker field to given value.

### HasIsTracker

`func (o *AssetListen) HasIsTracker() bool`

HasIsTracker returns a boolean if a field has been set.

### SetIsTrackerNil

`func (o *AssetListen) SetIsTrackerNil(b bool)`

 SetIsTrackerNil sets the value for IsTracker to be an explicit nil

### UnsetIsTracker
`func (o *AssetListen) UnsetIsTracker()`

UnsetIsTracker ensures that no value is present for IsTracker, not even an explicit nil
### GetTrackerId

`func (o *AssetListen) GetTrackerId() int32`

GetTrackerId returns the TrackerId field if non-nil, zero value otherwise.

### GetTrackerIdOk

`func (o *AssetListen) GetTrackerIdOk() (*int32, bool)`

GetTrackerIdOk returns a tuple with the TrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackerId

`func (o *AssetListen) SetTrackerId(v int32)`

SetTrackerId sets TrackerId field to given value.

### HasTrackerId

`func (o *AssetListen) HasTrackerId() bool`

HasTrackerId returns a boolean if a field has been set.

### SetTrackerIdNil

`func (o *AssetListen) SetTrackerIdNil(b bool)`

 SetTrackerIdNil sets the value for TrackerId to be an explicit nil

### UnsetTrackerId
`func (o *AssetListen) UnsetTrackerId()`

UnsetTrackerId ensures that no value is present for TrackerId, not even an explicit nil
### GetDescription

`func (o *AssetListen) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetListen) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetListen) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetListen) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AssetListen) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AssetListen) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentFunctionalAssetId

`func (o *AssetListen) GetParentFunctionalAssetId() int32`

GetParentFunctionalAssetId returns the ParentFunctionalAssetId field if non-nil, zero value otherwise.

### GetParentFunctionalAssetIdOk

`func (o *AssetListen) GetParentFunctionalAssetIdOk() (*int32, bool)`

GetParentFunctionalAssetIdOk returns a tuple with the ParentFunctionalAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFunctionalAssetId

`func (o *AssetListen) SetParentFunctionalAssetId(v int32)`

SetParentFunctionalAssetId sets ParentFunctionalAssetId field to given value.

### HasParentFunctionalAssetId

`func (o *AssetListen) HasParentFunctionalAssetId() bool`

HasParentFunctionalAssetId returns a boolean if a field has been set.

### SetParentFunctionalAssetIdNil

`func (o *AssetListen) SetParentFunctionalAssetIdNil(b bool)`

 SetParentFunctionalAssetIdNil sets the value for ParentFunctionalAssetId to be an explicit nil

### UnsetParentFunctionalAssetId
`func (o *AssetListen) UnsetParentFunctionalAssetId()`

UnsetParentFunctionalAssetId ensures that no value is present for ParentFunctionalAssetId, not even an explicit nil
### GetFunctionalAssetIdPath

`func (o *AssetListen) GetFunctionalAssetIdPath() []int32`

GetFunctionalAssetIdPath returns the FunctionalAssetIdPath field if non-nil, zero value otherwise.

### GetFunctionalAssetIdPathOk

`func (o *AssetListen) GetFunctionalAssetIdPathOk() (*[]int32, bool)`

GetFunctionalAssetIdPathOk returns a tuple with the FunctionalAssetIdPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionalAssetIdPath

`func (o *AssetListen) SetFunctionalAssetIdPath(v []int32)`

SetFunctionalAssetIdPath sets FunctionalAssetIdPath field to given value.

### HasFunctionalAssetIdPath

`func (o *AssetListen) HasFunctionalAssetIdPath() bool`

HasFunctionalAssetIdPath returns a boolean if a field has been set.

### SetFunctionalAssetIdPathNil

`func (o *AssetListen) SetFunctionalAssetIdPathNil(b bool)`

 SetFunctionalAssetIdPathNil sets the value for FunctionalAssetIdPath to be an explicit nil

### UnsetFunctionalAssetIdPath
`func (o *AssetListen) UnsetFunctionalAssetIdPath()`

UnsetFunctionalAssetIdPath ensures that no value is present for FunctionalAssetIdPath, not even an explicit nil
### GetParentLocationalAssetId

`func (o *AssetListen) GetParentLocationalAssetId() int32`

GetParentLocationalAssetId returns the ParentLocationalAssetId field if non-nil, zero value otherwise.

### GetParentLocationalAssetIdOk

`func (o *AssetListen) GetParentLocationalAssetIdOk() (*int32, bool)`

GetParentLocationalAssetIdOk returns a tuple with the ParentLocationalAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLocationalAssetId

`func (o *AssetListen) SetParentLocationalAssetId(v int32)`

SetParentLocationalAssetId sets ParentLocationalAssetId field to given value.

### HasParentLocationalAssetId

`func (o *AssetListen) HasParentLocationalAssetId() bool`

HasParentLocationalAssetId returns a boolean if a field has been set.

### SetParentLocationalAssetIdNil

`func (o *AssetListen) SetParentLocationalAssetIdNil(b bool)`

 SetParentLocationalAssetIdNil sets the value for ParentLocationalAssetId to be an explicit nil

### UnsetParentLocationalAssetId
`func (o *AssetListen) UnsetParentLocationalAssetId()`

UnsetParentLocationalAssetId ensures that no value is present for ParentLocationalAssetId, not even an explicit nil
### GetLocationalAssetIdPath

`func (o *AssetListen) GetLocationalAssetIdPath() []int32`

GetLocationalAssetIdPath returns the LocationalAssetIdPath field if non-nil, zero value otherwise.

### GetLocationalAssetIdPathOk

`func (o *AssetListen) GetLocationalAssetIdPathOk() (*[]int32, bool)`

GetLocationalAssetIdPathOk returns a tuple with the LocationalAssetIdPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationalAssetIdPath

`func (o *AssetListen) SetLocationalAssetIdPath(v []int32)`

SetLocationalAssetIdPath sets LocationalAssetIdPath field to given value.

### HasLocationalAssetIdPath

`func (o *AssetListen) HasLocationalAssetIdPath() bool`

HasLocationalAssetIdPath returns a boolean if a field has been set.

### SetLocationalAssetIdPathNil

`func (o *AssetListen) SetLocationalAssetIdPathNil(b bool)`

 SetLocationalAssetIdPathNil sets the value for LocationalAssetIdPath to be an explicit nil

### UnsetLocationalAssetIdPath
`func (o *AssetListen) UnsetLocationalAssetIdPath()`

UnsetLocationalAssetIdPath ensures that no value is present for LocationalAssetIdPath, not even an explicit nil
### GetParentFunctionalIdentifier

`func (o *AssetListen) GetParentFunctionalIdentifier() string`

GetParentFunctionalIdentifier returns the ParentFunctionalIdentifier field if non-nil, zero value otherwise.

### GetParentFunctionalIdentifierOk

`func (o *AssetListen) GetParentFunctionalIdentifierOk() (*string, bool)`

GetParentFunctionalIdentifierOk returns a tuple with the ParentFunctionalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFunctionalIdentifier

`func (o *AssetListen) SetParentFunctionalIdentifier(v string)`

SetParentFunctionalIdentifier sets ParentFunctionalIdentifier field to given value.

### HasParentFunctionalIdentifier

`func (o *AssetListen) HasParentFunctionalIdentifier() bool`

HasParentFunctionalIdentifier returns a boolean if a field has been set.

### SetParentFunctionalIdentifierNil

`func (o *AssetListen) SetParentFunctionalIdentifierNil(b bool)`

 SetParentFunctionalIdentifierNil sets the value for ParentFunctionalIdentifier to be an explicit nil

### UnsetParentFunctionalIdentifier
`func (o *AssetListen) UnsetParentFunctionalIdentifier()`

UnsetParentFunctionalIdentifier ensures that no value is present for ParentFunctionalIdentifier, not even an explicit nil
### GetParentLocationalIdentifier

`func (o *AssetListen) GetParentLocationalIdentifier() string`

GetParentLocationalIdentifier returns the ParentLocationalIdentifier field if non-nil, zero value otherwise.

### GetParentLocationalIdentifierOk

`func (o *AssetListen) GetParentLocationalIdentifierOk() (*string, bool)`

GetParentLocationalIdentifierOk returns a tuple with the ParentLocationalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLocationalIdentifier

`func (o *AssetListen) SetParentLocationalIdentifier(v string)`

SetParentLocationalIdentifier sets ParentLocationalIdentifier field to given value.

### HasParentLocationalIdentifier

`func (o *AssetListen) HasParentLocationalIdentifier() bool`

HasParentLocationalIdentifier returns a boolean if a field has been set.

### SetParentLocationalIdentifierNil

`func (o *AssetListen) SetParentLocationalIdentifierNil(b bool)`

 SetParentLocationalIdentifierNil sets the value for ParentLocationalIdentifier to be an explicit nil

### UnsetParentLocationalIdentifier
`func (o *AssetListen) UnsetParentLocationalIdentifier()`

UnsetParentLocationalIdentifier ensures that no value is present for ParentLocationalIdentifier, not even an explicit nil
### GetTags

`func (o *AssetListen) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetListen) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetListen) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssetListen) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AssetListen) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AssetListen) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAllowedInactivity

`func (o *AssetListen) GetAllowedInactivity() string`

GetAllowedInactivity returns the AllowedInactivity field if non-nil, zero value otherwise.

### GetAllowedInactivityOk

`func (o *AssetListen) GetAllowedInactivityOk() (*string, bool)`

GetAllowedInactivityOk returns a tuple with the AllowedInactivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedInactivity

`func (o *AssetListen) SetAllowedInactivity(v string)`

SetAllowedInactivity sets AllowedInactivity field to given value.

### HasAllowedInactivity

`func (o *AssetListen) HasAllowedInactivity() bool`

HasAllowedInactivity returns a boolean if a field has been set.

### SetAllowedInactivityNil

`func (o *AssetListen) SetAllowedInactivityNil(b bool)`

 SetAllowedInactivityNil sets the value for AllowedInactivity to be an explicit nil

### UnsetAllowedInactivity
`func (o *AssetListen) UnsetAllowedInactivity()`

UnsetAllowedInactivity ensures that no value is present for AllowedInactivity, not even an explicit nil
### GetChildrenInfo

`func (o *AssetListen) GetChildrenInfo() []Asset`

GetChildrenInfo returns the ChildrenInfo field if non-nil, zero value otherwise.

### GetChildrenInfoOk

`func (o *AssetListen) GetChildrenInfoOk() (*[]Asset, bool)`

GetChildrenInfoOk returns a tuple with the ChildrenInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenInfo

`func (o *AssetListen) SetChildrenInfo(v []Asset)`

SetChildrenInfo sets ChildrenInfo field to given value.

### HasChildrenInfo

`func (o *AssetListen) HasChildrenInfo() bool`

HasChildrenInfo returns a boolean if a field has been set.

### SetChildrenInfoNil

`func (o *AssetListen) SetChildrenInfoNil(b bool)`

 SetChildrenInfoNil sets the value for ChildrenInfo to be an explicit nil

### UnsetChildrenInfo
`func (o *AssetListen) UnsetChildrenInfo()`

UnsetChildrenInfo ensures that no value is present for ChildrenInfo, not even an explicit nil
### GetAttachments

`func (o *AssetListen) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AssetListen) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AssetListen) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *AssetListen) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *AssetListen) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *AssetListen) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetStatusCode

`func (o *AssetListen) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *AssetListen) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *AssetListen) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *AssetListen) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


