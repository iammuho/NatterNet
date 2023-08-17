package jwt

import "github.com/iammuho/natternet/pkg/errorhandler"

//go:generate mockgen -destination=mocks/mock_jwt_contexter.go -package=mockjwt -source=types.go

type JwtContext interface {
	CreatePair(claims map[string]interface{}) (*JWTResponse, *errorhandler.Response)
	ParseJWT(auth string) (map[string]interface{}, *errorhandler.Response)
}

type JWTResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}
