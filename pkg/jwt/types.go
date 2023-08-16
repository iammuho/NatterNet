package jwt

import "github.com/iammuho/natternet/pkg/errorhandler"

type JwtContext interface {
	CreatePair(claims map[string]interface{}) (*JWTResponse, *errorhandler.Response)
}

type JWTResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}
