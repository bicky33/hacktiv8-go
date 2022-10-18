package helper

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Email string `json:"email"`
	Role  string
	jwt.RegisteredClaims
}

// func GenerateToken(user dto.User, expiredTime time.Duration, privateKey string) (string, error) {
// 	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	claim := &JWTClaim{
// 		Email: user.Email,
// 		Role:  user.Role,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
// 	tokenString, err := token.SignedString(key)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// func ValidateToken(token string, expiredTime time.Duration, publicKey string) (*JWTClaim, error) {
// 	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
// 	if err != nil {
// 		fmt.Println("ttttt")
// 		return nil, err
// 	}
// 	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
// 	if err != nil {
// 		fmt.Println("rrrrr")
// 		return nil, err
// 	}
// 	contractToken, err := jwt.ParseWithClaims(
// 		token,
// 		&JWTClaim{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return key, nil
// 		},
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	claims := contractToken.Claims.(*JWTClaim)
// 	return claims, nil
// }
