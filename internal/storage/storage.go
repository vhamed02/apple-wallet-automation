package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Transaction struct {
	ID         string    `json:"id"`
	Amount     string    `json:"amount"`
	Card       string    `json:"card"`
	Merchant   string    `json:"merchant"`
	Category   string    `json:"category"`
	RecordedAt time.Time `json:"recorded_at"`
}

type UserStore struct {
	Transactions []Transaction `json:"transactions"`
}

type Storage struct {
	dataDir string
	mu      sync.Mutex
}

func New(dataDir string) (*Storage, error) {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}
	return &Storage{dataDir: dataDir}, nil
}

func (s *Storage) Save(username string, tx Transaction) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := s.load(username)
	if err != nil {
		return err
	}

	store.Transactions = append(store.Transactions, tx)

	return s.write(username, store)
}

func (s *Storage) load(username string) (UserStore, error) {
	path := s.filePath(username)

	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return UserStore{Transactions: []Transaction{}}, nil
	}
	if err != nil {
		return UserStore{}, err
	}
	defer f.Close()

	var store UserStore
	if err := json.NewDecoder(f).Decode(&store); err != nil {
		return UserStore{}, err
	}

	return store, nil
}

func (s *Storage) write(username string, store UserStore) error {
	path := s.filePath(username)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(store)
}

func (s *Storage) Read(username string) (UserStore, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.load(username)
}

func (s *Storage) filePath(username string) string {
	return filepath.Join(s.dataDir, username+".json")
}
