package jwtUtil

import (
	"fmt"
	"serverhealthcarepanel/utils/setting"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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
	timeNow := time.Now()
	//expire time
	expireTime := timeNow.Add(time.Hour * 24)
	claims := Claims{
		userClaim.UserId,
		userClaim.Username,
		userClaim.RoleKey,
		userClaim.IsAdmin,
		jwt.StandardClaims{
			Issuer:    "server-healthcare-panel",
			IssuedAt:  timeNow.Unix(),
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

func GetClaim(ctx echo.Context) Claims {
	claims := ctx.Get("claims")
	user := claims.(*Claims)
	return *user
}
