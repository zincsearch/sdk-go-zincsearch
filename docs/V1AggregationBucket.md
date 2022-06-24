# V1AggregationBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**map[string]V1AggregationResponse**](V1AggregationResponse.md) |  | [optional] 
**DocCount** | Pointer to **int32** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewV1AggregationBucket

`func NewV1AggregationBucket() *V1AggregationBucket`

NewV1AggregationBucket instantiates a new V1AggregationBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AggregationBucketWithDefaults

`func NewV1AggregationBucketWithDefaults() *V1AggregationBucket`

NewV1AggregationBucketWithDefaults instantiates a new V1AggregationBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *V1AggregationBucket) GetAggregations() map[string]V1AggregationResponse`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *V1AggregationBucket) GetAggregationsOk() (*map[string]V1AggregationResponse, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *V1AggregationBucket) SetAggregations(v map[string]V1AggregationResponse)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *V1AggregationBucket) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetDocCount

`func (o *V1AggregationBucket) GetDocCount() int32`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *V1AggregationBucket) GetDocCountOk() (*int32, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *V1AggregationBucket) SetDocCount(v int32)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *V1AggregationBucket) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetKey

`func (o *V1AggregationBucket) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V1AggregationBucket) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V1AggregationBucket) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *V1AggregationBucket) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


