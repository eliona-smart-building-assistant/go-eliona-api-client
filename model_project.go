/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.11
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the Project type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Project{}

// Project A project
type Project struct {
	// The internal ID of the project
	Id NullableString `json:"id,omitempty"`
	// The title of the project
	Title NullableString `json:"title,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject() *Project {
	this := Project{}
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Project) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Project) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *Project) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Project) UnsetId() {
	o.Id.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Project) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Project) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil
func (o *Project) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Project) UnsetTitle() {
	o.Title.Unset()
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Project) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return toSerialize, nil
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
