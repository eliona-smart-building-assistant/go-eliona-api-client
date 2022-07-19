/*
Eliona API

API to access Eliona Smart Building Assistant

API version: 2.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Attribute Named attribute to store data of assets
type Attribute struct {
	// The unique name for the asset type
	AssetTypeName *string `json:"assetTypeName,omitempty"`
	// Unique key of asset heap data
	Name    string      `json:"name"`
	Subtype HeapSubtype `json:"subtype"`
	// Name of the type for this attribute
	Type *string `json:"type,omitempty"`
	// Is heap data active or not
	Enable      *bool        `json:"enable,omitempty"`
	Translation *Translation `json:"translation,omitempty"`
	// Physical unit of numeric data
	Unit *string `json:"unit,omitempty"`
	// Number of decimal places
	Precision *int64 `json:"precision,omitempty"`
	// Lower limit
	Min *float64 `json:"min,omitempty"`
	// Upper limit
	Max      *float64           `json:"max,omitempty"`
	Pipeline *AttributePipeline `json:"pipeline,omitempty"`
	// Should the attribute be displayed in viewer
	Viewer *bool `json:"viewer,omitempty"`
	// Should the attribute be displayed in AR
	Ar *bool `json:"ar,omitempty"`
	// Sequence in AR display
	Sequence *int64 `json:"sequence,omitempty"`
	// Is the attribute virtual or not
	Virtual *bool `json:"virtual,omitempty"`
}

// NewAttribute instantiates a new Attribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttribute(name string, subtype HeapSubtype) *Attribute {
	this := Attribute{}
	this.Name = name
	this.Subtype = subtype
	var enable bool = true
	this.Enable = &enable
	var viewer bool = false
	this.Viewer = &viewer
	var ar bool = false
	this.Ar = &ar
	return &this
}

// NewAttributeWithDefaults instantiates a new Attribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeWithDefaults() *Attribute {
	this := Attribute{}
	var subtype HeapSubtype = INPUT
	this.Subtype = subtype
	var enable bool = true
	this.Enable = &enable
	var viewer bool = false
	this.Viewer = &viewer
	var ar bool = false
	this.Ar = &ar
	return &this
}

// GetAssetTypeName returns the AssetTypeName field value if set, zero value otherwise.
func (o *Attribute) GetAssetTypeName() string {
	if o == nil || o.AssetTypeName == nil {
		var ret string
		return ret
	}
	return *o.AssetTypeName
}

// GetAssetTypeNameOk returns a tuple with the AssetTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetAssetTypeNameOk() (*string, bool) {
	if o == nil || o.AssetTypeName == nil {
		return nil, false
	}
	return o.AssetTypeName, true
}

// HasAssetTypeName returns a boolean if a field has been set.
func (o *Attribute) HasAssetTypeName() bool {
	if o != nil && o.AssetTypeName != nil {
		return true
	}

	return false
}

// SetAssetTypeName gets a reference to the given string and assigns it to the AssetTypeName field.
func (o *Attribute) SetAssetTypeName(v string) {
	o.AssetTypeName = &v
}

// GetName returns the Name field value
func (o *Attribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Attribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Attribute) SetName(v string) {
	o.Name = v
}

// GetSubtype returns the Subtype field value
func (o *Attribute) GetSubtype() HeapSubtype {
	if o == nil {
		var ret HeapSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *Attribute) GetSubtypeOk() (*HeapSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *Attribute) SetSubtype(v HeapSubtype) {
	o.Subtype = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Attribute) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Attribute) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Attribute) SetType(v string) {
	o.Type = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *Attribute) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *Attribute) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *Attribute) SetEnable(v bool) {
	o.Enable = &v
}

// GetTranslation returns the Translation field value if set, zero value otherwise.
func (o *Attribute) GetTranslation() Translation {
	if o == nil || o.Translation == nil {
		var ret Translation
		return ret
	}
	return *o.Translation
}

// GetTranslationOk returns a tuple with the Translation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetTranslationOk() (*Translation, bool) {
	if o == nil || o.Translation == nil {
		return nil, false
	}
	return o.Translation, true
}

// HasTranslation returns a boolean if a field has been set.
func (o *Attribute) HasTranslation() bool {
	if o != nil && o.Translation != nil {
		return true
	}

	return false
}

// SetTranslation gets a reference to the given Translation and assigns it to the Translation field.
func (o *Attribute) SetTranslation(v Translation) {
	o.Translation = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Attribute) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Attribute) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Attribute) SetUnit(v string) {
	o.Unit = &v
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *Attribute) GetPrecision() int64 {
	if o == nil || o.Precision == nil {
		var ret int64
		return ret
	}
	return *o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetPrecisionOk() (*int64, bool) {
	if o == nil || o.Precision == nil {
		return nil, false
	}
	return o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *Attribute) HasPrecision() bool {
	if o != nil && o.Precision != nil {
		return true
	}

	return false
}

// SetPrecision gets a reference to the given int64 and assigns it to the Precision field.
func (o *Attribute) SetPrecision(v int64) {
	o.Precision = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Attribute) GetMin() float64 {
	if o == nil || o.Min == nil {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetMinOk() (*float64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Attribute) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *Attribute) SetMin(v float64) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Attribute) GetMax() float64 {
	if o == nil || o.Max == nil {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetMaxOk() (*float64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Attribute) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *Attribute) SetMax(v float64) {
	o.Max = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *Attribute) GetPipeline() AttributePipeline {
	if o == nil || o.Pipeline == nil {
		var ret AttributePipeline
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetPipelineOk() (*AttributePipeline, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *Attribute) HasPipeline() bool {
	if o != nil && o.Pipeline != nil {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given AttributePipeline and assigns it to the Pipeline field.
func (o *Attribute) SetPipeline(v AttributePipeline) {
	o.Pipeline = &v
}

// GetViewer returns the Viewer field value if set, zero value otherwise.
func (o *Attribute) GetViewer() bool {
	if o == nil || o.Viewer == nil {
		var ret bool
		return ret
	}
	return *o.Viewer
}

// GetViewerOk returns a tuple with the Viewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetViewerOk() (*bool, bool) {
	if o == nil || o.Viewer == nil {
		return nil, false
	}
	return o.Viewer, true
}

// HasViewer returns a boolean if a field has been set.
func (o *Attribute) HasViewer() bool {
	if o != nil && o.Viewer != nil {
		return true
	}

	return false
}

// SetViewer gets a reference to the given bool and assigns it to the Viewer field.
func (o *Attribute) SetViewer(v bool) {
	o.Viewer = &v
}

// GetAr returns the Ar field value if set, zero value otherwise.
func (o *Attribute) GetAr() bool {
	if o == nil || o.Ar == nil {
		var ret bool
		return ret
	}
	return *o.Ar
}

// GetArOk returns a tuple with the Ar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetArOk() (*bool, bool) {
	if o == nil || o.Ar == nil {
		return nil, false
	}
	return o.Ar, true
}

// HasAr returns a boolean if a field has been set.
func (o *Attribute) HasAr() bool {
	if o != nil && o.Ar != nil {
		return true
	}

	return false
}

// SetAr gets a reference to the given bool and assigns it to the Ar field.
func (o *Attribute) SetAr(v bool) {
	o.Ar = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *Attribute) GetSequence() int64 {
	if o == nil || o.Sequence == nil {
		var ret int64
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetSequenceOk() (*int64, bool) {
	if o == nil || o.Sequence == nil {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *Attribute) HasSequence() bool {
	if o != nil && o.Sequence != nil {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int64 and assigns it to the Sequence field.
func (o *Attribute) SetSequence(v int64) {
	o.Sequence = &v
}

// GetVirtual returns the Virtual field value if set, zero value otherwise.
func (o *Attribute) GetVirtual() bool {
	if o == nil || o.Virtual == nil {
		var ret bool
		return ret
	}
	return *o.Virtual
}

// GetVirtualOk returns a tuple with the Virtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetVirtualOk() (*bool, bool) {
	if o == nil || o.Virtual == nil {
		return nil, false
	}
	return o.Virtual, true
}

// HasVirtual returns a boolean if a field has been set.
func (o *Attribute) HasVirtual() bool {
	if o != nil && o.Virtual != nil {
		return true
	}

	return false
}

// SetVirtual gets a reference to the given bool and assigns it to the Virtual field.
func (o *Attribute) SetVirtual(v bool) {
	o.Virtual = &v
}

func (o Attribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetTypeName != nil {
		toSerialize["assetTypeName"] = o.AssetTypeName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.Translation != nil {
		toSerialize["translation"] = o.Translation
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Precision != nil {
		toSerialize["precision"] = o.Precision
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	if o.Viewer != nil {
		toSerialize["viewer"] = o.Viewer
	}
	if o.Ar != nil {
		toSerialize["ar"] = o.Ar
	}
	if o.Sequence != nil {
		toSerialize["sequence"] = o.Sequence
	}
	if o.Virtual != nil {
		toSerialize["virtual"] = o.Virtual
	}
	return json.Marshal(toSerialize)
}

type NullableAttribute struct {
	value *Attribute
	isSet bool
}

func (v NullableAttribute) Get() *Attribute {
	return v.value
}

func (v *NullableAttribute) Set(val *Attribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttribute(val *Attribute) *NullableAttribute {
	return &NullableAttribute{value: val, isSet: true}
}

func (v NullableAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
