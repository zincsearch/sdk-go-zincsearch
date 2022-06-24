# MetaPrefixQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boost** | Pointer to **float32** |  | [optional] 
**Value** | Pointer to **string** | You can speed up prefix queries using the index_prefixes mapping parameter. | [optional] 

## Methods

### NewMetaPrefixQuery

`func NewMetaPrefixQuery() *MetaPrefixQuery`

NewMetaPrefixQuery instantiates a new MetaPrefixQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaPrefixQueryWithDefaults

`func NewMetaPrefixQueryWithDefaults() *MetaPrefixQuery`

NewMetaPrefixQueryWithDefaults instantiates a new MetaPrefixQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoost

`func (o *MetaPrefixQuery) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *MetaPrefixQuery) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *MetaPrefixQuery) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *MetaPrefixQuery) HasBoost() bool`

HasBoost returns a boolean if a field has been set.

### GetValue

`func (o *MetaPrefixQuery) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetaPrefixQuery) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetaPrefixQuery) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetaPrefixQuery) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


