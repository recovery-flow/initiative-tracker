package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/initiative-tracker/internal/service/requests"
	"github.com/recovery-flow/initiative-tracker/internal/service/responses"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitiativeWalletsUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewWalletsUpdate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	iniId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "initiative_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse initiative id")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"initiative_id": validation.NewError("initiative_id", "Invalid initiative id"),
		})...)
		return
	}

	initiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	stmt := make(map[string]any)
	if req.Data.Attributes.BankAccount != nil {
		if req.Data.Attributes.BankAccount.Monobank != nil {
			stmt["bank_accounts.monobank"] = req.Data.Attributes.BankAccount.Monobank
		}
		if req.Data.Attributes.BankAccount.Privat != nil {
			stmt["bank_accounts.privat"] = req.Data.Attributes.BankAccount.Privat
		}
	}

	if req.Data.Attributes.PaymentSystem != nil {
		if req.Data.Attributes.PaymentSystem.GooglePay != nil {
			stmt["payment_systems.google_pay"] = req.Data.Attributes.PaymentSystem.GooglePay
		}
		if req.Data.Attributes.PaymentSystem.ApplePay != nil {
			stmt["payment_systems.apple_pay"] = req.Data.Attributes.PaymentSystem.ApplePay
		}
		if req.Data.Attributes.PaymentSystem.PayPal != nil {
			stmt["payment_systems.paypal"] = req.Data.Attributes.PaymentSystem.PayPal
		}
	}

	log.Infof("321312")
	if req.Data.Attributes.CryptoWallets != nil {
		if req.Data.Attributes.CryptoWallets.USDT != nil {
			stmt["crypto_wallets.usdt"] = req.Data.Attributes.CryptoWallets.USDT
		}
		if req.Data.Attributes.CryptoWallets.ETH != nil {
			stmt["crypto_wallets.eth"] = req.Data.Attributes.CryptoWallets.ETH
		}
		if req.Data.Attributes.CryptoWallets.BTC != nil {
			stmt["crypto_wallets.btc"] = req.Data.Attributes.CryptoWallets.BTC
		}
		if req.Data.Attributes.CryptoWallets.TON != nil {
			stmt["crypto_wallets.ton"] = req.Data.Attributes.CryptoWallets.TON
		}
		if req.Data.Attributes.CryptoWallets.SOL != nil {
			stmt["crypto_wallets.sol"] = req.Data.Attributes.CryptoWallets.SOL
		}
	}
	
	res, err := server.MongoDB.Initiative.New().FilterExact(map[string]any{
		"_id": iniId,
	}).UpdateWallets(r.Context(), stmt)
	if err != nil {
		log.WithError(err).Error("Failed to update wallets")
		httpkit.RenderErr(w, problems.InternalError("Failed to update wallets"))
		return
	}

	log.Infof("Wallets for initiative %s updated by %s", iniId, initiatorId)
	httpkit.Render(w, responses.Initiative(*res))
}
