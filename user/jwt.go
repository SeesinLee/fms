package user

import (
	"fms/moddle"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("user_key") //初始化key

type Claim struct {
	ID int
	jwt.StandardClaims
}

func ReleaseToken(u *moddle.UserInfo)(string,error) {
	expiresTime := time.Now().Add(time.Hour*24)		// 设置token过期时间
	claims := &Claim{
		ID: u.ID,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "user",
			Subject: "user_info",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err := token.SignedString(jwtKey)
	return tokenString,err
}

func ParseToken(tokenString string)(*jwt.Token,*Claim,error){
	claim := &Claim{}
	token,err := jwt.ParseWithClaims(tokenString,claim, func(token *jwt.Token) (interface{}, error) {
		return jwtKey,nil
	})
	return token,claim,err
}
