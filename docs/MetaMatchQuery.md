# MetaMatchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | Pointer to **string** |  | [optional] 
**Boost** | Pointer to **float32** |  | [optional] 
**Fuzziness** | Pointer to **map[string]interface{}** | auto, 1,2,3,n | [optional] 
**Operator** | Pointer to **string** | or(default), and | [optional] 
**PrefixLength** | Pointer to **float32** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaMatchQuery

`func NewMetaMatchQuery() *MetaMatchQuery`

NewMetaMatchQuery instantiates a new MetaMatchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaMatchQueryWithDefaults

`func NewMetaMatchQueryWithDefaults() *MetaMatchQuery`

NewMetaMatchQueryWithDefaults instantiates a new MetaMatchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *MetaMatchQuery) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *MetaMatchQuery) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *MetaMatchQuery) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *MetaMatchQuery) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.

### GetBoost

`func (o *MetaMatchQuery) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *MetaMatchQuery) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *MetaMatchQuery) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *MetaMatchQuery) HasBoost() bool`

HasBoost returns a boolean if a field has been set.

### GetFuzziness

`func (o *MetaMatchQuery) GetFuzziness() map[string]interface{}`

GetFuzziness returns the Fuzziness field if non-nil, zero value otherwise.

### GetFuzzinessOk

`func (o *MetaMatchQuery) GetFuzzinessOk() (*map[string]interface{}, bool)`

GetFuzzinessOk returns a tuple with the Fuzziness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuzziness

`func (o *MetaMatchQuery) SetFuzziness(v map[string]interface{})`

SetFuzziness sets Fuzziness field to given value.

### HasFuzziness

`func (o *MetaMatchQuery) HasFuzziness() bool`

HasFuzziness returns a boolean if a field has been set.

### GetOperator

`func (o *MetaMatchQuery) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MetaMatchQuery) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MetaMatchQuery) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *MetaMatchQuery) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetPrefixLength

`func (o *MetaMatchQuery) GetPrefixLength() float32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *MetaMatchQuery) GetPrefixLengthOk() (*float32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *MetaMatchQuery) SetPrefixLength(v float32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *MetaMatchQuery) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetQuery

`func (o *MetaMatchQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetaMatchQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetaMatchQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MetaMatchQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


