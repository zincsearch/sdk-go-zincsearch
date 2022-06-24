# MetaSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shards** | Pointer to [**MetaShards**](MetaShards.md) |  | [optional] 
**Aggregations** | Pointer to [**map[string]MetaAggregationResponse**](MetaAggregationResponse.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Hits** | Pointer to [**MetaHits**](MetaHits.md) |  | [optional] 
**TimedOut** | Pointer to **bool** |  | [optional] 
**Took** | Pointer to **int32** | Time it took to generate the response | [optional] 

## Methods

### NewMetaSearchResponse

`func NewMetaSearchResponse() *MetaSearchResponse`

NewMetaSearchResponse instantiates a new MetaSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaSearchResponseWithDefaults

`func NewMetaSearchResponseWithDefaults() *MetaSearchResponse`

NewMetaSearchResponseWithDefaults instantiates a new MetaSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShards

`func (o *MetaSearchResponse) GetShards() MetaShards`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *MetaSearchResponse) GetShardsOk() (*MetaShards, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *MetaSearchResponse) SetShards(v MetaShards)`

SetShards sets Shards field to given value.

### HasShards

`func (o *MetaSearchResponse) HasShards() bool`

HasShards returns a boolean if a field has been set.

### GetAggregations

`func (o *MetaSearchResponse) GetAggregations() map[string]MetaAggregationResponse`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *MetaSearchResponse) GetAggregationsOk() (*map[string]MetaAggregationResponse, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *MetaSearchResponse) SetAggregations(v map[string]MetaAggregationResponse)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *MetaSearchResponse) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetError

`func (o *MetaSearchResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MetaSearchResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MetaSearchResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MetaSearchResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHits

`func (o *MetaSearchResponse) GetHits() MetaHits`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *MetaSearchResponse) GetHitsOk() (*MetaHits, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *MetaSearchResponse) SetHits(v MetaHits)`

SetHits sets Hits field to given value.

### HasHits

`func (o *MetaSearchResponse) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetTimedOut

`func (o *MetaSearchResponse) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *MetaSearchResponse) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *MetaSearchResponse) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *MetaSearchResponse) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### GetTook

`func (o *MetaSearchResponse) GetTook() int32`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *MetaSearchResponse) GetTookOk() (*int32, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *MetaSearchResponse) SetTook(v int32)`

SetTook sets Took field to given value.

### HasTook

`func (o *MetaSearchResponse) HasTook() bool`

HasTook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


