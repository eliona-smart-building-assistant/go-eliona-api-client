/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.7.3
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AssetTypeAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetTypeAttribute{}

// AssetTypeAttribute Named attribute to store data of assets
type AssetTypeAttribute struct {
	// The unique name for the asset type
	AssetTypeName NullableString `json:"assetTypeName,omitempty"`
	// Unique key of asset data
	Name    string      `json:"name"`
	Subtype DataSubtype `json:"subtype"`
	// Name of the type for this attribute: air_quality, battery-voltage, brightness, co2, current, device-info, device-status, energy, flow, frequency, humidity, inputs-and-switches, level, motion, operating-status, people-count, power, presence, pressure, temperature, vehicle-detector, voltage, weather, voc
	Type NullableString `json:"type,omitempty"`
	// Is data active or not
	Enable      *bool               `json:"enable,omitempty"`
	Translation NullableTranslation `json:"translation,omitempty"`
	// Physical unit of numeric data
	Unit NullableString `json:"unit,omitempty"`
	// Number of decimal places
	Precision NullableInt64 `json:"precision,omitempty"`
	// Lower limit
	Min NullableFloat64 `json:"min,omitempty"`
	// Upper limit
	Max NullableFloat64 `json:"max,omitempty"`
	// Aggregation calculation mode
	AggregationMode    NullableString `json:"aggregationMode,omitempty"`
	AggregationRasters []string       `json:"aggregationRasters,omitempty"`
	// Should the attribute be displayed in viewer
	Viewer NullableBool `json:"viewer,omitempty"`
	// Should the attribute be displayed in AR
	Ar NullableBool `json:"ar,omitempty"`
	// Sequence in AR display
	Sequence NullableInt64 `json:"sequence,omitempty"`
	// Is the attribute virtual or not
	Virtual NullableBool `json:"virtual,omitempty"`
	// calculation rule to calculate the value for this attribute
	Formula NullableString `json:"formula,omitempty"`
	// value scale
	Scale NullableFloat32 `json:"scale,omitempty"`
	// value scale
	Zero NullableFloat32 `json:"zero,omitempty"`
	// list of mapping between value and custom text
	Map []map[string]interface{} `json:"map,omitempty"`
	// source path for attribute value
	SourcePath []string `json:"sourcePath,omitempty"`
	// is attribute digital
	IsDigital NullableBool `json:"isDigital,omitempty"`
}

type _AssetTypeAttribute AssetTypeAttribute

// NewAssetTypeAttribute instantiates a new AssetTypeAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTypeAttribute(name string, subtype DataSubtype) *AssetTypeAttribute {
	this := AssetTypeAttribute{}
	this.Name = name
	this.Subtype = subtype
	var enable bool = true
	this.Enable = &enable
	var viewer bool = false
	this.Viewer = *NewNullableBool(&viewer)
	var ar bool = false
	this.Ar = *NewNullableBool(&ar)
	return &this
}

// NewAssetTypeAttributeWithDefaults instantiates a new AssetTypeAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTypeAttributeWithDefaults() *AssetTypeAttribute {
	this := AssetTypeAttribute{}
	var subtype DataSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	var enable bool = true
	this.Enable = &enable
	var viewer bool = false
	this.Viewer = *NewNullableBool(&viewer)
	var ar bool = false
	this.Ar = *NewNullableBool(&ar)
	return &this
}

// GetAssetTypeName returns the AssetTypeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetAssetTypeName() string {
	if o == nil || IsNil(o.AssetTypeName.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTypeName.Get()
}

// GetAssetTypeNameOk returns a tuple with the AssetTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetAssetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTypeName.Get(), o.AssetTypeName.IsSet()
}

// HasAssetTypeName returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasAssetTypeName() bool {
	if o != nil && o.AssetTypeName.IsSet() {
		return true
	}

	return false
}

