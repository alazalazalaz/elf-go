package jwts

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaim struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

const JWT_KID = "INSD9HJ3NF"

func CreateJwtToken(id int) (string, error) {
	standardClaim := jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 60,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "elf-framework",
	}

	mc := MyClaim{
		id,
		standardClaim,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	if token == nil {
		return "", errors.New("token is nil")
	}

	jwtString, err := token.SignedString([]byte(JWT_KID))
	if err != nil {
		return "", err
	}

	return jwtString, nil
}

func ParseJwtToken(jwtString string) error {
	token, err := jwt.ParseWithClaims(jwtString, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_KID), nil
	})

	if err != nil {
		return err
	}

	if token == nil {
		return errors.New("token is nil pointer")
	}

	if _, ok := token.Claims.(*MyClaim); ok && token.Valid {
		//logs.Infof("elf parse JWT token success, expire at:%v", c.ExpiresAt)
	} else {
		return errors.New(fmt.Sprintf("elf parse JWT token error, err:%v", err))
	}

	return err
}
