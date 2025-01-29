# PointDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | Pointer to [**LinksDirect**](LinksDirect.md) |  | [optional] 
**PublishedBy** | [**LinksDirect**](LinksDirect.md) |  | 
**Initiative** | [**LinksDirect**](LinksDirect.md) |  | 
**Plan** | [**LinksDirect**](LinksDirect.md) |  | 

## Methods

### NewPointDataRelationships

`func NewPointDataRelationships(publishedBy LinksDirect, initiative LinksDirect, plan LinksDirect, ) *PointDataRelationships`

NewPointDataRelationships instantiates a new PointDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointDataRelationshipsWithDefaults

`func NewPointDataRelationshipsWithDefaults() *PointDataRelationships`

NewPointDataRelationshipsWithDefaults instantiates a new PointDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *PointDataRelationships) GetParent() LinksDirect`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PointDataRelationships) GetParentOk() (*LinksDirect, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PointDataRelationships) SetParent(v LinksDirect)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PointDataRelationships) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPublishedBy

`func (o *PointDataRelationships) GetPublishedBy() LinksDirect`

GetPublishedBy returns the PublishedBy field if non-nil, zero value otherwise.

### GetPublishedByOk

`func (o *PointDataRelationships) GetPublishedByOk() (*LinksDirect, bool)`

GetPublishedByOk returns a tuple with the PublishedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedBy

`func (o *PointDataRelationships) SetPublishedBy(v LinksDirect)`

SetPublishedBy sets PublishedBy field to given value.


### GetInitiative

`func (o *PointDataRelationships) GetInitiative() LinksDirect`

GetInitiative returns the Initiative field if non-nil, zero value otherwise.

### GetInitiativeOk

`func (o *PointDataRelationships) GetInitiativeOk() (*LinksDirect, bool)`

GetInitiativeOk returns a tuple with the Initiative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiative

`func (o *PointDataRelationships) SetInitiative(v LinksDirect)`

SetInitiative sets Initiative field to given value.


### GetPlan

`func (o *PointDataRelationships) GetPlan() LinksDirect`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PointDataRelationships) GetPlanOk() (*LinksDirect, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PointDataRelationships) SetPlan(v LinksDirect)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


