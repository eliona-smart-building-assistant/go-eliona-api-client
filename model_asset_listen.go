/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.7.3
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AssetListen type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetListen{}

// AssetListen struct for AssetListen
type AssetListen struct {
	// A custom identifier for the resource, which can be utilized to identify it for various operations such as updates, deletions, and other related tasks. If this identifier is not provided, the application will generate a new value for it automatically
	ResourceId NullableString `json:"resourceId,omitempty"`
	// The internal Id of asset
	Id NullableInt32 `json:"id,omitempty"`
	// A list of unique device ids
	DeviceIds []string `json:"deviceIds,omitempty"`
	// ID of the project to which the asset belongs
	ProjectId string `json:"projectId"`
	// Unique identifier for the asset
	GlobalAssetIdentifier string `json:"globalAssetIdentifier"`
	// Alternate text for the asset to display in frontend
	Name NullableString `json:"name,omitempty"`
	// Reference to asset type by name
	AssetType string `json:"assetType"`
	// Latitude coordinate (GPS) of the asset
	Latitude NullableFloat64 `json:"latitude,omitempty"`
	// Longitude coordinate (GPS) of the asset
	Longitude NullableFloat64 `json:"longitude,omitempty"`
	// Does the asset function as a tracker and capture its position by itself
	// Deprecated
	IsTracker NullableBool `json:"isTracker,omitempty"`
	// Tracker Asset Id
	TrackerId NullableInt32 `json:"trackerId,omitempty"`
	// Textual description for this asset
	Description NullableString `json:"description,omitempty"`
	// The id of an asset which groups this asset as a functional child
	ParentFunctionalAssetId NullableInt32 `json:"parentFunctionalAssetId,omitempty"`
	// The hierarchical path of functional ids of the asset
	FunctionalAssetIdPath []int32 `json:"functionalAssetIdPath,omitempty"`
	// The id of an asset which groups this asset as a locational child
	ParentLocationalAssetId NullableInt32 `json:"parentLocationalAssetId,omitempty"`
	// The hierarchical path of locational ids of the asset
	LocationalAssetIdPath []int32 `json:"locationalAssetIdPath,omitempty"`
	// The identifier specified by the identifiedBy parameter classifies this asset as a functional child.
	ParentFunctionalIdentifier NullableString `json:"parentFunctionalIdentifier,omitempty"`
	// The identifier specified by the identifiedBy parameter classifies this asset as a locational child
	ParentLocationalIdentifier NullableString `json:"parentLocationalIdentifier,omitempty"`
	// List of associated tags
	Tags []string `json:"tags,omitempty"`
	// List of children for this asset.
	ChildrenInfo []Asset `json:"childrenInfo,omitempty"`
	// A list of files attached to the asset
	Attachments []Attachment `json:"attachments,omitempty"`
	// The status code expecting when actually perform the operation. Some values are - 200: updated (ok)  - 201: created - 204: deleted (no content) - 304: unchanged (not modified) - 400: problem (bad request) - 404: not found - 409: duplicated (conflict) - 422: unprocessable
	StatusCode *int32 `json:"statusCode,omitempty"`
}

type _AssetListen AssetListen

// NewAssetListen instantiates a new AssetListen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetListen(projectId string, globalAssetIdentifier string, assetType string) *AssetListen {
	this := AssetListen{}
	this.ProjectId = projectId
	this.GlobalAssetIdentifier = globalAssetIdentifier
	this.AssetType = assetType
	var isTracker bool = false
	this.IsTracker = *NewNullableBool(&isTracker)
	return &this
}

// NewAssetListenWithDefaults instantiates a new AssetListen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetListenWithDefaults() *AssetListen {
	this := AssetListen{}
	var isTracker bool = false
	this.IsTracker = *NewNullableBool(&isTracker)
	return &this
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceId.Get()
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceId.Get(), o.ResourceId.IsSet()
}

