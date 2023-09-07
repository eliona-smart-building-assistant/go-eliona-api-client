/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.5.4
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Message{}

// Message A message
type Message struct {
	// Address of the sender, e.g. an e-mail address
	Sender NullableString `json:"sender,omitempty"`
	// A list of recipient addresses to receive this message
	Recipients []string `json:"recipients"`
	// A list of recipient addresses to receive this message as copy
	CopyRecipients []string `json:"copyRecipients,omitempty"`
	// A list of recipient addresses to receive this message as blind copy without any other recipient information
	BlindCopyRecipients []string `json:"blindCopyRecipients,omitempty"`
	// The subject for this message
	Subject *string `json:"subject,omitempty"`
	// The content of the message. If template is used, the content is embedded in the template.
	Content string `json:"content"`
	// A list of files attached to the message
	Attachments []Attachment `json:"attachments,omitempty"`
}

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage(recipients []string, content string) *Message {
	this := Message{}
	this.Recipients = recipients
	this.Content = content
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetSender returns the Sender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetSender() string {
	if o == nil || IsNil(o.Sender.Get()) {
		var ret string
		return ret
	}
	return *o.Sender.Get()
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sender.Get(), o.Sender.IsSet()
}

// HasSender returns a boolean if a field has been set.
func (o *Message) HasSender() bool {
	if o != nil && o.Sender.IsSet() {
		return true
	}

	return false
}

// SetSender gets a reference to the given NullableString and assigns it to the Sender field.
func (o *Message) SetSender(v string) {
	o.Sender.Set(&v)
}

// SetSenderNil sets the value for Sender to be an explicit nil
func (o *Message) SetSenderNil() {
	o.Sender.Set(nil)
}

// UnsetSender ensures that no value is present for Sender, not even an explicit nil
func (o *Message) UnsetSender() {
	o.Sender.Unset()
}

// GetRecipients returns the Recipients field value
func (o *Message) GetRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *Message) GetRecipientsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *Message) SetRecipients(v []string) {
	o.Recipients = v
}

// GetCopyRecipients returns the CopyRecipients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetCopyRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CopyRecipients
}

// GetCopyRecipientsOk returns a tuple with the CopyRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetCopyRecipientsOk() ([]string, bool) {
	if o == nil || IsNil(o.CopyRecipients) {
		return nil, false
	}
	return o.CopyRecipients, true
}

// HasCopyRecipients returns a boolean if a field has been set.
func (o *Message) HasCopyRecipients() bool {
	if o != nil && IsNil(o.CopyRecipients) {
		return true
	}

	return false
}

// SetCopyRecipients gets a reference to the given []string and assigns it to the CopyRecipients field.
func (o *Message) SetCopyRecipients(v []string) {
	o.CopyRecipients = v
}

// GetBlindCopyRecipients returns the BlindCopyRecipients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetBlindCopyRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BlindCopyRecipients
}

// GetBlindCopyRecipientsOk returns a tuple with the BlindCopyRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetBlindCopyRecipientsOk() ([]string, bool) {
	if o == nil || IsNil(o.BlindCopyRecipients) {
		return nil, false
	}
	return o.BlindCopyRecipients, true
}

// HasBlindCopyRecipients returns a boolean if a field has been set.
func (o *Message) HasBlindCopyRecipients() bool {
	if o != nil && IsNil(o.BlindCopyRecipients) {
		return true
	}

	return false
}

// SetBlindCopyRecipients gets a reference to the given []string and assigns it to the BlindCopyRecipients field.
func (o *Message) SetBlindCopyRecipients(v []string) {
	o.BlindCopyRecipients = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Message) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Message) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Message) SetSubject(v string) {
	o.Subject = &v
}

// GetContent returns the Content field value
func (o *Message) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Message) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *Message) SetContent(v string) {
	o.Content = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetAttachments() []Attachment {
	if o == nil {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *Message) HasAttachments() bool {
	if o != nil && IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *Message) SetAttachments(v []Attachment) {
	o.Attachments = v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Sender.IsSet() {
		toSerialize["sender"] = o.Sender.Get()
	}
	toSerialize["recipients"] = o.Recipients
	if o.CopyRecipients != nil {
		toSerialize["copyRecipients"] = o.CopyRecipients
	}
	if o.BlindCopyRecipients != nil {
		toSerialize["blindCopyRecipients"] = o.BlindCopyRecipients
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	toSerialize["content"] = o.Content
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
