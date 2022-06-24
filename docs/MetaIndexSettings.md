# MetaIndexSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analysis** | Pointer to [**MetaIndexAnalysis**](MetaIndexAnalysis.md) |  | [optional] 
**NumberOfReplicas** | Pointer to **int32** |  | [optional] 
**NumberOfShards** | Pointer to **int32** |  | [optional] 

## Methods

### NewMetaIndexSettings

`func NewMetaIndexSettings() *MetaIndexSettings`

NewMetaIndexSettings instantiates a new MetaIndexSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaIndexSettingsWithDefaults

`func NewMetaIndexSettingsWithDefaults() *MetaIndexSettings`

NewMetaIndexSettingsWithDefaults instantiates a new MetaIndexSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysis

`func (o *MetaIndexSettings) GetAnalysis() MetaIndexAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *MetaIndexSettings) GetAnalysisOk() (*MetaIndexAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *MetaIndexSettings) SetAnalysis(v MetaIndexAnalysis)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *MetaIndexSettings) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetNumberOfReplicas

`func (o *MetaIndexSettings) GetNumberOfReplicas() int32`

GetNumberOfReplicas returns the NumberOfReplicas field if non-nil, zero value otherwise.

### GetNumberOfReplicasOk

`func (o *MetaIndexSettings) GetNumberOfReplicasOk() (*int32, bool)`

GetNumberOfReplicasOk returns a tuple with the NumberOfReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfReplicas

`func (o *MetaIndexSettings) SetNumberOfReplicas(v int32)`

SetNumberOfReplicas sets NumberOfReplicas field to given value.

### HasNumberOfReplicas

`func (o *MetaIndexSettings) HasNumberOfReplicas() bool`

HasNumberOfReplicas returns a boolean if a field has been set.

### GetNumberOfShards

`func (o *MetaIndexSettings) GetNumberOfShards() int32`

GetNumberOfShards returns the NumberOfShards field if non-nil, zero value otherwise.

### GetNumberOfShardsOk

`func (o *MetaIndexSettings) GetNumberOfShardsOk() (*int32, bool)`

GetNumberOfShardsOk returns a tuple with the NumberOfShards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfShards

`func (o *MetaIndexSettings) SetNumberOfShards(v int32)`

SetNumberOfShards sets NumberOfShards field to given value.

### HasNumberOfShards

`func (o *MetaIndexSettings) HasNumberOfShards() bool`

HasNumberOfShards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