// SetAssetTypeName gets a reference to the given NullableString and assigns it to the AssetTypeName field.
func (o *AssetTypeAttribute) SetAssetTypeName(v string) {
	o.AssetTypeName.Set(&v)
}

// SetAssetTypeNameNil sets the value for AssetTypeName to be an explicit nil
func (o *AssetTypeAttribute) SetAssetTypeNameNil() {
	o.AssetTypeName.Set(nil)
}

// UnsetAssetTypeName ensures that no value is present for AssetTypeName, not even an explicit nil
func (o *AssetTypeAttribute) UnsetAssetTypeName() {
	o.AssetTypeName.Unset()
}

// GetName returns the Name field value
func (o *AssetTypeAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AssetTypeAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AssetTypeAttribute) SetName(v string) {
	o.Name = v
}

// GetSubtype returns the Subtype field value
func (o *AssetTypeAttribute) GetSubtype() DataSubtype {
	if o == nil {
		var ret DataSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *AssetTypeAttribute) GetSubtypeOk() (*DataSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *AssetTypeAttribute) SetSubtype(v DataSubtype) {
	o.Subtype = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *AssetTypeAttribute) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *AssetTypeAttribute) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *AssetTypeAttribute) UnsetType() {
	o.Type.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *AssetTypeAttribute) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTypeAttribute) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *AssetTypeAttribute) SetEnable(v bool) {
	o.Enable = &v
}

// GetTranslation returns the Translation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetTranslation() Translation {
	if o == nil || IsNil(o.Translation.Get()) {
		var ret Translation
		return ret
	}
	return *o.Translation.Get()
}

// GetTranslationOk returns a tuple with the Translation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetTranslationOk() (*Translation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Translation.Get(), o.Translation.IsSet()
}

// HasTranslation returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasTranslation() bool {
	if o != nil && o.Translation.IsSet() {
		return true
	}

	return false
}

// SetTranslation gets a reference to the given NullableTranslation and assigns it to the Translation field.
func (o *AssetTypeAttribute) SetTranslation(v Translation) {
	o.Translation.Set(&v)
}

// SetTranslationNil sets the value for Translation to be an explicit nil
func (o *AssetTypeAttribute) SetTranslationNil() {
	o.Translation.Set(nil)
}

// UnsetTranslation ensures that no value is present for Translation, not even an explicit nil
func (o *AssetTypeAttribute) UnsetTranslation() {
	o.Translation.Unset()
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetUnit() string {
	if o == nil || IsNil(o.Unit.Get()) {
		var ret string
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasUnit() bool {
	if o != nil && o.Unit.IsSet() {
		return true
	}

	return false
}

// SetUnit gets a reference to the given NullableString and assigns it to the Unit field.
func (o *AssetTypeAttribute) SetUnit(v string) {
	o.Unit.Set(&v)
}

// SetUnitNil sets the value for Unit to be an explicit nil
func (o *AssetTypeAttribute) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil
func (o *AssetTypeAttribute) UnsetUnit() {
	o.Unit.Unset()
}

// GetPrecision returns the Precision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetPrecision() int64 {
	if o == nil || IsNil(o.Precision.Get()) {
		var ret int64
		return ret
	}
	return *o.Precision.Get()
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetPrecisionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Precision.Get(), o.Precision.IsSet()
}

// HasPrecision returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasPrecision() bool {
	if o != nil && o.Precision.IsSet() {
		return true
	}

	return false
}

// SetPrecision gets a reference to the given NullableInt64 and assigns it to the Precision field.
func (o *AssetTypeAttribute) SetPrecision(v int64) {
	o.Precision.Set(&v)
}

// SetPrecisionNil sets the value for Precision to be an explicit nil
func (o *AssetTypeAttribute) SetPrecisionNil() {
	o.Precision.Set(nil)
}

// UnsetPrecision ensures that no value is present for Precision, not even an explicit nil
func (o *AssetTypeAttribute) UnsetPrecision() {
	o.Precision.Unset()
}

// GetMin returns the Min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetMin() float64 {
	if o == nil || IsNil(o.Min.Get()) {
		var ret float64
		return ret
	}
	return *o.Min.Get()
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetMinOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Min.Get(), o.Min.IsSet()
}

