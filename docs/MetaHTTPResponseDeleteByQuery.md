# MetaHTTPResponseDeleteByQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batches** | Pointer to **int32** |  | [optional] 
**Deleted** | Pointer to **int32** |  | [optional] 
**Failures** | Pointer to **[]string** |  | [optional] 
**Noops** | Pointer to **int32** |  | [optional] 
**RequestsPerSecond** | Pointer to **int32** |  | [optional] 
**Retries** | Pointer to [**MetaHttpRetriesResponse**](MetaHttpRetriesResponse.md) |  | [optional] 
**ThrottledMillis** | Pointer to **int32** |  | [optional] 
**ThrottledUntilMillis** | Pointer to **int32** |  | [optional] 
**TimeOut** | Pointer to **bool** |  | [optional] 
**Took** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**VersionConflicts** | Pointer to **int32** |  | [optional] 

## Methods

### NewMetaHTTPResponseDeleteByQuery

`func NewMetaHTTPResponseDeleteByQuery() *MetaHTTPResponseDeleteByQuery`

NewMetaHTTPResponseDeleteByQuery instantiates a new MetaHTTPResponseDeleteByQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaHTTPResponseDeleteByQueryWithDefaults

`func NewMetaHTTPResponseDeleteByQueryWithDefaults() *MetaHTTPResponseDeleteByQuery`

NewMetaHTTPResponseDeleteByQueryWithDefaults instantiates a new MetaHTTPResponseDeleteByQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatches

`func (o *MetaHTTPResponseDeleteByQuery) GetBatches() int32`

GetBatches returns the Batches field if non-nil, zero value otherwise.

### GetBatchesOk

`func (o *MetaHTTPResponseDeleteByQuery) GetBatchesOk() (*int32, bool)`

GetBatchesOk returns a tuple with the Batches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatches

`func (o *MetaHTTPResponseDeleteByQuery) SetBatches(v int32)`

SetBatches sets Batches field to given value.

### HasBatches

`func (o *MetaHTTPResponseDeleteByQuery) HasBatches() bool`

HasBatches returns a boolean if a field has been set.

### GetDeleted

