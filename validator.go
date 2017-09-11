package chaos

import (
	"regexp"
)

var (
	V_REGEXP_PHONE    = "^(1(([35][0-9])|[8][0-9]|[7][0-9]|[4][579]))\\d{8}$"
	V_REGEXP_USERNAME = "^[a-zA-Z0-9_]{4,16}$"
	V_REGEXP_PASSWORD = "^[a-zA-Z0-9]{6,16}$"
	V_REGEXP_NICK     = "^[\u4E00-\u9FA5A-Za-z0-9_]{2,12}$"
	V_REGEXP_EMAIL    = "[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?"
	V_REGEXP_CHINESE  = "^[\\u4e00-\\u9fa5]{0,}$"
)
//validator string、phone、email etc.
func IsPhone(phone string) bool {
	reg := regexp.MustCompile(V_REGEXP_PHONE)
	return reg.MatchString(phone)
}

func IsUserName(userName string) bool {
	reg := regexp.MustCompile(V_REGEXP_USERNAME)
	return reg.MatchString(userName)
}

func IsNick(nick string) bool {
	reg := regexp.MustCompile(V_REGEXP_NICK)
	return reg.MatchString(nick)
}

func IsEmail(mail string) bool {
	reg := regexp.MustCompile(V_REGEXP_EMAIL)
	return reg.MatchString(mail)
}

func IsChinese(chars string) bool {
	reg := regexp.MustCompile(V_REGEXP_CHINESE)
	return reg.MatchString(chars)
}

func IsNilString(s string) bool {
	if len(s) < 1 {
		return true
	}
	return false
}

func IsAllNilString(s ...string) bool {
	for _, v := range s {
		if len(v) > 1 {
			return false
		}
	}
	return true
}

func IsPassword(pwd string) bool {
	reg := regexp.MustCompile(V_REGEXP_PASSWORD)
	return reg.MatchString(pwd)
}

func IsASCII(s string) bool {
	for _, c := range s {
		if c >= 0x80 {
			return false
		}
	}
	return true
}
