# PointCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PointData**](PointData.md) |  | 
**Links** | [**LinksPagination**](LinksPagination.md) |  | 

## Methods

### NewPointCollection

`func NewPointCollection(data []PointData, links LinksPagination, ) *PointCollection`

NewPointCollection instantiates a new PointCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointCollectionWithDefaults

`func NewPointCollectionWithDefaults() *PointCollection`

NewPointCollectionWithDefaults instantiates a new PointCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PointCollection) GetData() []PointData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PointCollection) GetDataOk() (*[]PointData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PointCollection) SetData(v []PointData)`

SetData sets Data field to given value.


### GetLinks

`func (o *PointCollection) GetLinks() LinksPagination`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PointCollection) GetLinksOk() (*LinksPagination, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PointCollection) SetLinks(v LinksPagination)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


