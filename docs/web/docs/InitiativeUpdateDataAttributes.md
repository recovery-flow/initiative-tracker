# InitiativeUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name of initiative | [optional] 
**Desc** | Pointer to **string** | description of initiative | [optional] 
**Goal** | Pointer to **string** | goal of initiative | [optional] 
**Location** | Pointer to **string** | location of initiative | [optional] 
**Type** | Pointer to **string** | types of initiative | [optional] 
**Status** | Pointer to **string** | status of initiative | [optional] 
**FinalCost** | Pointer to **int64** | final cost of initiative | [optional] 
**Wallets** | Pointer to [**Wallets**](Wallets.md) |  | [optional] 
**OrgMembers** | Pointer to [**[]AddOrgMember**](AddOrgMember.md) |  | [optional] 

## Methods

### NewInitiativeUpdateDataAttributes

`func NewInitiativeUpdateDataAttributes() *InitiativeUpdateDataAttributes`

NewInitiativeUpdateDataAttributes instantiates a new InitiativeUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiativeUpdateDataAttributesWithDefaults

`func NewInitiativeUpdateDataAttributesWithDefaults() *InitiativeUpdateDataAttributes`

NewInitiativeUpdateDataAttributesWithDefaults instantiates a new InitiativeUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InitiativeUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InitiativeUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InitiativeUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InitiativeUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDesc

`func (o *InitiativeUpdateDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *InitiativeUpdateDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *InitiativeUpdateDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *InitiativeUpdateDataAttributes) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetGoal

`func (o *InitiativeUpdateDataAttributes) GetGoal() string`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *InitiativeUpdateDataAttributes) GetGoalOk() (*string, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *InitiativeUpdateDataAttributes) SetGoal(v string)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *InitiativeUpdateDataAttributes) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetLocation

`func (o *InitiativeUpdateDataAttributes) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InitiativeUpdateDataAttributes) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InitiativeUpdateDataAttributes) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InitiativeUpdateDataAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *InitiativeUpdateDataAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InitiativeUpdateDataAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InitiativeUpdateDataAttributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InitiativeUpdateDataAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *InitiativeUpdateDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InitiativeUpdateDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InitiativeUpdateDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InitiativeUpdateDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFinalCost

`func (o *InitiativeUpdateDataAttributes) GetFinalCost() int64`

GetFinalCost returns the FinalCost field if non-nil, zero value otherwise.

### GetFinalCostOk

`func (o *InitiativeUpdateDataAttributes) GetFinalCostOk() (*int64, bool)`

GetFinalCostOk returns a tuple with the FinalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCost

`func (o *InitiativeUpdateDataAttributes) SetFinalCost(v int64)`

SetFinalCost sets FinalCost field to given value.

### HasFinalCost

`func (o *InitiativeUpdateDataAttributes) HasFinalCost() bool`

HasFinalCost returns a boolean if a field has been set.

### GetWallets

`func (o *InitiativeUpdateDataAttributes) GetWallets() Wallets`

GetWallets returns the Wallets field if non-nil, zero value otherwise.

### GetWalletsOk

`func (o *InitiativeUpdateDataAttributes) GetWalletsOk() (*Wallets, bool)`

GetWalletsOk returns a tuple with the Wallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallets

`func (o *InitiativeUpdateDataAttributes) SetWallets(v Wallets)`

SetWallets sets Wallets field to given value.

### HasWallets

`func (o *InitiativeUpdateDataAttributes) HasWallets() bool`

HasWallets returns a boolean if a field has been set.

### GetOrgMembers

`func (o *InitiativeUpdateDataAttributes) GetOrgMembers() []AddOrgMember`

GetOrgMembers returns the OrgMembers field if non-nil, zero value otherwise.

### GetOrgMembersOk

`func (o *InitiativeUpdateDataAttributes) GetOrgMembersOk() (*[]AddOrgMember, bool)`

GetOrgMembersOk returns a tuple with the OrgMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMembers

`func (o *InitiativeUpdateDataAttributes) SetOrgMembers(v []AddOrgMember)`

SetOrgMembers sets OrgMembers field to given value.

### HasOrgMembers

`func (o *InitiativeUpdateDataAttributes) HasOrgMembers() bool`

HasOrgMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


