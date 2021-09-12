package api

import (
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const (
    jwtSecretKeyFilepath = "/opt/tedris/jwt/secret_key.pem"
)

// JWTSecretKey returns the file contents of the secret key file.
func JWTSecretKey(_ *jwt.Token) (interface{}, error) {
    f, err := ioutil.ReadFile(jwtSecretKeyFilepath)
    if err != nil {
        return nil, errors.Wrap(err, "unable to read secret key")
    }
    return jwt.ParseRSAPublicKeyFromPEM(f)
}
