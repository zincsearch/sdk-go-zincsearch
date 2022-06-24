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

// MetaHTTPResponseError struct for MetaHTTPResponseError
type MetaHTTPResponseError struct {
	Error *string `json:"error,omitempty"`
}

// NewMetaHTTPResponseError instantiates a new MetaHTTPResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaHTTPResponseError() *MetaHTTPResponseError {
	this := MetaHTTPResponseError{}
	return &this
}

// NewMetaHTTPResponseErrorWithDefaults instantiates a new MetaHTTPResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaHTTPResponseErrorWithDefaults() *MetaHTTPResponseError {
	this := MetaHTTPResponseError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MetaHTTPResponseError) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHTTPResponseError) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MetaHTTPResponseError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *MetaHTTPResponseError) SetError(v string) {
	o.Error = &v
}

func (o MetaHTTPResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMetaHTTPResponseError struct {
	value *MetaHTTPResponseError
	isSet bool
}

func (v NullableMetaHTTPResponseError) Get() *MetaHTTPResponseError {
	return v.value
}

func (v *NullableMetaHTTPResponseError) Set(val *MetaHTTPResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaHTTPResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaHTTPResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaHTTPResponseError(val *MetaHTTPResponseError) *NullableMetaHTTPResponseError {
	return &NullableMetaHTTPResponseError{value: val, isSet: true}
}

func (v NullableMetaHTTPResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaHTTPResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


