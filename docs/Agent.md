# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | Unique id for the agent | [optional] [readonly] 
**NodeId** | Pointer to **NullableString** | Id of the node where the agent is installed | [optional] 
**Class** | Pointer to [**NullableAgentClass**](AgentClass.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Descriptive text for the agent | [optional] 
**Enable** | Pointer to **bool** | Is the agent enabled or not | [optional] [default to false]
**Config** | Pointer to **map[string]interface{}** | Individual configuration depending on agent class | [optional] 
**AssetId** | Pointer to **NullableInt32** | ID of the corresponding asset | [optional] 

## Methods

### NewAgent

`func NewAgent() *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Agent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Agent) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Agent) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Agent) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNodeId

`func (o *Agent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Agent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Agent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *Agent) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *Agent) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *Agent) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetClass

`func (o *Agent) GetClass() AgentClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Agent) GetClassOk() (*AgentClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Agent) SetClass(v AgentClass)`

SetClass sets Class field to given value.

### HasClass

`func (o *Agent) HasClass() bool`

HasClass returns a boolean if a field has been set.

### SetClassNil

`func (o *Agent) SetClassNil(b bool)`

 SetClassNil sets the value for Class to be an explicit nil

### UnsetClass
`func (o *Agent) UnsetClass()`

UnsetClass ensures that no value is present for Class, not even an explicit nil
### GetDescription

`func (o *Agent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Agent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Agent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Agent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnable

`func (o *Agent) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Agent) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Agent) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *Agent) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetConfig

`func (o *Agent) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Agent) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Agent) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Agent) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *Agent) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *Agent) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetAssetId

`func (o *Agent) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Agent) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Agent) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Agent) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *Agent) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *Agent) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