`func (o *MetaHTTPResponseDeleteByQuery) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MetaHTTPResponseDeleteByQuery) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *MetaHTTPResponseDeleteByQuery) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *MetaHTTPResponseDeleteByQuery) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFailures

`func (o *MetaHTTPResponseDeleteByQuery) GetFailures() []string`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *MetaHTTPResponseDeleteByQuery) GetFailuresOk() (*[]string, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *MetaHTTPResponseDeleteByQuery) SetFailures(v []string)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *MetaHTTPResponseDeleteByQuery) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetNoops

`func (o *MetaHTTPResponseDeleteByQuery) GetNoops() int32`

GetNoops returns the Noops field if non-nil, zero value otherwise.

### GetNoopsOk

`func (o *MetaHTTPResponseDeleteByQuery) GetNoopsOk() (*int32, bool)`

GetNoopsOk returns a tuple with the Noops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoops

`func (o *MetaHTTPResponseDeleteByQuery) SetNoops(v int32)`

SetNoops sets Noops field to given value.

### HasNoops

`func (o *MetaHTTPResponseDeleteByQuery) HasNoops() bool`

HasNoops returns a boolean if a field has been set.

### GetRequestsPerSecond

`func (o *MetaHTTPResponseDeleteByQuery) GetRequestsPerSecond() int32`

GetRequestsPerSecond returns the RequestsPerSecond field if non-nil, zero value otherwise.

### GetRequestsPerSecondOk

`func (o *MetaHTTPResponseDeleteByQuery) GetRequestsPerSecondOk() (*int32, bool)`

GetRequestsPerSecondOk returns a tuple with the RequestsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPerSecond

`func (o *MetaHTTPResponseDeleteByQuery) SetRequestsPerSecond(v int32)`

SetRequestsPerSecond sets RequestsPerSecond field to given value.

### HasRequestsPerSecond

`func (o *MetaHTTPResponseDeleteByQuery) HasRequestsPerSecond() bool`

HasRequestsPerSecond returns a boolean if a field has been set.

### GetRetries

`func (o *MetaHTTPResponseDeleteByQuery) GetRetries() MetaHttpRetriesResponse`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *MetaHTTPResponseDeleteByQuery) GetRetriesOk() (*MetaHttpRetriesResponse, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *MetaHTTPResponseDeleteByQuery) SetRetries(v MetaHttpRetriesResponse)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *MetaHTTPResponseDeleteByQuery) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetThrottledMillis

`func (o *MetaHTTPResponseDeleteByQuery) GetThrottledMillis() int32`

GetThrottledMillis returns the ThrottledMillis field if non-nil, zero value otherwise.

### GetThrottledMillisOk

`func (o *MetaHTTPResponseDeleteByQuery) GetThrottledMillisOk() (*int32, bool)`

GetThrottledMillisOk returns a tuple with the ThrottledMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottledMillis

`func (o *MetaHTTPResponseDeleteByQuery) SetThrottledMillis(v int32)`

SetThrottledMillis sets ThrottledMillis field to given value.

### HasThrottledMillis

`func (o *MetaHTTPResponseDeleteByQuery) HasThrottledMillis() bool`

HasThrottledMillis returns a boolean if a field has been set.

### GetThrottledUntilMillis

`func (o *MetaHTTPResponseDeleteByQuery) GetThrottledUntilMillis() int32`

GetThrottledUntilMillis returns the ThrottledUntilMillis field if non-nil, zero value otherwise.

### GetThrottledUntilMillisOk

`func (o *MetaHTTPResponseDeleteByQuery) GetThrottledUntilMillisOk() (*int32, bool)`

GetThrottledUntilMillisOk returns a tuple with the ThrottledUntilMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottledUntilMillis

`func (o *MetaHTTPResponseDeleteByQuery) SetThrottledUntilMillis(v int32)`

SetThrottledUntilMillis sets ThrottledUntilMillis field to given value.

### HasThrottledUntilMillis

`func (o *MetaHTTPResponseDeleteByQuery) HasThrottledUntilMillis() bool`

HasThrottledUntilMillis returns a boolean if a field has been set.

### GetTimeOut

`func (o *MetaHTTPResponseDeleteByQuery) GetTimeOut() bool`

GetTimeOut returns the TimeOut field if non-nil, zero value otherwise.

### GetTimeOutOk

`func (o *MetaHTTPResponseDeleteByQuery) GetTimeOutOk() (*bool, bool)`

GetTimeOutOk returns a tuple with the TimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOut

`func (o *MetaHTTPResponseDeleteByQuery) SetTimeOut(v bool)`

SetTimeOut sets TimeOut field to given value.

### HasTimeOut

`func (o *MetaHTTPResponseDeleteByQuery) HasTimeOut() bool`

HasTimeOut returns a boolean if a field has been set.

### GetTook

`func (o *MetaHTTPResponseDeleteByQuery) GetTook() int32`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *MetaHTTPResponseDeleteByQuery) GetTookOk() (*int32, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *MetaHTTPResponseDeleteByQuery) SetTook(v int32)`

SetTook sets Took field to given value.

### HasTook

`func (o *MetaHTTPResponseDeleteByQuery) HasTook() bool`

HasTook returns a boolean if a field has been set.

### GetTotal

`func (o *MetaHTTPResponseDeleteByQuery) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MetaHTTPResponseDeleteByQuery) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MetaHTTPResponseDeleteByQuery) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MetaHTTPResponseDeleteByQuery) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetVersionConflicts

`func (o *MetaHTTPResponseDeleteByQuery) GetVersionConflicts() int32`

GetVersionConflicts returns the VersionConflicts field if non-nil, zero value otherwise.

### GetVersionConflictsOk

`func (o *MetaHTTPResponseDeleteByQuery) GetVersionConflictsOk() (*int32, bool)`

GetVersionConflictsOk returns a tuple with the VersionConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionConflicts

`func (o *MetaHTTPResponseDeleteByQuery) SetVersionConflicts(v int32)`

SetVersionConflicts sets VersionConflicts field to given value.

### HasVersionConflicts

`func (o *MetaHTTPResponseDeleteByQuery) HasVersionConflicts() bool`

HasVersionConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


