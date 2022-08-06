# MetaIndexSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mappings** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**MetaIndexSettings**](MetaIndexSettings.md) |  | [optional] 
**ShardNum** | Pointer to **int32** |  | [optional] 
**StorageType** | Pointer to **string** |  | [optional] 

## Methods

### NewMetaIndexSimple

`func NewMetaIndexSimple() *MetaIndexSimple`

NewMetaIndexSimple instantiates a new MetaIndexSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaIndexSimpleWithDefaults

`func NewMetaIndexSimpleWithDefaults() *MetaIndexSimple`

NewMetaIndexSimpleWithDefaults instantiates a new MetaIndexSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappings

`func (o *MetaIndexSimple) GetMappings() map[string]interface{}`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *MetaIndexSimple) GetMappingsOk() (*map[string]interface{}, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *MetaIndexSimple) SetMappings(v map[string]interface{})`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *MetaIndexSimple) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetName

`func (o *MetaIndexSimple) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaIndexSimple) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaIndexSimple) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaIndexSimple) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSettings

`func (o *MetaIndexSimple) GetSettings() MetaIndexSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MetaIndexSimple) GetSettingsOk() (*MetaIndexSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *MetaIndexSimple) SetSettings(v MetaIndexSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *MetaIndexSimple) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetShardNum

`func (o *MetaIndexSimple) GetShardNum() int32`

GetShardNum returns the ShardNum field if non-nil, zero value otherwise.

### GetShardNumOk

`func (o *MetaIndexSimple) GetShardNumOk() (*int32, bool)`

GetShardNumOk returns a tuple with the ShardNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardNum

`func (o *MetaIndexSimple) SetShardNum(v int32)`

SetShardNum sets ShardNum field to given value.

### HasShardNum

`func (o *MetaIndexSimple) HasShardNum() bool`

HasShardNum returns a boolean if a field has been set.

### GetStorageType

`func (o *MetaIndexSimple) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *MetaIndexSimple) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *MetaIndexSimple) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *MetaIndexSimple) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