// HasMin returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasMin() bool {
	if o != nil && o.Min.IsSet() {
		return true
	}

	return false
}

// SetMin gets a reference to the given NullableFloat64 and assigns it to the Min field.
func (o *AssetTypeAttribute) SetMin(v float64) {
	o.Min.Set(&v)
}

// SetMinNil sets the value for Min to be an explicit nil
func (o *AssetTypeAttribute) SetMinNil() {
	o.Min.Set(nil)
}

// UnsetMin ensures that no value is present for Min, not even an explicit nil
func (o *AssetTypeAttribute) UnsetMin() {
	o.Min.Unset()
}

// GetMax returns the Max field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetMax() float64 {
	if o == nil || IsNil(o.Max.Get()) {
		var ret float64
		return ret
	}
	return *o.Max.Get()
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Max.Get(), o.Max.IsSet()
}

// HasMax returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasMax() bool {
	if o != nil && o.Max.IsSet() {
		return true
	}

	return false
}

// SetMax gets a reference to the given NullableFloat64 and assigns it to the Max field.
func (o *AssetTypeAttribute) SetMax(v float64) {
	o.Max.Set(&v)
}

// SetMaxNil sets the value for Max to be an explicit nil
func (o *AssetTypeAttribute) SetMaxNil() {
	o.Max.Set(nil)
}

// UnsetMax ensures that no value is present for Max, not even an explicit nil
func (o *AssetTypeAttribute) UnsetMax() {
	o.Max.Unset()
}

// GetAggregationMode returns the AggregationMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetAggregationMode() string {
	if o == nil || IsNil(o.AggregationMode.Get()) {
		var ret string
		return ret
	}
	return *o.AggregationMode.Get()
}

// GetAggregationModeOk returns a tuple with the AggregationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetAggregationModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AggregationMode.Get(), o.AggregationMode.IsSet()
}

// HasAggregationMode returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasAggregationMode() bool {
	if o != nil && o.AggregationMode.IsSet() {
		return true
	}

	return false
}

// SetAggregationMode gets a reference to the given NullableString and assigns it to the AggregationMode field.
func (o *AssetTypeAttribute) SetAggregationMode(v string) {
	o.AggregationMode.Set(&v)
}

// SetAggregationModeNil sets the value for AggregationMode to be an explicit nil
func (o *AssetTypeAttribute) SetAggregationModeNil() {
	o.AggregationMode.Set(nil)
}

// UnsetAggregationMode ensures that no value is present for AggregationMode, not even an explicit nil
func (o *AssetTypeAttribute) UnsetAggregationMode() {
	o.AggregationMode.Unset()
}

// GetAggregationRasters returns the AggregationRasters field value if set, zero value otherwise.
func (o *AssetTypeAttribute) GetAggregationRasters() []string {
	if o == nil || IsNil(o.AggregationRasters) {
		var ret []string
		return ret
	}
	return o.AggregationRasters
}

// GetAggregationRastersOk returns a tuple with the AggregationRasters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTypeAttribute) GetAggregationRastersOk() ([]string, bool) {
	if o == nil || IsNil(o.AggregationRasters) {
		return nil, false
	}
	return o.AggregationRasters, true
}

// HasAggregationRasters returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasAggregationRasters() bool {
	if o != nil && !IsNil(o.AggregationRasters) {
		return true
	}

	return false
}

// SetAggregationRasters gets a reference to the given []string and assigns it to the AggregationRasters field.
func (o *AssetTypeAttribute) SetAggregationRasters(v []string) {
	o.AggregationRasters = v
}

