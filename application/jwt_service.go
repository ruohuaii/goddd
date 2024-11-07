package application

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/ruohuaii/goddd/infrastructure/cfg"
	"github.com/ruohuaii/goddd/infrastructure/log"

	"github.com/golang-jwt/jwt/v5"
)

// Claims 结构体
type Claims struct {
	MerchantID uint64 `json:"merchant_id"`
	jwt.RegisteredClaims
}

type JwtPayload struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type JwtService struct{}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (s *JwtService) Generate(merchantID uint64) (string, error) {
	privateKey, err := s.loadPrivateKey()
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &Claims{
		MerchantID: merchantID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "merchant server",
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "github.com/ruohuaii/goddd",
			Audience:  []string{"merchant"},
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	accessToken, err := t.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (s *JwtService) Parse(accessToken string) (*Claims, error) {
	publicKey, err := s.loadPublicKey()
	if err != nil {
		return nil, err
	}
	var claims Claims
	_, err = jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, nil
}

func (s *JwtService) loadPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyPem := cfg.GetString("JWT.PrivateKey")
	log.Debug("privateKey", fmt.Sprintf("%q", privateKeyPem))
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyPem))
	if err != nil {
		return nil, fmt.Errorf("not found private key:%v", err)
	}
	return privateKey, nil
}

func (s *JwtService) loadPublicKey() (*rsa.PublicKey, error) {
	publicKeyPem := cfg.GetString("JWT.PublicKey")
	log.Debug("publicKey", fmt.Sprintf("%q", publicKeyPem))
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyPem))
	if err != nil {
		return nil, fmt.Errorf("not found public key:%v", err)
	}
	return publicKey, nil
}
