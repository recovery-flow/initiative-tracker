# InitiativeCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of initiative | 
**Desc** | **string** | description of initiative | 
**Goal** | **string** | goal of initiative | 
**Location** | Pointer to **string** | location of initiative | [optional] 
**Type** | **string** | types of initiative | 
**Status** | **string** | status of initiative | 
**FinalCost** | **int64** | final cost of initiative | 
**OrgId** | **string** | organization id | 
**Wallets** | [**Wallets**](Wallets.md) |  | 

## Methods

### NewInitiativeCreateDataAttributes

`func NewInitiativeCreateDataAttributes(name string, desc string, goal string, type_ string, status string, finalCost int64, orgId string, wallets Wallets, ) *InitiativeCreateDataAttributes`

NewInitiativeCreateDataAttributes instantiates a new InitiativeCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiativeCreateDataAttributesWithDefaults

`func NewInitiativeCreateDataAttributesWithDefaults() *InitiativeCreateDataAttributes`

NewInitiativeCreateDataAttributesWithDefaults instantiates a new InitiativeCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InitiativeCreateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InitiativeCreateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InitiativeCreateDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetDesc

`func (o *InitiativeCreateDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *InitiativeCreateDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *InitiativeCreateDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.


### GetGoal

`func (o *InitiativeCreateDataAttributes) GetGoal() string`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *InitiativeCreateDataAttributes) GetGoalOk() (*string, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *InitiativeCreateDataAttributes) SetGoal(v string)`

SetGoal sets Goal field to given value.


### GetLocation

`func (o *InitiativeCreateDataAttributes) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InitiativeCreateDataAttributes) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InitiativeCreateDataAttributes) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InitiativeCreateDataAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *InitiativeCreateDataAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InitiativeCreateDataAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InitiativeCreateDataAttributes) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *InitiativeCreateDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InitiativeCreateDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InitiativeCreateDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFinalCost

`func (o *InitiativeCreateDataAttributes) GetFinalCost() int64`

GetFinalCost returns the FinalCost field if non-nil, zero value otherwise.

### GetFinalCostOk

`func (o *InitiativeCreateDataAttributes) GetFinalCostOk() (*int64, bool)`

GetFinalCostOk returns a tuple with the FinalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCost

`func (o *InitiativeCreateDataAttributes) SetFinalCost(v int64)`

SetFinalCost sets FinalCost field to given value.


### GetOrgId

`func (o *InitiativeCreateDataAttributes) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *InitiativeCreateDataAttributes) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *InitiativeCreateDataAttributes) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetWallets

`func (o *InitiativeCreateDataAttributes) GetWallets() Wallets`

GetWallets returns the Wallets field if non-nil, zero value otherwise.

### GetWalletsOk

`func (o *InitiativeCreateDataAttributes) GetWalletsOk() (*Wallets, bool)`

GetWalletsOk returns a tuple with the Wallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallets

`func (o *InitiativeCreateDataAttributes) SetWallets(v Wallets)`

SetWallets sets Wallets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


