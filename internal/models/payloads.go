package models

type JWE_01 struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type JWE_02 struct {
	Message string `json:"message"`
}

type EncryptRequest struct {
	Payload string `json:"payload"`
}

type EncryptResponse struct {
	JWE string `json:"jwe"`
}

type DecryptRequest struct {
	JWE string `json:"jwe"`
}

type DecryptResponse struct {
	Payload string `json:"payload"`
}
