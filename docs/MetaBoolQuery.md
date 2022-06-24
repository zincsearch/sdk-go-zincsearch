# MetaBoolQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**[]MetaQuery**](MetaQuery.md) | query, [query1, query2] | [optional] 
**MinimumShouldMatch** | Pointer to **float32** | only for should | [optional] 
**Must** | Pointer to [**[]MetaQuery**](MetaQuery.md) | query, [query1, query2] | [optional] 
**MustNot** | Pointer to [**[]MetaQuery**](MetaQuery.md) | query, [query1, query2] | [optional] 
**Should** | Pointer to [**[]MetaQuery**](MetaQuery.md) | query, [query1, query2] | [optional] 

## Methods

### NewMetaBoolQuery

`func NewMetaBoolQuery() *MetaBoolQuery`

NewMetaBoolQuery instantiates a new MetaBoolQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaBoolQueryWithDefaults

`func NewMetaBoolQueryWithDefaults() *MetaBoolQuery`

NewMetaBoolQueryWithDefaults instantiates a new MetaBoolQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *MetaBoolQuery) GetFilter() []MetaQuery`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetaBoolQuery) GetFilterOk() (*[]MetaQuery, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetaBoolQuery) SetFilter(v []MetaQuery)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MetaBoolQuery) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMinimumShouldMatch

`func (o *MetaBoolQuery) GetMinimumShouldMatch() float32`

GetMinimumShouldMatch returns the MinimumShouldMatch field if non-nil, zero value otherwise.

### GetMinimumShouldMatchOk

`func (o *MetaBoolQuery) GetMinimumShouldMatchOk() (*float32, bool)`

GetMinimumShouldMatchOk returns a tuple with the MinimumShouldMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumShouldMatch

`func (o *MetaBoolQuery) SetMinimumShouldMatch(v float32)`

SetMinimumShouldMatch sets MinimumShouldMatch field to given value.

### HasMinimumShouldMatch

`func (o *MetaBoolQuery) HasMinimumShouldMatch() bool`

HasMinimumShouldMatch returns a boolean if a field has been set.

### GetMust

`func (o *MetaBoolQuery) GetMust() []MetaQuery`

GetMust returns the Must field if non-nil, zero value otherwise.

### GetMustOk

`func (o *MetaBoolQuery) GetMustOk() (*[]MetaQuery, bool)`

GetMustOk returns a tuple with the Must field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMust

`func (o *MetaBoolQuery) SetMust(v []MetaQuery)`

SetMust sets Must field to given value.

### HasMust

`func (o *MetaBoolQuery) HasMust() bool`

HasMust returns a boolean if a field has been set.

### GetMustNot

`func (o *MetaBoolQuery) GetMustNot() []MetaQuery`

GetMustNot returns the MustNot field if non-nil, zero value otherwise.

### GetMustNotOk

`func (o *MetaBoolQuery) GetMustNotOk() (*[]MetaQuery, bool)`

GetMustNotOk returns a tuple with the MustNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustNot

`func (o *MetaBoolQuery) SetMustNot(v []MetaQuery)`

SetMustNot sets MustNot field to given value.

### HasMustNot

`func (o *MetaBoolQuery) HasMustNot() bool`

HasMustNot returns a boolean if a field has been set.

### GetShould

`func (o *MetaBoolQuery) GetShould() []MetaQuery`

GetShould returns the Should field if non-nil, zero value otherwise.

### GetShouldOk

`func (o *MetaBoolQuery) GetShouldOk() (*[]MetaQuery, bool)`

GetShouldOk returns a tuple with the Should field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShould

`func (o *MetaBoolQuery) SetShould(v []MetaQuery)`

SetShould sets Should field to given value.

### HasShould

`func (o *MetaBoolQuery) HasShould() bool`

HasShould returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