// HasResourceId returns a boolean if a field has been set.
func (o *AssetListen) HasResourceId() bool {
	if o != nil && o.ResourceId.IsSet() {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given NullableString and assigns it to the ResourceId field.
func (o *AssetListen) SetResourceId(v string) {
	o.ResourceId.Set(&v)
}

// SetResourceIdNil sets the value for ResourceId to be an explicit nil
func (o *AssetListen) SetResourceIdNil() {
	o.ResourceId.Set(nil)
}

// UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
func (o *AssetListen) UnsetResourceId() {
	o.ResourceId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *AssetListen) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *AssetListen) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *AssetListen) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *AssetListen) UnsetId() {
	o.Id.Unset()
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetDeviceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetDeviceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceIds) {
		return nil, false
	}
	return o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *AssetListen) HasDeviceIds() bool {
	if o != nil && IsNil(o.DeviceIds) {
		return true
	}

	return false
}

// SetDeviceIds gets a reference to the given []string and assigns it to the DeviceIds field.
func (o *AssetListen) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

// GetProjectId returns the ProjectId field value
func (o *AssetListen) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AssetListen) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AssetListen) SetProjectId(v string) {
	o.ProjectId = v
}

// GetGlobalAssetIdentifier returns the GlobalAssetIdentifier field value
func (o *AssetListen) GetGlobalAssetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalAssetIdentifier
}

// GetGlobalAssetIdentifierOk returns a tuple with the GlobalAssetIdentifier field value
// and a boolean to check if the value has been set.
func (o *AssetListen) GetGlobalAssetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalAssetIdentifier, true
}

// SetGlobalAssetIdentifier sets field value
func (o *AssetListen) SetGlobalAssetIdentifier(v string) {
	o.GlobalAssetIdentifier = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AssetListen) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AssetListen) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *AssetListen) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AssetListen) UnsetName() {
	o.Name.Unset()
}

// GetAssetType returns the AssetType field value
func (o *AssetListen) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *AssetListen) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *AssetListen) SetAssetType(v string) {
	o.AssetType = v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *AssetListen) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableFloat64 and assigns it to the Latitude field.
func (o *AssetListen) SetLatitude(v float64) {
	o.Latitude.Set(&v)
}

// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *AssetListen) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *AssetListen) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *AssetListen) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableFloat64 and assigns it to the Longitude field.
func (o *AssetListen) SetLongitude(v float64) {
	o.Longitude.Set(&v)
}

// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *AssetListen) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *AssetListen) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetIsTracker returns the IsTracker field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *AssetListen) GetIsTracker() bool {
	if o == nil || IsNil(o.IsTracker.Get()) {
		var ret bool
		return ret
	}
	return *o.IsTracker.Get()
}

// GetIsTrackerOk returns a tuple with the IsTracker field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *AssetListen) GetIsTrackerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsTracker.Get(), o.IsTracker.IsSet()
}

// HasIsTracker returns a boolean if a field has been set.
func (o *AssetListen) HasIsTracker() bool {
	if o != nil && o.IsTracker.IsSet() {
		return true
	}

	return false
}

// SetIsTracker gets a reference to the given NullableBool and assigns it to the IsTracker field.
// Deprecated
func (o *AssetListen) SetIsTracker(v bool) {
	o.IsTracker.Set(&v)
}

// SetIsTrackerNil sets the value for IsTracker to be an explicit nil
func (o *AssetListen) SetIsTrackerNil() {
	o.IsTracker.Set(nil)
}

// UnsetIsTracker ensures that no value is present for IsTracker, not even an explicit nil
func (o *AssetListen) UnsetIsTracker() {
	o.IsTracker.Unset()
}

// GetTrackerId returns the TrackerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetTrackerId() int32 {
	if o == nil || IsNil(o.TrackerId.Get()) {
		var ret int32
		return ret
	}
	return *o.TrackerId.Get()
}

// GetTrackerIdOk returns a tuple with the TrackerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetTrackerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackerId.Get(), o.TrackerId.IsSet()
}

