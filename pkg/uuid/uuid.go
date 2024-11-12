package uuid

import "github.com/google/uuid"

func GenerateUuid() (token string) {
	byteToken := uuid.New()
	token = byteToken.String()
	return
}
