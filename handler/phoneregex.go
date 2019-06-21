package handler

import "regexp"

// Regexphone 判断手机号码是否正确
func Regexphone(phone string) bool {
	right, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, phone)
	return right
}