// HasTrackerId returns a boolean if a field has been set.
func (o *AssetListen) HasTrackerId() bool {
	if o != nil && o.TrackerId.IsSet() {
		return true
	}

	return false
}

// SetTrackerId gets a reference to the given NullableInt32 and assigns it to the TrackerId field.
func (o *AssetListen) SetTrackerId(v int32) {
	o.TrackerId.Set(&v)
}

// SetTrackerIdNil sets the value for TrackerId to be an explicit nil
func (o *AssetListen) SetTrackerIdNil() {
	o.TrackerId.Set(nil)
}

// UnsetTrackerId ensures that no value is present for TrackerId, not even an explicit nil
func (o *AssetListen) UnsetTrackerId() {
	o.TrackerId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AssetListen) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AssetListen) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AssetListen) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AssetListen) UnsetDescription() {
	o.Description.Unset()
}

// GetParentFunctionalAssetId returns the ParentFunctionalAssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetParentFunctionalAssetId() int32 {
	if o == nil || IsNil(o.ParentFunctionalAssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.ParentFunctionalAssetId.Get()
}

// GetParentFunctionalAssetIdOk returns a tuple with the ParentFunctionalAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetParentFunctionalAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentFunctionalAssetId.Get(), o.ParentFunctionalAssetId.IsSet()
}

// HasParentFunctionalAssetId returns a boolean if a field has been set.
func (o *AssetListen) HasParentFunctionalAssetId() bool {
	if o != nil && o.ParentFunctionalAssetId.IsSet() {
		return true
	}

	return false
}

// SetParentFunctionalAssetId gets a reference to the given NullableInt32 and assigns it to the ParentFunctionalAssetId field.
func (o *AssetListen) SetParentFunctionalAssetId(v int32) {
	o.ParentFunctionalAssetId.Set(&v)
}

// SetParentFunctionalAssetIdNil sets the value for ParentFunctionalAssetId to be an explicit nil
func (o *AssetListen) SetParentFunctionalAssetIdNil() {
	o.ParentFunctionalAssetId.Set(nil)
}

// UnsetParentFunctionalAssetId ensures that no value is present for ParentFunctionalAssetId, not even an explicit nil
func (o *AssetListen) UnsetParentFunctionalAssetId() {
	o.ParentFunctionalAssetId.Unset()
}

// GetFunctionalAssetIdPath returns the FunctionalAssetIdPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetFunctionalAssetIdPath() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.FunctionalAssetIdPath
}

// GetFunctionalAssetIdPathOk returns a tuple with the FunctionalAssetIdPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetFunctionalAssetIdPathOk() ([]int32, bool) {
	if o == nil || IsNil(o.FunctionalAssetIdPath) {
		return nil, false
	}
	return o.FunctionalAssetIdPath, true
}

// HasFunctionalAssetIdPath returns a boolean if a field has been set.
func (o *AssetListen) HasFunctionalAssetIdPath() bool {
	if o != nil && IsNil(o.FunctionalAssetIdPath) {
		return true
	}

	return false
}

// SetFunctionalAssetIdPath gets a reference to the given []int32 and assigns it to the FunctionalAssetIdPath field.
func (o *AssetListen) SetFunctionalAssetIdPath(v []int32) {
	o.FunctionalAssetIdPath = v
}

// GetParentLocationalAssetId returns the ParentLocationalAssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetParentLocationalAssetId() int32 {
	if o == nil || IsNil(o.ParentLocationalAssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.ParentLocationalAssetId.Get()
}

// GetParentLocationalAssetIdOk returns a tuple with the ParentLocationalAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetParentLocationalAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentLocationalAssetId.Get(), o.ParentLocationalAssetId.IsSet()
}

// HasParentLocationalAssetId returns a boolean if a field has been set.
func (o *AssetListen) HasParentLocationalAssetId() bool {
	if o != nil && o.ParentLocationalAssetId.IsSet() {
		return true
	}

	return false
}

