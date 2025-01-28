# PointUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **string** | parent id | [optional] 
**Title** | **string** | title of point | 
**Desc** | **string** | description of point | 
**LocalCost** | **int32** | local cost | 

## Methods

### NewPointUpdateDataAttributes

`func NewPointUpdateDataAttributes(title string, desc string, localCost int32, ) *PointUpdateDataAttributes`

NewPointUpdateDataAttributes instantiates a new PointUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointUpdateDataAttributesWithDefaults

`func NewPointUpdateDataAttributesWithDefaults() *PointUpdateDataAttributes`

NewPointUpdateDataAttributesWithDefaults instantiates a new PointUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *PointUpdateDataAttributes) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PointUpdateDataAttributes) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PointUpdateDataAttributes) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PointUpdateDataAttributes) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetTitle

`func (o *PointUpdateDataAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PointUpdateDataAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PointUpdateDataAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDesc

`func (o *PointUpdateDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *PointUpdateDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *PointUpdateDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetLocalCost

`func (o *PointUpdateDataAttributes) GetLocalCost() int32`

GetLocalCost returns the LocalCost field if non-nil, zero value otherwise.

### GetLocalCostOk

`func (o *PointUpdateDataAttributes) GetLocalCostOk() (*int32, bool)`

GetLocalCostOk returns a tuple with the LocalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCost

`func (o *PointUpdateDataAttributes) SetLocalCost(v int32)`

SetLocalCost sets LocalCost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


