# MetaAggregationDateHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalendarInterval** | Pointer to **string** | minute,hour,day,week,month,quarter,year | [optional] 
**ExtendedBounds** | Pointer to [**AggregationHistogramBound**](AggregationHistogramBound.md) |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**FixedInterval** | Pointer to **string** | ms,s,m,h,d | [optional] 
**Format** | Pointer to **string** | format key_as_string | [optional] 
**HardBounds** | Pointer to [**AggregationHistogramBound**](AggregationHistogramBound.md) |  | [optional] 
**Interval** | Pointer to **string** | ms,s,m,h,d | [optional] 
**Keyed** | Pointer to **bool** |  | [optional] 
**MinDocCount** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**TimeZone** | Pointer to **string** | time_zone | [optional] 

## Methods

### NewMetaAggregationDateHistogram

`func NewMetaAggregationDateHistogram() *MetaAggregationDateHistogram`

NewMetaAggregationDateHistogram instantiates a new MetaAggregationDateHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaAggregationDateHistogramWithDefaults

`func NewMetaAggregationDateHistogramWithDefaults() *MetaAggregationDateHistogram`

NewMetaAggregationDateHistogramWithDefaults instantiates a new MetaAggregationDateHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalendarInterval

`func (o *MetaAggregationDateHistogram) GetCalendarInterval() string`

GetCalendarInterval returns the CalendarInterval field if non-nil, zero value otherwise.

### GetCalendarIntervalOk

`func (o *MetaAggregationDateHistogram) GetCalendarIntervalOk() (*string, bool)`

GetCalendarIntervalOk returns a tuple with the CalendarInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarInterval

`func (o *MetaAggregationDateHistogram) SetCalendarInterval(v string)`

SetCalendarInterval sets CalendarInterval field to given value.

### HasCalendarInterval

`func (o *MetaAggregationDateHistogram) HasCalendarInterval() bool`

HasCalendarInterval returns a boolean if a field has been set.

### GetExtendedBounds

`func (o *MetaAggregationDateHistogram) GetExtendedBounds() AggregationHistogramBound`

GetExtendedBounds returns the ExtendedBounds field if non-nil, zero value otherwise.

### GetExtendedBoundsOk

`func (o *MetaAggregationDateHistogram) GetExtendedBoundsOk() (*AggregationHistogramBound, bool)`

GetExtendedBoundsOk returns a tuple with the ExtendedBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedBounds

`func (o *MetaAggregationDateHistogram) SetExtendedBounds(v AggregationHistogramBound)`

SetExtendedBounds sets ExtendedBounds field to given value.

### HasExtendedBounds

`func (o *MetaAggregationDateHistogram) HasExtendedBounds() bool`

HasExtendedBounds returns a boolean if a field has been set.

### GetField

`func (o *MetaAggregationDateHistogram) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MetaAggregationDateHistogram) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MetaAggregationDateHistogram) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MetaAggregationDateHistogram) HasField() bool`

HasField returns a boolean if a field has been set.

### GetFixedInterval

`func (o *MetaAggregationDateHistogram) GetFixedInterval() string`

GetFixedInterval returns the FixedInterval field if non-nil, zero value otherwise.

### GetFixedIntervalOk

`func (o *MetaAggregationDateHistogram) GetFixedIntervalOk() (*string, bool)`

GetFixedIntervalOk returns a tuple with the FixedInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedInterval

`func (o *MetaAggregationDateHistogram) SetFixedInterval(v string)`

SetFixedInterval sets FixedInterval field to given value.

### HasFixedInterval

`func (o *MetaAggregationDateHistogram) HasFixedInterval() bool`

HasFixedInterval returns a boolean if a field has been set.

### GetFormat

`func (o *MetaAggregationDateHistogram) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MetaAggregationDateHistogram) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MetaAggregationDateHistogram) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MetaAggregationDateHistogram) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetHardBounds

`func (o *MetaAggregationDateHistogram) GetHardBounds() AggregationHistogramBound`

GetHardBounds returns the HardBounds field if non-nil, zero value otherwise.

### GetHardBoundsOk

`func (o *MetaAggregationDateHistogram) GetHardBoundsOk() (*AggregationHistogramBound, bool)`

GetHardBoundsOk returns a tuple with the HardBounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounds

`func (o *MetaAggregationDateHistogram) SetHardBounds(v AggregationHistogramBound)`

SetHardBounds sets HardBounds field to given value.

### HasHardBounds

`func (o *MetaAggregationDateHistogram) HasHardBounds() bool`

HasHardBounds returns a boolean if a field has been set.

### GetInterval

`func (o *MetaAggregationDateHistogram) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MetaAggregationDateHistogram) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MetaAggregationDateHistogram) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MetaAggregationDateHistogram) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetKeyed

`func (o *MetaAggregationDateHistogram) GetKeyed() bool`

GetKeyed returns the Keyed field if non-nil, zero value otherwise.

### GetKeyedOk

`func (o *MetaAggregationDateHistogram) GetKeyedOk() (*bool, bool)`

GetKeyedOk returns a tuple with the Keyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyed

`func (o *MetaAggregationDateHistogram) SetKeyed(v bool)`

SetKeyed sets Keyed field to given value.

### HasKeyed

`func (o *MetaAggregationDateHistogram) HasKeyed() bool`

HasKeyed returns a boolean if a field has been set.

### GetMinDocCount

`func (o *MetaAggregationDateHistogram) GetMinDocCount() int32`

GetMinDocCount returns the MinDocCount field if non-nil, zero value otherwise.

### GetMinDocCountOk

`func (o *MetaAggregationDateHistogram) GetMinDocCountOk() (*int32, bool)`

GetMinDocCountOk returns a tuple with the MinDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDocCount

`func (o *MetaAggregationDateHistogram) SetMinDocCount(v int32)`

SetMinDocCount sets MinDocCount field to given value.

### HasMinDocCount

`func (o *MetaAggregationDateHistogram) HasMinDocCount() bool`

HasMinDocCount returns a boolean if a field has been set.

### GetSize

`func (o *MetaAggregationDateHistogram) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MetaAggregationDateHistogram) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MetaAggregationDateHistogram) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MetaAggregationDateHistogram) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTimeZone

`func (o *MetaAggregationDateHistogram) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MetaAggregationDateHistogram) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MetaAggregationDateHistogram) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MetaAggregationDateHistogram) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


