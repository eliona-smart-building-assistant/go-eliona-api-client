/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.9.4
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Attachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Attachment{}

// Attachment A binary attachment
type Attachment struct {
	// The name for the attachment, e.g. the filename
	Name string `json:"name"`
	// The type of binary data
	ContentType NullableString `json:"contentType,omitempty"`
	// The encoding of content
	Encoding NullableString `json:"encoding,omitempty"`
	// The binary data as encoded string
	Content *string `json:"content,omitempty"`
}

type _Attachment Attachment

// NewAttachment instantiates a new Attachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachment(name string) *Attachment {
	this := Attachment{}
	this.Name = name
	return &this
}

// NewAttachmentWithDefaults instantiates a new Attachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentWithDefaults() *Attachment {
	this := Attachment{}
	return &this
}

// GetName returns the Name field value
func (o *Attachment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Attachment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Attachment) SetName(v string) {
	o.Name = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attachment) GetContentType() string {
	if o == nil || IsNil(o.ContentType.Get()) {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attachment) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *Attachment) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *Attachment) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *Attachment) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *Attachment) UnsetContentType() {
	o.ContentType.Unset()
}

// GetEncoding returns the Encoding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attachment) GetEncoding() string {
	if o == nil || IsNil(o.Encoding.Get()) {
		var ret string
		return ret
	}
	return *o.Encoding.Get()
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attachment) GetEncodingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Encoding.Get(), o.Encoding.IsSet()
}

// HasEncoding returns a boolean if a field has been set.
func (o *Attachment) HasEncoding() bool {
	if o != nil && o.Encoding.IsSet() {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given NullableString and assigns it to the Encoding field.
func (o *Attachment) SetEncoding(v string) {
	o.Encoding.Set(&v)
}

// SetEncodingNil sets the value for Encoding to be an explicit nil
func (o *Attachment) SetEncodingNil() {
	o.Encoding.Set(nil)
}

// UnsetEncoding ensures that no value is present for Encoding, not even an explicit nil
func (o *Attachment) UnsetEncoding() {
	o.Encoding.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Attachment) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attachment) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Attachment) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Attachment) SetContent(v string) {
	o.Content = &v
}

func (o Attachment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.ContentType.IsSet() {
		toSerialize["contentType"] = o.ContentType.Get()
	}
	if o.Encoding.IsSet() {
		toSerialize["encoding"] = o.Encoding.Get()
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

func (o *Attachment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAttachment := _Attachment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttachment)

	if err != nil {
		return err
	}

	*o = Attachment(varAttachment)

	return err
}

type NullableAttachment struct {
	value *Attachment
	isSet bool
}

func (v NullableAttachment) Get() *Attachment {
	return v.value
}

func (v *NullableAttachment) Set(val *Attachment) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachment(val *Attachment) *NullableAttachment {
	return &NullableAttachment{value: val, isSet: true}
}

func (v NullableAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
