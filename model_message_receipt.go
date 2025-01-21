/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.8.1
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the MessageReceipt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageReceipt{}

// MessageReceipt A receipt for a message or notification
type MessageReceipt struct {
	// Identifies the message or notification
	Id string `json:"id"`
	// When the message or notification is scheduled
	ScheduledAt NullableTime `json:"scheduledAt,omitempty"`
	// Status of communication processing
	Status *string `json:"status,omitempty"`
}

type _MessageReceipt MessageReceipt

// NewMessageReceipt instantiates a new MessageReceipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageReceipt(id string) *MessageReceipt {
	this := MessageReceipt{}
	this.Id = id
	return &this
}

// NewMessageReceiptWithDefaults instantiates a new MessageReceipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageReceiptWithDefaults() *MessageReceipt {
	this := MessageReceipt{}
	return &this
}

// GetId returns the Id field value
func (o *MessageReceipt) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MessageReceipt) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MessageReceipt) SetId(v string) {
	o.Id = v
}

// GetScheduledAt returns the ScheduledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageReceipt) GetScheduledAt() time.Time {
	if o == nil || IsNil(o.ScheduledAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledAt.Get()
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageReceipt) GetScheduledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledAt.Get(), o.ScheduledAt.IsSet()
}

// HasScheduledAt returns a boolean if a field has been set.
func (o *MessageReceipt) HasScheduledAt() bool {
	if o != nil && o.ScheduledAt.IsSet() {
		return true
	}

	return false
}

// SetScheduledAt gets a reference to the given NullableTime and assigns it to the ScheduledAt field.
func (o *MessageReceipt) SetScheduledAt(v time.Time) {
	o.ScheduledAt.Set(&v)
}

// SetScheduledAtNil sets the value for ScheduledAt to be an explicit nil
func (o *MessageReceipt) SetScheduledAtNil() {
	o.ScheduledAt.Set(nil)
}

// UnsetScheduledAt ensures that no value is present for ScheduledAt, not even an explicit nil
func (o *MessageReceipt) UnsetScheduledAt() {
	o.ScheduledAt.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MessageReceipt) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageReceipt) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MessageReceipt) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MessageReceipt) SetStatus(v string) {
	o.Status = &v
}

func (o MessageReceipt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageReceipt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.ScheduledAt.IsSet() {
		toSerialize["scheduledAt"] = o.ScheduledAt.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *MessageReceipt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varMessageReceipt := _MessageReceipt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageReceipt)

	if err != nil {
		return err
	}

	*o = MessageReceipt(varMessageReceipt)

	return err
}

type NullableMessageReceipt struct {
	value *MessageReceipt
	isSet bool
}

func (v NullableMessageReceipt) Get() *MessageReceipt {
	return v.value
}

func (v *NullableMessageReceipt) Set(val *MessageReceipt) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageReceipt(val *MessageReceipt) *NullableMessageReceipt {
	return &NullableMessageReceipt{value: val, isSet: true}
}

func (v NullableMessageReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
