# MetaAggregationMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**WeightField** | Pointer to **string** | Field name to be used for setting weight for primary field for weighted average aggregation | [optional] 

## Methods

### NewMetaAggregationMetric

`func NewMetaAggregationMetric() *MetaAggregationMetric`

NewMetaAggregationMetric instantiates a new MetaAggregationMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaAggregationMetricWithDefaults

`func NewMetaAggregationMetricWithDefaults() *MetaAggregationMetric`

NewMetaAggregationMetricWithDefaults instantiates a new MetaAggregationMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *MetaAggregationMetric) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MetaAggregationMetric) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MetaAggregationMetric) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MetaAggregationMetric) HasField() bool`

HasField returns a boolean if a field has been set.

### GetWeightField

`func (o *MetaAggregationMetric) GetWeightField() string`

GetWeightField returns the WeightField field if non-nil, zero value otherwise.

### GetWeightFieldOk

`func (o *MetaAggregationMetric) GetWeightFieldOk() (*string, bool)`

GetWeightFieldOk returns a tuple with the WeightField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightField

`func (o *MetaAggregationMetric) SetWeightField(v string)`

SetWeightField sets WeightField field to given value.

### HasWeightField

`func (o *MetaAggregationMetric) HasWeightField() bool`

HasWeightField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


