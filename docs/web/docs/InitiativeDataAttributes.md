# InitiativeDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of initiative | 
**Desc** | **string** | description of initiative | 
**Goal** | **string** | goal of initiative | 
**Verified** | **bool** | verified status | 
**Location** | Pointer to **string** | location of initiative | [optional] 
**Status** | **string** | status of initiative | 
**Likes** | **int64** | likes of initiative | 
**Reposts** | **int64** | reposts of initiative | 
**Reports** | **int64** | reports of initiative | 
**CreatedAt** | **time.Time** | Initiative created at | 
**UpdatedAt** | Pointer to **time.Time** | Initiative updated at | [optional] 
**ClosedAt** | Pointer to **time.Time** | Initiative closed at | [optional] 

## Methods

### NewInitiativeDataAttributes

`func NewInitiativeDataAttributes(name string, desc string, goal string, verified bool, status string, likes int64, reposts int64, reports int64, createdAt time.Time, ) *InitiativeDataAttributes`

NewInitiativeDataAttributes instantiates a new InitiativeDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiativeDataAttributesWithDefaults

`func NewInitiativeDataAttributesWithDefaults() *InitiativeDataAttributes`

NewInitiativeDataAttributesWithDefaults instantiates a new InitiativeDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InitiativeDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InitiativeDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InitiativeDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *InitiativeDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *InitiativeDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *InitiativeDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetGoal

`func (o *InitiativeDataAttributes) GetGoal() string`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *InitiativeDataAttributes) GetGoalOk() (*string, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *InitiativeDataAttributes) SetGoal(v string)`

SetGoal sets Goal field to given value.


### GetVerified

`func (o *InitiativeDataAttributes) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *InitiativeDataAttributes) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *InitiativeDataAttributes) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetLocation

`func (o *InitiativeDataAttributes) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InitiativeDataAttributes) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InitiativeDataAttributes) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InitiativeDataAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatus

`func (o *InitiativeDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InitiativeDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InitiativeDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLikes

`func (o *InitiativeDataAttributes) GetLikes() int64`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *InitiativeDataAttributes) GetLikesOk() (*int64, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *InitiativeDataAttributes) SetLikes(v int64)`

SetLikes sets Likes field to given value.


### GetReposts

`func (o *InitiativeDataAttributes) GetReposts() int64`

GetReposts returns the Reposts field if non-nil, zero value otherwise.

### GetRepostsOk

`func (o *InitiativeDataAttributes) GetRepostsOk() (*int64, bool)`

GetRepostsOk returns a tuple with the Reposts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposts

`func (o *InitiativeDataAttributes) SetReposts(v int64)`

SetReposts sets Reposts field to given value.


### GetReports

`func (o *InitiativeDataAttributes) GetReports() int64`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *InitiativeDataAttributes) GetReportsOk() (*int64, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *InitiativeDataAttributes) SetReports(v int64)`

SetReports sets Reports field to given value.


### GetCreatedAt

`func (o *InitiativeDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InitiativeDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InitiativeDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InitiativeDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InitiativeDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InitiativeDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InitiativeDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetClosedAt

`func (o *InitiativeDataAttributes) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *InitiativeDataAttributes) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *InitiativeDataAttributes) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *InitiativeDataAttributes) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


