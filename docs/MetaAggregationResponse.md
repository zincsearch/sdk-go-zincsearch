# MetaAggregationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to **map[string]interface{}** | slice or map | [optional] 
**Interval** | Pointer to **string** | support for auto_date_histogram_aggregation | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMetaAggregationResponse

`func NewMetaAggregationResponse() *MetaAggregationResponse`

NewMetaAggregationResponse instantiates a new MetaAggregationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaAggregationResponseWithDefaults

`func NewMetaAggregationResponseWithDefaults() *MetaAggregationResponse`

NewMetaAggregationResponseWithDefaults instantiates a new MetaAggregationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *MetaAggregationResponse) GetBuckets() map[string]interface{}`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *MetaAggregationResponse) GetBucketsOk() (*map[string]interface{}, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *MetaAggregationResponse) SetBuckets(v map[string]interface{})`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *MetaAggregationResponse) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetInterval

`func (o *MetaAggregationResponse) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MetaAggregationResponse) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MetaAggregationResponse) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MetaAggregationResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetValue

`func (o *MetaAggregationResponse) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetaAggregationResponse) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetaAggregationResponse) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *MetaAggregationResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


