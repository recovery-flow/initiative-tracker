# OrgMemberData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**OrgMemberDataAttributes**](OrgMemberDataAttributes.md) |  | 

## Methods

### NewOrgMemberData

`func NewOrgMemberData(type_ string, attributes OrgMemberDataAttributes, ) *OrgMemberData`

NewOrgMemberData instantiates a new OrgMemberData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgMemberDataWithDefaults

`func NewOrgMemberDataWithDefaults() *OrgMemberData`

NewOrgMemberDataWithDefaults instantiates a new OrgMemberData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrgMemberData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgMemberData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgMemberData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OrgMemberData) GetAttributes() OrgMemberDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrgMemberData) GetAttributesOk() (*OrgMemberDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrgMemberData) SetAttributes(v OrgMemberDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


