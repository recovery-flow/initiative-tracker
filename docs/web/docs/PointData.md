# PointData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | point id | 
**Type** | **string** |  | 
**Attributes** | [**PointAttributes**](PointAttributes.md) |  | 
**Relationships** | [**PointDataRelationships**](PointDataRelationships.md) |  | 

## Methods

### NewPointData

`func NewPointData(id string, type_ string, attributes PointAttributes, relationships PointDataRelationships, ) *PointData`

NewPointData instantiates a new PointData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointDataWithDefaults

`func NewPointDataWithDefaults() *PointData`

NewPointDataWithDefaults instantiates a new PointData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PointData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PointData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PointData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PointData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PointData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PointData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PointData) GetAttributes() PointAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PointData) GetAttributesOk() (*PointAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PointData) SetAttributes(v PointAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PointData) GetRelationships() PointDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PointData) GetRelationshipsOk() (*PointDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PointData) SetRelationships(v PointDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


