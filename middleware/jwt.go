package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"healthcare-panel/common"
	"healthcare-panel/utils/code"
	"healthcare-panel/utils/jwt"
	redisUtil "healthcare-panel/utils/redis"
	"net/http"
	"strings"
)

func JWTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		g := common.Gin{C: c}

		var rCode int
		var data interface{}

		rCode = code.SUCCESS

		var token = g.C.Query("token")
		var authorization = g.C.GetHeader("Authorization")
		if authorization != "" {
			parts := strings.SplitN(authorization, " ", 2)
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			}
		}

		var claims *jwtUtil.Claims
		var err error

		if token == "" {
			rCode = code.TokenInvalid
		} else {
			claims, err = jwtUtil.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					rCode = code.ErrorAuthCheckTokenTimeout
				default:
					rCode = code.ErrorAuthCheckTokenFail
				}
			}
		}

		// check token on the redis
		if rCode != code.SUCCESS && rCode != code.TokenInvalid && claims != nil {
			_, _ = redisUtil.Delete(claims.Issuer)
		} else if claims != nil {
			isExist, err := redisUtil.Exists(claims.Issuer)
			if err != nil {
				rCode = code.ErrorAuthCheckTokenTimeout
				data = err.Error()
			}
			if isExist == false {
				rCode = code.ErrorAuthCheckTokenTimeout
				data = redisUtil.Error{Msg: "Token on the redis is not found!"}
			}
		}

		if rCode != code.SUCCESS {
			g.Error(http.StatusUnauthorized,
				rCode,
				code.GetMsg(rCode),
				data)
			g.C.Abort()
		} else {
			// Store login user information
			g.C.Set("claims", claims)
		}
		g.C.Next()
	}
}
