package storage_test

import (
	"testing"
	"time"

	"github.com/apple-wallet-automation/internal/storage"
)

func newTestStorage(t *testing.T) *storage.Storage {
	t.Helper()
	s, err := storage.New(t.TempDir())
	if err != nil {
		t.Fatalf("storage.New() error: %v", err)
	}
	return s
}

func TestSave_CreatesFileOnFirstTransaction(t *testing.T) {
	s := newTestStorage(t)

	tx := storage.Transaction{
		ID:         "tx-001",
		Amount:     "֏5 000,00",
		Card:       "Visa Classic",
		Merchant:   "KFC",
		Category:   "Restaurant",
		RecordedAt: time.Now().UTC(),
	}

	if err := s.Save("alice", tx); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	store, err := s.Read("alice")
	if err != nil {
		t.Fatalf("Read() error: %v", err)
	}

	if len(store.Transactions) != 1 {
		t.Fatalf("expected 1 transaction, got %d", len(store.Transactions))
	}
	if store.Transactions[0].ID != "tx-001" {
		t.Errorf("expected tx-001, got %s", store.Transactions[0].ID)
	}
}

func TestSave_AppendsMultipleTransactions(t *testing.T) {
	s := newTestStorage(t)

	for i, merchant := range []string{"KFC", "Yerevan City", "Amazon"} {
		tx := storage.Transaction{
			ID:         "tx-" + string(rune('0'+i+1)),
			Amount:     "֏1 000,00",
			Card:       "Visa",
			Merchant:   merchant,
			Category:   "Other",
			RecordedAt: time.Now().UTC(),
		}
		if err := s.Save("alice", tx); err != nil {
			t.Fatalf("Save() error on tx %d: %v", i, err)
		}
	}

	store, err := s.Read("alice")
	if err != nil {
		t.Fatalf("Read() error: %v", err)
	}

	if len(store.Transactions) != 3 {
		t.Errorf("expected 3 transactions, got %d", len(store.Transactions))
	}
}

func TestSave_IsolatesUserData(t *testing.T) {
	s := newTestStorage(t)

	txAlice := storage.Transaction{ID: "alice-tx", Amount: "$10", Card: "Visa", Merchant: "KFC", Category: "Restaurant", RecordedAt: time.Now().UTC()}
	txBob := storage.Transaction{ID: "bob-tx", Amount: "$20", Card: "Mastercard", Merchant: "Uber", Category: "Transport", RecordedAt: time.Now().UTC()}

	if err := s.Save("alice", txAlice); err != nil {
		t.Fatalf("Save alice error: %v", err)
	}
	if err := s.Save("bob", txBob); err != nil {
		t.Fatalf("Save bob error: %v", err)
	}

	aliceStore, _ := s.Read("alice")
	bobStore, _ := s.Read("bob")

	if len(aliceStore.Transactions) != 1 || aliceStore.Transactions[0].ID != "alice-tx" {
		t.Errorf("alice store unexpected: %+v", aliceStore)
	}
	if len(bobStore.Transactions) != 1 || bobStore.Transactions[0].ID != "bob-tx" {
		t.Errorf("bob store unexpected: %+v", bobStore)
	}
}

func TestRead_EmptyForNewUser(t *testing.T) {
	s := newTestStorage(t)

	store, err := s.Read("nonexistent")
	if err != nil {
		t.Fatalf("Read() error: %v", err)
	}
	if len(store.Transactions) != 0 {
		t.Errorf("expected 0 transactions for new user, got %d", len(store.Transactions))
	}
}

func TestSave_PersistsAllFields(t *testing.T) {
	s := newTestStorage(t)
	now := time.Now().UTC().Truncate(time.Second)

	tx := storage.Transaction{
		ID:         "full-tx",
		Amount:     "֏26 307,00",
		Card:       "Visa Classic",
		Merchant:   "Yerevan City Komitas",
		Category:   "Groceries",
		RecordedAt: now,
	}

	if err := s.Save("vhamed32", tx); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	store, err := s.Read("vhamed32")
	if err != nil {
		t.Fatalf("Read() error: %v", err)
	}

	got := store.Transactions[0]
	if got.Amount != tx.Amount {
		t.Errorf("Amount: got %q, want %q", got.Amount, tx.Amount)
	}
	if got.Card != tx.Card {
		t.Errorf("Card: got %q, want %q", got.Card, tx.Card)
	}
	if got.Merchant != tx.Merchant {
		t.Errorf("Merchant: got %q, want %q", got.Merchant, tx.Merchant)
	}
	if got.Category != tx.Category {
		t.Errorf("Category: got %q, want %q", got.Category, tx.Category)
	}
	if !got.RecordedAt.Equal(now) {
		t.Errorf("RecordedAt: got %v, want %v", got.RecordedAt, now)
	}
}
