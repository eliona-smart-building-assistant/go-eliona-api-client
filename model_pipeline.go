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

// Pipeline struct for Pipeline
type Pipeline struct {
	// Pipeline calculation mode
	Mode *string `json:"mode,omitempty"`
	// Pipeline calculation intervals
	Raster []string `json:"raster,omitempty"`
}

// NewPipeline instantiates a new Pipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipeline() *Pipeline {
	this := Pipeline{}
	return &this
}

// NewPipelineWithDefaults instantiates a new Pipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineWithDefaults() *Pipeline {
	this := Pipeline{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *Pipeline) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pipeline) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *Pipeline) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *Pipeline) SetMode(v string) {
	o.Mode = &v
}

// GetRaster returns the Raster field value if set, zero value otherwise.
func (o *Pipeline) GetRaster() []string {
	if o == nil || o.Raster == nil {
		var ret []string
		return ret
	}
	return o.Raster
}

// GetRasterOk returns a tuple with the Raster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pipeline) GetRasterOk() ([]string, bool) {
	if o == nil || o.Raster == nil {
		return nil, false
	}
	return o.Raster, true
}

// HasRaster returns a boolean if a field has been set.
func (o *Pipeline) HasRaster() bool {
	if o != nil && o.Raster != nil {
		return true
	}

	return false
}

// SetRaster gets a reference to the given []string and assigns it to the Raster field.
func (o *Pipeline) SetRaster(v []string) {
	o.Raster = v
}

func (o Pipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Raster != nil {
		toSerialize["raster"] = o.Raster
	}
	return json.Marshal(toSerialize)
}

type NullablePipeline struct {
	value *Pipeline
	isSet bool
}

func (v NullablePipeline) Get() *Pipeline {
	return v.value
}

func (v *NullablePipeline) Set(val *Pipeline) {
	v.value = val
	v.isSet = true
}

func (v NullablePipeline) IsSet() bool {
	return v.isSet
}

func (v *NullablePipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipeline(val *Pipeline) *NullablePipeline {
	return &NullablePipeline{value: val, isSet: true}
}

func (v NullablePipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

