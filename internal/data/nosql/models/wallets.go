package models

type Wallets struct {
	BankAccounts   *BankAccounts   `bson:"bank_accounts,omitempty" json:"bank_accounts,omitempty"`
	PaymentSystems *PaymentSystems `bson:"payment_systems,omitempty" json:"payment_systems,omitempty"`
	CryptoWallets  *CryptoWallets  `bson:"crypto_wallets,omitempty" json:"crypto_wallets,omitempty"`
}

type BankAccounts struct {
	Monobank *string `bson:"monobank,omitempty" json:"monobank,omitempty"`
	Privat   *string `bson:"privat,omitempty"   json:"privat,omitempty"`
}

type PaymentSystems struct {
	GooglePay *string `bson:"google_pay,omitempty" json:"google_pay,omitempty"`
	ApplePay  *string `bson:"apple_pay,omitempty"  json:"apple_pay,omitempty"`
	PayPal    *string `bson:"paypal,omitempty"     json:"paypal,omitempty"`
}

type CryptoWallets struct {
	USDT *string `bson:"usdt,omitempty" json:"usdt,omitempty"`
	ETH  *string `bson:"eth,omitempty"  json:"eth,omitempty"`
	BTC  *string `bson:"btc,omitempty"  json:"btc,omitempty"`
	TON  *string `bson:"ton,omitempty"  json:"ton,omitempty"`
	SOL  *string `bson:"sol,omitempty"  json:"sol,omitempty"`
}