// GetViewer returns the Viewer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetViewer() bool {
	if o == nil || IsNil(o.Viewer.Get()) {
		var ret bool
		return ret
	}
	return *o.Viewer.Get()
}

// GetViewerOk returns a tuple with the Viewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetViewerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Viewer.Get(), o.Viewer.IsSet()
}

// HasViewer returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasViewer() bool {
	if o != nil && o.Viewer.IsSet() {
		return true
	}

	return false
}

// SetViewer gets a reference to the given NullableBool and assigns it to the Viewer field.
func (o *AssetTypeAttribute) SetViewer(v bool) {
	o.Viewer.Set(&v)
}

// SetViewerNil sets the value for Viewer to be an explicit nil
func (o *AssetTypeAttribute) SetViewerNil() {
	o.Viewer.Set(nil)
}

// UnsetViewer ensures that no value is present for Viewer, not even an explicit nil
func (o *AssetTypeAttribute) UnsetViewer() {
	o.Viewer.Unset()
}

// GetAr returns the Ar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetAr() bool {
	if o == nil || IsNil(o.Ar.Get()) {
		var ret bool
		return ret
	}
	return *o.Ar.Get()
}

// GetArOk returns a tuple with the Ar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetArOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ar.Get(), o.Ar.IsSet()
}

// HasAr returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasAr() bool {
	if o != nil && o.Ar.IsSet() {
		return true
	}

	return false
}

// SetAr gets a reference to the given NullableBool and assigns it to the Ar field.
func (o *AssetTypeAttribute) SetAr(v bool) {
	o.Ar.Set(&v)
}

// SetArNil sets the value for Ar to be an explicit nil
func (o *AssetTypeAttribute) SetArNil() {
	o.Ar.Set(nil)
}

// UnsetAr ensures that no value is present for Ar, not even an explicit nil
func (o *AssetTypeAttribute) UnsetAr() {
	o.Ar.Unset()
}

// GetSequence returns the Sequence field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetSequence() int64 {
	if o == nil || IsNil(o.Sequence.Get()) {
		var ret int64
		return ret
	}
	return *o.Sequence.Get()
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetSequenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sequence.Get(), o.Sequence.IsSet()
}

// HasSequence returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasSequence() bool {
	if o != nil && o.Sequence.IsSet() {
		return true
	}

	return false
}

// SetSequence gets a reference to the given NullableInt64 and assigns it to the Sequence field.
func (o *AssetTypeAttribute) SetSequence(v int64) {
	o.Sequence.Set(&v)
}

// SetSequenceNil sets the value for Sequence to be an explicit nil
func (o *AssetTypeAttribute) SetSequenceNil() {
	o.Sequence.Set(nil)
}

// UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
func (o *AssetTypeAttribute) UnsetSequence() {
	o.Sequence.Unset()
}

// GetVirtual returns the Virtual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetVirtual() bool {
	if o == nil || IsNil(o.Virtual.Get()) {
		var ret bool
		return ret
	}
	return *o.Virtual.Get()
}

// GetVirtualOk returns a tuple with the Virtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetVirtualOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Virtual.Get(), o.Virtual.IsSet()
}

// HasVirtual returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasVirtual() bool {
	if o != nil && o.Virtual.IsSet() {
		return true
	}

	return false
}

// SetVirtual gets a reference to the given NullableBool and assigns it to the Virtual field.
func (o *AssetTypeAttribute) SetVirtual(v bool) {
	o.Virtual.Set(&v)
}

// SetVirtualNil sets the value for Virtual to be an explicit nil
func (o *AssetTypeAttribute) SetVirtualNil() {
	o.Virtual.Set(nil)
}

// UnsetVirtual ensures that no value is present for Virtual, not even an explicit nil
func (o *AssetTypeAttribute) UnsetVirtual() {
	o.Virtual.Unset()
}