// SetParentLocationalAssetId gets a reference to the given NullableInt32 and assigns it to the ParentLocationalAssetId field.
func (o *AssetListen) SetParentLocationalAssetId(v int32) {
	o.ParentLocationalAssetId.Set(&v)
}

// SetParentLocationalAssetIdNil sets the value for ParentLocationalAssetId to be an explicit nil
func (o *AssetListen) SetParentLocationalAssetIdNil() {
	o.ParentLocationalAssetId.Set(nil)
}

// UnsetParentLocationalAssetId ensures that no value is present for ParentLocationalAssetId, not even an explicit nil
func (o *AssetListen) UnsetParentLocationalAssetId() {
	o.ParentLocationalAssetId.Unset()
}

// GetLocationalAssetIdPath returns the LocationalAssetIdPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetLocationalAssetIdPath() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.LocationalAssetIdPath
}

// GetLocationalAssetIdPathOk returns a tuple with the LocationalAssetIdPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetLocationalAssetIdPathOk() ([]int32, bool) {
	if o == nil || IsNil(o.LocationalAssetIdPath) {
		return nil, false
	}
	return o.LocationalAssetIdPath, true
}

// HasLocationalAssetIdPath returns a boolean if a field has been set.
func (o *AssetListen) HasLocationalAssetIdPath() bool {
	if o != nil && IsNil(o.LocationalAssetIdPath) {
		return true
	}

	return false
}

// SetLocationalAssetIdPath gets a reference to the given []int32 and assigns it to the LocationalAssetIdPath field.
func (o *AssetListen) SetLocationalAssetIdPath(v []int32) {
	o.LocationalAssetIdPath = v
}

// GetParentFunctionalIdentifier returns the ParentFunctionalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetParentFunctionalIdentifier() string {
	if o == nil || IsNil(o.ParentFunctionalIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.ParentFunctionalIdentifier.Get()
}

// GetParentFunctionalIdentifierOk returns a tuple with the ParentFunctionalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetParentFunctionalIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentFunctionalIdentifier.Get(), o.ParentFunctionalIdentifier.IsSet()
}

// HasParentFunctionalIdentifier returns a boolean if a field has been set.
func (o *AssetListen) HasParentFunctionalIdentifier() bool {
	if o != nil && o.ParentFunctionalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetParentFunctionalIdentifier gets a reference to the given NullableString and assigns it to the ParentFunctionalIdentifier field.
func (o *AssetListen) SetParentFunctionalIdentifier(v string) {
	o.ParentFunctionalIdentifier.Set(&v)
}

// SetParentFunctionalIdentifierNil sets the value for ParentFunctionalIdentifier to be an explicit nil
func (o *AssetListen) SetParentFunctionalIdentifierNil() {
	o.ParentFunctionalIdentifier.Set(nil)
}

// UnsetParentFunctionalIdentifier ensures that no value is present for ParentFunctionalIdentifier, not even an explicit nil
func (o *AssetListen) UnsetParentFunctionalIdentifier() {
	o.ParentFunctionalIdentifier.Unset()
}

// GetParentLocationalIdentifier returns the ParentLocationalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetParentLocationalIdentifier() string {
	if o == nil || IsNil(o.ParentLocationalIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.ParentLocationalIdentifier.Get()
}

// GetParentLocationalIdentifierOk returns a tuple with the ParentLocationalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetParentLocationalIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentLocationalIdentifier.Get(), o.ParentLocationalIdentifier.IsSet()
}

