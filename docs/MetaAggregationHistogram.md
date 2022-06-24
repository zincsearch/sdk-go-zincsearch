# MetaAggregationHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtendedBounds** | Pointer to [**AggregationHistogramBound**](AggregationHistogramBound.md) |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**HardBounds** | Pointer to [**AggregationHistogramBound**](AggregationHistogramBound.md) |  | [optional] 
**Interval** | Pointer to **float32** |  | [optional] 
**Keyed** | Pointer to **bool** |  | [optional] 
**MinDocCount** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **float32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewMetaAggregationHistogram

`func NewMetaAggregationHistogram() *MetaAggregationHistogram`

NewMetaAggregationHistogram instantiates a new MetaAggregationHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaAggregationHistogramWithDefaults

`func NewMetaAggregationHistogramWithDefaults() *MetaAggregationHistogram`

NewMetaAggregationHistogramWithDefaults instantiates a new MetaAggregationHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtendedBounds

`func (o *MetaAggregationHistogram) GetExtendedBounds() AggregationHistogramBound`

GetExtendedBounds returns the ExtendedBounds field if non-nil, zero value otherwise.

### GetExtendedBoundsOk

`func (o *MetaAggregationHistogram) GetExtendedBoundsOk() (*AggregationHistogramBound, bool)`

GetExtendedBoundsOk returns a tuple with the ExtendedBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedBounds

`func (o *MetaAggregationHistogram) SetExtendedBounds(v AggregationHistogramBound)`

SetExtendedBounds sets ExtendedBounds field to given value.

### HasExtendedBounds

`func (o *MetaAggregationHistogram) HasExtendedBounds() bool`

HasExtendedBounds returns a boolean if a field has been set.

### GetField

`func (o *MetaAggregationHistogram) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MetaAggregationHistogram) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MetaAggregationHistogram) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MetaAggregationHistogram) HasField() bool`

HasField returns a boolean if a field has been set.

### GetHardBounds

`func (o *MetaAggregationHistogram) GetHardBounds() AggregationHistogramBound`

GetHardBounds returns the HardBounds field if non-nil, zero value otherwise.

### GetHardBoundsOk

`func (o *MetaAggregationHistogram) GetHardBoundsOk() (*AggregationHistogramBound, bool)`

GetHardBoundsOk returns a tuple with the HardBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounds

`func (o *MetaAggregationHistogram) SetHardBounds(v AggregationHistogramBound)`

SetHardBounds sets HardBounds field to given value.

### HasHardBounds

`func (o *MetaAggregationHistogram) HasHardBounds() bool`

HasHardBounds returns a boolean if a field has been set.

### GetInterval

`func (o *MetaAggregationHistogram) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MetaAggregationHistogram) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MetaAggregationHistogram) SetInterval(v float32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MetaAggregationHistogram) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetKeyed

`func (o *MetaAggregationHistogram) GetKeyed() bool`

GetKeyed returns the Keyed field if non-nil, zero value otherwise.

### GetKeyedOk

`func (o *MetaAggregationHistogram) GetKeyedOk() (*bool, bool)`

GetKeyedOk returns a tuple with the Keyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyed

`func (o *MetaAggregationHistogram) SetKeyed(v bool)`

SetKeyed sets Keyed field to given value.

### HasKeyed

`func (o *MetaAggregationHistogram) HasKeyed() bool`

HasKeyed returns a boolean if a field has been set.

### GetMinDocCount

`func (o *MetaAggregationHistogram) GetMinDocCount() int32`

GetMinDocCount returns the MinDocCount field if non-nil, zero value otherwise.

### GetMinDocCountOk

`func (o *MetaAggregationHistogram) GetMinDocCountOk() (*int32, bool)`

GetMinDocCountOk returns a tuple with the MinDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDocCount

`func (o *MetaAggregationHistogram) SetMinDocCount(v int32)`

SetMinDocCount sets MinDocCount field to given value.

### HasMinDocCount

`func (o *MetaAggregationHistogram) HasMinDocCount() bool`

HasMinDocCount returns a boolean if a field has been set.

### GetOffset

`func (o *MetaAggregationHistogram) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MetaAggregationHistogram) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MetaAggregationHistogram) SetOffset(v float32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *MetaAggregationHistogram) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSize

`func (o *MetaAggregationHistogram) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MetaAggregationHistogram) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MetaAggregationHistogram) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MetaAggregationHistogram) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


