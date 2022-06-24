# MetaFuzzyQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boost** | Pointer to **float32** |  | [optional] 
**Fuzziness** | Pointer to **map[string]interface{}** | auto, 1,2,3,n | [optional] 
**PrefixLength** | Pointer to **float32** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaFuzzyQuery

`func NewMetaFuzzyQuery() *MetaFuzzyQuery`

NewMetaFuzzyQuery instantiates a new MetaFuzzyQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaFuzzyQueryWithDefaults

`func NewMetaFuzzyQueryWithDefaults() *MetaFuzzyQuery`

NewMetaFuzzyQueryWithDefaults instantiates a new MetaFuzzyQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoost

`func (o *MetaFuzzyQuery) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *MetaFuzzyQuery) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *MetaFuzzyQuery) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *MetaFuzzyQuery) HasBoost() bool`

HasBoost returns a boolean if a field has been set.

### GetFuzziness

`func (o *MetaFuzzyQuery) GetFuzziness() map[string]interface{}`

GetFuzziness returns the Fuzziness field if non-nil, zero value otherwise.

### GetFuzzinessOk

`func (o *MetaFuzzyQuery) GetFuzzinessOk() (*map[string]interface{}, bool)`

GetFuzzinessOk returns a tuple with the Fuzziness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuzziness

`func (o *MetaFuzzyQuery) SetFuzziness(v map[string]interface{})`

SetFuzziness sets Fuzziness field to given value.

### HasFuzziness

`func (o *MetaFuzzyQuery) HasFuzziness() bool`

HasFuzziness returns a boolean if a field has been set.

### GetPrefixLength

`func (o *MetaFuzzyQuery) GetPrefixLength() float32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *MetaFuzzyQuery) GetPrefixLengthOk() (*float32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *MetaFuzzyQuery) SetPrefixLength(v float32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *MetaFuzzyQuery) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetValue

`func (o *MetaFuzzyQuery) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetaFuzzyQuery) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetaFuzzyQuery) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetaFuzzyQuery) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


