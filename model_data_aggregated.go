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
	"time"
)

// checks if the DataAggregated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataAggregated{}

// DataAggregated Aggregated data combines multiple data points for a periodical raster
type DataAggregated struct {
	// ID of the aggregation
	AggregationId *int32 `json:"aggregationId,omitempty"`
	// ID of the corresponding asset
	AssetId int32       `json:"assetId"`
	Subtype DataSubtype `json:"subtype"`
	// Name of the attribute which holds the data points
	Attribute *string `json:"attribute,omitempty"`
	// Calculation intervals.
	Raster string `json:"raster"`
	// Timestamp of this aggregated data set. The timestamp when the timeslot based on raster starts.
	Timestamp NullableTime `json:"timestamp,omitempty"`
	// Count of data points in this aggregated data set
	Count NullableFloat64 `json:"count,omitempty"`
	// Average of all data points for this aggregated data set
	Average NullableFloat64 `json:"average,omitempty"`
	// Sum of all data points for this aggregated data set
	Sum NullableFloat64 `json:"sum,omitempty"`
	// First data point in this aggregated data set
	First NullableFloat64 `json:"first,omitempty"`
	// Data point with the most minimal value in this aggregated data set
	Min NullableFloat64 `json:"min,omitempty"`
	// Data point with the most maximal value in this aggregated data set
	Max NullableFloat64 `json:"max,omitempty"`
	// Latest data point in this aggregated data set
	Last NullableFloat64 `json:"last,omitempty"`
	// Timestamp of the latest data point
	LastTimestamp NullableTime `json:"lastTimestamp,omitempty"`
	// The name of the corresponding asset type
	AssetTypeName NullableString `json:"assetTypeName,omitempty"`
}

type _DataAggregated DataAggregated

// NewDataAggregated instantiates a new DataAggregated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataAggregated(assetId int32, subtype DataSubtype, raster string) *DataAggregated {
	this := DataAggregated{}
	this.AssetId = assetId
	this.Subtype = subtype
	this.Raster = raster
	return &this
}

// NewDataAggregatedWithDefaults instantiates a new DataAggregated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataAggregatedWithDefaults() *DataAggregated {
	this := DataAggregated{}
	var subtype DataSubtype = SUBTYPE_INPUT
	this.Subtype = subtype
	return &this
}

// GetAggregationId returns the AggregationId field value if set, zero value otherwise.
func (o *DataAggregated) GetAggregationId() int32 {
	if o == nil || IsNil(o.AggregationId) {
		var ret int32
		return ret
	}
	return *o.AggregationId
}

// GetAggregationIdOk returns a tuple with the AggregationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAggregated) GetAggregationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AggregationId) {
		return nil, false
	}
	return o.AggregationId, true
}

// HasAggregationId returns a boolean if a field has been set.
func (o *DataAggregated) HasAggregationId() bool {
	if o != nil && !IsNil(o.AggregationId) {
		return true
	}

	return false
}

// SetAggregationId gets a reference to the given int32 and assigns it to the AggregationId field.
func (o *DataAggregated) SetAggregationId(v int32) {
	o.AggregationId = &v
}

// GetAssetId returns the AssetId field value
func (o *DataAggregated) GetAssetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *DataAggregated) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *DataAggregated) SetAssetId(v int32) {
	o.AssetId = v
}

// GetSubtype returns the Subtype field value
func (o *DataAggregated) GetSubtype() DataSubtype {
	if o == nil {
		var ret DataSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *DataAggregated) GetSubtypeOk() (*DataSubtype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *DataAggregated) SetSubtype(v DataSubtype) {
	o.Subtype = v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *DataAggregated) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAggregated) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *DataAggregated) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *DataAggregated) SetAttribute(v string) {
	o.Attribute = &v
}

// GetRaster returns the Raster field value
func (o *DataAggregated) GetRaster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Raster
}

// GetRasterOk returns a tuple with the Raster field value
// and a boolean to check if the value has been set.
func (o *DataAggregated) GetRasterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Raster, true
}

// SetRaster sets field value
func (o *DataAggregated) SetRaster(v string) {
	o.Raster = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataAggregated) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataAggregated) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DataAggregated) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *DataAggregated) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *DataAggregated) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *DataAggregated) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataAggregated) GetCount() float64 {
	if o == nil || IsNil(o.Count.Get()) {
		var ret float64
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataAggregated) GetCountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *DataAggregated) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableFloat64 and assigns it to the Count field.
func (o *DataAggregated) SetCount(v float64) {
	o.Count.Set(&v)
}

// SetCountNil sets the value for Count to be an explicit nil
func (o *DataAggregated) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *DataAggregated) UnsetCount() {
	o.Count.Unset()
}

