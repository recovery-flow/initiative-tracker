# Initiative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**InitiativeData**](InitiativeData.md) |  | 
**Included** | Pointer to [**InitiativeIncluded**](InitiativeIncluded.md) |  | [optional] 

## Methods

### NewInitiative

`func NewInitiative(data InitiativeData, ) *Initiative`

NewInitiative instantiates a new Initiative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiativeWithDefaults

`func NewInitiativeWithDefaults() *Initiative`

NewInitiativeWithDefaults instantiates a new Initiative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Initiative) GetData() InitiativeData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Initiative) GetDataOk() (*InitiativeData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Initiative) SetData(v InitiativeData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *Initiative) GetIncluded() InitiativeIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *Initiative) GetIncludedOk() (*InitiativeIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *Initiative) SetIncluded(v InitiativeIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *Initiative) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


