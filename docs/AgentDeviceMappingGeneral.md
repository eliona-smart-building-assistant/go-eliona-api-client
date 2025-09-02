# AgentDeviceMappingGeneral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to [**NullableAgentClass**](AgentClass.md) |  | [optional] 
**Id** | Pointer to **NullableInt32** | Unique id for the mapping | [optional] [readonly] 
**DeviceId** | Pointer to **NullableInt32** | The id of the device the mapping belongs to | [optional] [readonly] 
**Enable** | Pointer to **bool** | Is the mapping enabled or not | [optional] [default to true]
**AssetId** | Pointer to **NullableInt32** | ID of the corresponding asset | [optional] 
**Subtype** | [**DataSubtype**](DataSubtype.md) |  | [default to INPUT]
**Attribute** | **string** | Name of the attribute to map | 

## Methods

### NewAgentDeviceMappingGeneral

`func NewAgentDeviceMappingGeneral(subtype DataSubtype, attribute string, ) *AgentDeviceMappingGeneral`

NewAgentDeviceMappingGeneral instantiates a new AgentDeviceMappingGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentDeviceMappingGeneralWithDefaults

`func NewAgentDeviceMappingGeneralWithDefaults() *AgentDeviceMappingGeneral`

NewAgentDeviceMappingGeneralWithDefaults instantiates a new AgentDeviceMappingGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *AgentDeviceMappingGeneral) GetClass() AgentClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *AgentDeviceMappingGeneral) GetClassOk() (*AgentClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *AgentDeviceMappingGeneral) SetClass(v AgentClass)`

SetClass sets Class field to given value.

### HasClass

`func (o *AgentDeviceMappingGeneral) HasClass() bool`

HasClass returns a boolean if a field has been set.

### SetClassNil

`func (o *AgentDeviceMappingGeneral) SetClassNil(b bool)`

 SetClassNil sets the value for Class to be an explicit nil

### UnsetClass
`func (o *AgentDeviceMappingGeneral) UnsetClass()`

UnsetClass ensures that no value is present for Class, not even an explicit nil
### GetId

`func (o *AgentDeviceMappingGeneral) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentDeviceMappingGeneral) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentDeviceMappingGeneral) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgentDeviceMappingGeneral) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AgentDeviceMappingGeneral) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AgentDeviceMappingGeneral) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDeviceId

`func (o *AgentDeviceMappingGeneral) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AgentDeviceMappingGeneral) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AgentDeviceMappingGeneral) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AgentDeviceMappingGeneral) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *AgentDeviceMappingGeneral) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *AgentDeviceMappingGeneral) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetEnable

`func (o *AgentDeviceMappingGeneral) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AgentDeviceMappingGeneral) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AgentDeviceMappingGeneral) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AgentDeviceMappingGeneral) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetAssetId

`func (o *AgentDeviceMappingGeneral) GetAssetId() int32`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AgentDeviceMappingGeneral) GetAssetIdOk() (*int32, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AgentDeviceMappingGeneral) SetAssetId(v int32)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AgentDeviceMappingGeneral) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *AgentDeviceMappingGeneral) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *AgentDeviceMappingGeneral) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
### GetSubtype

`func (o *AgentDeviceMappingGeneral) GetSubtype() DataSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AgentDeviceMappingGeneral) GetSubtypeOk() (*DataSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AgentDeviceMappingGeneral) SetSubtype(v DataSubtype)`

SetSubtype sets Subtype field to given value.


### GetAttribute

`func (o *AgentDeviceMappingGeneral) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AgentDeviceMappingGeneral) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AgentDeviceMappingGeneral) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


