package tool

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

//token
func GetJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

//md5 加密
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

//手机号验证
func TelephoneVerified(mobile string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

//字符串的长度判断
func StringLenVerified(v string, n int) bool {
	//lengths := len([]rune(v))
	//lengths := bytes.Count([]byte(v), nil)
	lengths := strings.Count(v, "") - 1
	//lengths := utf8.RuneCountInString(v)
	if lengths != n {
		return false
	} else {
		return true
	}
}

//邮箱的验证
func EmailVerified(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
