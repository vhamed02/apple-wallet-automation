package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/apple-wallet-automation/internal/categorizer"
	"github.com/apple-wallet-automation/internal/config"
	"github.com/apple-wallet-automation/internal/handler"
	"github.com/apple-wallet-automation/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type recordResponse struct {
	ID          string    `json:"id"`
	Category    string    `json:"category"`
	RecordedAt  time.Time `json:"recorded_at"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func newTestApp(t *testing.T) (*fiber.App, *storage.Storage) {
	t.Helper()

	cfg := &config.Config{
		Users: []config.User{
			{Username: "vhamed32", APIKey: "valid-key-001"},
			{Username: "alice", APIKey: "valid-key-002"},
		},
		Categories: map[string][]string{
			"Groceries":  {"yerevan city", "carrefour", "supermarket"},
			"Restaurant": {"kfc", "mcdonald", "starbucks", "cafe", "restaurant"},
			"Transport":  {"uber", "bolt", "taxi", "parking"},
			"Shopping":   {"amazon", "zara", "nike"},
			"Health":     {"pharmacy", "hospital"},
		},
	}

	store, err := storage.New(t.TempDir())
	if err != nil {
		t.Fatalf("storage.New() error: %v", err)
	}

	cat := categorizer.New(cfg.Categories)
	h := handler.NewRecordHandler(cfg, store, cat)

	app := fiber.New()
	app.Post("/record/", h.Handle)
	app.Post("/record", h.Handle)

	return app, store
}

func doRequest(t *testing.T, app *fiber.App, method, path, apiKey string, body any) *http.Response {
	t.Helper()

	b, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("json.Marshal error: %v", err)
	}

	req := httptest.NewRequest(method, path, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Set("X-Api-Key", apiKey)
	}

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test() error: %v", err)
	}
	return resp
}

func TestFeature_SuccessfulTransaction_Groceries(t *testing.T) {
	app, store := newTestApp(t)

	body := map[string]string{
		"amount":   "֏26 307,00",
		"card":     "Visa Classic",
		"merchant": "Yerevan City Komitas",
	}

	resp := doRequest(t, app, http.MethodPost, "/record/", "valid-key-001", body)

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	var result recordResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("decode error: %v", err)
	}

	if result.ID == "" {
		t.Error("expected non-empty id")
	}
	if result.Category != "Groceries" {
		t.Errorf("expected Groceries, got %s", result.Category)
	}
	if result.RecordedAt.IsZero() {
		t.Error("expected non-zero recorded_at")
	}

	saved, err := store.Read("vhamed32")
	if err != nil {
		t.Fatalf("Read() error: %v", err)
	}
	if len(saved.Transactions) != 1 {
		t.Fatalf("expected 1 saved transaction, got %d", len(saved.Transactions))
	}
	tx := saved.Transactions[0]
	if tx.Category != "Groceries" {
		t.Errorf("saved category: got %s, want Groceries", tx.Category)
	}
	if tx.Amount != "֏26 307,00" {
		t.Errorf("saved amount: got %s, want ֏26 307,00", tx.Amount)
	}
	if tx.Card != "Visa Classic" {
		t.Errorf("saved card: got %s, want Visa Classic", tx.Card)
	}
	if tx.Merchant != "Yerevan City Komitas" {
		t.Errorf("saved merchant: got %s, want Yerevan City Komitas", tx.Merchant)
	}
}

func TestFeature_SuccessfulTransaction_Restaurant(t *testing.T) {
	app, store := newTestApp(t)

	body := map[string]string{
		"amount":   "$12.50",
		"card":     "Mastercard",
		"merchant": "KFC Downtown",
	}

	resp := doRequest(t, app, http.MethodPost, "/record/", "valid-key-001", body)

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	var result recordResponse
	json.NewDecoder(resp.Body).Decode(&result)

	if result.Category != "Restaurant" {
		t.Errorf("expected Restaurant, got %s", result.Category)
	}

	saved, _ := store.Read("vhamed32")
	if saved.Transactions[0].Category != "Restaurant" {
		t.Errorf("saved category should be Restaurant")
	}
}

func TestFeature_SuccessfulTransaction_FallsBackToOther(t *testing.T) {
	app, store := newTestApp(t)

	body := map[string]string{
		"amount":   "$5.00",
		"card":     "Visa",
		"merchant": "Random Unknown Vendor",
	}

	resp := doRequest(t, app, http.MethodPost, "/record/", "valid-key-001", body)

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	var result recordResponse
	json.NewDecoder(resp.Body).Decode(&result)

	if result.Category != "Other" {
		t.Errorf("expected Other, got %s", result.Category)
	}

	saved, _ := store.Read("vhamed32")
	if saved.Transactions[0].Category != "Other" {
		t.Error("saved category should be Other")
	}
}

func TestFeature_MissingAPIKey_Returns401_DoesNotRecord(t *testing.T) {
	app, store := newTestApp(t)

	body := map[string]string{
		"amount":   "$10.00",
		"card":     "Visa",
		"merchant": "KFC",
	}

	resp := doRequest(t, app, http.MethodPost, "/record/", "", body)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", resp.StatusCode)
	}

	var result errorResponse
	json.NewDecoder(resp.Body).Decode(&result)
	if result.Error != "missing api key" {
		t.Errorf("unexpected error message: %s", result.Error)
	}

	saved, _ := store.Read("vhamed32")
	if len(saved.Transactions) != 0 {
		t.Errorf("expected 0 saved transactions, got %d", len(saved.Transactions))
	}
}

func TestFeature_InvalidAPIKey_Returns401_DoesNotRecord(t *testing.T) {
	app, store := newTestApp(t)

	body := map[string]string{
		"amount":   "$10.00",
		"card":     "Visa",
		"merchant": "KFC",
	}

	resp := doRequest(t, app, http.MethodPost, "/record/", "totally-wrong-key", body)

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", resp.StatusCode)
	}

	var result errorResponse
	json.NewDecoder(resp.Body).Decode(&result)
	if result.Error != "invalid api key" {
		t.Errorf("unexpected error message: %s", result.Error)
	}

	saved, _ := store.Read("vhamed32")
	if len(saved.Transactions) != 0 {
		t.Errorf("expected 0 saved transactions, got %d", len(saved.Transactions))
	}
}

func TestFeature_MissingAmount_Returns400_DoesNotRecord(t *testing.T) {
	app, store := newTestApp(t)

	body := map[string]string{
		"card":     "Visa",
		"merchant": "KFC",
	}

	resp := doRequest(t, app, http.MethodPost, "/record/", "valid-key-001", body)

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}

	saved, _ := store.Read("vhamed32")
	if len(saved.Transactions) != 0 {
		t.Errorf("expected 0 saved transactions, got %d", len(saved.Transactions))
	}
}

func TestFeature_MissingMerchant_Returns400_DoesNotRecord(t *testing.T) {
	app, store := newTestApp(t)

	body := map[string]string{
		"amount": "$10.00",
		"card":   "Visa",
	}

	resp := doRequest(t, app, http.MethodPost, "/record/", "valid-key-001", body)

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}

	saved, _ := store.Read("vhamed32")
	if len(saved.Transactions) != 0 {
		t.Errorf("expected 0 saved transactions, got %d", len(saved.Transactions))
	}
}

func TestFeature_EmptyBody_Returns400(t *testing.T) {
	app, _ := newTestApp(t)

	req := httptest.NewRequest(http.MethodPost, "/record/", bytes.NewReader([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", "valid-key-001")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test() error: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestFeature_MultipleTransactions_AllSaved(t *testing.T) {
	app, store := newTestApp(t)

	transactions := []map[string]string{
		{"amount": "֏26 307,00", "card": "Visa Classic", "merchant": "Yerevan City Komitas"},
		{"amount": "$12.50", "card": "Mastercard", "merchant": "KFC"},
		{"amount": "$5.00", "card": "Visa", "merchant": "Random Unknown"},
	}

	for _, tx := range transactions {
		resp := doRequest(t, app, http.MethodPost, "/record/", "valid-key-001", tx)
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected 201 for merchant %q, got %d", tx["merchant"], resp.StatusCode)
		}
	}

	saved, err := store.Read("vhamed32")
	if err != nil {
		t.Fatalf("Read() error: %v", err)
	}

	if len(saved.Transactions) != 3 {
		t.Fatalf("expected 3 saved transactions, got %d", len(saved.Transactions))
	}

	expectedCategories := []string{"Groceries", "Restaurant", "Other"}
	for i, tx := range saved.Transactions {
		if tx.Category != expectedCategories[i] {
			t.Errorf("tx[%d] category: got %s, want %s", i, tx.Category, expectedCategories[i])
		}
	}
}

func TestFeature_DifferentUsers_DataIsolated(t *testing.T) {
	app, store := newTestApp(t)

	doRequest(t, app, http.MethodPost, "/record/", "valid-key-001", map[string]string{
		"amount": "$10.00", "card": "Visa", "merchant": "KFC",
	})
	doRequest(t, app, http.MethodPost, "/record/", "valid-key-002", map[string]string{
		"amount": "$20.00", "card": "Mastercard", "merchant": "Amazon",
	})

	vhamedStore, _ := store.Read("vhamed32")
	aliceStore, _ := store.Read("alice")

	if len(vhamedStore.Transactions) != 1 {
		t.Errorf("vhamed32: expected 1 transaction, got %d", len(vhamedStore.Transactions))
	}
	if len(aliceStore.Transactions) != 1 {
		t.Errorf("alice: expected 1 transaction, got %d", len(aliceStore.Transactions))
	}
	if vhamedStore.Transactions[0].Merchant != "KFC" {
		t.Errorf("vhamed32 merchant: got %s, want KFC", vhamedStore.Transactions[0].Merchant)
	}
	if aliceStore.Transactions[0].Merchant != "Amazon" {
		t.Errorf("alice merchant: got %s, want Amazon", aliceStore.Transactions[0].Merchant)
	}
}

func TestFeature_ResponseContainsUniqueIDs(t *testing.T) {
	app, _ := newTestApp(t)

	ids := make(map[string]bool)
	for i := 0; i < 5; i++ {
		resp := doRequest(t, app, http.MethodPost, "/record/", "valid-key-001", map[string]string{
			"amount": "$1.00", "card": "Visa", "merchant": "KFC",
		})
		var result recordResponse
		json.NewDecoder(resp.Body).Decode(&result)
		if ids[result.ID] {
			t.Errorf("duplicate transaction id: %s", result.ID)
		}
		ids[result.ID] = true
	}
}

func TestFeature_TrailingSlashAndWithout_BothWork(t *testing.T) {
	app, _ := newTestApp(t)

	body := map[string]string{
		"amount": "$5.00", "card": "Visa", "merchant": "Starbucks",
	}

	for _, path := range []string{"/record/", "/record"} {
		resp := doRequest(t, app, http.MethodPost, path, "valid-key-001", body)
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("path %s: expected 201, got %d", path, resp.StatusCode)
		}
	}
}

func TestFeature_TransactionIDInResponseMatchesSavedID(t *testing.T) {
	app, store := newTestApp(t)

	body := map[string]string{
		"amount": "֏5 000,00", "card": "Visa", "merchant": "Carrefour",
	}

	resp := doRequest(t, app, http.MethodPost, "/record/", "valid-key-001", body)

	var result recordResponse
	json.NewDecoder(resp.Body).Decode(&result)

	saved, _ := store.Read("vhamed32")
	if saved.Transactions[0].ID != result.ID {
		t.Errorf("ID mismatch: response has %s, saved has %s", result.ID, saved.Transactions[0].ID)
	}
}
