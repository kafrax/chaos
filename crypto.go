package chaos

import (
	"io"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"crypto/sha1"
	"encoding/base64"
)

//commonly used secret function

func MD5BySalt(src,salt string)string{
	hash := md5.New()
	io.WriteString(hash, src)
	io.WriteString(hash, salt)
	return hex.EncodeToString(hash.Sum(nil))
}

func MD5(src string)string{
	hash := md5.New()
	io.WriteString(hash, src)
	return hex.EncodeToString(hash.Sum(nil))
}

func SHA1BySalt(src,salt string)(string,error){
	t := sha1.New()
	_, err := io.WriteString(t, src)
	if err != nil {
		return "",err
	}
	_,err =io.WriteString(t,salt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", t.Sum(nil)), nil
}

func SHA1(src  string)(string,error){
	t := sha1.New()
	_, err := io.WriteString(t, src)
	if err != nil {
		return "",err
	}
	return fmt.Sprintf("%x", t.Sum(nil)), nil
}

type B64Encoding=base64.Encoding
func B64NewEncoding(s string)*B64Encoding{
	return base64.NewEncoding(s)
}

func (b *B64Encoding)Encode(s string) string {
	return b.EncodeToString([]byte(s))
}

func (b *B64Encoding)Decode(s string)string{
	result, err := b.DecodeString(s)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(result)
}
