/*
Zinc Search engine API

Zinc Search engine API documents https://docs.zincsearch.com

API version: 0.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// MetaAggregationDateHistogram struct for MetaAggregationDateHistogram
type MetaAggregationDateHistogram struct {
	// minute,hour,day,week,month,quarter,year
	CalendarInterval *string `json:"calendar_interval,omitempty"`
	ExtendedBounds *AggregationHistogramBound `json:"extended_bounds,omitempty"`
	Field *string `json:"field,omitempty"`
	// ms,s,m,h,d
	FixedInterval *string `json:"fixed_interval,omitempty"`
	// format key_as_string
	Format *string `json:"format,omitempty"`
	HardBounds *AggregationHistogramBound `json:"hard_bounds,omitempty"`
	// ms,s,m,h,d
	Interval *string `json:"interval,omitempty"`
	Keyed *bool `json:"keyed,omitempty"`
	MinDocCount *int32 `json:"min_doc_count,omitempty"`
	Size *int32 `json:"size,omitempty"`
	// time_zone
	TimeZone *string `json:"time_zone,omitempty"`
}

// NewMetaAggregationDateHistogram instantiates a new MetaAggregationDateHistogram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaAggregationDateHistogram() *MetaAggregationDateHistogram {
	this := MetaAggregationDateHistogram{}
	return &this
}

// NewMetaAggregationDateHistogramWithDefaults instantiates a new MetaAggregationDateHistogram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaAggregationDateHistogramWithDefaults() *MetaAggregationDateHistogram {
	this := MetaAggregationDateHistogram{}
	return &this
}

// GetCalendarInterval returns the CalendarInterval field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetCalendarInterval() string {
	if o == nil || o.CalendarInterval == nil {
		var ret string
		return ret
	}
	return *o.CalendarInterval
}

// GetCalendarIntervalOk returns a tuple with the CalendarInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetCalendarIntervalOk() (*string, bool) {
	if o == nil || o.CalendarInterval == nil {
		return nil, false
	}
	return o.CalendarInterval, true
}

// HasCalendarInterval returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasCalendarInterval() bool {
	if o != nil && o.CalendarInterval != nil {
		return true
	}

	return false
}

// SetCalendarInterval gets a reference to the given string and assigns it to the CalendarInterval field.
func (o *MetaAggregationDateHistogram) SetCalendarInterval(v string) {
	o.CalendarInterval = &v
}

// GetExtendedBounds returns the ExtendedBounds field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetExtendedBounds() AggregationHistogramBound {
	if o == nil || o.ExtendedBounds == nil {
		var ret AggregationHistogramBound
		return ret
	}
	return *o.ExtendedBounds
}

// GetExtendedBoundsOk returns a tuple with the ExtendedBounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetExtendedBoundsOk() (*AggregationHistogramBound, bool) {
	if o == nil || o.ExtendedBounds == nil {
		return nil, false
	}
	return o.ExtendedBounds, true
}

// HasExtendedBounds returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasExtendedBounds() bool {
	if o != nil && o.ExtendedBounds != nil {
		return true
	}

	return false
}

// SetExtendedBounds gets a reference to the given AggregationHistogramBound and assigns it to the ExtendedBounds field.
func (o *MetaAggregationDateHistogram) SetExtendedBounds(v AggregationHistogramBound) {
	o.ExtendedBounds = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *MetaAggregationDateHistogram) SetField(v string) {
	o.Field = &v
}

// GetFixedInterval returns the FixedInterval field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetFixedInterval() string {
	if o == nil || o.FixedInterval == nil {
		var ret string
		return ret
	}
	return *o.FixedInterval
}

// GetFixedIntervalOk returns a tuple with the FixedInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetFixedIntervalOk() (*string, bool) {
	if o == nil || o.FixedInterval == nil {
		return nil, false
	}
	return o.FixedInterval, true
}

// HasFixedInterval returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasFixedInterval() bool {
	if o != nil && o.FixedInterval != nil {
		return true
	}

	return false
}

// SetFixedInterval gets a reference to the given string and assigns it to the FixedInterval field.
func (o *MetaAggregationDateHistogram) SetFixedInterval(v string) {
	o.FixedInterval = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *MetaAggregationDateHistogram) SetFormat(v string) {
	o.Format = &v
}

// GetHardBounds returns the HardBounds field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetHardBounds() AggregationHistogramBound {
	if o == nil || o.HardBounds == nil {
		var ret AggregationHistogramBound
		return ret
	}
	return *o.HardBounds
}

// GetHardBoundsOk returns a tuple with the HardBounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetHardBoundsOk() (*AggregationHistogramBound, bool) {
	if o == nil || o.HardBounds == nil {
		return nil, false
	}
	return o.HardBounds, true
}

// HasHardBounds returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasHardBounds() bool {
	if o != nil && o.HardBounds != nil {
		return true
	}

	return false
}

// SetHardBounds gets a reference to the given AggregationHistogramBound and assigns it to the HardBounds field.
func (o *MetaAggregationDateHistogram) SetHardBounds(v AggregationHistogramBound) {
	o.HardBounds = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *MetaAggregationDateHistogram) SetInterval(v string) {
	o.Interval = &v
}

// GetKeyed returns the Keyed field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetKeyed() bool {
	if o == nil || o.Keyed == nil {
		var ret bool
		return ret
	}
	return *o.Keyed
}

// GetKeyedOk returns a tuple with the Keyed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetKeyedOk() (*bool, bool) {
	if o == nil || o.Keyed == nil {
		return nil, false
	}
	return o.Keyed, true
}

// HasKeyed returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasKeyed() bool {
	if o != nil && o.Keyed != nil {
		return true
	}

	return false
}

// SetKeyed gets a reference to the given bool and assigns it to the Keyed field.
func (o *MetaAggregationDateHistogram) SetKeyed(v bool) {
	o.Keyed = &v
}

// GetMinDocCount returns the MinDocCount field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetMinDocCount() int32 {
	if o == nil || o.MinDocCount == nil {
		var ret int32
		return ret
	}
	return *o.MinDocCount
}

// GetMinDocCountOk returns a tuple with the MinDocCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetMinDocCountOk() (*int32, bool) {
	if o == nil || o.MinDocCount == nil {
		return nil, false
	}
	return o.MinDocCount, true
}

// HasMinDocCount returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasMinDocCount() bool {
	if o != nil && o.MinDocCount != nil {
		return true
	}

	return false
}

// SetMinDocCount gets a reference to the given int32 and assigns it to the MinDocCount field.
func (o *MetaAggregationDateHistogram) SetMinDocCount(v int32) {
	o.MinDocCount = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *MetaAggregationDateHistogram) SetSize(v int32) {
	o.Size = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *MetaAggregationDateHistogram) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationDateHistogram) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *MetaAggregationDateHistogram) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *MetaAggregationDateHistogram) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o MetaAggregationDateHistogram) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CalendarInterval != nil {
		toSerialize["calendar_interval"] = o.CalendarInterval
	}
	if o.ExtendedBounds != nil {
		toSerialize["extended_bounds"] = o.ExtendedBounds
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.FixedInterval != nil {
		toSerialize["fixed_interval"] = o.FixedInterval
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.HardBounds != nil {
		toSerialize["hard_bounds"] = o.HardBounds
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Keyed != nil {
		toSerialize["keyed"] = o.Keyed
	}
	if o.MinDocCount != nil {
		toSerialize["min_doc_count"] = o.MinDocCount
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.TimeZone != nil {
		toSerialize["time_zone"] = o.TimeZone
	}
	return json.Marshal(toSerialize)
}

type NullableMetaAggregationDateHistogram struct {
	value *MetaAggregationDateHistogram
	isSet bool
}

func (v NullableMetaAggregationDateHistogram) Get() *MetaAggregationDateHistogram {
	return v.value
}

func (v *NullableMetaAggregationDateHistogram) Set(val *MetaAggregationDateHistogram) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaAggregationDateHistogram) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaAggregationDateHistogram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaAggregationDateHistogram(val *MetaAggregationDateHistogram) *NullableMetaAggregationDateHistogram {
	return &NullableMetaAggregationDateHistogram{value: val, isSet: true}
}

func (v NullableMetaAggregationDateHistogram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaAggregationDateHistogram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

