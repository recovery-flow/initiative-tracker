# Point

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**PointData**](PointData.md) |  | 
**Included** | Pointer to [**[]JarData**](JarData.md) |  | [optional] 

## Methods

### NewPoint

`func NewPoint(data PointData, ) *Point`

NewPoint instantiates a new Point object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointWithDefaults

`func NewPointWithDefaults() *Point`

NewPointWithDefaults instantiates a new Point object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Point) GetData() PointData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Point) GetDataOk() (*PointData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Point) SetData(v PointData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *Point) GetIncluded() []JarData`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *Point) GetIncludedOk() (*[]JarData, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *Point) SetIncluded(v []JarData)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *Point) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


