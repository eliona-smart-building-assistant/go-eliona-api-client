/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Dashboard A frontend dashboard
type Dashboard struct {
	// The internal Id of dashboard
	Id NullableInt32 `json:"Id,omitempty"`
	// The name for this dashboard
	Name string `json:"name"`
	// ID of the project to which the dashboard belongs
	ProjectId string `json:"projectId"`
	// ID of the user who owns the dashboard
	UserId string `json:"userId"`
	// The sequence of the dashboard
	Sequence NullableInt32 `json:"sequence,omitempty"`
}

// NewDashboard instantiates a new Dashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboard(name string, projectId string, userId string) *Dashboard {
	this := Dashboard{}
	this.Name = name
	this.ProjectId = projectId
	this.UserId = userId
	var sequence int32 = 0
	this.Sequence = *NewNullableInt32(&sequence)
	return &this
}

// NewDashboardWithDefaults instantiates a new Dashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardWithDefaults() *Dashboard {
	this := Dashboard{}
	var sequence int32 = 0
	this.Sequence = *NewNullableInt32(&sequence)
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dashboard) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dashboard) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Dashboard) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Dashboard) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Dashboard) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Dashboard) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value
func (o *Dashboard) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Dashboard) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *Dashboard) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Dashboard) SetProjectId(v string) {
	o.ProjectId = v
}

// GetUserId returns the UserId field value
func (o *Dashboard) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *Dashboard) SetUserId(v string) {
	o.UserId = v
}

// GetSequence returns the Sequence field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dashboard) GetSequence() int32 {
	if o == nil || o.Sequence.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Sequence.Get()
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dashboard) GetSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sequence.Get(), o.Sequence.IsSet()
}

// HasSequence returns a boolean if a field has been set.
func (o *Dashboard) HasSequence() bool {
	if o != nil && o.Sequence.IsSet() {
		return true
	}

	return false
}

// SetSequence gets a reference to the given NullableInt32 and assigns it to the Sequence field.
func (o *Dashboard) SetSequence(v int32) {
	o.Sequence.Set(&v)
}
// SetSequenceNil sets the value for Sequence to be an explicit nil
func (o *Dashboard) SetSequenceNil() {
	o.Sequence.Set(nil)
}

// UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
func (o *Dashboard) UnsetSequence() {
	o.Sequence.Unset()
}

func (o Dashboard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["projectId"] = o.ProjectId
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if o.Sequence.IsSet() {
		toSerialize["sequence"] = o.Sequence.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDashboard struct {
	value *Dashboard
	isSet bool
}

func (v NullableDashboard) Get() *Dashboard {
	return v.value
}

func (v *NullableDashboard) Set(val *Dashboard) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboard(val *Dashboard) *NullableDashboard {
	return &NullableDashboard{value: val, isSet: true}
}

func (v NullableDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


