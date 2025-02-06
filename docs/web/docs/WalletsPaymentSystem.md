# WalletsPaymentSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayPal** | Pointer to **string** | paypal account | [optional] 
**ApplePay** | Pointer to **string** | apple pay account | [optional] 
**GooglePay** | Pointer to **string** | google pay account | [optional] 

## Methods

### NewWalletsPaymentSystem

`func NewWalletsPaymentSystem() *WalletsPaymentSystem`

NewWalletsPaymentSystem instantiates a new WalletsPaymentSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletsPaymentSystemWithDefaults

`func NewWalletsPaymentSystemWithDefaults() *WalletsPaymentSystem`

NewWalletsPaymentSystemWithDefaults instantiates a new WalletsPaymentSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayPal

`func (o *WalletsPaymentSystem) GetPayPal() string`

GetPayPal returns the PayPal field if non-nil, zero value otherwise.

### GetPayPalOk

`func (o *WalletsPaymentSystem) GetPayPalOk() (*string, bool)`

GetPayPalOk returns a tuple with the PayPal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPal

`func (o *WalletsPaymentSystem) SetPayPal(v string)`

SetPayPal sets PayPal field to given value.

### HasPayPal

`func (o *WalletsPaymentSystem) HasPayPal() bool`

HasPayPal returns a boolean if a field has been set.

### GetApplePay

`func (o *WalletsPaymentSystem) GetApplePay() string`

GetApplePay returns the ApplePay field if non-nil, zero value otherwise.

### GetApplePayOk

`func (o *WalletsPaymentSystem) GetApplePayOk() (*string, bool)`

GetApplePayOk returns a tuple with the ApplePay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplePay

`func (o *WalletsPaymentSystem) SetApplePay(v string)`

SetApplePay sets ApplePay field to given value.

### HasApplePay

`func (o *WalletsPaymentSystem) HasApplePay() bool`

HasApplePay returns a boolean if a field has been set.

### GetGooglePay

`func (o *WalletsPaymentSystem) GetGooglePay() string`

GetGooglePay returns the GooglePay field if non-nil, zero value otherwise.

### GetGooglePayOk

`func (o *WalletsPaymentSystem) GetGooglePayOk() (*string, bool)`

GetGooglePayOk returns a tuple with the GooglePay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePay

`func (o *WalletsPaymentSystem) SetGooglePay(v string)`

SetGooglePay sets GooglePay field to given value.

### HasGooglePay

`func (o *WalletsPaymentSystem) HasGooglePay() bool`

HasGooglePay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


