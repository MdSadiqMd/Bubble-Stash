package vault

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/MdSadiqMd/Bubble-Stash/internal/encrypt"
)

type Vault struct {
	mutex       sync.Mutex
	encodingKey string
	filePath    string
	keyValues   map[string]string
}

func File(encodingKey, filePath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filePath:    filePath,
	}
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}

func (v *Vault) load() error {
	f, err := os.Open(v.filePath)
	if os.IsNotExist(err) {
		v.keyValues = make(map[string]string)
		return nil
	}
	if err != nil {
		return err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if fi.Size() == 0 {
		v.keyValues = make(map[string]string)
		return nil
	}

	r, err := encrypt.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.readKeyValues(r)
}

func (v *Vault) save() error {
	f, err := os.OpenFile(v.filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := encrypt.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return "", err
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("vault: no value found for that")
	}
	return value, nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}

	v.keyValues[key] = value

	err = v.save()
	if err != nil {
		return err
	}
	return nil
}
