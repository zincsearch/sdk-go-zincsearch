# MetaHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Source** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to **map[string]interface{}** |  | [optional] 
**Highlight** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMetaHit

`func NewMetaHit() *MetaHit`

NewMetaHit instantiates a new MetaHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaHitWithDefaults

`func NewMetaHitWithDefaults() *MetaHit`

NewMetaHitWithDefaults instantiates a new MetaHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *MetaHit) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MetaHit) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MetaHit) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MetaHit) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetId

`func (o *MetaHit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetaHit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetaHit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetaHit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *MetaHit) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MetaHit) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MetaHit) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MetaHit) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetScore

`func (o *MetaHit) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *MetaHit) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *MetaHit) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *MetaHit) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSource

`func (o *MetaHit) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MetaHit) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MetaHit) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *MetaHit) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *MetaHit) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetaHit) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetaHit) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetaHit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFields

`func (o *MetaHit) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MetaHit) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MetaHit) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *MetaHit) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHighlight

`func (o *MetaHit) GetHighlight() map[string]interface{}`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *MetaHit) GetHighlightOk() (*map[string]interface{}, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *MetaHit) SetHighlight(v map[string]interface{})`

SetHighlight sets Highlight field to given value.

### HasHighlight

`func (o *MetaHit) HasHighlight() bool`

HasHighlight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


