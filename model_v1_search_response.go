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

// V1SearchResponse struct for V1SearchResponse
type V1SearchResponse struct {
	Aggregations *map[string]V1AggregationResponse `json:"aggregations,omitempty"`
	Error *string `json:"error,omitempty"`
	Hits *V1Hits `json:"hits,omitempty"`
	TimedOut *bool `json:"timed_out,omitempty"`
	// Time it took to generate the response
	Took *int32 `json:"took,omitempty"`
}

// NewV1SearchResponse instantiates a new V1SearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SearchResponse() *V1SearchResponse {
	this := V1SearchResponse{}
	return &this
}

// NewV1SearchResponseWithDefaults instantiates a new V1SearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SearchResponseWithDefaults() *V1SearchResponse {
	this := V1SearchResponse{}
	return &this
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *V1SearchResponse) GetAggregations() map[string]V1AggregationResponse {
	if o == nil || o.Aggregations == nil {
		var ret map[string]V1AggregationResponse
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SearchResponse) GetAggregationsOk() (*map[string]V1AggregationResponse, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *V1SearchResponse) HasAggregations() bool {
	if o != nil && o.Aggregations != nil {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given map[string]V1AggregationResponse and assigns it to the Aggregations field.
func (o *V1SearchResponse) SetAggregations(v map[string]V1AggregationResponse) {
	o.Aggregations = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V1SearchResponse) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SearchResponse) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V1SearchResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V1SearchResponse) SetError(v string) {
	o.Error = &v
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *V1SearchResponse) GetHits() V1Hits {
	if o == nil || o.Hits == nil {
		var ret V1Hits
		return ret
	}
	return *o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SearchResponse) GetHitsOk() (*V1Hits, bool) {
	if o == nil || o.Hits == nil {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *V1SearchResponse) HasHits() bool {
	if o != nil && o.Hits != nil {
		return true
	}

	return false
}

// SetHits gets a reference to the given V1Hits and assigns it to the Hits field.
func (o *V1SearchResponse) SetHits(v V1Hits) {
	o.Hits = &v
}

// GetTimedOut returns the TimedOut field value if set, zero value otherwise.
func (o *V1SearchResponse) GetTimedOut() bool {
	if o == nil || o.TimedOut == nil {
		var ret bool
		return ret
	}
	return *o.TimedOut
}

// GetTimedOutOk returns a tuple with the TimedOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SearchResponse) GetTimedOutOk() (*bool, bool) {
	if o == nil || o.TimedOut == nil {
		return nil, false
	}
	return o.TimedOut, true
}

// HasTimedOut returns a boolean if a field has been set.
func (o *V1SearchResponse) HasTimedOut() bool {
	if o != nil && o.TimedOut != nil {
		return true
	}

	return false
}

// SetTimedOut gets a reference to the given bool and assigns it to the TimedOut field.
func (o *V1SearchResponse) SetTimedOut(v bool) {
	o.TimedOut = &v
}

// GetTook returns the Took field value if set, zero value otherwise.
func (o *V1SearchResponse) GetTook() int32 {
	if o == nil || o.Took == nil {
		var ret int32
		return ret
	}
	return *o.Took
}

// GetTookOk returns a tuple with the Took field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SearchResponse) GetTookOk() (*int32, bool) {
	if o == nil || o.Took == nil {
		return nil, false
	}
	return o.Took, true
}

// HasTook returns a boolean if a field has been set.
func (o *V1SearchResponse) HasTook() bool {
	if o != nil && o.Took != nil {
		return true
	}

	return false
}

// SetTook gets a reference to the given int32 and assigns it to the Took field.
func (o *V1SearchResponse) SetTook(v int32) {
	o.Took = &v
}

func (o V1SearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Aggregations != nil {
		toSerialize["aggregations"] = o.Aggregations
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Hits != nil {
		toSerialize["hits"] = o.Hits
	}
	if o.TimedOut != nil {
		toSerialize["timed_out"] = o.TimedOut
	}
	if o.Took != nil {
		toSerialize["took"] = o.Took
	}
	return json.Marshal(toSerialize)
}

type NullableV1SearchResponse struct {
	value *V1SearchResponse
	isSet bool
}

func (v NullableV1SearchResponse) Get() *V1SearchResponse {
	return v.value
}

func (v *NullableV1SearchResponse) Set(val *V1SearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SearchResponse(val *V1SearchResponse) *NullableV1SearchResponse {
	return &NullableV1SearchResponse{value: val, isSet: true}
}

func (v NullableV1SearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


