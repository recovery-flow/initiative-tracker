package handlers

import (
	"net/http"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/recovery-flow/initiative-tracker/internal/service/requests"
	"github.com/recovery-flow/initiative-tracker/internal/service/responses"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitiativeCreate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewInitiativeCreate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	iniStatus := models.GetIniStatus(req.Data.Attributes.Status)
	if iniStatus == nil {
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"status": validation.NewError("status", "invalid status"),
		})...)
		return
	}

	iniType := models.GetIniType(req.Data.Attributes.Type)
	if iniType == nil {
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"type": validation.NewError("type", "invalid type"),
		})...)
		return
	}

	orgId, err := primitive.ObjectIDFromHex(req.Data.Attributes.OrgId)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"org_id": validation.NewError("org_id", "invalid organization id"),
		})...)
		return
	}

	initiative := models.Initiative{
		ID:       primitive.NewObjectID(),
		Name:     req.Data.Attributes.Name,
		Desc:     req.Data.Attributes.Desc,
		Goal:     req.Data.Attributes.Goal,
		Verified: false,
		Location: req.Data.Attributes.Location,
		Type:     *iniType,
		Status:   *iniStatus,

		ChatID: primitive.NilObjectID,
		Organizations: []models.OrgMember{
			{
				ID:     orgId,
				Status: models.StatusOrgFounder,
				Since:  primitive.NewDateTimeFromTime(time.Now().UTC()),
			},
		},

		FinalCost:    int(req.Data.Attributes.FinalCost),
		CollectedSum: 0,

		Wallets: models.Wallets{
			BankAccounts: &models.BankAccounts{
				Monobank: req.Data.Attributes.Wallets.BankAccount.Monobank,
				Privat:   req.Data.Attributes.Wallets.BankAccount.Privat,
			},
			PaymentSystems: &models.PaymentSystems{
				GooglePay: req.Data.Attributes.Wallets.PaymentSystem.GooglePay,
				ApplePay:  req.Data.Attributes.Wallets.PaymentSystem.ApplePay,
				PayPal:    req.Data.Attributes.Wallets.PaymentSystem.PayPal,
			},
			CryptoWallets: &models.CryptoWallets{
				USDT: req.Data.Attributes.Wallets.CryptoWallets.USDT,
				ETH:  req.Data.Attributes.Wallets.CryptoWallets.ETH,
				BTC:  req.Data.Attributes.Wallets.CryptoWallets.BTC,
				TON:  req.Data.Attributes.Wallets.CryptoWallets.TON,
				SOL:  req.Data.Attributes.Wallets.CryptoWallets.SOL,
			},
		},

		Likes:   0,
		Reposts: 0,
		Reports: 0,
	}

	res, err := server.MongoDB.Initiative.New().Insert(r.Context(), initiative)
	if err != nil {
		if strings.HasPrefix(err.Error(), "failed to insert initiative") {
			httpkit.RenderErr(w, problems.InternalError("Failed to insert initiative"))
			return
		}
		log.Infof("Failed to create initiative: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	log.Infof("Initiative created: %v by user %s", res.ID, userId)
	httpkit.Render(w, responses.Initiative(*res))
}
