# MetaQueryStringQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | Pointer to **string** |  | [optional] 
**Boost** | Pointer to **float32** |  | [optional] 
**DefaultField** | Pointer to **string** |  | [optional] 
**DefaultOperator** | Pointer to **string** | or(default), and | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaQueryStringQuery

`func NewMetaQueryStringQuery() *MetaQueryStringQuery`

NewMetaQueryStringQuery instantiates a new MetaQueryStringQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaQueryStringQueryWithDefaults

`func NewMetaQueryStringQueryWithDefaults() *MetaQueryStringQuery`

NewMetaQueryStringQueryWithDefaults instantiates a new MetaQueryStringQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *MetaQueryStringQuery) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *MetaQueryStringQuery) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *MetaQueryStringQuery) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *MetaQueryStringQuery) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.

### GetBoost

`func (o *MetaQueryStringQuery) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *MetaQueryStringQuery) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *MetaQueryStringQuery) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *MetaQueryStringQuery) HasBoost() bool`

HasBoost returns a boolean if a field has been set.

### GetDefaultField

`func (o *MetaQueryStringQuery) GetDefaultField() string`

GetDefaultField returns the DefaultField field if non-nil, zero value otherwise.

### GetDefaultFieldOk

`func (o *MetaQueryStringQuery) GetDefaultFieldOk() (*string, bool)`

GetDefaultFieldOk returns a tuple with the DefaultField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultField

`func (o *MetaQueryStringQuery) SetDefaultField(v string)`

SetDefaultField sets DefaultField field to given value.

### HasDefaultField

`func (o *MetaQueryStringQuery) HasDefaultField() bool`

HasDefaultField returns a boolean if a field has been set.

### GetDefaultOperator

`func (o *MetaQueryStringQuery) GetDefaultOperator() string`

GetDefaultOperator returns the DefaultOperator field if non-nil, zero value otherwise.

### GetDefaultOperatorOk

`func (o *MetaQueryStringQuery) GetDefaultOperatorOk() (*string, bool)`

GetDefaultOperatorOk returns a tuple with the DefaultOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOperator

`func (o *MetaQueryStringQuery) SetDefaultOperator(v string)`

SetDefaultOperator sets DefaultOperator field to given value.

### HasDefaultOperator

`func (o *MetaQueryStringQuery) HasDefaultOperator() bool`

HasDefaultOperator returns a boolean if a field has been set.

### GetFields

`func (o *MetaQueryStringQuery) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MetaQueryStringQuery) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MetaQueryStringQuery) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MetaQueryStringQuery) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetQuery

`func (o *MetaQueryStringQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetaQueryStringQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetaQueryStringQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MetaQueryStringQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


