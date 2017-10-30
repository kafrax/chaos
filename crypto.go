package chaos

import (
	"io"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"crypto/sha1"
	"encoding/base64"
	"strings"

	"sort"
	"github.com/kafrax/logx"
)

//commonly used secret function
func MD5BySalt(src, salt string) string {
	hash := md5.New()
	io.WriteString(hash, src)
	io.WriteString(hash, salt)
	return hex.EncodeToString(hash.Sum(nil))
}

func MD5(src string) string {
	hash := md5.New()
	io.WriteString(hash, src)
	return hex.EncodeToString(hash.Sum(nil))
}

func SHA1BySalt(src, salt string) (string, error) {
	t := sha1.New()
	_, err := io.WriteString(t, src)
	if err != nil {
		return "", err
	}
	_, err = io.WriteString(t, salt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", t.Sum(nil)), nil
}

func SHA1(src string) (string, error) {
	t := sha1.New()
	_, err := io.WriteString(t, src)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", t.Sum(nil)), nil
}

type B64Encoding struct{ b *base64.Encoding }

func B64NewEncoding(s string) *B64Encoding {
	return &B64Encoding{b: base64.NewEncoding(s)}
}

func (b *B64Encoding) B64Encode(s string) string {
	return b.b.EncodeToString([]byte(s))
}

func (b *B64Encoding) B64Decode(s string) string {
	result, err := b.b.DecodeString(s)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(result)
}


//--------------------- for wechat pay -----------------------
func CheckSign(msg map[string]interface{}, key, sign string) bool {
	signCalc := CalcSign(msg, key)
	if sign == signCalc {
		return true
	}
	return false
}

//CalcSign
func CalcSign(mReq map[string]interface{}, key string) (sign string) {
	sortedKeys := make([]string, 0)
	for k, _ := range mReq {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	var signStrings string
	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	signStrings = signStrings + "key=" + key
	logx.Debugf("CalcSign |signStr=%v", signStrings)
	md5Ctx := md5.New()
	md5Ctx.Write(String2Byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}
