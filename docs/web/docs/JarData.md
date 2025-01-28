# JarData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | point id | 
**Type** | **string** |  | 
**Attributes** | [**JarAttributes**](JarAttributes.md) |  | 
**Relationships** | [**JarDataRelationships**](JarDataRelationships.md) |  | 

## Methods

### NewJarData

`func NewJarData(id string, type_ string, attributes JarAttributes, relationships JarDataRelationships, ) *JarData`

NewJarData instantiates a new JarData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJarDataWithDefaults

`func NewJarDataWithDefaults() *JarData`

NewJarDataWithDefaults instantiates a new JarData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JarData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JarData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JarData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *JarData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JarData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JarData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *JarData) GetAttributes() JarAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *JarData) GetAttributesOk() (*JarAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *JarData) SetAttributes(v JarAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *JarData) GetRelationships() JarDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *JarData) GetRelationshipsOk() (*JarDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *JarData) SetRelationships(v JarDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


