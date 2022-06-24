# CoreIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateAt** | Pointer to **string** |  | [optional] 
**DocNum** | Pointer to **int32** |  | [optional] 
**DocTimeMax** | Pointer to **int32** |  | [optional] 
**DocTimeMin** | Pointer to **int32** |  | [optional] 
**Mappings** | Pointer to [**MetaMappings**](MetaMappings.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**MetaIndexSettings**](MetaIndexSettings.md) |  | [optional] 
**ShardNum** | Pointer to **int32** |  | [optional] 
**Shards** | Pointer to [**[]MetaIndexShard**](MetaIndexShard.md) |  | [optional] 
**StorageSize** | Pointer to **int32** |  | [optional] 
**StorageType** | Pointer to **string** |  | [optional] 
**UpdateAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCoreIndex

`func NewCoreIndex() *CoreIndex`

NewCoreIndex instantiates a new CoreIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreIndexWithDefaults

`func NewCoreIndexWithDefaults() *CoreIndex`

NewCoreIndexWithDefaults instantiates a new CoreIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateAt

`func (o *CoreIndex) GetCreateAt() string`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *CoreIndex) GetCreateAtOk() (*string, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *CoreIndex) SetCreateAt(v string)`

SetCreateAt sets CreateAt field to given value.

### HasCreateAt

`func (o *CoreIndex) HasCreateAt() bool`

HasCreateAt returns a boolean if a field has been set.

### GetDocNum

`func (o *CoreIndex) GetDocNum() int32`

GetDocNum returns the DocNum field if non-nil, zero value otherwise.

### GetDocNumOk

`func (o *CoreIndex) GetDocNumOk() (*int32, bool)`

GetDocNumOk returns a tuple with the DocNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocNum

`func (o *CoreIndex) SetDocNum(v int32)`

SetDocNum sets DocNum field to given value.

### HasDocNum

`func (o *CoreIndex) HasDocNum() bool`

HasDocNum returns a boolean if a field has been set.

### GetDocTimeMax

`func (o *CoreIndex) GetDocTimeMax() int32`

GetDocTimeMax returns the DocTimeMax field if non-nil, zero value otherwise.

### GetDocTimeMaxOk

`func (o *CoreIndex) GetDocTimeMaxOk() (*int32, bool)`

GetDocTimeMaxOk returns a tuple with the DocTimeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocTimeMax

`func (o *CoreIndex) SetDocTimeMax(v int32)`

SetDocTimeMax sets DocTimeMax field to given value.

### HasDocTimeMax

`func (o *CoreIndex) HasDocTimeMax() bool`

HasDocTimeMax returns a boolean if a field has been set.

### GetDocTimeMin

`func (o *CoreIndex) GetDocTimeMin() int32`

GetDocTimeMin returns the DocTimeMin field if non-nil, zero value otherwise.

### GetDocTimeMinOk

`func (o *CoreIndex) GetDocTimeMinOk() (*int32, bool)`

GetDocTimeMinOk returns a tuple with the DocTimeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocTimeMin

`func (o *CoreIndex) SetDocTimeMin(v int32)`

SetDocTimeMin sets DocTimeMin field to given value.

### HasDocTimeMin

`func (o *CoreIndex) HasDocTimeMin() bool`

HasDocTimeMin returns a boolean if a field has been set.

### GetMappings

`func (o *CoreIndex) GetMappings() MetaMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *CoreIndex) GetMappingsOk() (*MetaMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *CoreIndex) SetMappings(v MetaMappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *CoreIndex) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetName

`func (o *CoreIndex) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreIndex) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreIndex) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CoreIndex) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *CoreIndex) GetSettings() MetaIndexSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CoreIndex) GetSettingsOk() (*MetaIndexSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CoreIndex) SetSettings(v MetaIndexSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CoreIndex) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetShardNum

`func (o *CoreIndex) GetShardNum() int32`

GetShardNum returns the ShardNum field if non-nil, zero value otherwise.

### GetShardNumOk

`func (o *CoreIndex) GetShardNumOk() (*int32, bool)`

GetShardNumOk returns a tuple with the ShardNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardNum

`func (o *CoreIndex) SetShardNum(v int32)`

SetShardNum sets ShardNum field to given value.

### HasShardNum

`func (o *CoreIndex) HasShardNum() bool`

HasShardNum returns a boolean if a field has been set.

### GetShards

`func (o *CoreIndex) GetShards() []MetaIndexShard`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *CoreIndex) GetShardsOk() (*[]MetaIndexShard, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *CoreIndex) SetShards(v []MetaIndexShard)`

SetShards sets Shards field to given value.

### HasShards

`func (o *CoreIndex) HasShards() bool`

HasShards returns a boolean if a field has been set.

### GetStorageSize

`func (o *CoreIndex) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *CoreIndex) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *CoreIndex) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *CoreIndex) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetStorageType

`func (o *CoreIndex) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *CoreIndex) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *CoreIndex) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *CoreIndex) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetUpdateAt

`func (o *CoreIndex) GetUpdateAt() string`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *CoreIndex) GetUpdateAtOk() (*string, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *CoreIndex) SetUpdateAt(v string)`

SetUpdateAt sets UpdateAt field to given value.

### HasUpdateAt

`func (o *CoreIndex) HasUpdateAt() bool`

HasUpdateAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


