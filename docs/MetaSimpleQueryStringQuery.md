# MetaSimpleQueryStringQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllFields** | Pointer to **bool** |  | [optional] 
**Analyzer** | Pointer to **string** |  | [optional] 
**Boost** | Pointer to **float32** |  | [optional] 
**DefaultOperator** | Pointer to **string** | or(default), and | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaSimpleQueryStringQuery

`func NewMetaSimpleQueryStringQuery() *MetaSimpleQueryStringQuery`

NewMetaSimpleQueryStringQuery instantiates a new MetaSimpleQueryStringQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaSimpleQueryStringQueryWithDefaults

`func NewMetaSimpleQueryStringQueryWithDefaults() *MetaSimpleQueryStringQuery`

NewMetaSimpleQueryStringQueryWithDefaults instantiates a new MetaSimpleQueryStringQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllFields

`func (o *MetaSimpleQueryStringQuery) GetAllFields() bool`

GetAllFields returns the AllFields field if non-nil, zero value otherwise.

### GetAllFieldsOk

`func (o *MetaSimpleQueryStringQuery) GetAllFieldsOk() (*bool, bool)`

GetAllFieldsOk returns a tuple with the AllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllFields

`func (o *MetaSimpleQueryStringQuery) SetAllFields(v bool)`

SetAllFields sets AllFields field to given value.

### HasAllFields

`func (o *MetaSimpleQueryStringQuery) HasAllFields() bool`

HasAllFields returns a boolean if a field has been set.

### GetAnalyzer

`func (o *MetaSimpleQueryStringQuery) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *MetaSimpleQueryStringQuery) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *MetaSimpleQueryStringQuery) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *MetaSimpleQueryStringQuery) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.

### GetBoost

`func (o *MetaSimpleQueryStringQuery) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *MetaSimpleQueryStringQuery) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *MetaSimpleQueryStringQuery) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *MetaSimpleQueryStringQuery) HasBoost() bool`

HasBoost returns a boolean if a field has been set.

### GetDefaultOperator

`func (o *MetaSimpleQueryStringQuery) GetDefaultOperator() string`

GetDefaultOperator returns the DefaultOperator field if non-nil, zero value otherwise.

### GetDefaultOperatorOk

`func (o *MetaSimpleQueryStringQuery) GetDefaultOperatorOk() (*string, bool)`

GetDefaultOperatorOk returns a tuple with the DefaultOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOperator

`func (o *MetaSimpleQueryStringQuery) SetDefaultOperator(v string)`

SetDefaultOperator sets DefaultOperator field to given value.

### HasDefaultOperator

`func (o *MetaSimpleQueryStringQuery) HasDefaultOperator() bool`

HasDefaultOperator returns a boolean if a field has been set.

### GetFields

`func (o *MetaSimpleQueryStringQuery) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MetaSimpleQueryStringQuery) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MetaSimpleQueryStringQuery) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MetaSimpleQueryStringQuery) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetQuery

`func (o *MetaSimpleQueryStringQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetaSimpleQueryStringQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetaSimpleQueryStringQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MetaSimpleQueryStringQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


