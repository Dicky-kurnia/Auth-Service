package helper

import (
	"errors"
	"os"
	"service-producer/exception"
	"service-producer/model"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(request model.JwtPayload) *model.TokenDetails {
	accessExpired, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_MINUTE"))

	td := &model.TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * time.Duration(accessExpired)).Unix()

	keyAccess, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_ACCESS_PRIVATE_KEY")))
	exception.PanicIfNeeded(err)

	now := time.Now().UTC()

	atClaims := jwt.MapClaims{}
	atClaims["agent_id"] = request.AgentId
	atClaims["exp"] = td.AtExpires
	atClaims["iat"] = now.Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	at.Header["topindopay"] = "jwt"
	td.AccessToken, err = at.SignedString(keyAccess)

	if err != nil {
		panic(errors.New(model.AUTHENTICATION_FAILURE_ERR_TYPE))
	}

	return td

}
