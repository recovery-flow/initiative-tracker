package repositories

import (
	"context"
	"fmt"

	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (i *initiatives) UpdateWallets(ctx context.Context, walletFields map[string]any) (*models.Initiative, error) {
	// Проверяем, что есть основной фильтр по _id (или другим полям),
	// иначе не знаем, какую инициативу обновлять.
	if i.filters == nil || i.filters["_id"] == nil {
		return nil, fmt.Errorf("initiative filter is missing (need at least '_id')")
	}

	// Допустимые ключи (dot-notation) в кошельках:
	validWalletKeys := map[string]bool{
		// BankAccounts
		"bank_accounts.monobank": true,
		"bank_accounts.privat":   true,

		// PaymentSystems
		"payment_systems.google_pay": true,
		"payment_systems.apple_pay":  true,
		"payment_systems.paypal":     true,

		// CryptoWallets
		"crypto_wallets.usdt": true,
		"crypto_wallets.eth":  true,
		"crypto_wallets.btc":  true,
		"crypto_wallets.ton":  true,
		"crypto_wallets.sol":  true,
	}

	// Собираем поля для $set
	updateFields := bson.M{}

	for key, val := range walletFields {
		if validWalletKeys[key] {
			// Формируем: "wallets.bank_accounts.monobank" -> <val>
			fieldPath := fmt.Sprintf("wallets.%s", key)
			updateFields[fieldPath] = val
		} else {
			logrus.Warnf("UpdateWallets: invalid or unknown wallet key '%s', skipping", key)
		}
	}

	if len(updateFields) == 0 {
		return nil, fmt.Errorf("no valid wallet fields provided to update")
	}

	// Готовим запрос
	update := bson.M{"$set": updateFields}
	filter := i.filters // копируем или используем напрямую

	// Делаем UpdateOne
	res, err := i.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update wallets: %w", err)
	}

	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("no initiative found with the given criteria")
	}

	// Если хотим проверить, что что-то реально изменилось,
	// можем смотреть res.ModifiedCount; но это опционально.

	// Возвращаем обновлённый документ
	var updated models.Initiative
	err = i.collection.FindOne(ctx, filter).Decode(&updated)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated initiative: %w", err)
	}

	return &updated, nil
}
