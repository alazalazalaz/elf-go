package jwts

import (
	"elf-go/components/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaim struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

const JWT_KID = "INSD9HJ3NF"

func CreateJwtToken(id int) string {
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
	jwtString, err := token.SignedString([]byte(JWT_KID))
	if err != nil {
		logs.Errorf("token SignedString err:%v", err)
	}
	logs.Infof("elf create JWT token success :%s", jwtString)

	return jwtString
}

func ParseJwtToken(jwtString string) error {
	token, err := jwt.ParseWithClaims(jwtString, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_KID), nil
	})

	if c, ok := token.Claims.(*MyClaim); ok && token.Valid {
		logs.Infof("elf parse JWT token success, expire at:%v", c.ExpiresAt)
	} else {
		logs.Errorf("elf parse JWT token error, err:%v", err)
	}

	return err
}
