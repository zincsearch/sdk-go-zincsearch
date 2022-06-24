# AuthLoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**AuthLoginUser**](AuthLoginUser.md) |  | [optional] 
**Validated** | Pointer to **bool** |  | [optional] 

## Methods

### NewAuthLoginResponse

`func NewAuthLoginResponse() *AuthLoginResponse`

NewAuthLoginResponse instantiates a new AuthLoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLoginResponseWithDefaults

`func NewAuthLoginResponseWithDefaults() *AuthLoginResponse`

NewAuthLoginResponseWithDefaults instantiates a new AuthLoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *AuthLoginResponse) GetUser() AuthLoginUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthLoginResponse) GetUserOk() (*AuthLoginUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthLoginResponse) SetUser(v AuthLoginUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthLoginResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetValidated

`func (o *AuthLoginResponse) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *AuthLoginResponse) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *AuthLoginResponse) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *AuthLoginResponse) HasValidated() bool`

HasValidated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


