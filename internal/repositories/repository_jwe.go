package repositories

import (
	"crypto/rand"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
)

const tag string = "JWERepository"

type CryptoRepository interface {
	Encrypt(payload string) (string, error)
	Decrypt(token string) (string, error)
}

type cryptoRepository struct {
	key []byte
}

func NewCryptoRepository(key []byte) CryptoRepository {
	return &cryptoRepository{key: key}
}

func (r *cryptoRepository) Encrypt(payload string) (string, error) {
	encrypted, err := jwe.Encrypt(
		[]byte(payload),
		jwe.WithKey(jwa.DIRECT, r.key),
		jwe.WithContentEncryption(jwa.A256GCM),
	)
	if err != nil {
		return "", err
	}

	return string(encrypted), nil
}

func (r *cryptoRepository) Decrypt(token string) (string, error) {
	decrypted, err := jwe.Decrypt(
		[]byte(token),
		jwe.WithKey(jwa.DIRECT, r.key),
	)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

// Geração de chave simétrica (32 bytes)
func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar chave: %w", err)
	}
	return key, nil
}