// GetAverage returns the Average field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataAggregated) GetAverage() float64 {
	if o == nil || IsNil(o.Average.Get()) {
		var ret float64
		return ret
	}
	return *o.Average.Get()
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataAggregated) GetAverageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Average.Get(), o.Average.IsSet()
}

// HasAverage returns a boolean if a field has been set.
func (o *DataAggregated) HasAverage() bool {
	if o != nil && o.Average.IsSet() {
		return true
	}

	return false
}

// SetAverage gets a reference to the given NullableFloat64 and assigns it to the Average field.
func (o *DataAggregated) SetAverage(v float64) {
	o.Average.Set(&v)
}

// SetAverageNil sets the value for Average to be an explicit nil
func (o *DataAggregated) SetAverageNil() {
	o.Average.Set(nil)
}

// UnsetAverage ensures that no value is present for Average, not even an explicit nil
func (o *DataAggregated) UnsetAverage() {
	o.Average.Unset()
}

// GetSum returns the Sum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataAggregated) GetSum() float64 {
	if o == nil || IsNil(o.Sum.Get()) {
		var ret float64
		return ret
	}
	return *o.Sum.Get()
}

// GetSumOk returns a tuple with the Sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataAggregated) GetSumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sum.Get(), o.Sum.IsSet()
}

// HasSum returns a boolean if a field has been set.
func (o *DataAggregated) HasSum() bool {
	if o != nil && o.Sum.IsSet() {
		return true
	}

	return false
}

// SetSum gets a reference to the given NullableFloat64 and assigns it to the Sum field.
func (o *DataAggregated) SetSum(v float64) {
	o.Sum.Set(&v)
}

// SetSumNil sets the value for Sum to be an explicit nil
func (o *DataAggregated) SetSumNil() {
	o.Sum.Set(nil)
}

// UnsetSum ensures that no value is present for Sum, not even an explicit nil
func (o *DataAggregated) UnsetSum() {
	o.Sum.Unset()
}

// GetFirst returns the First field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataAggregated) GetFirst() float64 {
	if o == nil || IsNil(o.First.Get()) {
		var ret float64
		return ret
	}
	return *o.First.Get()
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataAggregated) GetFirstOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.First.Get(), o.First.IsSet()
}

// HasFirst returns a boolean if a field has been set.
func (o *DataAggregated) HasFirst() bool {
	if o != nil && o.First.IsSet() {
		return true
	}

	return false
}

// SetFirst gets a reference to the given NullableFloat64 and assigns it to the First field.
func (o *DataAggregated) SetFirst(v float64) {
	o.First.Set(&v)
}

// SetFirstNil sets the value for First to be an explicit nil
func (o *DataAggregated) SetFirstNil() {
	o.First.Set(nil)
}

// UnsetFirst ensures that no value is present for First, not even an explicit nil
func (o *DataAggregated) UnsetFirst() {
	o.First.Unset()
}

// GetMin returns the Min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataAggregated) GetMin() float64 {
	if o == nil || IsNil(o.Min.Get()) {
		var ret float64
		return ret
	}
	return *o.Min.Get()
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataAggregated) GetMinOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Min.Get(), o.Min.IsSet()
}

// HasMin returns a boolean if a field has been set.
func (o *DataAggregated) HasMin() bool {
	if o != nil && o.Min.IsSet() {
		return true
	}

	return false
}

// SetMin gets a reference to the given NullableFloat64 and assigns it to the Min field.
func (o *DataAggregated) SetMin(v float64) {
	o.Min.Set(&v)
}

// SetMinNil sets the value for Min to be an explicit nil
func (o *DataAggregated) SetMinNil() {
	o.Min.Set(nil)
}

// UnsetMin ensures that no value is present for Min, not even an explicit nil
func (o *DataAggregated) UnsetMin() {
	o.Min.Unset()
}

// GetMax returns the Max field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataAggregated) GetMax() float64 {
	if o == nil || IsNil(o.Max.Get()) {
		var ret float64
		return ret
	}
	return *o.Max.Get()
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataAggregated) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Max.Get(), o.Max.IsSet()
}

// HasMax returns a boolean if a field has been set.
func (o *DataAggregated) HasMax() bool {
	if o != nil && o.Max.IsSet() {
		return true
	}

	return false
}

// SetMax gets a reference to the given NullableFloat64 and assigns it to the Max field.
func (o *DataAggregated) SetMax(v float64) {
	o.Max.Set(&v)
}

// SetMaxNil sets the value for Max to be an explicit nil
func (o *DataAggregated) SetMaxNil() {
	o.Max.Set(nil)
}

// UnsetMax ensures that no value is present for Max, not even an explicit nil
func (o *DataAggregated) UnsetMax() {
	o.Max.Unset()
}

// GetLast returns the Last field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataAggregated) GetLast() float64 {
	if o == nil || IsNil(o.Last.Get()) {
		var ret float64
		return ret
	}
	return *o.Last.Get()
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataAggregated) GetLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Last.Get(), o.Last.IsSet()
}

