# V1AggregationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]V1AggregationBucket**](V1AggregationBucket.md) |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewV1AggregationResponse

`func NewV1AggregationResponse() *V1AggregationResponse`

NewV1AggregationResponse instantiates a new V1AggregationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AggregationResponseWithDefaults

`func NewV1AggregationResponseWithDefaults() *V1AggregationResponse`

NewV1AggregationResponseWithDefaults instantiates a new V1AggregationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *V1AggregationResponse) GetBuckets() []V1AggregationBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *V1AggregationResponse) GetBucketsOk() (*[]V1AggregationBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *V1AggregationResponse) SetBuckets(v []V1AggregationBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *V1AggregationResponse) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetValue

`func (o *V1AggregationResponse) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V1AggregationResponse) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V1AggregationResponse) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *V1AggregationResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


