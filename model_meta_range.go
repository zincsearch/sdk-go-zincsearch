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

// MetaRange struct for MetaRange
type MetaRange struct {
	From *float32 `json:"from,omitempty"`
	To *float32 `json:"to,omitempty"`
}

// NewMetaRange instantiates a new MetaRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaRange() *MetaRange {
	this := MetaRange{}
	return &this
}

// NewMetaRangeWithDefaults instantiates a new MetaRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaRangeWithDefaults() *MetaRange {
	this := MetaRange{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *MetaRange) GetFrom() float32 {
	if o == nil || o.From == nil {
		var ret float32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRange) GetFromOk() (*float32, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *MetaRange) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given float32 and assigns it to the From field.
func (o *MetaRange) SetFrom(v float32) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *MetaRange) GetTo() float32 {
	if o == nil || o.To == nil {
		var ret float32
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRange) GetToOk() (*float32, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *MetaRange) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given float32 and assigns it to the To field.
func (o *MetaRange) SetTo(v float32) {
	o.To = &v
}

func (o MetaRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableMetaRange struct {
	value *MetaRange
	isSet bool
}

func (v NullableMetaRange) Get() *MetaRange {
	return v.value
}

func (v *NullableMetaRange) Set(val *MetaRange) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaRange) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaRange(val *MetaRange) *NullableMetaRange {
	return &NullableMetaRange{value: val, isSet: true}
}

func (v NullableMetaRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


