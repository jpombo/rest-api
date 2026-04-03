package usecases

import "rest-api/internal/repositories"

const tagUsecaseJwe string = "UsecaseJwe"

type CryptoUseCase interface {
	EncryptJSON(json string) (string, error)
	DecryptJWE(jwe string) (string, error)
}

type cryptoUseCase struct {
	repo repositories.CryptoRepository
}

func NewCryptoUseCase(repo repositories.CryptoRepository) CryptoUseCase {
	return &cryptoUseCase{repo: repo}
}

func (u *cryptoUseCase) EncryptJSON(json string) (string, error) {
	return u.repo.Encrypt(json)
}

func (u *cryptoUseCase) DecryptJWE(jwe string) (string, error) {
	return u.repo.Decrypt(jwe)
}
