# ParticipantDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | user id | 
**FirstName** | **string** | first name of participant | 
**SecondName** | **string** | second name of participant | 
**ThirdName** | Pointer to **string** | third name of participant | [optional] 
**DisplayName** | **string** | name of participant | 
**Position** | **string** | position in the company | 
**Verified** | **bool** | verified status | 
**Desc** | Pointer to **string** | description of participant | [optional] 
**Role** | **string** | User role | 
**UpdatedAt** | Pointer to **time.Time** | User updated at | [optional] 
**CreatedAt** | **time.Time** | User created at | 

## Methods

### NewParticipantDataAttributes

`func NewParticipantDataAttributes(userId string, firstName string, secondName string, displayName string, position string, verified bool, role string, createdAt time.Time, ) *ParticipantDataAttributes`

NewParticipantDataAttributes instantiates a new ParticipantDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantDataAttributesWithDefaults

`func NewParticipantDataAttributesWithDefaults() *ParticipantDataAttributes`

NewParticipantDataAttributesWithDefaults instantiates a new ParticipantDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ParticipantDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ParticipantDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ParticipantDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetFirstName

`func (o *ParticipantDataAttributes) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ParticipantDataAttributes) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ParticipantDataAttributes) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetSecondName

`func (o *ParticipantDataAttributes) GetSecondName() string`

GetSecondName returns the SecondName field if non-nil, zero value otherwise.

### GetSecondNameOk

`func (o *ParticipantDataAttributes) GetSecondNameOk() (*string, bool)`

GetSecondNameOk returns a tuple with the SecondName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondName

`func (o *ParticipantDataAttributes) SetSecondName(v string)`

SetSecondName sets SecondName field to given value.


### GetThirdName

`func (o *ParticipantDataAttributes) GetThirdName() string`

GetThirdName returns the ThirdName field if non-nil, zero value otherwise.

### GetThirdNameOk

`func (o *ParticipantDataAttributes) GetThirdNameOk() (*string, bool)`

GetThirdNameOk returns a tuple with the ThirdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdName

`func (o *ParticipantDataAttributes) SetThirdName(v string)`

SetThirdName sets ThirdName field to given value.

### HasThirdName

`func (o *ParticipantDataAttributes) HasThirdName() bool`

HasThirdName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ParticipantDataAttributes) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ParticipantDataAttributes) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ParticipantDataAttributes) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPosition

`func (o *ParticipantDataAttributes) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ParticipantDataAttributes) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ParticipantDataAttributes) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetVerified

`func (o *ParticipantDataAttributes) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ParticipantDataAttributes) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ParticipantDataAttributes) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetDesc

`func (o *ParticipantDataAttributes) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ParticipantDataAttributes) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ParticipantDataAttributes) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *ParticipantDataAttributes) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetRole

`func (o *ParticipantDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ParticipantDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ParticipantDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.


### GetUpdatedAt

`func (o *ParticipantDataAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ParticipantDataAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ParticipantDataAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ParticipantDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ParticipantDataAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ParticipantDataAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ParticipantDataAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


