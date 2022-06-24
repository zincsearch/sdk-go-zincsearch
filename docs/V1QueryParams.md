# V1QueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boost** | Pointer to **int32** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Term** | Pointer to **string** |  | [optional] 
**Terms** | Pointer to **[][]string** | For multi phrase query | [optional] 

## Methods

### NewV1QueryParams

`func NewV1QueryParams() *V1QueryParams`

NewV1QueryParams instantiates a new V1QueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1QueryParamsWithDefaults

`func NewV1QueryParamsWithDefaults() *V1QueryParams`

NewV1QueryParamsWithDefaults instantiates a new V1QueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoost

`func (o *V1QueryParams) GetBoost() int32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *V1QueryParams) GetBoostOk() (*int32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *V1QueryParams) SetBoost(v int32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *V1QueryParams) HasBoost() bool`

HasBoost returns a boolean if a field has been set.

### GetEndTime

`func (o *V1QueryParams) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *V1QueryParams) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *V1QueryParams) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *V1QueryParams) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetField

`func (o *V1QueryParams) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *V1QueryParams) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *V1QueryParams) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *V1QueryParams) HasField() bool`

HasField returns a boolean if a field has been set.

### GetStartTime

`func (o *V1QueryParams) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V1QueryParams) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V1QueryParams) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V1QueryParams) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTerm

`func (o *V1QueryParams) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *V1QueryParams) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *V1QueryParams) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *V1QueryParams) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### GetTerms

`func (o *V1QueryParams) GetTerms() [][]string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *V1QueryParams) GetTermsOk() (*[][]string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *V1QueryParams) SetTerms(v [][]string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *V1QueryParams) HasTerms() bool`

HasTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


