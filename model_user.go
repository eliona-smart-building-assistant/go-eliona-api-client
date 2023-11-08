/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.5.5
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User An user
type User struct {
	// The internal ID of user
	Id NullableString `json:"id,omitempty"`
	// The user's firstname
	Firstname NullableString `json:"firstname,omitempty"`
	// The user's lastname
	Lastname NullableString `json:"lastname,omitempty"`
	// Address of the user
	Email string `json:"email"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(email string) *User {
	this := User{}
	this.Email = email
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *User) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *User) UnsetId() {
	o.Id.Unset()
}

// GetFirstname returns the Firstname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetFirstname() string {
	if o == nil || IsNil(o.Firstname.Get()) {
		var ret string
		return ret
	}
	return *o.Firstname.Get()
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Firstname.Get(), o.Firstname.IsSet()
}

// HasFirstname returns a boolean if a field has been set.
func (o *User) HasFirstname() bool {
	if o != nil && o.Firstname.IsSet() {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given NullableString and assigns it to the Firstname field.
func (o *User) SetFirstname(v string) {
	o.Firstname.Set(&v)
}

// SetFirstnameNil sets the value for Firstname to be an explicit nil
func (o *User) SetFirstnameNil() {
	o.Firstname.Set(nil)
}

// UnsetFirstname ensures that no value is present for Firstname, not even an explicit nil
func (o *User) UnsetFirstname() {
	o.Firstname.Unset()
}

// GetLastname returns the Lastname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLastname() string {
	if o == nil || IsNil(o.Lastname.Get()) {
		var ret string
		return ret
	}
	return *o.Lastname.Get()
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lastname.Get(), o.Lastname.IsSet()
}

// HasLastname returns a boolean if a field has been set.
func (o *User) HasLastname() bool {
	if o != nil && o.Lastname.IsSet() {
		return true
	}

	return false
}

// SetLastname gets a reference to the given NullableString and assigns it to the Lastname field.
func (o *User) SetLastname(v string) {
	o.Lastname.Set(&v)
}

// SetLastnameNil sets the value for Lastname to be an explicit nil
func (o *User) SetLastnameNil() {
	o.Lastname.Set(nil)
}

// UnsetLastname ensures that no value is present for Lastname, not even an explicit nil
func (o *User) UnsetLastname() {
	o.Lastname.Unset()
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Firstname.IsSet() {
		toSerialize["firstname"] = o.Firstname.Get()
	}
	if o.Lastname.IsSet() {
		toSerialize["lastname"] = o.Lastname.Get()
	}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