// HasLast returns a boolean if a field has been set.
func (o *DataAggregated) HasLast() bool {
	if o != nil && o.Last.IsSet() {
		return true
	}

	return false
}

// SetLast gets a reference to the given NullableFloat64 and assigns it to the Last field.
func (o *DataAggregated) SetLast(v float64) {
	o.Last.Set(&v)
}

// SetLastNil sets the value for Last to be an explicit nil
func (o *DataAggregated) SetLastNil() {
	o.Last.Set(nil)
}

// UnsetLast ensures that no value is present for Last, not even an explicit nil
func (o *DataAggregated) UnsetLast() {
	o.Last.Unset()
}

// GetLastTimestamp returns the LastTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataAggregated) GetLastTimestamp() time.Time {
	if o == nil || IsNil(o.LastTimestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastTimestamp.Get()
}

// GetLastTimestampOk returns a tuple with the LastTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataAggregated) GetLastTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTimestamp.Get(), o.LastTimestamp.IsSet()
}

// HasLastTimestamp returns a boolean if a field has been set.
func (o *DataAggregated) HasLastTimestamp() bool {
	if o != nil && o.LastTimestamp.IsSet() {
		return true
	}

	return false
}

// SetLastTimestamp gets a reference to the given NullableTime and assigns it to the LastTimestamp field.
func (o *DataAggregated) SetLastTimestamp(v time.Time) {
	o.LastTimestamp.Set(&v)
}

// SetLastTimestampNil sets the value for LastTimestamp to be an explicit nil
func (o *DataAggregated) SetLastTimestampNil() {
	o.LastTimestamp.Set(nil)
}

// UnsetLastTimestamp ensures that no value is present for LastTimestamp, not even an explicit nil
func (o *DataAggregated) UnsetLastTimestamp() {
	o.LastTimestamp.Unset()
}

// GetAssetTypeName returns the AssetTypeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataAggregated) GetAssetTypeName() string {
	if o == nil || IsNil(o.AssetTypeName.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTypeName.Get()
}

// GetAssetTypeNameOk returns a tuple with the AssetTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataAggregated) GetAssetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTypeName.Get(), o.AssetTypeName.IsSet()
}

// HasAssetTypeName returns a boolean if a field has been set.
func (o *DataAggregated) HasAssetTypeName() bool {
	if o != nil && o.AssetTypeName.IsSet() {
		return true
	}

	return false
}

// SetAssetTypeName gets a reference to the given NullableString and assigns it to the AssetTypeName field.
func (o *DataAggregated) SetAssetTypeName(v string) {
	o.AssetTypeName.Set(&v)
}

// SetAssetTypeNameNil sets the value for AssetTypeName to be an explicit nil
func (o *DataAggregated) SetAssetTypeNameNil() {
	o.AssetTypeName.Set(nil)
}

// UnsetAssetTypeName ensures that no value is present for AssetTypeName, not even an explicit nil
func (o *DataAggregated) UnsetAssetTypeName() {
	o.AssetTypeName.Unset()
}

func (o DataAggregated) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataAggregated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationId) {
		toSerialize["aggregationId"] = o.AggregationId
	}
	toSerialize["assetId"] = o.AssetId
	toSerialize["subtype"] = o.Subtype
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	toSerialize["raster"] = o.Raster
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.Average.IsSet() {
		toSerialize["average"] = o.Average.Get()
	}
	if o.Sum.IsSet() {
		toSerialize["sum"] = o.Sum.Get()
	}
	if o.First.IsSet() {
		toSerialize["first"] = o.First.Get()
	}
	if o.Min.IsSet() {
		toSerialize["min"] = o.Min.Get()
	}
	if o.Max.IsSet() {
		toSerialize["max"] = o.Max.Get()
	}
	if o.Last.IsSet() {
		toSerialize["last"] = o.Last.Get()
	}
	if o.LastTimestamp.IsSet() {
		toSerialize["lastTimestamp"] = o.LastTimestamp.Get()
	}
	if o.AssetTypeName.IsSet() {
		toSerialize["assetTypeName"] = o.AssetTypeName.Get()
	}
	return toSerialize, nil
}

func (o *DataAggregated) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assetId",
		"subtype",
		"raster",
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

	varDataAggregated := _DataAggregated{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataAggregated)

	if err != nil {
		return err
	}

	*o = DataAggregated(varDataAggregated)

	return err
}

type NullableDataAggregated struct {
	value *DataAggregated
	isSet bool
}

func (v NullableDataAggregated) Get() *DataAggregated {
	return v.value
}

func (v *NullableDataAggregated) Set(val *DataAggregated) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAggregated) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAggregated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAggregated(val *DataAggregated) *NullableDataAggregated {
	return &NullableDataAggregated{value: val, isSet: true}
}

func (v NullableDataAggregated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAggregated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
