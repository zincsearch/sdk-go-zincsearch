# MetaAggregationDateRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** | format the &#x60;to&#x60; and &#x60;from&#x60; values to &#x60;_as_string&#x60;, and used to format &#x60;keyed response&#x60; | [optional] 
**Keyed** | Pointer to **bool** |  | [optional] 
**Ranges** | Pointer to [**[]MetaDateRange**](MetaDateRange.md) | refer | [optional] 
**TimeZone** | Pointer to **string** | refer | [optional] 

## Methods

### NewMetaAggregationDateRange

`func NewMetaAggregationDateRange() *MetaAggregationDateRange`

NewMetaAggregationDateRange instantiates a new MetaAggregationDateRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaAggregationDateRangeWithDefaults

`func NewMetaAggregationDateRangeWithDefaults() *MetaAggregationDateRange`

NewMetaAggregationDateRangeWithDefaults instantiates a new MetaAggregationDateRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *MetaAggregationDateRange) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MetaAggregationDateRange) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MetaAggregationDateRange) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MetaAggregationDateRange) HasField() bool`

HasField returns a boolean if a field has been set.

### GetFormat

`func (o *MetaAggregationDateRange) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MetaAggregationDateRange) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MetaAggregationDateRange) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MetaAggregationDateRange) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetKeyed

`func (o *MetaAggregationDateRange) GetKeyed() bool`

GetKeyed returns the Keyed field if non-nil, zero value otherwise.

### GetKeyedOk

`func (o *MetaAggregationDateRange) GetKeyedOk() (*bool, bool)`

GetKeyedOk returns a tuple with the Keyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyed

`func (o *MetaAggregationDateRange) SetKeyed(v bool)`

SetKeyed sets Keyed field to given value.

### HasKeyed

`func (o *MetaAggregationDateRange) HasKeyed() bool`

HasKeyed returns a boolean if a field has been set.

### GetRanges

`func (o *MetaAggregationDateRange) GetRanges() []MetaDateRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *MetaAggregationDateRange) GetRangesOk() (*[]MetaDateRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *MetaAggregationDateRange) SetRanges(v []MetaDateRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *MetaAggregationDateRange) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetTimeZone

`func (o *MetaAggregationDateRange) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MetaAggregationDateRange) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MetaAggregationDateRange) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MetaAggregationDateRange) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


