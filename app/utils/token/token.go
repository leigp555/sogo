package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"sogo/app/global/variable"
	"time"
)

type AccessClaims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}
type RefreshClaims struct {
	jwt.RegisteredClaims
}

var (
	accessTokenKey  string
	accessTokenExp  int
	refreshTokenKey string
	refreshTokenExp int
)

func init() {
	accessTokenKey = variable.Config.GetString("token.accessToken.signingKey")
	accessTokenExp = variable.Config.GetInt("token.accessToken.expiresTime")
	refreshTokenKey = variable.Config.GetString("token.refreshToken.signingKey")
	refreshTokenExp = variable.Config.GetInt("token.refreshToken.expiresTime")
}

func AccessToken(uid string) (string, error) {
	claims := AccessClaims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(accessTokenExp) * time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "lgp",      //颁发者
			Subject:   "somebody", //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessTokenKey))
}

func RefreshToken() (string, error) {
	claims := RefreshClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(refreshTokenExp) * time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "lgp",      //颁发者
			Subject:   "somebody", //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshTokenKey))
}

func ParseAccessToken(tokenStr string) (uid string, err error) {
	claims := new(AccessClaims)
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessTokenKey), nil
	})
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			err = errors.New("token is expired")
			return "", err
		} else {
			err = errors.New("invalid token")
			return "", err
		}
	}
	if token.Valid {
		return claims.Uid, nil
	}
	return "", errors.New("invalid token")
}

func ParseRefreshToken(tokenStr string) (err error) {
	claims := new(RefreshClaims)
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshTokenKey), nil
	})
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			err = errors.New("token is expired")
			return err
		} else {
			err = errors.New("invalid token")
			return err
		}
	}
	if token.Valid {
		return nil
	}
	err = errors.New("invalid token")
	return err
}

func FlushToken(accessToken string, refreshToken string) (newToken string, err error) {
	//判断refreshToken是否过期或者有效
	err = ParseRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	//判断accessToken是否过期或者有效
	claims := new(AccessClaims)
	_, err = jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessTokenKey), nil
	})
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			newToken, err := AccessToken(claims.Uid)
			if err != nil {
				return "", err
			}
			return newToken, nil
		} else {
			err = errors.New("invalid token")
			return "", err
		}
	}
	return "", errors.New("token invalid don't need generate new token")
}
