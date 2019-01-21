package oauthstore

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"golang.org/x/oauth2"
)

// FileStorage satisfies our storage interface
type FileStorage struct {
	Path string
	mu   sync.RWMutex
}

// GetToken retrieves a token from a file
func (f *FileStorage) GetToken() (*oauth2.Token, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	in, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	var t *oauth2.Token
	data := json.NewDecoder(in)
	return t, data.Decode(&t)
}

// SetToken creates, truncates, then stores a token
// in a file
func (f *FileStorage) SetToken(t *oauth2.Token) error {
	if t == nil || !t.Valid() {
		return errors.New("bad token")
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	out, err := os.OpenFile(f.Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()
	data, err := json.Marshal(&t)
	if err != nil {
		return err
	}

	_, err = out.Write(data)
	return err
}
