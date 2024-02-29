/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.5
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the Translation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Translation{}

// Translation Readable text to display in frontend
type Translation struct {
	// German text
	De *string `json:"de,omitempty"`
	// English text
	En *string `json:"en,omitempty"`
	// French text
	Fr *string `json:"fr,omitempty"`
	// Italian text
	It *string `json:"it,omitempty"`
}

// NewTranslation instantiates a new Translation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslation() *Translation {
	this := Translation{}
	return &this
}

// NewTranslationWithDefaults instantiates a new Translation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationWithDefaults() *Translation {
	this := Translation{}
	return &this
}

// GetDe returns the De field value if set, zero value otherwise.
func (o *Translation) GetDe() string {
	if o == nil || IsNil(o.De) {
		var ret string
		return ret
	}
	return *o.De
}

// GetDeOk returns a tuple with the De field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetDeOk() (*string, bool) {
	if o == nil || IsNil(o.De) {
		return nil, false
	}
	return o.De, true
}

// HasDe returns a boolean if a field has been set.
func (o *Translation) HasDe() bool {
	if o != nil && !IsNil(o.De) {
		return true
	}

	return false
}

// SetDe gets a reference to the given string and assigns it to the De field.
func (o *Translation) SetDe(v string) {
	o.De = &v
}

// GetEn returns the En field value if set, zero value otherwise.
func (o *Translation) GetEn() string {
	if o == nil || IsNil(o.En) {
		var ret string
		return ret
	}
	return *o.En
}

// GetEnOk returns a tuple with the En field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetEnOk() (*string, bool) {
	if o == nil || IsNil(o.En) {
		return nil, false
	}
	return o.En, true
}

// HasEn returns a boolean if a field has been set.
func (o *Translation) HasEn() bool {
	if o != nil && !IsNil(o.En) {
		return true
	}

	return false
}

// SetEn gets a reference to the given string and assigns it to the En field.
func (o *Translation) SetEn(v string) {
	o.En = &v
}

// GetFr returns the Fr field value if set, zero value otherwise.
func (o *Translation) GetFr() string {
	if o == nil || IsNil(o.Fr) {
		var ret string
		return ret
	}
	return *o.Fr
}

// GetFrOk returns a tuple with the Fr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetFrOk() (*string, bool) {
	if o == nil || IsNil(o.Fr) {
		return nil, false
	}
	return o.Fr, true
}

// HasFr returns a boolean if a field has been set.
func (o *Translation) HasFr() bool {
	if o != nil && !IsNil(o.Fr) {
		return true
	}

	return false
}

// SetFr gets a reference to the given string and assigns it to the Fr field.
func (o *Translation) SetFr(v string) {
	o.Fr = &v
}

// GetIt returns the It field value if set, zero value otherwise.
func (o *Translation) GetIt() string {
	if o == nil || IsNil(o.It) {
		var ret string
		return ret
	}
	return *o.It
}

// GetItOk returns a tuple with the It field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Translation) GetItOk() (*string, bool) {
	if o == nil || IsNil(o.It) {
		return nil, false
	}
	return o.It, true
}

// HasIt returns a boolean if a field has been set.
func (o *Translation) HasIt() bool {
	if o != nil && !IsNil(o.It) {
		return true
	}

	return false
}

// SetIt gets a reference to the given string and assigns it to the It field.
func (o *Translation) SetIt(v string) {
	o.It = &v
}

func (o Translation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Translation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.De) {
		toSerialize["de"] = o.De
	}
	if !IsNil(o.En) {
		toSerialize["en"] = o.En
	}
	if !IsNil(o.Fr) {
		toSerialize["fr"] = o.Fr
	}
	if !IsNil(o.It) {
		toSerialize["it"] = o.It
	}
	return toSerialize, nil
}

type NullableTranslation struct {
	value *Translation
	isSet bool
}

func (v NullableTranslation) Get() *Translation {
	return v.value
}

func (v *NullableTranslation) Set(val *Translation) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslation) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslation(val *Translation) *NullableTranslation {
	return &NullableTranslation{value: val, isSet: true}
}

func (v NullableTranslation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
