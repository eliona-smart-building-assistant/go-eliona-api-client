/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.4.11
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the Asset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Asset{}

// Asset An asset
type Asset struct {
	// The internal Id of asset
	Id NullableInt32 `json:"id,omitempty"`
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
	// Textual description for this asset
	Description NullableString `json:"description,omitempty"`
	// The id of an asset which groups this asset as a functional child
	ParentFunctionalAssetId NullableInt32 `json:"parentFunctionalAssetId,omitempty"`
	// The id of an asset which groups this asset as a locational child
	ParentLocationalAssetId NullableInt32 `json:"parentLocationalAssetId,omitempty"`
	// List of associated tags
	Tags []string `json:"tags,omitempty"`
	// List of children for this asset.
	ChildrenInfo []Asset `json:"childrenInfo,omitempty"`
}

// NewAsset instantiates a new Asset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsset(projectId string, globalAssetIdentifier string, assetType string) *Asset {
	this := Asset{}
	this.ProjectId = projectId
	this.GlobalAssetIdentifier = globalAssetIdentifier
	this.AssetType = assetType
	return &this
}

// NewAssetWithDefaults instantiates a new Asset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWithDefaults() *Asset {
	this := Asset{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetId() int32 {
	if o == nil || isNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Asset) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Asset) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *Asset) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Asset) UnsetId() {
	o.Id.Unset()
}

// GetProjectId returns the ProjectId field value
func (o *Asset) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Asset) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Asset) SetProjectId(v string) {
	o.ProjectId = v
}

// GetGlobalAssetIdentifier returns the GlobalAssetIdentifier field value
func (o *Asset) GetGlobalAssetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalAssetIdentifier
}

// GetGlobalAssetIdentifierOk returns a tuple with the GlobalAssetIdentifier field value
// and a boolean to check if the value has been set.
func (o *Asset) GetGlobalAssetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalAssetIdentifier, true
}

// SetGlobalAssetIdentifier sets field value
func (o *Asset) SetGlobalAssetIdentifier(v string) {
	o.GlobalAssetIdentifier = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Asset) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Asset) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *Asset) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Asset) UnsetName() {
	o.Name.Unset()
}

// GetAssetType returns the AssetType field value
func (o *Asset) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *Asset) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *Asset) SetAssetType(v string) {
	o.AssetType = v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetLatitude() float64 {
	if o == nil || isNil(o.Latitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *Asset) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableFloat64 and assigns it to the Latitude field.
func (o *Asset) SetLatitude(v float64) {
	o.Latitude.Set(&v)
}

// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *Asset) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *Asset) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetLongitude() float64 {
	if o == nil || isNil(o.Longitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *Asset) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableFloat64 and assigns it to the Longitude field.
func (o *Asset) SetLongitude(v float64) {
	o.Longitude.Set(&v)
}

// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *Asset) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *Asset) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Asset) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Asset) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Asset) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Asset) UnsetDescription() {
	o.Description.Unset()
}

// GetParentFunctionalAssetId returns the ParentFunctionalAssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetParentFunctionalAssetId() int32 {
	if o == nil || isNil(o.ParentFunctionalAssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.ParentFunctionalAssetId.Get()
}

// GetParentFunctionalAssetIdOk returns a tuple with the ParentFunctionalAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetParentFunctionalAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentFunctionalAssetId.Get(), o.ParentFunctionalAssetId.IsSet()
}

// HasParentFunctionalAssetId returns a boolean if a field has been set.
func (o *Asset) HasParentFunctionalAssetId() bool {
	if o != nil && o.ParentFunctionalAssetId.IsSet() {
		return true
	}

	return false
}

// SetParentFunctionalAssetId gets a reference to the given NullableInt32 and assigns it to the ParentFunctionalAssetId field.
func (o *Asset) SetParentFunctionalAssetId(v int32) {
	o.ParentFunctionalAssetId.Set(&v)
}

// SetParentFunctionalAssetIdNil sets the value for ParentFunctionalAssetId to be an explicit nil
func (o *Asset) SetParentFunctionalAssetIdNil() {
	o.ParentFunctionalAssetId.Set(nil)
}

// UnsetParentFunctionalAssetId ensures that no value is present for ParentFunctionalAssetId, not even an explicit nil
func (o *Asset) UnsetParentFunctionalAssetId() {
	o.ParentFunctionalAssetId.Unset()
}

// GetParentLocationalAssetId returns the ParentLocationalAssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetParentLocationalAssetId() int32 {
	if o == nil || isNil(o.ParentLocationalAssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.ParentLocationalAssetId.Get()
}

// GetParentLocationalAssetIdOk returns a tuple with the ParentLocationalAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetParentLocationalAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentLocationalAssetId.Get(), o.ParentLocationalAssetId.IsSet()
}

// HasParentLocationalAssetId returns a boolean if a field has been set.
func (o *Asset) HasParentLocationalAssetId() bool {
	if o != nil && o.ParentLocationalAssetId.IsSet() {
		return true
	}

	return false
}

// SetParentLocationalAssetId gets a reference to the given NullableInt32 and assigns it to the ParentLocationalAssetId field.
func (o *Asset) SetParentLocationalAssetId(v int32) {
	o.ParentLocationalAssetId.Set(&v)
}

// SetParentLocationalAssetIdNil sets the value for ParentLocationalAssetId to be an explicit nil
func (o *Asset) SetParentLocationalAssetIdNil() {
	o.ParentLocationalAssetId.Set(nil)
}

// UnsetParentLocationalAssetId ensures that no value is present for ParentLocationalAssetId, not even an explicit nil
func (o *Asset) UnsetParentLocationalAssetId() {
	o.ParentLocationalAssetId.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Asset) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Asset) SetTags(v []string) {
	o.Tags = v
}

// GetChildrenInfo returns the ChildrenInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Asset) GetChildrenInfo() []Asset {
	if o == nil {
		var ret []Asset
		return ret
	}
	return o.ChildrenInfo
}

// GetChildrenInfoOk returns a tuple with the ChildrenInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Asset) GetChildrenInfoOk() ([]Asset, bool) {
	if o == nil || isNil(o.ChildrenInfo) {
		return nil, false
	}
	return o.ChildrenInfo, true
}

// HasChildrenInfo returns a boolean if a field has been set.
func (o *Asset) HasChildrenInfo() bool {
	if o != nil && isNil(o.ChildrenInfo) {
		return true
	}

	return false
}

// SetChildrenInfo gets a reference to the given []Asset and assigns it to the ChildrenInfo field.
func (o *Asset) SetChildrenInfo(v []Asset) {
	o.ChildrenInfo = v
}

func (o Asset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Asset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
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
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ParentFunctionalAssetId.IsSet() {
		toSerialize["parentFunctionalAssetId"] = o.ParentFunctionalAssetId.Get()
	}
	if o.ParentLocationalAssetId.IsSet() {
		toSerialize["parentLocationalAssetId"] = o.ParentLocationalAssetId.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ChildrenInfo != nil {
		toSerialize["childrenInfo"] = o.ChildrenInfo
	}
	return toSerialize, nil
}

type NullableAsset struct {
	value *Asset
	isSet bool
}

func (v NullableAsset) Get() *Asset {
	return v.value
}

func (v *NullableAsset) Set(val *Asset) {
	v.value = val
	v.isSet = true
}

func (v NullableAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsset(val *Asset) *NullableAsset {
	return &NullableAsset{value: val, isSet: true}
}

func (v NullableAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