// GetFormula returns the Formula field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetFormula() string {
	if o == nil || IsNil(o.Formula.Get()) {
		var ret string
		return ret
	}
	return *o.Formula.Get()
}

// GetFormulaOk returns a tuple with the Formula field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetFormulaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Formula.Get(), o.Formula.IsSet()
}

// HasFormula returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasFormula() bool {
	if o != nil && o.Formula.IsSet() {
		return true
	}

	return false
}

// SetFormula gets a reference to the given NullableString and assigns it to the Formula field.
func (o *AssetTypeAttribute) SetFormula(v string) {
	o.Formula.Set(&v)
}

// SetFormulaNil sets the value for Formula to be an explicit nil
func (o *AssetTypeAttribute) SetFormulaNil() {
	o.Formula.Set(nil)
}

// UnsetFormula ensures that no value is present for Formula, not even an explicit nil
func (o *AssetTypeAttribute) UnsetFormula() {
	o.Formula.Unset()
}

// GetScale returns the Scale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetScale() float32 {
	if o == nil || IsNil(o.Scale.Get()) {
		var ret float32
		return ret
	}
	return *o.Scale.Get()
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetScaleOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scale.Get(), o.Scale.IsSet()
}

// HasScale returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasScale() bool {
	if o != nil && o.Scale.IsSet() {
		return true
	}

	return false
}

// SetScale gets a reference to the given NullableFloat32 and assigns it to the Scale field.
func (o *AssetTypeAttribute) SetScale(v float32) {
	o.Scale.Set(&v)
}

// SetScaleNil sets the value for Scale to be an explicit nil
func (o *AssetTypeAttribute) SetScaleNil() {
	o.Scale.Set(nil)
}

// UnsetScale ensures that no value is present for Scale, not even an explicit nil
func (o *AssetTypeAttribute) UnsetScale() {
	o.Scale.Unset()
}

// GetZero returns the Zero field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetZero() float32 {
	if o == nil || IsNil(o.Zero.Get()) {
		var ret float32
		return ret
	}
	return *o.Zero.Get()
}

// GetZeroOk returns a tuple with the Zero field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetZeroOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zero.Get(), o.Zero.IsSet()
}

// HasZero returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasZero() bool {
	if o != nil && o.Zero.IsSet() {
		return true
	}

	return false
}

// SetZero gets a reference to the given NullableFloat32 and assigns it to the Zero field.
func (o *AssetTypeAttribute) SetZero(v float32) {
	o.Zero.Set(&v)
}

// SetZeroNil sets the value for Zero to be an explicit nil
func (o *AssetTypeAttribute) SetZeroNil() {
	o.Zero.Set(nil)
}

// UnsetZero ensures that no value is present for Zero, not even an explicit nil
func (o *AssetTypeAttribute) UnsetZero() {
	o.Zero.Unset()
}

// GetMap returns the Map field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetMap() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Map
}

// GetMapOk returns a tuple with the Map field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetMapOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Map) {
		return nil, false
	}
	return o.Map, true
}

// HasMap returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasMap() bool {
	if o != nil && IsNil(o.Map) {
		return true
	}

	return false
}

// SetMap gets a reference to the given []map[string]interface{} and assigns it to the Map field.
func (o *AssetTypeAttribute) SetMap(v []map[string]interface{}) {
	o.Map = v
}

// GetSourcePath returns the SourcePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetSourcePath() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SourcePath
}

// GetSourcePathOk returns a tuple with the SourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetSourcePathOk() ([]string, bool) {
	if o == nil || IsNil(o.SourcePath) {
		return nil, false
	}
	return o.SourcePath, true
}

// HasSourcePath returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasSourcePath() bool {
	if o != nil && IsNil(o.SourcePath) {
		return true
	}

	return false
}

// SetSourcePath gets a reference to the given []string and assigns it to the SourcePath field.
func (o *AssetTypeAttribute) SetSourcePath(v []string) {
	o.SourcePath = v
}

