# MetaAggregations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggs** | Pointer to [**map[string]MetaAggregations**](MetaAggregations.md) | nested aggregations | [optional] 
**AutoDateHistogram** | Pointer to [**MetaAggregationAutoDateHistogram**](MetaAggregationAutoDateHistogram.md) |  | [optional] 
**Avg** | Pointer to [**MetaAggregationMetric**](MetaAggregationMetric.md) |  | [optional] 
**Cardinality** | Pointer to [**MetaAggregationMetric**](MetaAggregationMetric.md) |  | [optional] 
**Count** | Pointer to [**MetaAggregationMetric**](MetaAggregationMetric.md) |  | [optional] 
**DateHistogram** | Pointer to [**MetaAggregationDateHistogram**](MetaAggregationDateHistogram.md) |  | [optional] 
**DateRange** | Pointer to [**MetaAggregationDateRange**](MetaAggregationDateRange.md) |  | [optional] 
**Histogram** | Pointer to [**MetaAggregationHistogram**](MetaAggregationHistogram.md) |  | [optional] 
**IpRange** | Pointer to [**MetaAggregationIPRange**](MetaAggregationIPRange.md) |  | [optional] 
**Max** | Pointer to [**MetaAggregationMetric**](MetaAggregationMetric.md) |  | [optional] 
**Min** | Pointer to [**MetaAggregationMetric**](MetaAggregationMetric.md) |  | [optional] 
**Range** | Pointer to [**MetaAggregationRange**](MetaAggregationRange.md) |  | [optional] 
**Sum** | Pointer to [**MetaAggregationMetric**](MetaAggregationMetric.md) |  | [optional] 
**Terms** | Pointer to [**MetaAggregationsTerms**](MetaAggregationsTerms.md) |  | [optional] 
**WeightedAvg** | Pointer to [**MetaAggregationMetric**](MetaAggregationMetric.md) |  | [optional] 

## Methods

### NewMetaAggregations

`func NewMetaAggregations() *MetaAggregations`

NewMetaAggregations instantiates a new MetaAggregations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaAggregationsWithDefaults

`func NewMetaAggregationsWithDefaults() *MetaAggregations`

NewMetaAggregationsWithDefaults instantiates a new MetaAggregations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggs

`func (o *MetaAggregations) GetAggs() map[string]MetaAggregations`

GetAggs returns the Aggs field if non-nil, zero value otherwise.

### GetAggsOk

`func (o *MetaAggregations) GetAggsOk() (*map[string]MetaAggregations, bool)`

GetAggsOk returns a tuple with the Aggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggs

`func (o *MetaAggregations) SetAggs(v map[string]MetaAggregations)`

SetAggs sets Aggs field to given value.

### HasAggs

`func (o *MetaAggregations) HasAggs() bool`

HasAggs returns a boolean if a field has been set.

### GetAutoDateHistogram

`func (o *MetaAggregations) GetAutoDateHistogram() MetaAggregationAutoDateHistogram`

GetAutoDateHistogram returns the AutoDateHistogram field if non-nil, zero value otherwise.

### GetAutoDateHistogramOk

`func (o *MetaAggregations) GetAutoDateHistogramOk() (*MetaAggregationAutoDateHistogram, bool)`

GetAutoDateHistogramOk returns a tuple with the AutoDateHistogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDateHistogram

`func (o *MetaAggregations) SetAutoDateHistogram(v MetaAggregationAutoDateHistogram)`

SetAutoDateHistogram sets AutoDateHistogram field to given value.

### HasAutoDateHistogram

`func (o *MetaAggregations) HasAutoDateHistogram() bool`

HasAutoDateHistogram returns a boolean if a field has been set.

### GetAvg

`func (o *MetaAggregations) GetAvg() MetaAggregationMetric`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *MetaAggregations) GetAvgOk() (*MetaAggregationMetric, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *MetaAggregations) SetAvg(v MetaAggregationMetric)`

SetAvg sets Avg field to given value.

### HasAvg

`func (o *MetaAggregations) HasAvg() bool`

HasAvg returns a boolean if a field has been set.

### GetCardinality

`func (o *MetaAggregations) GetCardinality() MetaAggregationMetric`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *MetaAggregations) GetCardinalityOk() (*MetaAggregationMetric, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *MetaAggregations) SetCardinality(v MetaAggregationMetric)`

