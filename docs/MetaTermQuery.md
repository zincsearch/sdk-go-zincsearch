# MetaTermQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boost** | Pointer to **float32** |  | [optional] 
**CaseInsensitive** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaTermQuery

`func NewMetaTermQuery() *MetaTermQuery`

NewMetaTermQuery instantiates a new MetaTermQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaTermQueryWithDefaults

`func NewMetaTermQueryWithDefaults() *MetaTermQuery`

NewMetaTermQueryWithDefaults instantiates a new MetaTermQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoost

`func (o *MetaTermQuery) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *MetaTermQuery) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *MetaTermQuery) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *MetaTermQuery) HasBoost() bool`

HasBoost returns a boolean if a field has been set.

### GetCaseInsensitive

`func (o *MetaTermQuery) GetCaseInsensitive() bool`

GetCaseInsensitive returns the CaseInsensitive field if non-nil, zero value otherwise.

### GetCaseInsensitiveOk

`func (o *MetaTermQuery) GetCaseInsensitiveOk() (*bool, bool)`

GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitive

`func (o *MetaTermQuery) SetCaseInsensitive(v bool)`

SetCaseInsensitive sets CaseInsensitive field to given value.

### HasCaseInsensitive

`func (o *MetaTermQuery) HasCaseInsensitive() bool`

HasCaseInsensitive returns a boolean if a field has been set.

### GetValue

`func (o *MetaTermQuery) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetaTermQuery) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetaTermQuery) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetaTermQuery) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


