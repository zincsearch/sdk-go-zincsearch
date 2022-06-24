# MetaAnalyzer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CharFilter** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to **[]string** | compatibility with es, alias for TokenFilter | [optional] 
**Lowercase** | Pointer to **bool** | for type&#x3D;pattern | [optional] 
**Pattern** | Pointer to **string** | for type&#x3D;pattern | [optional] 
**Stopwords** | Pointer to **[]string** | for type&#x3D;pattern,standard,stop | [optional] 
**TokenFilter** | Pointer to **[]string** |  | [optional] 
**Tokenizer** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | options for compatible | [optional] 

## Methods

### NewMetaAnalyzer

`func NewMetaAnalyzer() *MetaAnalyzer`

NewMetaAnalyzer instantiates a new MetaAnalyzer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaAnalyzerWithDefaults

`func NewMetaAnalyzerWithDefaults() *MetaAnalyzer`

NewMetaAnalyzerWithDefaults instantiates a new MetaAnalyzer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharFilter

`func (o *MetaAnalyzer) GetCharFilter() []string`

GetCharFilter returns the CharFilter field if non-nil, zero value otherwise.

### GetCharFilterOk

`func (o *MetaAnalyzer) GetCharFilterOk() (*[]string, bool)`

GetCharFilterOk returns a tuple with the CharFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharFilter

`func (o *MetaAnalyzer) SetCharFilter(v []string)`

SetCharFilter sets CharFilter field to given value.

### HasCharFilter

`func (o *MetaAnalyzer) HasCharFilter() bool`

HasCharFilter returns a boolean if a field has been set.

### GetFilter

`func (o *MetaAnalyzer) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetaAnalyzer) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetaAnalyzer) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MetaAnalyzer) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLowercase

`func (o *MetaAnalyzer) GetLowercase() bool`

GetLowercase returns the Lowercase field if non-nil, zero value otherwise.

### GetLowercaseOk

`func (o *MetaAnalyzer) GetLowercaseOk() (*bool, bool)`

GetLowercaseOk returns a tuple with the Lowercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowercase

`func (o *MetaAnalyzer) SetLowercase(v bool)`

SetLowercase sets Lowercase field to given value.

### HasLowercase

`func (o *MetaAnalyzer) HasLowercase() bool`

HasLowercase returns a boolean if a field has been set.

### GetPattern

`func (o *MetaAnalyzer) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *MetaAnalyzer) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *MetaAnalyzer) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *MetaAnalyzer) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetStopwords

`func (o *MetaAnalyzer) GetStopwords() []string`

GetStopwords returns the Stopwords field if non-nil, zero value otherwise.

### GetStopwordsOk

`func (o *MetaAnalyzer) GetStopwordsOk() (*[]string, bool)`

GetStopwordsOk returns a tuple with the Stopwords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopwords

`func (o *MetaAnalyzer) SetStopwords(v []string)`

SetStopwords sets Stopwords field to given value.

### HasStopwords

`func (o *MetaAnalyzer) HasStopwords() bool`

HasStopwords returns a boolean if a field has been set.

### GetTokenFilter

`func (o *MetaAnalyzer) GetTokenFilter() []string`

GetTokenFilter returns the TokenFilter field if non-nil, zero value otherwise.

### GetTokenFilterOk

`func (o *MetaAnalyzer) GetTokenFilterOk() (*[]string, bool)`

GetTokenFilterOk returns a tuple with the TokenFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFilter

`func (o *MetaAnalyzer) SetTokenFilter(v []string)`

SetTokenFilter sets TokenFilter field to given value.

### HasTokenFilter

`func (o *MetaAnalyzer) HasTokenFilter() bool`

HasTokenFilter returns a boolean if a field has been set.

### GetTokenizer

`func (o *MetaAnalyzer) GetTokenizer() string`

GetTokenizer returns the Tokenizer field if non-nil, zero value otherwise.

### GetTokenizerOk

`func (o *MetaAnalyzer) GetTokenizerOk() (*string, bool)`

GetTokenizerOk returns a tuple with the Tokenizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizer

`func (o *MetaAnalyzer) SetTokenizer(v string)`

SetTokenizer sets Tokenizer field to given value.

### HasTokenizer

`func (o *MetaAnalyzer) HasTokenizer() bool`

HasTokenizer returns a boolean if a field has been set.

### GetType

`func (o *MetaAnalyzer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetaAnalyzer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetaAnalyzer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetaAnalyzer) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


