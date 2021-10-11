package authentication

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wasanx25/go_nextjs_authentication/config"
	"github.com/wasanx25/go_nextjs_authentication/model"
	"github.com/wasanx25/go_nextjs_authentication/repository"
)

type JWTCustomClaims struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Email   string `json:"email"`
	jwt.StandardClaims
}

func JWTConfig(conf config.Config, userRepository repository.UserRepositoryInterface) middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte("secret"),
		KeyFunc: func(token *jwt.Token) (interface{}, error) {
			return getKey(token, conf)
		},
		SigningMethod: jwt.SigningMethodRS256.Name,
		SuccessHandler: func(c echo.Context) {
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(*JWTCustomClaims)

			user := model.User{Username: claims.Name, Auth0UserID: claims.Subject}
			err := userRepository.CreateIfNotExists(user)

			if err != nil {
				fmt.Println(err)
				panic(err)
			}
		},
	}
}

func getKey(token *jwt.Token, conf config.Config) (interface{}, error) {
	aud := conf.Auth0Audience
	claims := token.Claims.(*JWTCustomClaims)
	checkAud := claims.VerifyAudience(aud, false)
	if !checkAud {
		return token, errors.New("Invalid audience.")
	}

	iss := conf.Auth0Issuer
	checkIss := claims.VerifyIssuer(iss, false)
	if !checkIss {
		return token, errors.New("Invalid issuer.")
	}

	cert, err := getPemCert(token, conf)
	if err != nil {
		panic(err.Error())
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func getPemCert(token *jwt.Token, conf config.Config) (string, error) {
	cert := ""
	resp, err := http.Get(conf.AUTH0JWKSURL)
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	return cert, nil
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}
