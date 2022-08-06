# MetaProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregatable** | Pointer to **bool** |  | [optional] 
**Analyzer** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**map[string]MetaProperty**](MetaProperty.md) | Fields allow the same string value to be indexed in multiple ways for different purposes, such as one field for search and a multi-field for sorting and aggregations, or the same string value analyzed by different analyzers. If the Fields property is defined within a sub-field, it will be ignored.  Currently, only \&quot;text\&quot; fields support the Fields parameter. | [optional] 
**Format** | Pointer to **string** | date format yyyy-MM-dd HH:mm:ss || yyyy-MM-dd || epoch_millis | [optional] 
**Highlightable** | Pointer to **bool** |  | [optional] 
**Index** | Pointer to **bool** |  | [optional] 
**SearchAnalyzer** | Pointer to **string** |  | [optional] 
**Sortable** | Pointer to **bool** |  | [optional] 
**Store** | Pointer to **bool** |  | [optional] 
**TimeZone** | Pointer to **string** | date format time_zone | [optional] 
**Type** | Pointer to **string** | text, keyword, date, numeric, boolean, geo_point | [optional] 

## Methods

### NewMetaProperty

`func NewMetaProperty() *MetaProperty`

NewMetaProperty instantiates a new MetaProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaPropertyWithDefaults

`func NewMetaPropertyWithDefaults() *MetaProperty`

NewMetaPropertyWithDefaults instantiates a new MetaProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregatable

`func (o *MetaProperty) GetAggregatable() bool`

GetAggregatable returns the Aggregatable field if non-nil, zero value otherwise.

### GetAggregatableOk

`func (o *MetaProperty) GetAggregatableOk() (*bool, bool)`

GetAggregatableOk returns a tuple with the Aggregatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatable

`func (o *MetaProperty) SetAggregatable(v bool)`

SetAggregatable sets Aggregatable field to given value.

### HasAggregatable

`func (o *MetaProperty) HasAggregatable() bool`

HasAggregatable returns a boolean if a field has been set.

### GetAnalyzer

`func (o *MetaProperty) GetAnalyzer() string`

GetAnalyzer returns the Analyzer field if non-nil, zero value otherwise.

### GetAnalyzerOk

`func (o *MetaProperty) GetAnalyzerOk() (*string, bool)`

GetAnalyzerOk returns a tuple with the Analyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzer

`func (o *MetaProperty) SetAnalyzer(v string)`

SetAnalyzer sets Analyzer field to given value.

### HasAnalyzer

`func (o *MetaProperty) HasAnalyzer() bool`

HasAnalyzer returns a boolean if a field has been set.

### GetFields

`func (o *MetaProperty) GetFields() map[string]MetaProperty`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MetaProperty) GetFieldsOk() (*map[string]MetaProperty, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MetaProperty) SetFields(v map[string]MetaProperty)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MetaProperty) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFormat

`func (o *MetaProperty) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MetaProperty) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MetaProperty) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MetaProperty) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetHighlightable

`func (o *MetaProperty) GetHighlightable() bool`

GetHighlightable returns the Highlightable field if non-nil, zero value otherwise.

### GetHighlightableOk

`func (o *MetaProperty) GetHighlightableOk() (*bool, bool)`

GetHighlightableOk returns a tuple with the Highlightable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightable

`func (o *MetaProperty) SetHighlightable(v bool)`

SetHighlightable sets Highlightable field to given value.

### HasHighlightable

`func (o *MetaProperty) HasHighlightable() bool`

HasHighlightable returns a boolean if a field has been set.

### GetIndex

`func (o *MetaProperty) GetIndex() bool`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MetaProperty) GetIndexOk() (*bool, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MetaProperty) SetIndex(v bool)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MetaProperty) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetSearchAnalyzer

`func (o *MetaProperty) GetSearchAnalyzer() string`

GetSearchAnalyzer returns the SearchAnalyzer field if non-nil, zero value otherwise.

### GetSearchAnalyzerOk

`func (o *MetaProperty) GetSearchAnalyzerOk() (*string, bool)`

GetSearchAnalyzerOk returns a tuple with the SearchAnalyzer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAnalyzer

`func (o *MetaProperty) SetSearchAnalyzer(v string)`

SetSearchAnalyzer sets SearchAnalyzer field to given value.

### HasSearchAnalyzer

`func (o *MetaProperty) HasSearchAnalyzer() bool`

HasSearchAnalyzer returns a boolean if a field has been set.

### GetSortable

`func (o *MetaProperty) GetSortable() bool`

GetSortable returns the Sortable field if non-nil, zero value otherwise.

### GetSortableOk

`func (o *MetaProperty) GetSortableOk() (*bool, bool)`

GetSortableOk returns a tuple with the Sortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortable

`func (o *MetaProperty) SetSortable(v bool)`

SetSortable sets Sortable field to given value.

### HasSortable

`func (o *MetaProperty) HasSortable() bool`

HasSortable returns a boolean if a field has been set.

### GetStore

`func (o *MetaProperty) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *MetaProperty) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *MetaProperty) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *MetaProperty) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTimeZone

`func (o *MetaProperty) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MetaProperty) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MetaProperty) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MetaProperty) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetType

`func (o *MetaProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetaProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetaProperty) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetaProperty) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


