/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.10
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Notification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Notification{}

// Notification A notification
type Notification struct {
	// E-Mail address or internal ID of the Eliona user
	User string `json:"user"`
	// ID of the project the notification should appear
	ProjectId NullableString      `json:"projectId,omitempty"`
	Message   NullableTranslation `json:"message"`
}

type _Notification Notification

// NewNotification instantiates a new Notification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotification(user string, message NullableTranslation) *Notification {
	this := Notification{}
	this.User = user
	this.Message = message
	return &this
}

// NewNotificationWithDefaults instantiates a new Notification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWithDefaults() *Notification {
	this := Notification{}
	return &this
}

// GetUser returns the User field value
func (o *Notification) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Notification) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Notification) SetUser(v string) {
	o.User = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Notification) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Notification) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *Notification) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableString and assigns it to the ProjectId field.
func (o *Notification) SetProjectId(v string) {
	o.ProjectId.Set(&v)
}

// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *Notification) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *Notification) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for Translation will be returned
func (o *Notification) GetMessage() Translation {
	if o == nil || o.Message.Get() == nil {
		var ret Translation
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Notification) GetMessageOk() (*Translation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *Notification) SetMessage(v Translation) {
	o.Message.Set(&v)
}

func (o Notification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Notification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	toSerialize["message"] = o.Message.Get()
	return toSerialize, nil
}

func (o *Notification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
		"message",
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

	varNotification := _Notification{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotification)

	if err != nil {
		return err
	}

	*o = Notification(varNotification)

	return err
}

type NullableNotification struct {
	value *Notification
	isSet bool
}

func (v NullableNotification) Get() *Notification {
	return v.value
}

func (v *NullableNotification) Set(val *Notification) {
	v.value = val
	v.isSet = true
}

func (v NullableNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotification(val *Notification) *NullableNotification {
	return &NullableNotification{value: val, isSet: true}
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
