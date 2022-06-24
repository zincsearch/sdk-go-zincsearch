# MetaIndexTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**IndexPatterns** | Pointer to **[]string** |  | [optional] 
**Priority** | Pointer to **int32** | highest priority is chosen | [optional] 
**Template** | Pointer to [**MetaTemplateTemplate**](MetaTemplateTemplate.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaIndexTemplate

`func NewMetaIndexTemplate() *MetaIndexTemplate`

NewMetaIndexTemplate instantiates a new MetaIndexTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaIndexTemplateWithDefaults

`func NewMetaIndexTemplateWithDefaults() *MetaIndexTemplate`

NewMetaIndexTemplateWithDefaults instantiates a new MetaIndexTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MetaIndexTemplate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetaIndexTemplate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetaIndexTemplate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MetaIndexTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIndexPatterns

`func (o *MetaIndexTemplate) GetIndexPatterns() []string`

GetIndexPatterns returns the IndexPatterns field if non-nil, zero value otherwise.

### GetIndexPatternsOk

`func (o *MetaIndexTemplate) GetIndexPatternsOk() (*[]string, bool)`

GetIndexPatternsOk returns a tuple with the IndexPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPatterns

`func (o *MetaIndexTemplate) SetIndexPatterns(v []string)`

SetIndexPatterns sets IndexPatterns field to given value.

### HasIndexPatterns

`func (o *MetaIndexTemplate) HasIndexPatterns() bool`

HasIndexPatterns returns a boolean if a field has been set.

### GetPriority

`func (o *MetaIndexTemplate) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MetaIndexTemplate) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MetaIndexTemplate) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MetaIndexTemplate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTemplate

`func (o *MetaIndexTemplate) GetTemplate() MetaTemplateTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *MetaIndexTemplate) GetTemplateOk() (*MetaTemplateTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *MetaIndexTemplate) SetTemplate(v MetaTemplateTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *MetaIndexTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MetaIndexTemplate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MetaIndexTemplate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MetaIndexTemplate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MetaIndexTemplate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


