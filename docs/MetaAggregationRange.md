# MetaAggregationRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Keyed** | Pointer to **bool** |  | [optional] 
**Ranges** | Pointer to [**[]MetaRange**](MetaRange.md) |  | [optional] 

## Methods

### NewMetaAggregationRange

`func NewMetaAggregationRange() *MetaAggregationRange`

NewMetaAggregationRange instantiates a new MetaAggregationRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaAggregationRangeWithDefaults

`func NewMetaAggregationRangeWithDefaults() *MetaAggregationRange`

NewMetaAggregationRangeWithDefaults instantiates a new MetaAggregationRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *MetaAggregationRange) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MetaAggregationRange) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MetaAggregationRange) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MetaAggregationRange) HasField() bool`

HasField returns a boolean if a field has been set.

### GetKeyed

`func (o *MetaAggregationRange) GetKeyed() bool`

GetKeyed returns the Keyed field if non-nil, zero value otherwise.

### GetKeyedOk

`func (o *MetaAggregationRange) GetKeyedOk() (*bool, bool)`

GetKeyedOk returns a tuple with the Keyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyed

`func (o *MetaAggregationRange) SetKeyed(v bool)`

SetKeyed sets Keyed field to given value.

### HasKeyed

`func (o *MetaAggregationRange) HasKeyed() bool`

HasKeyed returns a boolean if a field has been set.

### GetRanges

`func (o *MetaAggregationRange) GetRanges() []MetaRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *MetaAggregationRange) GetRangesOk() (*[]MetaRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *MetaAggregationRange) SetRanges(v []MetaRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *MetaAggregationRange) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


