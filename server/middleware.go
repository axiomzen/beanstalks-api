package server

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	jwt "github.com/dgrijalva/jwt-go"
)

type claimsCtxKeyType string

var claimsCtxKey = claimsCtxKeyType("claims")

func parseToken(token, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
}

type Middler func(http.HandlerFunc) http.HandlerFunc

func wrap(f http.HandlerFunc, middleware ...Middler) http.HandlerFunc {
	for _, wrapper := range middleware {
		f = wrapper(f)
	}
	return f
}

func (s *Server) logIt(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		s.log.Infof("%s %s", req.Method, req.URL.Path)
		next(res, req)
	}
}

func (s *Server) recover(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				s.log.Errorf("recovered from error in request handler %v\n", r)
				debug.PrintStack()
				res.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next(res, req)
	}
}

func (s *Server) authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("az-auth-token")
		parsedToken, err := verifyJWT(token, s.config.Secret)
		if err != nil {
			s.log.WithError(err).Error("failed to parse user auth token")
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			s.log.WithError(err).Error("user auth token contains invalid claims")
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		// Attach the token to the request context
		req = req.WithContext(context.WithValue(req.Context(), claimsCtxKey, claims))
		next(res, req)
	}
}