SetCardinality sets Cardinality field to given value.

### HasCardinality

`func (o *MetaAggregations) HasCardinality() bool`

HasCardinality returns a boolean if a field has been set.

### GetCount

`func (o *MetaAggregations) GetCount() MetaAggregationMetric`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MetaAggregations) GetCountOk() (*MetaAggregationMetric, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MetaAggregations) SetCount(v MetaAggregationMetric)`

SetCount sets Count field to given value.

### HasCount

`func (o *MetaAggregations) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDateHistogram

`func (o *MetaAggregations) GetDateHistogram() MetaAggregationDateHistogram`

GetDateHistogram returns the DateHistogram field if non-nil, zero value otherwise.

### GetDateHistogramOk

`func (o *MetaAggregations) GetDateHistogramOk() (*MetaAggregationDateHistogram, bool)`

GetDateHistogramOk returns a tuple with the DateHistogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateHistogram

`func (o *MetaAggregations) SetDateHistogram(v MetaAggregationDateHistogram)`

SetDateHistogram sets DateHistogram field to given value.

### HasDateHistogram

`func (o *MetaAggregations) HasDateHistogram() bool`

HasDateHistogram returns a boolean if a field has been set.

### GetDateRange

`func (o *MetaAggregations) GetDateRange() MetaAggregationDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *MetaAggregations) GetDateRangeOk() (*MetaAggregationDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *MetaAggregations) SetDateRange(v MetaAggregationDateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *MetaAggregations) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetHistogram

`func (o *MetaAggregations) GetHistogram() MetaAggregationHistogram`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *MetaAggregations) GetHistogramOk() (*MetaAggregationHistogram, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *MetaAggregations) SetHistogram(v MetaAggregationHistogram)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *MetaAggregations) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetIpRange

`func (o *MetaAggregations) GetIpRange() MetaAggregationIPRange`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *MetaAggregations) GetIpRangeOk() (*MetaAggregationIPRange, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *MetaAggregations) SetIpRange(v MetaAggregationIPRange)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *MetaAggregations) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetMax

`func (o *MetaAggregations) GetMax() MetaAggregationMetric`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *MetaAggregations) GetMaxOk() (*MetaAggregationMetric, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *MetaAggregations) SetMax(v MetaAggregationMetric)`

SetMax sets Max field to given value.

### HasMax

`func (o *MetaAggregations) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *MetaAggregations) GetMin() MetaAggregationMetric`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *MetaAggregations) GetMinOk() (*MetaAggregationMetric, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *MetaAggregations) SetMin(v MetaAggregationMetric)`

SetMin sets Min field to given value.

### HasMin

`func (o *MetaAggregations) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetRange

`func (o *MetaAggregations) GetRange() MetaAggregationRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *MetaAggregations) GetRangeOk() (*MetaAggregationRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *MetaAggregations) SetRange(v MetaAggregationRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *MetaAggregations) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSum

`func (o *MetaAggregations) GetSum() MetaAggregationMetric`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *MetaAggregations) GetSumOk() (*MetaAggregationMetric, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *MetaAggregations) SetSum(v MetaAggregationMetric)`

SetSum sets Sum field to given value.

### HasSum

`func (o *MetaAggregations) HasSum() bool`

HasSum returns a boolean if a field has been set.

### GetTerms

`func (o *MetaAggregations) GetTerms() MetaAggregationsTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *MetaAggregations) GetTermsOk() (*MetaAggregationsTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *MetaAggregations) SetTerms(v MetaAggregationsTerms)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *MetaAggregations) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetWeightedAvg

`func (o *MetaAggregations) GetWeightedAvg() MetaAggregationMetric`

GetWeightedAvg returns the WeightedAvg field if non-nil, zero value otherwise.

### GetWeightedAvgOk

`func (o *MetaAggregations) GetWeightedAvgOk() (*MetaAggregationMetric, bool)`

GetWeightedAvgOk returns a tuple with the WeightedAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAvg

`func (o *MetaAggregations) SetWeightedAvg(v MetaAggregationMetric)`

SetWeightedAvg sets WeightedAvg field to given value.

### HasWeightedAvg

`func (o *MetaAggregations) HasWeightedAvg() bool`

HasWeightedAvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


