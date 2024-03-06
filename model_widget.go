/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.6.7
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Widget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Widget{}

// Widget A widget on a frontend dashboard
type Widget struct {
	// The internal Id of widget
	Id NullableInt32 `json:"id,omitempty"`
	// The name for the type of this widget
	WidgetTypeName string `json:"widgetTypeName"`
	// Detailed configuration depending on the widget type
	Details map[string]interface{} `json:"details,omitempty"`
	// The master asset id of this widget
	AssetId NullableInt32 `json:"assetId,omitempty"`
	// Placement order on dashboard; if not set the index in array is taken
	Sequence NullableInt32 `json:"sequence,omitempty"`
	// List of data for the elements of widget
	Data []WidgetData `json:"data,omitempty"`
}

type _Widget Widget

// NewWidget instantiates a new Widget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidget(widgetTypeName string) *Widget {
	this := Widget{}
	this.WidgetTypeName = widgetTypeName
	return &this
}

// NewWidgetWithDefaults instantiates a new Widget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetWithDefaults() *Widget {
	this := Widget{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Widget) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Widget) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Widget) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Widget) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *Widget) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Widget) UnsetId() {
	o.Id.Unset()
}

// GetWidgetTypeName returns the WidgetTypeName field value
func (o *Widget) GetWidgetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WidgetTypeName
}

// GetWidgetTypeNameOk returns a tuple with the WidgetTypeName field value
// and a boolean to check if the value has been set.
func (o *Widget) GetWidgetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WidgetTypeName, true
}

// SetWidgetTypeName sets field value
func (o *Widget) SetWidgetTypeName(v string) {
	o.WidgetTypeName = v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Widget) GetDetails() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Widget) GetDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return map[string]interface{}{}, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Widget) HasDetails() bool {
	if o != nil && IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *Widget) SetDetails(v map[string]interface{}) {
	o.Details = v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Widget) GetAssetId() int32 {
	if o == nil || IsNil(o.AssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.AssetId.Get()
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Widget) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetId.Get(), o.AssetId.IsSet()
}

// HasAssetId returns a boolean if a field has been set.
func (o *Widget) HasAssetId() bool {
	if o != nil && o.AssetId.IsSet() {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given NullableInt32 and assigns it to the AssetId field.
func (o *Widget) SetAssetId(v int32) {
	o.AssetId.Set(&v)
}

// SetAssetIdNil sets the value for AssetId to be an explicit nil
func (o *Widget) SetAssetIdNil() {
	o.AssetId.Set(nil)
}

// UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
func (o *Widget) UnsetAssetId() {
	o.AssetId.Unset()
}

// GetSequence returns the Sequence field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Widget) GetSequence() int32 {
	if o == nil || IsNil(o.Sequence.Get()) {
		var ret int32
		return ret
	}
	return *o.Sequence.Get()
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Widget) GetSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sequence.Get(), o.Sequence.IsSet()
}

// HasSequence returns a boolean if a field has been set.
func (o *Widget) HasSequence() bool {
	if o != nil && o.Sequence.IsSet() {
		return true
	}

	return false
}

// SetSequence gets a reference to the given NullableInt32 and assigns it to the Sequence field.
func (o *Widget) SetSequence(v int32) {
	o.Sequence.Set(&v)
}

// SetSequenceNil sets the value for Sequence to be an explicit nil
func (o *Widget) SetSequenceNil() {
	o.Sequence.Set(nil)
}

// UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
func (o *Widget) UnsetSequence() {
	o.Sequence.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Widget) GetData() []WidgetData {
	if o == nil {
		var ret []WidgetData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Widget) GetDataOk() ([]WidgetData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Widget) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []WidgetData and assigns it to the Data field.
func (o *Widget) SetData(v []WidgetData) {
	o.Data = v
}

func (o Widget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Widget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["widgetTypeName"] = o.WidgetTypeName
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.AssetId.IsSet() {
		toSerialize["assetId"] = o.AssetId.Get()
	}
	if o.Sequence.IsSet() {
		toSerialize["sequence"] = o.Sequence.Get()
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

func (o *Widget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"widgetTypeName",
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

	varWidget := _Widget{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWidget)

	if err != nil {
		return err
	}

	*o = Widget(varWidget)

	return err
}

type NullableWidget struct {
	value *Widget
	isSet bool
}

func (v NullableWidget) Get() *Widget {
	return v.value
}

func (v *NullableWidget) Set(val *Widget) {
	v.value = val
	v.isSet = true
}

func (v NullableWidget) IsSet() bool {
	return v.isSet
}

func (v *NullableWidget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidget(val *Widget) *NullableWidget {
	return &NullableWidget{value: val, isSet: true}
}

func (v NullableWidget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
