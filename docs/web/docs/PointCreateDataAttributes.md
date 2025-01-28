# PointCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitiativeId** | **string** | initiative id | 
**ParentId** | Pointer to **string** | parent id | [optional] 
**Title** | **string** | title of point | 
**Desc** | **string** | description of point | 
**LocalCost** | **int32** | local cost | 
**Jar** | [**JarAttributes**](JarAttributes.md) |  | 

## Methods

### NewPointCreateDataAttributes

`func NewPointCreateDataAttributes(initiativeId string, title string, desc string, localCost int32, jar JarAttributes, ) *PointCreateDataAttributes`

NewPointCreateDataAttributes instantiates a new PointCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointCreateDataAttributesWithDefaults

`func NewPointCreateDataAttributesWithDefaults() *PointCreateDataAttributes`

NewPointCreateDataAttributesWithDefaults instantiates a new PointCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiativeId

`func (o *PointCreateDataAttributes) GetInitiativeId() string`

GetInitiativeId returns the InitiativeId field if non-nil, zero value otherwise.

### GetInitiativeIdOk

`func (o *PointCreateDataAttributes) GetInitiativeIdOk() (*string, bool)`

GetInitiativeIdOk returns a tuple with the InitiativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiativeId

`func (o *PointCreateDataAttributes) SetInitiativeId(v string)`

SetInitiativeId sets InitiativeId field to given value.


### GetParentId

`func (o *PointCreateDataAttributes) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PointCreateDataAttributes) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PointCreateDataAttributes) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PointCreateDataAttributes) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetTitle

`func (o *PointCreateDataAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PointCreateDataAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PointCreateDataAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDesc

`func (o *PointCreateDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *PointCreateDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *PointCreateDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetLocalCost

`func (o *PointCreateDataAttributes) GetLocalCost() int32`

GetLocalCost returns the LocalCost field if non-nil, zero value otherwise.

### GetLocalCostOk

`func (o *PointCreateDataAttributes) GetLocalCostOk() (*int32, bool)`

GetLocalCostOk returns a tuple with the LocalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCost

`func (o *PointCreateDataAttributes) SetLocalCost(v int32)`

SetLocalCost sets LocalCost field to given value.


### GetJar

`func (o *PointCreateDataAttributes) GetJar() JarAttributes`

GetJar returns the Jar field if non-nil, zero value otherwise.

### GetJarOk

`func (o *PointCreateDataAttributes) GetJarOk() (*JarAttributes, bool)`

GetJarOk returns a tuple with the Jar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJar

`func (o *PointCreateDataAttributes) SetJar(v JarAttributes)`

SetJar sets Jar field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


