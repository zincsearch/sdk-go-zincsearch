# V1AggregationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggType** | Pointer to **string** |  | [optional] 
**Aggs** | Pointer to [**map[string]V1AggregationParams**](V1AggregationParams.md) |  | [optional] 
**DateRanges** | Pointer to [**[]V1AggregationDateRange**](V1AggregationDateRange.md) |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**Ranges** | Pointer to [**[]V1AggregationNumberRange**](V1AggregationNumberRange.md) |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**WeightField** | Pointer to **string** | Field name to be used for setting weight for primary field for weighted average aggregation | [optional] 

## Methods

### NewV1AggregationParams

`func NewV1AggregationParams() *V1AggregationParams`

NewV1AggregationParams instantiates a new V1AggregationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AggregationParamsWithDefaults

`func NewV1AggregationParamsWithDefaults() *V1AggregationParams`

NewV1AggregationParamsWithDefaults instantiates a new V1AggregationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggType

`func (o *V1AggregationParams) GetAggType() string`

GetAggType returns the AggType field if non-nil, zero value otherwise.

### GetAggTypeOk

`func (o *V1AggregationParams) GetAggTypeOk() (*string, bool)`

GetAggTypeOk returns a tuple with the AggType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggType

`func (o *V1AggregationParams) SetAggType(v string)`

SetAggType sets AggType field to given value.

### HasAggType

`func (o *V1AggregationParams) HasAggType() bool`

HasAggType returns a boolean if a field has been set.

### GetAggs

`func (o *V1AggregationParams) GetAggs() map[string]V1AggregationParams`

GetAggs returns the Aggs field if non-nil, zero value otherwise.

### GetAggsOk

`func (o *V1AggregationParams) GetAggsOk() (*map[string]V1AggregationParams, bool)`

GetAggsOk returns a tuple with the Aggs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggs

`func (o *V1AggregationParams) SetAggs(v map[string]V1AggregationParams)`

SetAggs sets Aggs field to given value.

### HasAggs

`func (o *V1AggregationParams) HasAggs() bool`

HasAggs returns a boolean if a field has been set.

### GetDateRanges

`func (o *V1AggregationParams) GetDateRanges() []V1AggregationDateRange`

GetDateRanges returns the DateRanges field if non-nil, zero value otherwise.

### GetDateRangesOk

`func (o *V1AggregationParams) GetDateRangesOk() (*[]V1AggregationDateRange, bool)`

GetDateRangesOk returns a tuple with the DateRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRanges

`func (o *V1AggregationParams) SetDateRanges(v []V1AggregationDateRange)`

SetDateRanges sets DateRanges field to given value.

### HasDateRanges

`func (o *V1AggregationParams) HasDateRanges() bool`

HasDateRanges returns a boolean if a field has been set.

### GetField

`func (o *V1AggregationParams) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *V1AggregationParams) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *V1AggregationParams) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *V1AggregationParams) HasField() bool`

HasField returns a boolean if a field has been set.

### GetRanges

`func (o *V1AggregationParams) GetRanges() []V1AggregationNumberRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *V1AggregationParams) GetRangesOk() (*[]V1AggregationNumberRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *V1AggregationParams) SetRanges(v []V1AggregationNumberRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *V1AggregationParams) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetSize

`func (o *V1AggregationParams) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *V1AggregationParams) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *V1AggregationParams) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *V1AggregationParams) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetWeightField

`func (o *V1AggregationParams) GetWeightField() string`

GetWeightField returns the WeightField field if non-nil, zero value otherwise.

### GetWeightFieldOk

`func (o *V1AggregationParams) GetWeightFieldOk() (*string, bool)`

GetWeightFieldOk returns a tuple with the WeightField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightField

`func (o *V1AggregationParams) SetWeightField(v string)`

SetWeightField sets WeightField field to given value.

### HasWeightField

`func (o *V1AggregationParams) HasWeightField() bool`

HasWeightField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


