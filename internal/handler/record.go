package handler

import (
	"time"

	"github.com/apple-wallet-automation/internal/categorizer"
	"github.com/apple-wallet-automation/internal/config"
	"github.com/apple-wallet-automation/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RecordRequest struct {
	Amount   string `json:"amount"`
	Card     string `json:"card"`
	Merchant string `json:"merchant"`
}

type RecordHandler struct {
	cfg         *config.Config
	store       *storage.Storage
	categorizer *categorizer.Categorizer
}

func NewRecordHandler(cfg *config.Config, store *storage.Storage, cat *categorizer.Categorizer) *RecordHandler {
	return &RecordHandler{cfg: cfg, store: store, categorizer: cat}
}

func (h *RecordHandler) Handle(c *fiber.Ctx) error {
	apiKey := c.Get("X-Api-Key")
	if apiKey == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing api key",
		})
	}

	user, ok := h.cfg.FindUserByAPIKey(apiKey)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid api key",
		})
	}

	var req RecordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.Amount == "" || req.Merchant == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "amount and merchant are required",
		})
	}

	category := h.categorizer.Categorize(req.Merchant)

	tx := storage.Transaction{
		ID:         uuid.New().String(),
		Amount:     req.Amount,
		Card:       req.Card,
		Merchant:   req.Merchant,
		Category:   category,
		RecordedAt: time.Now().UTC(),
	}

	if err := h.store.Save(user.Username, tx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save transaction",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":         tx.ID,
		"category":   tx.Category,
		"recorded_at": tx.RecordedAt,
	})
}
