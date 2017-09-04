package chaos

import "unsafe"

//data convert
func Byte2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func String2Byte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

