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

// MetaAggregationIPRange struct for MetaAggregationIPRange
type MetaAggregationIPRange struct {
	Field *string `json:"field,omitempty"`
	Keyed *bool `json:"keyed,omitempty"`
	Ranges []MetaIPRange `json:"ranges,omitempty"`
}

// NewMetaAggregationIPRange instantiates a new MetaAggregationIPRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaAggregationIPRange() *MetaAggregationIPRange {
	this := MetaAggregationIPRange{}
	return &this
}

// NewMetaAggregationIPRangeWithDefaults instantiates a new MetaAggregationIPRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaAggregationIPRangeWithDefaults() *MetaAggregationIPRange {
	this := MetaAggregationIPRange{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *MetaAggregationIPRange) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationIPRange) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *MetaAggregationIPRange) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *MetaAggregationIPRange) SetField(v string) {
	o.Field = &v
}

// GetKeyed returns the Keyed field value if set, zero value otherwise.
func (o *MetaAggregationIPRange) GetKeyed() bool {
	if o == nil || o.Keyed == nil {
		var ret bool
		return ret
	}
	return *o.Keyed
}

// GetKeyedOk returns a tuple with the Keyed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationIPRange) GetKeyedOk() (*bool, bool) {
	if o == nil || o.Keyed == nil {
		return nil, false
	}
	return o.Keyed, true
}

// HasKeyed returns a boolean if a field has been set.
func (o *MetaAggregationIPRange) HasKeyed() bool {
	if o != nil && o.Keyed != nil {
		return true
	}

	return false
}

// SetKeyed gets a reference to the given bool and assigns it to the Keyed field.
func (o *MetaAggregationIPRange) SetKeyed(v bool) {
	o.Keyed = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *MetaAggregationIPRange) GetRanges() []MetaIPRange {
	if o == nil || o.Ranges == nil {
		var ret []MetaIPRange
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationIPRange) GetRangesOk() ([]MetaIPRange, bool) {
	if o == nil || o.Ranges == nil {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *MetaAggregationIPRange) HasRanges() bool {
	if o != nil && o.Ranges != nil {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []MetaIPRange and assigns it to the Ranges field.
func (o *MetaAggregationIPRange) SetRanges(v []MetaIPRange) {
	o.Ranges = v
}

func (o MetaAggregationIPRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Keyed != nil {
		toSerialize["keyed"] = o.Keyed
	}
	if o.Ranges != nil {
		toSerialize["ranges"] = o.Ranges
	}
	return json.Marshal(toSerialize)
}

type NullableMetaAggregationIPRange struct {
	value *MetaAggregationIPRange
	isSet bool
}

func (v NullableMetaAggregationIPRange) Get() *MetaAggregationIPRange {
	return v.value
}

func (v *NullableMetaAggregationIPRange) Set(val *MetaAggregationIPRange) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaAggregationIPRange) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaAggregationIPRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaAggregationIPRange(val *MetaAggregationIPRange) *NullableMetaAggregationIPRange {
	return &NullableMetaAggregationIPRange{value: val, isSet: true}
}

func (v NullableMetaAggregationIPRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaAggregationIPRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


