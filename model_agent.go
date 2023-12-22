/*
Eliona REST API

The Eliona REST API enables unified access to the resources and data of an Eliona environment.

API version: 2.5.7
Contact: hello@eliona.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the Agent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Agent{}

// Agent An agent installed on an edge node
type Agent struct {
	// Unique id for the agent
	Id NullableInt32 `json:"id,omitempty"`
	// Id of the node where the agent is installed
	NodeId NullableString `json:"nodeId,omitempty"`
	// ID of the corresponding asset
	AssetId NullableInt32      `json:"assetId,omitempty"`
	Class   NullableAgentClass `json:"class,omitempty"`
	// Descriptive text for the agent
	Description NullableString `json:"description,omitempty"`
	// Is the agent enabled or not
	Enable *bool `json:"enable,omitempty"`
	// Individual configuration depending on agent class
	Config map[string]interface{} `json:"config,omitempty"`
}

// NewAgent instantiates a new Agent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgent() *Agent {
	this := Agent{}
	var enable bool = false
	this.Enable = &enable
	return &this
}

// NewAgentWithDefaults instantiates a new Agent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWithDefaults() *Agent {
	this := Agent{}
	var enable bool = false
	this.Enable = &enable
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agent) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agent) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Agent) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Agent) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *Agent) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Agent) UnsetId() {
	o.Id.Unset()
}

// GetNodeId returns the NodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agent) GetNodeId() string {
	if o == nil || IsNil(o.NodeId.Get()) {
		var ret string
		return ret
	}
	return *o.NodeId.Get()
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agent) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeId.Get(), o.NodeId.IsSet()
}

// HasNodeId returns a boolean if a field has been set.
func (o *Agent) HasNodeId() bool {
	if o != nil && o.NodeId.IsSet() {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given NullableString and assigns it to the NodeId field.
func (o *Agent) SetNodeId(v string) {
	o.NodeId.Set(&v)
}

// SetNodeIdNil sets the value for NodeId to be an explicit nil
func (o *Agent) SetNodeIdNil() {
	o.NodeId.Set(nil)
}

// UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
func (o *Agent) UnsetNodeId() {
	o.NodeId.Unset()
}

// GetAssetId returns the AssetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agent) GetAssetId() int32 {
	if o == nil || IsNil(o.AssetId.Get()) {
		var ret int32
		return ret
	}
	return *o.AssetId.Get()
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agent) GetAssetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetId.Get(), o.AssetId.IsSet()
}

// HasAssetId returns a boolean if a field has been set.
func (o *Agent) HasAssetId() bool {
	if o != nil && o.AssetId.IsSet() {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given NullableInt32 and assigns it to the AssetId field.
func (o *Agent) SetAssetId(v int32) {
	o.AssetId.Set(&v)
}

// SetAssetIdNil sets the value for AssetId to be an explicit nil
func (o *Agent) SetAssetIdNil() {
	o.AssetId.Set(nil)
}

// UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
func (o *Agent) UnsetAssetId() {
	o.AssetId.Unset()
}

// GetClass returns the Class field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agent) GetClass() AgentClass {
	if o == nil || IsNil(o.Class.Get()) {
		var ret AgentClass
		return ret
	}
	return *o.Class.Get()
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agent) GetClassOk() (*AgentClass, bool) {
	if o == nil {
		return nil, false
	}
	return o.Class.Get(), o.Class.IsSet()
}

// HasClass returns a boolean if a field has been set.
func (o *Agent) HasClass() bool {
	if o != nil && o.Class.IsSet() {
		return true
	}

	return false
}

// SetClass gets a reference to the given NullableAgentClass and assigns it to the Class field.
func (o *Agent) SetClass(v AgentClass) {
	o.Class.Set(&v)
}

// SetClassNil sets the value for Class to be an explicit nil
func (o *Agent) SetClassNil() {
	o.Class.Set(nil)
}

// UnsetClass ensures that no value is present for Class, not even an explicit nil
func (o *Agent) UnsetClass() {
	o.Class.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agent) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agent) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Agent) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Agent) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Agent) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Agent) UnsetDescription() {
	o.Description.Unset()
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *Agent) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *Agent) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *Agent) SetEnable(v bool) {
	o.Enable = &v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Agent) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Agent) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Agent) HasConfig() bool {
	if o != nil && IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *Agent) SetConfig(v map[string]interface{}) {
	o.Config = v
}

func (o Agent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Agent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.NodeId.IsSet() {
		toSerialize["nodeId"] = o.NodeId.Get()
	}
	if o.AssetId.IsSet() {
		toSerialize["assetId"] = o.AssetId.Get()
	}
	if o.Class.IsSet() {
		toSerialize["class"] = o.Class.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableAgent struct {
	value *Agent
	isSet bool
}

func (v NullableAgent) Get() *Agent {
	return v.value
}

func (v *NullableAgent) Set(val *Agent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgent(val *Agent) *NullableAgent {
	return &NullableAgent{value: val, isSet: true}
}

func (v NullableAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
