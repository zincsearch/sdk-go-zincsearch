# MetaMatchPhrasePrefixQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | Pointer to **string** |  | [optional] 
**Boost** | Pointer to **float32** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaMatchPhrasePrefixQuery

`func NewMetaMatchPhrasePrefixQuery() *MetaMatchPhrasePrefixQuery`

NewMetaMatchPhrasePrefixQuery instantiates a new MetaMatchPhrasePrefixQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaMatchPhrasePrefixQueryWithDefaults

`func NewMetaMatchPhrasePrefixQueryWithDefaults() *MetaMatchPhrasePrefixQuery`

NewMetaMatchPhrasePrefixQueryWithDefaults instantiates a new MetaMatchPhrasePrefixQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *MetaMatchPhrasePrefixQuery) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *MetaMatchPhrasePrefixQuery) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *MetaMatchPhrasePrefixQuery) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *MetaMatchPhrasePrefixQuery) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.

### GetBoost

`func (o *MetaMatchPhrasePrefixQuery) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *MetaMatchPhrasePrefixQuery) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *MetaMatchPhrasePrefixQuery) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *MetaMatchPhrasePrefixQuery) HasBoost() bool`

HasBoost returns a boolean if a field has been set.

### GetQuery

`func (o *MetaMatchPhrasePrefixQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetaMatchPhrasePrefixQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetaMatchPhrasePrefixQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MetaMatchPhrasePrefixQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


