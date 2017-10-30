package chaos

import (
	"encoding/xml"
	"unsafe"
	"bytes"
	"github.com/json-iterator/go"
)

//data convert
func Byte2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String2Byte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

//to ascii
func String2ASCII(s string) string {
	if IsASCII(s) {
		return s
	}
	var buf bytes.Buffer
	for _, c := range s {
		if c < 0x80 {
			buf.WriteByte(byte(c))
		}
	}
	return buf.String()
}

func MustMarshal2String(v interface{}) string {
	b, _ := jsoniter.MarshalToString(v)
	return b
}

func MustXMll2Byte(v interface{}) []byte {
	b, _ := xml.Marshal(v)
	return b
}