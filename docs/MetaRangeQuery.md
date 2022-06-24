# MetaRangeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boost** | Pointer to **float32** |  | [optional] 
**Format** | Pointer to **string** | Date format used to convert date values in the query. | [optional] 
**Gt** | Pointer to **string** | string, float64 | [optional] 
**Gte** | Pointer to **string** | string, float64 | [optional] 
**Lt** | Pointer to **string** | string, float64 | [optional] 
**Lte** | Pointer to **string** | string, float64 | [optional] 
**TimeZone** | Pointer to **string** | used to convert date values in the query to UTC. | [optional] 

## Methods

### NewMetaRangeQuery

`func NewMetaRangeQuery() *MetaRangeQuery`

NewMetaRangeQuery instantiates a new MetaRangeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaRangeQueryWithDefaults

`func NewMetaRangeQueryWithDefaults() *MetaRangeQuery`

NewMetaRangeQueryWithDefaults instantiates a new MetaRangeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoost

`func (o *MetaRangeQuery) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *MetaRangeQuery) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *MetaRangeQuery) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *MetaRangeQuery) HasBoost() bool`

HasBoost returns a boolean if a field has been set.

### GetFormat

`func (o *MetaRangeQuery) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MetaRangeQuery) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MetaRangeQuery) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MetaRangeQuery) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGt

`func (o *MetaRangeQuery) GetGt() string`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *MetaRangeQuery) GetGtOk() (*string, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *MetaRangeQuery) SetGt(v string)`

SetGt sets Gt field to given value.

### HasGt

`func (o *MetaRangeQuery) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetGte

`func (o *MetaRangeQuery) GetGte() string`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *MetaRangeQuery) GetGteOk() (*string, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *MetaRangeQuery) SetGte(v string)`

SetGte sets Gte field to given value.

### HasGte

`func (o *MetaRangeQuery) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetLt

`func (o *MetaRangeQuery) GetLt() string`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *MetaRangeQuery) GetLtOk() (*string, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *MetaRangeQuery) SetLt(v string)`

SetLt sets Lt field to given value.

### HasLt

`func (o *MetaRangeQuery) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetLte

`func (o *MetaRangeQuery) GetLte() string`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *MetaRangeQuery) GetLteOk() (*string, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *MetaRangeQuery) SetLte(v string)`

SetLte sets Lte field to given value.

### HasLte

`func (o *MetaRangeQuery) HasLte() bool`

HasLte returns a boolean if a field has been set.

### GetTimeZone

`func (o *MetaRangeQuery) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MetaRangeQuery) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MetaRangeQuery) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MetaRangeQuery) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


