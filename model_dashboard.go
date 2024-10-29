/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.7.2
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Dashboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dashboard{}

// Dashboard A frontend dashboard
type Dashboard struct {
	// The internal Id of dashboard
	Id NullableInt32 `json:"id,omitempty"`
	// The name for this dashboard
	Name string `json:"name"`
	// ID of the project to which the dashboard belongs
	ProjectId string `json:"projectId"`
	// ID of the user who owns the dashboard
	UserId string `json:"userId"`
	// The sequence of the dashboard
	// Deprecated
	Sequence NullableInt32 `json:"sequence,omitempty"`
	// List of widgets on this dashboard (order matches the order of widgets on the dashboard)
	Widgets []Widget `json:"widgets,omitempty"`
	// Is the dashboard public and not bound to a dedicated user
	Public NullableBool `json:"public,omitempty"`
}

type _Dashboard Dashboard

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
	var public bool = false
	this.Public = *NewNullableBool(&public)
	return &this
}

// NewDashboardWithDefaults instantiates a new Dashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardWithDefaults() *Dashboard {
	this := Dashboard{}
	var sequence int32 = 0
	this.Sequence = *NewNullableInt32(&sequence)
	var public bool = false
	this.Public = *NewNullableBool(&public)
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dashboard) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
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
// Deprecated
func (o *Dashboard) GetSequence() int32 {
	if o == nil || IsNil(o.Sequence.Get()) {
		var ret int32
		return ret
	}
	return *o.Sequence.Get()
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
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
// Deprecated
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

// GetWidgets returns the Widgets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dashboard) GetWidgets() []Widget {
	if o == nil {
		var ret []Widget
		return ret
	}
	return o.Widgets
}

// GetWidgetsOk returns a tuple with the Widgets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dashboard) GetWidgetsOk() ([]Widget, bool) {
	if o == nil || IsNil(o.Widgets) {
		return nil, false
	}
	return o.Widgets, true
}

// HasWidgets returns a boolean if a field has been set.
func (o *Dashboard) HasWidgets() bool {
	if o != nil && IsNil(o.Widgets) {
		return true
	}

	return false
}

// SetWidgets gets a reference to the given []Widget and assigns it to the Widgets field.
func (o *Dashboard) SetWidgets(v []Widget) {
	o.Widgets = v
}

// GetPublic returns the Public field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Dashboard) GetPublic() bool {
	if o == nil || IsNil(o.Public.Get()) {
		var ret bool
		return ret
	}
	return *o.Public.Get()
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Dashboard) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Public.Get(), o.Public.IsSet()
}

// HasPublic returns a boolean if a field has been set.
func (o *Dashboard) HasPublic() bool {
	if o != nil && o.Public.IsSet() {
		return true
	}

	return false
}

// SetPublic gets a reference to the given NullableBool and assigns it to the Public field.
func (o *Dashboard) SetPublic(v bool) {
	o.Public.Set(&v)
}

// SetPublicNil sets the value for Public to be an explicit nil
func (o *Dashboard) SetPublicNil() {
	o.Public.Set(nil)
}

// UnsetPublic ensures that no value is present for Public, not even an explicit nil
func (o *Dashboard) UnsetPublic() {
	o.Public.Unset()
}

func (o Dashboard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dashboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["projectId"] = o.ProjectId
	toSerialize["userId"] = o.UserId
	if o.Sequence.IsSet() {
		toSerialize["sequence"] = o.Sequence.Get()
	}
	if o.Widgets != nil {
		toSerialize["widgets"] = o.Widgets
	}
	if o.Public.IsSet() {
		toSerialize["public"] = o.Public.Get()
	}
	return toSerialize, nil
}

func (o *Dashboard) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"projectId",
		"userId",
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

	varDashboard := _Dashboard{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDashboard)

	if err != nil {
		return err
	}

	*o = Dashboard(varDashboard)

	return err
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
