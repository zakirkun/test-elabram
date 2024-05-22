package jwt

type JWT interface {
	GenerateToken(data map[string]interface{}) (string, error)
	ValidateToken(token string) (bool, error)
	ParseToken(tokenString string) (map[string]interface{}, error)
}

type JWTImpl struct {
	SignatureKey string
	Expiration   int
}
