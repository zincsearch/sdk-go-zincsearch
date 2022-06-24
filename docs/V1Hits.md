# V1Hits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]V1Hit**](V1Hit.md) |  | [optional] 
**MaxScore** | Pointer to **float32** |  | [optional] 
**Total** | Pointer to [**V1Total**](V1Total.md) |  | [optional] 

## Methods

### NewV1Hits

`func NewV1Hits() *V1Hits`

NewV1Hits instantiates a new V1Hits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HitsWithDefaults

`func NewV1HitsWithDefaults() *V1Hits`

NewV1HitsWithDefaults instantiates a new V1Hits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *V1Hits) GetHits() []V1Hit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *V1Hits) GetHitsOk() (*[]V1Hit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *V1Hits) SetHits(v []V1Hit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *V1Hits) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetMaxScore

`func (o *V1Hits) GetMaxScore() float32`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *V1Hits) GetMaxScoreOk() (*float32, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *V1Hits) SetMaxScore(v float32)`

SetMaxScore sets MaxScore field to given value.

### HasMaxScore

`func (o *V1Hits) HasMaxScore() bool`

HasMaxScore returns a boolean if a field has been set.

### GetTotal

`func (o *V1Hits) GetTotal() V1Total`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *V1Hits) GetTotalOk() (*V1Total, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *V1Hits) SetTotal(v V1Total)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *V1Hits) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


