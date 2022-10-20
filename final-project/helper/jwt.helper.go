package helper

import (
	"encoding/base64"
	"final-project/dto"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	ID       int32  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(user dto.UserCreateResponse, expiredTime time.Duration, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", err
	}
	claim := &JWTClaim{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiredTime).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string, publicKey string) (*JWTClaim, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, err
	}
	contractToken, err := jwt.ParseWithClaims(
		token,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims := contractToken.Claims.(*JWTClaim)
	return claims, nil
}
