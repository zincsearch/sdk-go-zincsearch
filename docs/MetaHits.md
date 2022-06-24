# MetaHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]MetaHit**](MetaHit.md) |  | [optional] 
**MaxScore** | Pointer to **float32** |  | [optional] 
**Total** | Pointer to [**MetaTotal**](MetaTotal.md) |  | [optional] 

## Methods

### NewMetaHits

`func NewMetaHits() *MetaHits`

NewMetaHits instantiates a new MetaHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaHitsWithDefaults

`func NewMetaHitsWithDefaults() *MetaHits`

NewMetaHitsWithDefaults instantiates a new MetaHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *MetaHits) GetHits() []MetaHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *MetaHits) GetHitsOk() (*[]MetaHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *MetaHits) SetHits(v []MetaHit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *MetaHits) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetMaxScore

`func (o *MetaHits) GetMaxScore() float32`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *MetaHits) GetMaxScoreOk() (*float32, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *MetaHits) SetMaxScore(v float32)`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *MetaHits) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### GetTotal

`func (o *MetaHits) GetTotal() MetaTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MetaHits) GetTotalOk() (*MetaTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MetaHits) SetTotal(v MetaTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MetaHits) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


