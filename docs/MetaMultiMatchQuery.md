# MetaMultiMatchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | Pointer to **string** |  | [optional] 
**Boost** | Pointer to **float32** |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**MinimumShouldMatch** | Pointer to **float32** |  | [optional] 
**Operator** | Pointer to **string** | or(default), and | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | best_fields(default), most_fields, cross_fields, phrase, phrase_prefix, bool_prefix | [optional] 

## Methods

### NewMetaMultiMatchQuery

`func NewMetaMultiMatchQuery() *MetaMultiMatchQuery`

NewMetaMultiMatchQuery instantiates a new MetaMultiMatchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaMultiMatchQueryWithDefaults

`func NewMetaMultiMatchQueryWithDefaults() *MetaMultiMatchQuery`

NewMetaMultiMatchQueryWithDefaults instantiates a new MetaMultiMatchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *MetaMultiMatchQuery) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *MetaMultiMatchQuery) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *MetaMultiMatchQuery) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *MetaMultiMatchQuery) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.

### GetBoost

`func (o *MetaMultiMatchQuery) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *MetaMultiMatchQuery) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *MetaMultiMatchQuery) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *MetaMultiMatchQuery) HasBoost() bool`

HasBoost returns a boolean if a field has been set.

### GetFields

`func (o *MetaMultiMatchQuery) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MetaMultiMatchQuery) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MetaMultiMatchQuery) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MetaMultiMatchQuery) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMinimumShouldMatch

`func (o *MetaMultiMatchQuery) GetMinimumShouldMatch() float32`

GetMinimumShouldMatch returns the MinimumShouldMatch field if non-nil, zero value otherwise.

### GetMinimumShouldMatchOk

`func (o *MetaMultiMatchQuery) GetMinimumShouldMatchOk() (*float32, bool)`

GetMinimumShouldMatchOk returns a tuple with the MinimumShouldMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumShouldMatch

`func (o *MetaMultiMatchQuery) SetMinimumShouldMatch(v float32)`

SetMinimumShouldMatch sets MinimumShouldMatch field to given value.

### HasMinimumShouldMatch

`func (o *MetaMultiMatchQuery) HasMinimumShouldMatch() bool`

HasMinimumShouldMatch returns a boolean if a field has been set.

### GetOperator

`func (o *MetaMultiMatchQuery) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MetaMultiMatchQuery) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MetaMultiMatchQuery) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *MetaMultiMatchQuery) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetQuery

`func (o *MetaMultiMatchQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetaMultiMatchQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetaMultiMatchQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MetaMultiMatchQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *MetaMultiMatchQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetaMultiMatchQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetaMultiMatchQuery) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetaMultiMatchQuery) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


