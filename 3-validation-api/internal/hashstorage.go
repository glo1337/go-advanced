package internal

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"os"
)

type StorageItem struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type HashStorage struct {
	Path  string
	Items []StorageItem
}

func (hs *HashStorage) ReadItems() ([]StorageItem, error) {
	file, err := os.Open(hs.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return []StorageItem{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var items []StorageItem
	err = json.NewDecoder(file).Decode(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (hs *HashStorage) WriteItems(items []StorageItem) error {
	file, err := os.Create(hs.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(items)
}

func GenerateHash() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func (hs *HashStorage) FindByHash(hash string) *StorageItem {
	for i := range hs.Items {
		if hs.Items[i].Hash == hash {
			return &hs.Items[i]
		}
	}
	return nil
}
