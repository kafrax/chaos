package chaos

import (
	"unsafe"
	"bytes"
)

//data convert
func Byte2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String2Byte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

//转换成ascii
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