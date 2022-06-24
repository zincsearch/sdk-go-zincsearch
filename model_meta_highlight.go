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

// MetaHighlight struct for MetaHighlight
type MetaHighlight struct {
	Fields *map[string]MetaHighlight `json:"fields,omitempty"`
	FragmentSize *int32 `json:"fragment_size,omitempty"`
	NumberOfFragments *int32 `json:"number_of_fragments,omitempty"`
	PostTags []string `json:"post_tags,omitempty"`
	PreTags []string `json:"pre_tags,omitempty"`
}

// NewMetaHighlight instantiates a new MetaHighlight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaHighlight() *MetaHighlight {
	this := MetaHighlight{}
	return &this
}

// NewMetaHighlightWithDefaults instantiates a new MetaHighlight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaHighlightWithDefaults() *MetaHighlight {
	this := MetaHighlight{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MetaHighlight) GetFields() map[string]MetaHighlight {
	if o == nil || o.Fields == nil {
		var ret map[string]MetaHighlight
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHighlight) GetFieldsOk() (*map[string]MetaHighlight, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MetaHighlight) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]MetaHighlight and assigns it to the Fields field.
func (o *MetaHighlight) SetFields(v map[string]MetaHighlight) {
	o.Fields = &v
}

// GetFragmentSize returns the FragmentSize field value if set, zero value otherwise.
func (o *MetaHighlight) GetFragmentSize() int32 {
	if o == nil || o.FragmentSize == nil {
		var ret int32
		return ret
	}
	return *o.FragmentSize
}

// GetFragmentSizeOk returns a tuple with the FragmentSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHighlight) GetFragmentSizeOk() (*int32, bool) {
	if o == nil || o.FragmentSize == nil {
		return nil, false
	}
	return o.FragmentSize, true
}

// HasFragmentSize returns a boolean if a field has been set.
func (o *MetaHighlight) HasFragmentSize() bool {
	if o != nil && o.FragmentSize != nil {
		return true
	}

	return false
}

// SetFragmentSize gets a reference to the given int32 and assigns it to the FragmentSize field.
func (o *MetaHighlight) SetFragmentSize(v int32) {
	o.FragmentSize = &v
}

// GetNumberOfFragments returns the NumberOfFragments field value if set, zero value otherwise.
func (o *MetaHighlight) GetNumberOfFragments() int32 {
	if o == nil || o.NumberOfFragments == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfFragments
}

// GetNumberOfFragmentsOk returns a tuple with the NumberOfFragments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHighlight) GetNumberOfFragmentsOk() (*int32, bool) {
	if o == nil || o.NumberOfFragments == nil {
		return nil, false
	}
	return o.NumberOfFragments, true
}

// HasNumberOfFragments returns a boolean if a field has been set.
func (o *MetaHighlight) HasNumberOfFragments() bool {
	if o != nil && o.NumberOfFragments != nil {
		return true
	}

	return false
}

// SetNumberOfFragments gets a reference to the given int32 and assigns it to the NumberOfFragments field.
func (o *MetaHighlight) SetNumberOfFragments(v int32) {
	o.NumberOfFragments = &v
}

// GetPostTags returns the PostTags field value if set, zero value otherwise.
func (o *MetaHighlight) GetPostTags() []string {
	if o == nil || o.PostTags == nil {
		var ret []string
		return ret
	}
	return o.PostTags
}

// GetPostTagsOk returns a tuple with the PostTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHighlight) GetPostTagsOk() ([]string, bool) {
	if o == nil || o.PostTags == nil {
		return nil, false
	}
	return o.PostTags, true
}

// HasPostTags returns a boolean if a field has been set.
func (o *MetaHighlight) HasPostTags() bool {
	if o != nil && o.PostTags != nil {
		return true
	}

	return false
}

// SetPostTags gets a reference to the given []string and assigns it to the PostTags field.
func (o *MetaHighlight) SetPostTags(v []string) {
	o.PostTags = v
}

// GetPreTags returns the PreTags field value if set, zero value otherwise.
func (o *MetaHighlight) GetPreTags() []string {
	if o == nil || o.PreTags == nil {
		var ret []string
		return ret
	}
	return o.PreTags
}

// GetPreTagsOk returns a tuple with the PreTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHighlight) GetPreTagsOk() ([]string, bool) {
	if o == nil || o.PreTags == nil {
		return nil, false
	}
	return o.PreTags, true
}

// HasPreTags returns a boolean if a field has been set.
func (o *MetaHighlight) HasPreTags() bool {
	if o != nil && o.PreTags != nil {
		return true
	}

	return false
}

// SetPreTags gets a reference to the given []string and assigns it to the PreTags field.
func (o *MetaHighlight) SetPreTags(v []string) {
	o.PreTags = v
}

func (o MetaHighlight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.FragmentSize != nil {
		toSerialize["fragment_size"] = o.FragmentSize
	}
	if o.NumberOfFragments != nil {
		toSerialize["number_of_fragments"] = o.NumberOfFragments
	}
	if o.PostTags != nil {
		toSerialize["post_tags"] = o.PostTags
	}
	if o.PreTags != nil {
		toSerialize["pre_tags"] = o.PreTags
	}
	return json.Marshal(toSerialize)
}

type NullableMetaHighlight struct {
	value *MetaHighlight
	isSet bool
}

func (v NullableMetaHighlight) Get() *MetaHighlight {
	return v.value
}

func (v *NullableMetaHighlight) Set(val *MetaHighlight) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaHighlight) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaHighlight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaHighlight(val *MetaHighlight) *NullableMetaHighlight {
	return &NullableMetaHighlight{value: val, isSet: true}
}

func (v NullableMetaHighlight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaHighlight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