// GetIsDigital returns the IsDigital field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetTypeAttribute) GetIsDigital() bool {
	if o == nil || IsNil(o.IsDigital.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDigital.Get()
}

// GetIsDigitalOk returns a tuple with the IsDigital field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetTypeAttribute) GetIsDigitalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDigital.Get(), o.IsDigital.IsSet()
}

// HasIsDigital returns a boolean if a field has been set.
func (o *AssetTypeAttribute) HasIsDigital() bool {
	if o != nil && o.IsDigital.IsSet() {
		return true
	}

	return false
}

// SetIsDigital gets a reference to the given NullableBool and assigns it to the IsDigital field.
func (o *AssetTypeAttribute) SetIsDigital(v bool) {
	o.IsDigital.Set(&v)
}

// SetIsDigitalNil sets the value for IsDigital to be an explicit nil
func (o *AssetTypeAttribute) SetIsDigitalNil() {
	o.IsDigital.Set(nil)
}

// UnsetIsDigital ensures that no value is present for IsDigital, not even an explicit nil
func (o *AssetTypeAttribute) UnsetIsDigital() {
	o.IsDigital.Unset()
}

func (o AssetTypeAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetTypeAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetTypeName.IsSet() {
		toSerialize["assetTypeName"] = o.AssetTypeName.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["subtype"] = o.Subtype
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if o.Translation.IsSet() {
		toSerialize["translation"] = o.Translation.Get()
	}
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}
	if o.Precision.IsSet() {
		toSerialize["precision"] = o.Precision.Get()
	}
	if o.Min.IsSet() {
		toSerialize["min"] = o.Min.Get()
	}
	if o.Max.IsSet() {
		toSerialize["max"] = o.Max.Get()
	}
	if o.AggregationMode.IsSet() {
		toSerialize["aggregationMode"] = o.AggregationMode.Get()
	}
	if !IsNil(o.AggregationRasters) {
		toSerialize["aggregationRasters"] = o.AggregationRasters
	}
	if o.Viewer.IsSet() {
		toSerialize["viewer"] = o.Viewer.Get()
	}
	if o.Ar.IsSet() {
		toSerialize["ar"] = o.Ar.Get()
	}
	if o.Sequence.IsSet() {
		toSerialize["sequence"] = o.Sequence.Get()
	}
	if o.Virtual.IsSet() {
		toSerialize["virtual"] = o.Virtual.Get()
	}
	if o.Formula.IsSet() {
		toSerialize["formula"] = o.Formula.Get()
	}
	if o.Scale.IsSet() {
		toSerialize["scale"] = o.Scale.Get()
	}
	if o.Zero.IsSet() {
		toSerialize["zero"] = o.Zero.Get()
	}
	if o.Map != nil {
		toSerialize["map"] = o.Map
	}
	if o.SourcePath != nil {
		toSerialize["sourcePath"] = o.SourcePath
	}
	if o.IsDigital.IsSet() {
		toSerialize["isDigital"] = o.IsDigital.Get()
	}
	return toSerialize, nil
}

func (o *AssetTypeAttribute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"subtype",
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

	varAssetTypeAttribute := _AssetTypeAttribute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssetTypeAttribute)

	if err != nil {
		return err
	}

	*o = AssetTypeAttribute(varAssetTypeAttribute)

	return err
}

type NullableAssetTypeAttribute struct {
	value *AssetTypeAttribute
	isSet bool
}

func (v NullableAssetTypeAttribute) Get() *AssetTypeAttribute {
	return v.value
}

func (v *NullableAssetTypeAttribute) Set(val *AssetTypeAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTypeAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTypeAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTypeAttribute(val *AssetTypeAttribute) *NullableAssetTypeAttribute {
	return &NullableAssetTypeAttribute{value: val, isSet: true}
}

func (v NullableAssetTypeAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTypeAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
