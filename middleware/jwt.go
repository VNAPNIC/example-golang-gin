package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"serverhealthcarepanel/utils"
	"serverhealthcarepanel/utils/code"
	"serverhealthcarepanel/utils/response"
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

			var claims *utils.Claims
			var err error

			if token == "" {
				rCode = code.TokenInvalid
			} else {
				claims, err = utils.ParseToken(token)
				if err != nil {
					switch err.(*jwt.ValidationError).Errors {
					case jwt.ValidationErrorExpired:
						rCode = code.ErrorAuthCheckTokenTimeout
					default:
						rCode = code.ErrorAuthCheckTokenFail
					}
				}
			}

			//jwtCount, _ := userService.InBlockList(token)
			//if jwtCount >= 1 {
			//	rCode = code.AuthTokenInBlockList
			//}

			if rCode != code.SUCCESS {
				return response.Response(ctx,
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
