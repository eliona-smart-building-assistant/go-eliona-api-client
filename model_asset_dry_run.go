/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.5.10
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AssetDryRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetDryRun{}

// AssetDryRun struct for AssetDryRun
type AssetDryRun struct {
	// Unique identifier (textual or numeric) of resources
	Identifier *string `json:"identifier,omitempty"`
	// The status code expecting when actually perform the operation. Some values are - 200: updated (ok)  - 201: created - 204: deleted (no content) - 304: unchanged (not modified) - 400: problem (bad request) - 404: not found - 409: duplicated (conflict) - 422: unprocessable
	StatusCode *int32 `json:"statusCode,omitempty"`
	// The error message expecting when actually perform the operation.
	ErrorMessage NullableString   `json:"errorMessage,omitempty"`
	IdentifiedBy *AssetIdentifyBy `json:"identifiedBy,omitempty"`
	Before       NullableAsset    `json:"before,omitempty"`
	After        NullableAsset    `json:"after,omitempty"`
}

// NewAssetDryRun instantiates a new AssetDryRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDryRun() *AssetDryRun {
	this := AssetDryRun{}
	return &this
}

// NewAssetDryRunWithDefaults instantiates a new AssetDryRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDryRunWithDefaults() *AssetDryRun {
	this := AssetDryRun{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *AssetDryRun) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDryRun) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *AssetDryRun) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *AssetDryRun) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *AssetDryRun) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDryRun) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *AssetDryRun) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *AssetDryRun) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDryRun) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDryRun) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AssetDryRun) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *AssetDryRun) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *AssetDryRun) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *AssetDryRun) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetIdentifiedBy returns the IdentifiedBy field value if set, zero value otherwise.
func (o *AssetDryRun) GetIdentifiedBy() AssetIdentifyBy {
	if o == nil || IsNil(o.IdentifiedBy) {
		var ret AssetIdentifyBy
		return ret
	}
	return *o.IdentifiedBy
}

// GetIdentifiedByOk returns a tuple with the IdentifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDryRun) GetIdentifiedByOk() (*AssetIdentifyBy, bool) {
	if o == nil || IsNil(o.IdentifiedBy) {
		return nil, false
	}
	return o.IdentifiedBy, true
}

// HasIdentifiedBy returns a boolean if a field has been set.
func (o *AssetDryRun) HasIdentifiedBy() bool {
	if o != nil && !IsNil(o.IdentifiedBy) {
		return true
	}

	return false
}

// SetIdentifiedBy gets a reference to the given AssetIdentifyBy and assigns it to the IdentifiedBy field.
func (o *AssetDryRun) SetIdentifiedBy(v AssetIdentifyBy) {
	o.IdentifiedBy = &v
}

// GetBefore returns the Before field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDryRun) GetBefore() Asset {
	if o == nil || IsNil(o.Before.Get()) {
		var ret Asset
		return ret
	}
	return *o.Before.Get()
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDryRun) GetBeforeOk() (*Asset, bool) {
	if o == nil {
		return nil, false
	}
	return o.Before.Get(), o.Before.IsSet()
}

// HasBefore returns a boolean if a field has been set.
func (o *AssetDryRun) HasBefore() bool {
	if o != nil && o.Before.IsSet() {
		return true
	}

	return false
}

// SetBefore gets a reference to the given NullableAsset and assigns it to the Before field.
func (o *AssetDryRun) SetBefore(v Asset) {
	o.Before.Set(&v)
}

// SetBeforeNil sets the value for Before to be an explicit nil
func (o *AssetDryRun) SetBeforeNil() {
	o.Before.Set(nil)
}

// UnsetBefore ensures that no value is present for Before, not even an explicit nil
func (o *AssetDryRun) UnsetBefore() {
	o.Before.Unset()
}

// GetAfter returns the After field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDryRun) GetAfter() Asset {
	if o == nil || IsNil(o.After.Get()) {
		var ret Asset
		return ret
	}
	return *o.After.Get()
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDryRun) GetAfterOk() (*Asset, bool) {
	if o == nil {
		return nil, false
	}
	return o.After.Get(), o.After.IsSet()
}

// HasAfter returns a boolean if a field has been set.
func (o *AssetDryRun) HasAfter() bool {
	if o != nil && o.After.IsSet() {
		return true
	}

	return false
}

// SetAfter gets a reference to the given NullableAsset and assigns it to the After field.
func (o *AssetDryRun) SetAfter(v Asset) {
	o.After.Set(&v)
}

// SetAfterNil sets the value for After to be an explicit nil
func (o *AssetDryRun) SetAfterNil() {
	o.After.Set(nil)
}

// UnsetAfter ensures that no value is present for After, not even an explicit nil
func (o *AssetDryRun) UnsetAfter() {
	o.After.Unset()
}

func (o AssetDryRun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetDryRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if !IsNil(o.IdentifiedBy) {
		toSerialize["identifiedBy"] = o.IdentifiedBy
	}
	if o.Before.IsSet() {
		toSerialize["before"] = o.Before.Get()
	}
	if o.After.IsSet() {
		toSerialize["after"] = o.After.Get()
	}
	return toSerialize, nil
}

type NullableAssetDryRun struct {
	value *AssetDryRun
	isSet bool
}

func (v NullableAssetDryRun) Get() *AssetDryRun {
	return v.value
}

func (v *NullableAssetDryRun) Set(val *AssetDryRun) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDryRun) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDryRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDryRun(val *AssetDryRun) *NullableAssetDryRun {
	return &NullableAssetDryRun{value: val, isSet: true}
}

func (v NullableAssetDryRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDryRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
