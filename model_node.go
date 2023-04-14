/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.4.14
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the Node type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Node{}

// Node An edge node
type Node struct {
	// Unique id for the edge node
	Id NullableString `json:"id,omitempty"`
	// UUID to identify the edge node
	Ident NullableString `json:"ident,omitempty"`
	// Password with which the node identifies itself
	// Deprecated
	Password NullableString `json:"password,omitempty"`
	// ID of the corresponding asset
	AssetId NullableInt32 `json:"assetId,omitempty"`
	// Vendor name
	Vendor NullableString `json:"vendor,omitempty"`
	// Model name
	Model NullableString `json:"model,omitempty"`
	// Descriptive text for the edge node
	Description NullableString `json:"description,omitempty"`
	// Is the node enabled or not
	Enable *bool `json:"enable,omitempty"`
}

// NewNode instantiates a new Node object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNode() *Node {
	this := Node{}
	var enable bool = false
	this.Enable = &enable
	return &this
}

// NewNodeWithDefaults instantiates a new Node object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeWithDefaults() *Node {
	this := Node{}
	var enable bool = false
	this.Enable = &enable
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Node) GetId() string {
	if o == nil || isNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Node) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Node) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Node) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *Node) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Node) UnsetId() {
	o.Id.Unset()
}

// GetIdent returns the Ident field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Node) GetIdent() string {
	if o == nil || isNil(o.Ident.Get()) {
		var ret string
		return ret
	}
	return *o.Ident.Get()
}

// GetIdentOk returns a tuple with the Ident field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Node) GetIdentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ident.Get(), o.Ident.IsSet()
}

// HasIdent returns a boolean if a field has been set.
func (o *Node) HasIdent() bool {
	if o != nil && o.Ident.IsSet() {
		return true
	}

	return false
}

// SetIdent gets a reference to the given NullableString and assigns it to the Ident field.
func (o *Node) SetIdent(v string) {
	o.Ident.Set(&v)
}

// SetIdentNil sets the value for Ident to be an explicit nil
func (o *Node) SetIdentNil() {
	o.Ident.Set(nil)
}

// UnsetIdent ensures that no value is present for Ident, not even an explicit nil
func (o *Node) UnsetIdent() {
	o.Ident.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *Node) GetPassword() string {
	if o == nil || isNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *Node) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *Node) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
// Deprecated
func (o *Node) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *Node) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *Node) UnsetPassword() {
	o.Password.Unset()
}

// GetAssetId returns the AssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Node) GetAssetId() int32 {
	if o == nil || isNil(o.AssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.AssetId.Get()
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Node) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetId.Get(), o.AssetId.IsSet()
}

// HasAssetId returns a boolean if a field has been set.
func (o *Node) HasAssetId() bool {
	if o != nil && o.AssetId.IsSet() {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given NullableInt32 and assigns it to the AssetId field.
func (o *Node) SetAssetId(v int32) {
	o.AssetId.Set(&v)
}

// SetAssetIdNil sets the value for AssetId to be an explicit nil
func (o *Node) SetAssetIdNil() {
	o.AssetId.Set(nil)
}

// UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
func (o *Node) UnsetAssetId() {
	o.AssetId.Unset()
}

// GetVendor returns the Vendor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Node) GetVendor() string {
	if o == nil || isNil(o.Vendor.Get()) {
		var ret string
		return ret
	}
	return *o.Vendor.Get()
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Node) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vendor.Get(), o.Vendor.IsSet()
}

// HasVendor returns a boolean if a field has been set.
func (o *Node) HasVendor() bool {
	if o != nil && o.Vendor.IsSet() {
		return true
	}

	return false
}

// SetVendor gets a reference to the given NullableString and assigns it to the Vendor field.
func (o *Node) SetVendor(v string) {
	o.Vendor.Set(&v)
}

// SetVendorNil sets the value for Vendor to be an explicit nil
func (o *Node) SetVendorNil() {
	o.Vendor.Set(nil)
}

// UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
func (o *Node) UnsetVendor() {
	o.Vendor.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Node) GetModel() string {
	if o == nil || isNil(o.Model.Get()) {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Node) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *Node) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *Node) SetModel(v string) {
	o.Model.Set(&v)
}

// SetModelNil sets the value for Model to be an explicit nil
func (o *Node) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *Node) UnsetModel() {
	o.Model.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Node) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Node) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Node) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Node) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Node) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Node) UnsetDescription() {
	o.Description.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *Node) GetEnable() bool {
	if o == nil || isNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetEnableOk() (*bool, bool) {
	if o == nil || isNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *Node) HasEnable() bool {
	if o != nil && !isNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *Node) SetEnable(v bool) {
	o.Enable = &v
}

func (o Node) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Node) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Ident.IsSet() {
		toSerialize["ident"] = o.Ident.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.AssetId.IsSet() {
		toSerialize["assetId"] = o.AssetId.Get()
	}
	if o.Vendor.IsSet() {
		toSerialize["vendor"] = o.Vendor.Get()
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	return toSerialize, nil
}

type NullableNode struct {
	value *Node
	isSet bool
}

func (v NullableNode) Get() *Node {
	return v.value
}

func (v *NullableNode) Set(val *Node) {
	v.value = val
	v.isSet = true
}

func (v NullableNode) IsSet() bool {
	return v.isSet
}

func (v *NullableNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNode(val *Node) *NullableNode {
	return &NullableNode{value: val, isSet: true}
}

func (v NullableNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
