package jwt

import (
	"crypto/rsa"
	"strings"
	"time"

	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/gofiber/fiber/v2"
	jwtPKG "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwtPKG.Claims
	CustomClaims map[string]interface{} `json:"customClaims"`
}

type jwt struct {
	options    JwtOptions
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewJwtContext is the constructor for the jwt context
func NewJWT(opts ...Option) (JwtContext, error) {
	jwtOptions := JwtOptions{}
	for _, o := range opts {
		o(&jwtOptions)
	}

	j := &jwt{
		options: jwtOptions,
	}

	// Load the keys
	j.loadPublicKey()
	j.loadPrivateKey()

	return j, nil
}

// CreatePair aims to create a new JSON WEB TOKEN for access/refresh
func (j *jwt) CreatePair(claims map[string]interface{}) (*JWTResponse, *errorhandler.Response) {
	expireAccess := time.Now().Add(time.Hour)
	accessToken, err := j.createJWT(claims, true, expireAccess)
	if err != nil {
		return nil, &errorhandler.Response{Code: errorhandler.InternalSystemErrorCode, Message: err.Error(), StatusCode: fiber.StatusBadGateway}
	}

	expireRefresh := time.Now().AddDate(2, 0, 0)
	refreshToken, err := j.createJWT(claims, true, expireRefresh)
	if err != nil {
		return nil, &errorhandler.Response{Code: errorhandler.InternalSystemErrorCode, Message: err.Error(), StatusCode: fiber.StatusBadGateway}
	}

	return &JWTResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expireAccess.Unix(),
	}, nil
}

// CreateJWT aims to create a new JSON WEB TOKEN
func (j *jwt) createJWT(claims map[string]interface{}, expire bool, expireDate time.Time) (string, error) {
	// create a new JWT with the claims and the signing method
	t := jwtPKG.NewWithClaims(jwtPKG.SigningMethodRS256, // (1)!
		jwtPKG.MapClaims{ // (2)!
			"iss": j.options.Issuer,
			"sub": j.options.Subject,
			"exp": expireDate.Unix(),
			"iat": time.Now().Unix(),
			"kid": j.options.Kid,
			// add the custom claims
			"custom": claims,
		})

	ss, err := t.SignedString(j.privateKey)

	if err != nil {
		return "", err
	}

	return ss, nil
}

// ParseJWT parses the given JWT token into struct
func (j *jwt) ParseJWT(auth string) (map[string]interface{}, *errorhandler.Response) {
	// Parse the JWT Token
	token, err := jwtPKG.Parse(parseBearerAuth(auth), func(token *jwtPKG.Token) (interface{}, error) {
		return j.publicKey, nil
	})

	if err != nil {
		return nil, &errorhandler.Response{Code: errorhandler.InvalidAccessTokenErrorCode, Message: errorhandler.InvalidAccessTokenMessage, StatusCode: fiber.StatusBadRequest}
	}

	out := Claims{}

	if claims, ok := token.Claims.(jwtPKG.MapClaims); ok && token.Valid {
		if claims["custom"] == nil {
			return nil, &errorhandler.Response{Code: errorhandler.InvalidAccessTokenErrorCode, Message: errorhandler.InvalidAccessTokenMessage, StatusCode: fiber.StatusBadRequest}
		}

		out.CustomClaims = claims["custom"].(map[string]interface{})
	} else {
		return nil, &errorhandler.Response{Code: errorhandler.InvalidAccessTokenErrorCode, Message: errorhandler.InvalidAccessTokenMessage, StatusCode: fiber.StatusBadRequest}
	}

	return out.CustomClaims, nil
}

// parseBearerAuth parses the header to get JWT token
func parseBearerAuth(auth string) string {
	if strings.HasPrefix(auth, "Bearer ") {
		bearer := auth[7:]
		return bearer
	}
	return auth
}
