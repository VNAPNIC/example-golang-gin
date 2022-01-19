package jwtUtil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	redisUtil "healthcare-panel/utils/redis"
	"healthcare-panel/utils/setting"
	"log"
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

	if err != nil {
		log.Println(err)
		return token, err
	}

	// set token to the redis
	successful, err := redisUtil.Set(claims.Issuer, token, time.Hour*24)
	if err != nil {
		log.Println(err)
		return token, err
	}
	if !successful {
		return token, redisUtil.Error{Msg: "Can't set token to redis"}
	}

	return token, nil
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

func GetClaim(ctx *gin.Context) (Claims, bool) {
	claims, exists := ctx.Get("claims")
	user := claims.(*Claims)
	return *user, exists
}
