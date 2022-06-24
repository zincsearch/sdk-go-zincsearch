# MetaIndexAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analyzer** | Pointer to [**map[string]MetaAnalyzer**](MetaAnalyzer.md) |  | [optional] 
**CharFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**Filter** | Pointer to **map[string]interface{}** | compatibility with es, alias for TokenFilter | [optional] 
**TokenFilter** | Pointer to **map[string]interface{}** |  | [optional] 
**Tokenizer** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMetaIndexAnalysis

`func NewMetaIndexAnalysis() *MetaIndexAnalysis`

NewMetaIndexAnalysis instantiates a new MetaIndexAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaIndexAnalysisWithDefaults

`func NewMetaIndexAnalysisWithDefaults() *MetaIndexAnalysis`

NewMetaIndexAnalysisWithDefaults instantiates a new MetaIndexAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzer

`func (o *MetaIndexAnalysis) GetAnalyzer() map[string]MetaAnalyzer`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *MetaIndexAnalysis) GetAnalyzerOk() (*map[string]MetaAnalyzer, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *MetaIndexAnalysis) SetAnalyzer(v map[string]MetaAnalyzer)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *MetaIndexAnalysis) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.

### GetCharFilter

`func (o *MetaIndexAnalysis) GetCharFilter() map[string]interface{}`

GetCharFilter returns the CharFilter field if non-nil, zero value otherwise.

### GetCharFilterOk

`func (o *MetaIndexAnalysis) GetCharFilterOk() (*map[string]interface{}, bool)`

GetCharFilterOk returns a tuple with the CharFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharFilter

`func (o *MetaIndexAnalysis) SetCharFilter(v map[string]interface{})`

SetCharFilter sets CharFilter field to given value.

### HasCharFilter

`func (o *MetaIndexAnalysis) HasCharFilter() bool`

HasCharFilter returns a boolean if a field has been set.

### GetFilter

`func (o *MetaIndexAnalysis) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetaIndexAnalysis) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetaIndexAnalysis) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MetaIndexAnalysis) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetTokenFilter

`func (o *MetaIndexAnalysis) GetTokenFilter() map[string]interface{}`

GetTokenFilter returns the TokenFilter field if non-nil, zero value otherwise.

### GetTokenFilterOk

`func (o *MetaIndexAnalysis) GetTokenFilterOk() (*map[string]interface{}, bool)`

GetTokenFilterOk returns a tuple with the TokenFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFilter

`func (o *MetaIndexAnalysis) SetTokenFilter(v map[string]interface{})`

SetTokenFilter sets TokenFilter field to given value.

### HasTokenFilter

`func (o *MetaIndexAnalysis) HasTokenFilter() bool`

HasTokenFilter returns a boolean if a field has been set.

### GetTokenizer

`func (o *MetaIndexAnalysis) GetTokenizer() map[string]interface{}`

GetTokenizer returns the Tokenizer field if non-nil, zero value otherwise.

### GetTokenizerOk

`func (o *MetaIndexAnalysis) GetTokenizerOk() (*map[string]interface{}, bool)`

GetTokenizerOk returns a tuple with the Tokenizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizer

`func (o *MetaIndexAnalysis) SetTokenizer(v map[string]interface{})`

SetTokenizer sets Tokenizer field to given value.

### HasTokenizer

`func (o *MetaIndexAnalysis) HasTokenizer() bool`

HasTokenizer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