// HasParentLocationalIdentifier returns a boolean if a field has been set.
func (o *AssetListen) HasParentLocationalIdentifier() bool {
	if o != nil && o.ParentLocationalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetParentLocationalIdentifier gets a reference to the given NullableString and assigns it to the ParentLocationalIdentifier field.
func (o *AssetListen) SetParentLocationalIdentifier(v string) {
	o.ParentLocationalIdentifier.Set(&v)
}

// SetParentLocationalIdentifierNil sets the value for ParentLocationalIdentifier to be an explicit nil
func (o *AssetListen) SetParentLocationalIdentifierNil() {
	o.ParentLocationalIdentifier.Set(nil)
}

// UnsetParentLocationalIdentifier ensures that no value is present for ParentLocationalIdentifier, not even an explicit nil
func (o *AssetListen) UnsetParentLocationalIdentifier() {
	o.ParentLocationalIdentifier.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AssetListen) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AssetListen) SetTags(v []string) {
	o.Tags = v
}

// GetChildrenInfo returns the ChildrenInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetChildrenInfo() []Asset {
	if o == nil {
		var ret []Asset
		return ret
	}
	return o.ChildrenInfo
}

// GetChildrenInfoOk returns a tuple with the ChildrenInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetChildrenInfoOk() ([]Asset, bool) {
	if o == nil || IsNil(o.ChildrenInfo) {
		return nil, false
	}
	return o.ChildrenInfo, true
}

// HasChildrenInfo returns a boolean if a field has been set.
func (o *AssetListen) HasChildrenInfo() bool {
	if o != nil && IsNil(o.ChildrenInfo) {
		return true
	}

	return false
}

// SetChildrenInfo gets a reference to the given []Asset and assigns it to the ChildrenInfo field.
func (o *AssetListen) SetChildrenInfo(v []Asset) {
	o.ChildrenInfo = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetListen) GetAttachments() []Attachment {
	if o == nil {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetListen) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AssetListen) HasAttachments() bool {
	if o != nil && IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *AssetListen) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *AssetListen) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetListen) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *AssetListen) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *AssetListen) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o AssetListen) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetListen) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceId.IsSet() {
		toSerialize["resourceId"] = o.ResourceId.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.DeviceIds != nil {
		toSerialize["deviceIds"] = o.DeviceIds
	}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["globalAssetIdentifier"] = o.GlobalAssetIdentifier
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["assetType"] = o.AssetType
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	if o.IsTracker.IsSet() {
		toSerialize["isTracker"] = o.IsTracker.Get()
	}
	if o.TrackerId.IsSet() {
		toSerialize["trackerId"] = o.TrackerId.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ParentFunctionalAssetId.IsSet() {
		toSerialize["parentFunctionalAssetId"] = o.ParentFunctionalAssetId.Get()
	}
	if o.FunctionalAssetIdPath != nil {
		toSerialize["functionalAssetIdPath"] = o.FunctionalAssetIdPath
	}
	if o.ParentLocationalAssetId.IsSet() {
		toSerialize["parentLocationalAssetId"] = o.ParentLocationalAssetId.Get()
	}
	if o.LocationalAssetIdPath != nil {
		toSerialize["locationalAssetIdPath"] = o.LocationalAssetIdPath
	}
	if o.ParentFunctionalIdentifier.IsSet() {
		toSerialize["parentFunctionalIdentifier"] = o.ParentFunctionalIdentifier.Get()
	}
	if o.ParentLocationalIdentifier.IsSet() {
		toSerialize["parentLocationalIdentifier"] = o.ParentLocationalIdentifier.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ChildrenInfo != nil {
		toSerialize["childrenInfo"] = o.ChildrenInfo
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	return toSerialize, nil
}

func (o *AssetListen) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projectId",
		"globalAssetIdentifier",
		"assetType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAssetListen := _AssetListen{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssetListen)

	if err != nil {
		return err
	}

	*o = AssetListen(varAssetListen)

	return err
}

type NullableAssetListen struct {
	value *AssetListen
	isSet bool
}

func (v NullableAssetListen) Get() *AssetListen {
	return v.value
}

func (v *NullableAssetListen) Set(val *AssetListen) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetListen) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetListen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetListen(val *AssetListen) *NullableAssetListen {
	return &NullableAssetListen{value: val, isSet: true}
}

func (v NullableAssetListen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetListen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
