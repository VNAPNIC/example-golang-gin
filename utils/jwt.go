package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"serverhealthcarepanel/utils/setting"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"user_name"`
	RoleKey  string `json:"role_key"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

// GenerateToken Generate Token used for auth
func GenerateToken(userClaim Claims) (string, error) {
	nowTime := time.Now()
	zeroTime := time.Date(nowTime.YearDay(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.Location())
	//expire time
	expireTime := zeroTime.Add(time.Hour * 24)
	claims := Claims{
		userClaim.UserId,
		userClaim.Username,
		userClaim.RoleKey,
		userClaim.IsAdmin,
		jwt.StandardClaims{
			Issuer:    "server_healthcare_panel",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaim.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return jwtSecret, nil
	}

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}