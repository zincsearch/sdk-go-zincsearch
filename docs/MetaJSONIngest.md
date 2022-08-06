# MetaJSONIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **string** |  | [optional] 
**Records** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewMetaJSONIngest

`func NewMetaJSONIngest() *MetaJSONIngest`

NewMetaJSONIngest instantiates a new MetaJSONIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaJSONIngestWithDefaults

`func NewMetaJSONIngestWithDefaults() *MetaJSONIngest`

NewMetaJSONIngestWithDefaults instantiates a new MetaJSONIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *MetaJSONIngest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MetaJSONIngest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MetaJSONIngest) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MetaJSONIngest) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetRecords

`func (o *MetaJSONIngest) GetRecords() []map[string]interface{}`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *MetaJSONIngest) GetRecordsOk() (*[]map[string]interface{}, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *MetaJSONIngest) SetRecords(v []map[string]interface{})`

SetRecords sets Records field to given value.

### HasRecords

`func (o *MetaJSONIngest) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


