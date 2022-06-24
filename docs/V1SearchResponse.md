# V1SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**map[string]V1AggregationResponse**](V1AggregationResponse.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Hits** | Pointer to [**V1Hits**](V1Hits.md) |  | [optional] 
**TimedOut** | Pointer to **bool** |  | [optional] 
**Took** | Pointer to **int32** | Time it took to generate the response | [optional] 

## Methods

### NewV1SearchResponse

`func NewV1SearchResponse() *V1SearchResponse`

NewV1SearchResponse instantiates a new V1SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SearchResponseWithDefaults

`func NewV1SearchResponseWithDefaults() *V1SearchResponse`

NewV1SearchResponseWithDefaults instantiates a new V1SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *V1SearchResponse) GetAggregations() map[string]V1AggregationResponse`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *V1SearchResponse) GetAggregationsOk() (*map[string]V1AggregationResponse, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *V1SearchResponse) SetAggregations(v map[string]V1AggregationResponse)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *V1SearchResponse) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetError

`func (o *V1SearchResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V1SearchResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V1SearchResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V1SearchResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHits

`func (o *V1SearchResponse) GetHits() V1Hits`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *V1SearchResponse) GetHitsOk() (*V1Hits, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *V1SearchResponse) SetHits(v V1Hits)`

SetHits sets Hits field to given value.

### HasHits

`func (o *V1SearchResponse) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetTimedOut

`func (o *V1SearchResponse) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *V1SearchResponse) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *V1SearchResponse) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *V1SearchResponse) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### GetTook

`func (o *V1SearchResponse) GetTook() int32`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *V1SearchResponse) GetTookOk() (*int32, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *V1SearchResponse) SetTook(v int32)`

SetTook sets Took field to given value.

### HasTook

`func (o *V1SearchResponse) HasTook() bool`

HasTook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


