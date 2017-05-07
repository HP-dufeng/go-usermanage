package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"crypto/rsa"

	"github.com/dgrijalva/jwt-go/request"
)

const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

// private key for signing and public key for verification
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func initKeys() {
	var err error
	var signBytes, verifyBytes []byte

	signBytes, err = ioutil.ReadFile(privKeyPath)
	fatalf(err)
	verifyBytes, err = ioutil.ReadFile(pubKeyPath)
	fatalf(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatalf(err)
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatalf(err)
}

func fatalf(err error) {
	if err != nil {
		log.Fatalf("initKeys: %s\n", err)
		panic(err)
	}
}

func GenerateJWT(name, role string) (string, error) {
	claims := jwt.MapClaims{
		"iss": "admin",
		"UserInfo": struct {
			Name string
			Role string
		}{name, role},
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	}

	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Middleware for validating jwt tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// validate the token
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		// Verify the token with public key
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return
			default:
				DisplayAppError(w,
					err,
					"Error while parsing the Access Token!",
					500,
				)
				return
			}
		default:
			DisplayAppError(w,
				err,
				"Error while parsing Access Token!",
				500)
			return
		}
	}

	if token.Valid {
		next(w, r)
	} else {
		DisplayAppError(w,
			err,
			"Invalid Access Token",
			401)
	}
}
