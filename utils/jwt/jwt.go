package jwt

import "github.com/golang-jwt/jwt/v5"

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	strToken, err := token.SignedString([]byte("$!1gnK3yyy!!!"))
	if err != nil {
		return "", err
	}

	return strToken, nil
}
