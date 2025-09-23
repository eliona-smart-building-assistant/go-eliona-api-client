# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key | [optional] 
**ValidUntil** | Pointer to **NullableTime** | Date until the key is valid | [optional] 
**WriteAccess** | Pointer to **NullableBool** | Does the key has write access | [optional] 

## Methods

### NewApiKey

`func NewApiKey() *ApiKey`

NewApiKey instantiates a new ApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyWithDefaults

`func NewApiKeyWithDefaults() *ApiKey`

NewApiKeyWithDefaults instantiates a new ApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApiKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValidUntil

`func (o *ApiKey) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *ApiKey) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *ApiKey) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *ApiKey) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *ApiKey) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *ApiKey) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
### GetWriteAccess

`func (o *ApiKey) GetWriteAccess() bool`

GetWriteAccess returns the WriteAccess field if non-nil, zero value otherwise.

### GetWriteAccessOk

`func (o *ApiKey) GetWriteAccessOk() (*bool, bool)`

GetWriteAccessOk returns a tuple with the WriteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAccess

`func (o *ApiKey) SetWriteAccess(v bool)`

SetWriteAccess sets WriteAccess field to given value.

### HasWriteAccess

`func (o *ApiKey) HasWriteAccess() bool`

HasWriteAccess returns a boolean if a field has been set.

### SetWriteAccessNil

`func (o *ApiKey) SetWriteAccessNil(b bool)`

 SetWriteAccessNil sets the value for WriteAccess to be an explicit nil

### UnsetWriteAccess
`func (o *ApiKey) UnsetWriteAccess()`

UnsetWriteAccess ensures that no value is present for WriteAccess, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


