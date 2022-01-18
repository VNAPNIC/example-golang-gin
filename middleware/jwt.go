package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"serverhealthcarepanel/dto"
	"serverhealthcarepanel/utils/code"
	"serverhealthcarepanel/utils/jwt"
	redisUtil "serverhealthcarepanel/utils/redis"
	"strings"
)

func JWTHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			var rCode int
			var data interface{}

			rCode = code.SUCCESS

			var token = ctx.QueryParam("token")
			var authorization = ctx.Request().Header.Get("Authorization")

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
				return dto.Error(ctx,
					http.StatusUnauthorized,
					rCode,
					code.GetMsg(rCode),
					data,
				)
			} else {
				// Store login user information
				ctx.Set("claims", claims)
			}
			return next(ctx)
		}
	}
}
