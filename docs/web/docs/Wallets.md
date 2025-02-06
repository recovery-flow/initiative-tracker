# Wallets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccount** | Pointer to [**WalletsBankAccount**](WalletsBankAccount.md) |  | [optional] 
**PaymentSystem** | Pointer to [**WalletsPaymentSystem**](WalletsPaymentSystem.md) |  | [optional] 
**CryptoWallets** | Pointer to [**WalletsCryptoWallets**](WalletsCryptoWallets.md) |  | [optional] 

## Methods

### NewWallets

`func NewWallets() *Wallets`

NewWallets instantiates a new Wallets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletsWithDefaults

`func NewWalletsWithDefaults() *Wallets`

NewWalletsWithDefaults instantiates a new Wallets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccount

`func (o *Wallets) GetBankAccount() WalletsBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *Wallets) GetBankAccountOk() (*WalletsBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *Wallets) SetBankAccount(v WalletsBankAccount)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *Wallets) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetPaymentSystem

`func (o *Wallets) GetPaymentSystem() WalletsPaymentSystem`

GetPaymentSystem returns the PaymentSystem field if non-nil, zero value otherwise.

### GetPaymentSystemOk

`func (o *Wallets) GetPaymentSystemOk() (*WalletsPaymentSystem, bool)`

GetPaymentSystemOk returns a tuple with the PaymentSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSystem

`func (o *Wallets) SetPaymentSystem(v WalletsPaymentSystem)`

SetPaymentSystem sets PaymentSystem field to given value.

### HasPaymentSystem

`func (o *Wallets) HasPaymentSystem() bool`

HasPaymentSystem returns a boolean if a field has been set.

### GetCryptoWallets

`func (o *Wallets) GetCryptoWallets() WalletsCryptoWallets`

GetCryptoWallets returns the CryptoWallets field if non-nil, zero value otherwise.

### GetCryptoWalletsOk

`func (o *Wallets) GetCryptoWalletsOk() (*WalletsCryptoWallets, bool)`

GetCryptoWalletsOk returns a tuple with the CryptoWallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoWallets

`func (o *Wallets) SetCryptoWallets(v WalletsCryptoWallets)`

SetCryptoWallets sets CryptoWallets field to given value.

### HasCryptoWallets

`func (o *Wallets) HasCryptoWallets() bool`

HasCryptoWallets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


