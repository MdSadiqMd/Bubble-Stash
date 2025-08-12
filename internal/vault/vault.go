package vault

import (
	"errors"

	"github.com/MdSadiqMd/Bubble-Stash/internal/encrypt"
)

type Vault struct {
	encodingKey string
	keyValues   map[string]string
}

func Memory(encodingKey string) Vault {
	return Vault{
		encodingKey: encodingKey,
		keyValues:   make(map[string]string),
	}
}

func (v *Vault) Get(key string) (string, error) {
	encryptedValue, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("vault: no value found for that")
	}

	decryptedValue, err := encrypt.Decrypt(v.encodingKey, encryptedValue)
	if err != nil {
		return "", err
	}
	return decryptedValue, nil
}

func (v *Vault) Set(key, value string) error {
	encryptedValue, err := encrypt.Encrypt(v.encodingKey, value)
	if err != nil {
		return err
	}

	v.keyValues[key] = encryptedValue
	return nil
}
