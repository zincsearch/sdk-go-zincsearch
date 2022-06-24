# MetaAggregationsTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **map[string]string** | { \&quot;_count\&quot;: \&quot;asc\&quot; } | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewMetaAggregationsTerms

`func NewMetaAggregationsTerms() *MetaAggregationsTerms`

NewMetaAggregationsTerms instantiates a new MetaAggregationsTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaAggregationsTermsWithDefaults

`func NewMetaAggregationsTermsWithDefaults() *MetaAggregationsTerms`

NewMetaAggregationsTermsWithDefaults instantiates a new MetaAggregationsTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *MetaAggregationsTerms) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MetaAggregationsTerms) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MetaAggregationsTerms) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MetaAggregationsTerms) HasField() bool`

HasField returns a boolean if a field has been set.

### GetOrder

`func (o *MetaAggregationsTerms) GetOrder() map[string]string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MetaAggregationsTerms) GetOrderOk() (*map[string]string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MetaAggregationsTerms) SetOrder(v map[string]string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MetaAggregationsTerms) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSize

`func (o *MetaAggregationsTerms) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MetaAggregationsTerms) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MetaAggregationsTerms) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MetaAggregationsTerms) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


