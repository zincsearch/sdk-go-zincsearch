# MetaTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexTemplate** | Pointer to [**MetaIndexTemplate**](MetaIndexTemplate.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaTemplate

`func NewMetaTemplate() *MetaTemplate`

NewMetaTemplate instantiates a new MetaTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaTemplateWithDefaults

`func NewMetaTemplateWithDefaults() *MetaTemplate`

NewMetaTemplateWithDefaults instantiates a new MetaTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexTemplate

`func (o *MetaTemplate) GetIndexTemplate() MetaIndexTemplate`

GetIndexTemplate returns the IndexTemplate field if non-nil, zero value otherwise.

### GetIndexTemplateOk

`func (o *MetaTemplate) GetIndexTemplateOk() (*MetaIndexTemplate, bool)`

GetIndexTemplateOk returns a tuple with the IndexTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexTemplate

`func (o *MetaTemplate) SetIndexTemplate(v MetaIndexTemplate)`

SetIndexTemplate sets IndexTemplate field to given value.

### HasIndexTemplate

`func (o *MetaTemplate) HasIndexTemplate() bool`

HasIndexTemplate returns a boolean if a field has been set.

### GetName

`func (o *MetaTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaTemplate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


