package jwt

import (
	"crypto/rsa"
	"time"

	"github.com/iammuho/natternet/pkg/errorhandler"

	"github.com/gofiber/fiber/v2"
	jose "gopkg.in/square/go-jose.v2"
	jwtPKG "gopkg.in/square/go-jose.v2/jwt"
)

type Claims struct {
	*jwtPKG.Claims
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
	// create Square.jose signing key
	key := jose.SigningKey{Algorithm: jose.RS256, Key: j.privateKey}

	// create a Square.jose RSA signer, used to sign the JWT
	var signerOpts = jose.SignerOptions{}
	signerOpts.WithType("JWT")
	signerOpts.WithHeader("kid", j.options.Kid)
	rsaSigner, err := jose.NewSigner(key, &signerOpts)

	if err != nil {
		return "", err
	}

	// create an instance of Builder that uses the rsa signer
	builder := jwtPKG.Signed(rsaSigner)

	// public claims
	pubClaims := &jwtPKG.Claims{
		Issuer:   j.options.Issuer,
		Subject:  j.options.Subject,
		IssuedAt: jwtPKG.NewNumericDate(time.Now()),
		Expiry:   jwtPKG.NewNumericDate(expireDate),
	}

	c := Claims{
		pubClaims,
		claims,
	}

	// Add the claims. Note Claims returns a Builder so can chain
	builder = builder.Claims(c)

	// validate all ok, sign with the RSA key, and return a compact JWT
	rawJWT, err := builder.CompactSerialize()
	if err != nil {
		return "", err
	}

	return rawJWT, nil
}
