# PointAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitiativeId** | **string** | initiative id | 
**ParentId** | Pointer to **string** | parent id | [optional] 
**Title** | **string** | title of point | 
**Desc** | **string** | description of point | 
**PublishedBy** | **string** | published by | 
**LocalCost** | **int32** | local cost | 
**LocalCollected** | **int32** | local collected | 
**CumulativeCost** | **int32** | cumulative cost | 
**CumulativeCollected** | **int32** | cumulative collected | 
**Status** | **string** | status of point | 
**CreatedAt** | **time.Time** | point creation timestamp | 
**UpdatedAt** | Pointer to **time.Time** | point updated timestamp | [optional] 

## Methods

### NewPointAttributes

`func NewPointAttributes(initiativeId string, title string, desc string, publishedBy string, localCost int32, localCollected int32, cumulativeCost int32, cumulativeCollected int32, status string, createdAt time.Time, ) *PointAttributes`

NewPointAttributes instantiates a new PointAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointAttributesWithDefaults

`func NewPointAttributesWithDefaults() *PointAttributes`

NewPointAttributesWithDefaults instantiates a new PointAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiativeId

`func (o *PointAttributes) GetInitiativeId() string`

GetInitiativeId returns the InitiativeId field if non-nil, zero value otherwise.

### GetInitiativeIdOk

`func (o *PointAttributes) GetInitiativeIdOk() (*string, bool)`

GetInitiativeIdOk returns a tuple with the InitiativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiativeId

`func (o *PointAttributes) SetInitiativeId(v string)`

SetInitiativeId sets InitiativeId field to given value.


### GetParentId

`func (o *PointAttributes) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PointAttributes) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PointAttributes) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PointAttributes) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetTitle

`func (o *PointAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PointAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PointAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDesc

`func (o *PointAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *PointAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *PointAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetPublishedBy

`func (o *PointAttributes) GetPublishedBy() string`

GetPublishedBy returns the PublishedBy field if non-nil, zero value otherwise.

### GetPublishedByOk

`func (o *PointAttributes) GetPublishedByOk() (*string, bool)`

GetPublishedByOk returns a tuple with the PublishedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedBy

`func (o *PointAttributes) SetPublishedBy(v string)`

SetPublishedBy sets PublishedBy field to given value.


### GetLocalCost

`func (o *PointAttributes) GetLocalCost() int32`

GetLocalCost returns the LocalCost field if non-nil, zero value otherwise.

### GetLocalCostOk

`func (o *PointAttributes) GetLocalCostOk() (*int32, bool)`

GetLocalCostOk returns a tuple with the LocalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCost

`func (o *PointAttributes) SetLocalCost(v int32)`

SetLocalCost sets LocalCost field to given value.


### GetLocalCollected

`func (o *PointAttributes) GetLocalCollected() int32`

GetLocalCollected returns the LocalCollected field if non-nil, zero value otherwise.

### GetLocalCollectedOk

`func (o *PointAttributes) GetLocalCollectedOk() (*int32, bool)`

GetLocalCollectedOk returns a tuple with the LocalCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCollected

`func (o *PointAttributes) SetLocalCollected(v int32)`

SetLocalCollected sets LocalCollected field to given value.


### GetCumulativeCost

`func (o *PointAttributes) GetCumulativeCost() int32`

GetCumulativeCost returns the CumulativeCost field if non-nil, zero value otherwise.

### GetCumulativeCostOk

`func (o *PointAttributes) GetCumulativeCostOk() (*int32, bool)`

GetCumulativeCostOk returns a tuple with the CumulativeCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeCost

`func (o *PointAttributes) SetCumulativeCost(v int32)`

SetCumulativeCost sets CumulativeCost field to given value.


### GetCumulativeCollected

`func (o *PointAttributes) GetCumulativeCollected() int32`

GetCumulativeCollected returns the CumulativeCollected field if non-nil, zero value otherwise.

### GetCumulativeCollectedOk

`func (o *PointAttributes) GetCumulativeCollectedOk() (*int32, bool)`

GetCumulativeCollectedOk returns a tuple with the CumulativeCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeCollected

`func (o *PointAttributes) SetCumulativeCollected(v int32)`

SetCumulativeCollected sets CumulativeCollected field to given value.


### GetStatus

`func (o *PointAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PointAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PointAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *PointAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PointAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PointAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PointAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PointAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PointAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PointAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


