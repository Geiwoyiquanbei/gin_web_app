package JWT

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2

var MyScrect = []byte("长亭外古道边")
var FreshToken string = "default"

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userid int64, username string) (atoken string, rtoken string, err error) {
	// 创建一个我们自己的声明
	claims := MyClaims{
		userid, // 自定义字段
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "my-project", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(MyScrect)
	rtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "web_app",
	}).SignedString(MyScrect)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return atoken, rtoken, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return MyScrect, nil
	})

	v, _ := err.(jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		return nil, errors.New("token pass time")
	}
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func RefreshToken(atoken, rtoken string) (newAtoken, newRtoken string, err error) {

	if _, err = jwt.Parse(rtoken, func(token *jwt.Token) (interface{}, error) {
		return MyScrect, nil
	}); err != nil {
		return
	}
	var claims MyClaims
	_, err = jwt.ParseWithClaims(atoken, &claims, func(token *jwt.Token) (interface{}, error) {
		return MyScrect, nil
	})
	v, _ := err.(jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID, claims.UserName)
	}
	return
}
