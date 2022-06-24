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

// MetaTermQuery struct for MetaTermQuery
type MetaTermQuery struct {
	Boost *float32 `json:"boost,omitempty"`
	CaseInsensitive *bool `json:"case_insensitive,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewMetaTermQuery instantiates a new MetaTermQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaTermQuery() *MetaTermQuery {
	this := MetaTermQuery{}
	return &this
}

// NewMetaTermQueryWithDefaults instantiates a new MetaTermQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaTermQueryWithDefaults() *MetaTermQuery {
	this := MetaTermQuery{}
	return &this
}

// GetBoost returns the Boost field value if set, zero value otherwise.
func (o *MetaTermQuery) GetBoost() float32 {
	if o == nil || o.Boost == nil {
		var ret float32
		return ret
	}
	return *o.Boost
}

// GetBoostOk returns a tuple with the Boost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaTermQuery) GetBoostOk() (*float32, bool) {
	if o == nil || o.Boost == nil {
		return nil, false
	}
	return o.Boost, true
}

// HasBoost returns a boolean if a field has been set.
func (o *MetaTermQuery) HasBoost() bool {
	if o != nil && o.Boost != nil {
		return true
	}

	return false
}

// SetBoost gets a reference to the given float32 and assigns it to the Boost field.
func (o *MetaTermQuery) SetBoost(v float32) {
	o.Boost = &v
}

// GetCaseInsensitive returns the CaseInsensitive field value if set, zero value otherwise.
func (o *MetaTermQuery) GetCaseInsensitive() bool {
	if o == nil || o.CaseInsensitive == nil {
		var ret bool
		return ret
	}
	return *o.CaseInsensitive
}

// GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaTermQuery) GetCaseInsensitiveOk() (*bool, bool) {
	if o == nil || o.CaseInsensitive == nil {
		return nil, false
	}
	return o.CaseInsensitive, true
}

// HasCaseInsensitive returns a boolean if a field has been set.
func (o *MetaTermQuery) HasCaseInsensitive() bool {
	if o != nil && o.CaseInsensitive != nil {
		return true
	}

	return false
}

// SetCaseInsensitive gets a reference to the given bool and assigns it to the CaseInsensitive field.
func (o *MetaTermQuery) SetCaseInsensitive(v bool) {
	o.CaseInsensitive = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetaTermQuery) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaTermQuery) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetaTermQuery) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MetaTermQuery) SetValue(v string) {
	o.Value = &v
}

func (o MetaTermQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Boost != nil {
		toSerialize["boost"] = o.Boost
	}
	if o.CaseInsensitive != nil {
		toSerialize["case_insensitive"] = o.CaseInsensitive
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableMetaTermQuery struct {
	value *MetaTermQuery
	isSet bool
}

func (v NullableMetaTermQuery) Get() *MetaTermQuery {
	return v.value
}

func (v *NullableMetaTermQuery) Set(val *MetaTermQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaTermQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaTermQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaTermQuery(val *MetaTermQuery) *NullableMetaTermQuery {
	return &NullableMetaTermQuery{value: val, isSet: true}
}

func (v NullableMetaTermQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaTermQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

