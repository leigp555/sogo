package token

import "github.com/golang-jwt/jwt/v4"

//生成一个token

func GenerateToken(claims jwt.Claims, key string) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(key))
}

//解析token

func ParseToken(token string, key string) {
	claims := new(CustomClaims)

	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		v, ok := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			err = errors.New("token is expired")
			return "", err
		} else {
			err = errors.New("invalid token")
			return "", err
		}
	}
}
