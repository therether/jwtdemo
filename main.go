package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	//生成token
	JwtKey := []byte("suhfiiafhaohi") //加密,密钥
	//MyClaims的实例化
	c := MyClaims{
		Username: "admin",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,      //生效时间
			ExpiresAt: time.Now().Unix() + 60*60*2, //过期时间，失效时间
			Issuer:    "admin",                     //签发人
		},
	}
	//t:生成的字符串
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//加密token
	//tokenString:加密后的字符串
	tokenString, err := t.SignedString(JwtKey)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(tokenString)

	//解析token
	//token:解密后的token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	fmt.Println(err)
	fmt.Println(token)
}
