# IndexIndexListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Page** | Pointer to [**MetaPage**](MetaPage.md) |  | [optional] 

## Methods

### NewIndexIndexListResponse

`func NewIndexIndexListResponse() *IndexIndexListResponse`

NewIndexIndexListResponse instantiates a new IndexIndexListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexIndexListResponseWithDefaults

`func NewIndexIndexListResponseWithDefaults() *IndexIndexListResponse`

NewIndexIndexListResponseWithDefaults instantiates a new IndexIndexListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *IndexIndexListResponse) GetList() []map[string]interface{}`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *IndexIndexListResponse) GetListOk() (*[]map[string]interface{}, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *IndexIndexListResponse) SetList(v []map[string]interface{})`

SetList sets List field to given value.

### HasList

`func (o *IndexIndexListResponse) HasList() bool`

HasList returns a boolean if a field has been set.

### GetPage

`func (o *IndexIndexListResponse) GetPage() MetaPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *IndexIndexListResponse) GetPageOk() (*MetaPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *IndexIndexListResponse) SetPage(v MetaPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *IndexIndexListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


