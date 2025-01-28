# InitiativeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | initiative id | 
**Type** | **string** |  | 
**Attributes** | [**InitiativeDataAttributes**](InitiativeDataAttributes.md) |  | 
**Links** | [**LinkSelf**](LinkSelf.md) |  | 
**Relationships** | [**InitiativeDataRelationships**](InitiativeDataRelationships.md) |  | 

## Methods

### NewInitiativeData

`func NewInitiativeData(id string, type_ string, attributes InitiativeDataAttributes, links LinkSelf, relationships InitiativeDataRelationships, ) *InitiativeData`

NewInitiativeData instantiates a new InitiativeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiativeDataWithDefaults

`func NewInitiativeDataWithDefaults() *InitiativeData`

NewInitiativeDataWithDefaults instantiates a new InitiativeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InitiativeData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InitiativeData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InitiativeData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *InitiativeData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InitiativeData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InitiativeData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InitiativeData) GetAttributes() InitiativeDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InitiativeData) GetAttributesOk() (*InitiativeDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InitiativeData) SetAttributes(v InitiativeDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *InitiativeData) GetLinks() LinkSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InitiativeData) GetLinksOk() (*LinkSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InitiativeData) SetLinks(v LinkSelf)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *InitiativeData) GetRelationships() InitiativeDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InitiativeData) GetRelationshipsOk() (*InitiativeDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InitiativeData) SetRelationships(v InitiativeDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


