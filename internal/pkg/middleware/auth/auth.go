package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

var currentUserKey struct{}

type CurrentUser struct {
	UserID string
}

func GenerateToken(secret string, userid string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userid,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				auths := strings.SplitN(tokenString, " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], "Token") {
					return nil, errors.New("jwt token missing")
				}

				token, err := jwt.Parse(auths[1], func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(secret), nil
				})

				if err != nil {
					return nil, err
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					// put CurrentUser into ctx
					if u, ok := claims["userid"]; ok {
						ctx = WithContext(ctx, &CurrentUser{UserID: string(u.(string))})
					}
				} else {
					return nil, errors.New("Token Invalid")
				}
			}
			return handler(ctx, req)
		}
	}
}

//func NewHeaderServer() func(handler middleware.Handler) middleware.Handler {
//	return func(handler middleware.Handler) middleware.Handler {
//		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
//			var orgnizationId int64
//			var uhasVisit int64
//			var districtID int64
//			if md, ok := metadata.FromServerContext(ctx); ok {
//				extra := md.Get("U-OrGniZaTiOn")
//				fmt.Println(extra)
//			}
//			if header, ok := transport.FromServerContext(ctx); ok {
//				orgnizationId, _ = strconv.ParseInt(header.RequestHeader().Get("U-OrGniZaTiOn"), 10, 64)
//				uhasVisit, _ = strconv.ParseInt(header.RequestHeader().Get("U-HASVISIT"), 10, 64)
//				//U-DistrictID
//				districtID, _ = strconv.ParseInt(header.RequestHeader().Get("U-DistrictID"), 10, 64)
//			} else {
//				return nil, errors.New("U-Token missing")
//			}
//			ctx = context.WithValue(ctx, OrgnizationId, orgnizationId)
//			ctx = context.WithValue(ctx, HasVisit, uhasVisit == 1)
//			ctx = context.WithValue(ctx, DistrictID, districtID)
//			reply, err = handler(ctx, req)
//			return
//		}
//	}
//}

func FromContext(ctx context.Context) *CurrentUser {
	return ctx.Value(currentUserKey).(*CurrentUser)
}

func WithContext(ctx context.Context, user *CurrentUser) context.Context {
	return context.WithValue(ctx, currentUserKey, user)
}
