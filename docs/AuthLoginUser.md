# AuthLoginUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthLoginUser

`func NewAuthLoginUser() *AuthLoginUser`

NewAuthLoginUser instantiates a new AuthLoginUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLoginUserWithDefaults

`func NewAuthLoginUserWithDefaults() *AuthLoginUser`

NewAuthLoginUserWithDefaults instantiates a new AuthLoginUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthLoginUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthLoginUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthLoginUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthLoginUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthLoginUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthLoginUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthLoginUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthLoginUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *AuthLoginUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AuthLoginUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AuthLoginUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AuthLoginUser) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


