package chaos

import (
	"strconv"
	"time"
	"io"
	"crypto/rand"

	"github.com/google/uuid"
)

//uuid+unix time
func RandId() string {
	return strconv.FormatInt(int64(uuid.New().Time()/10000000000)*10000000000+time.Now().Unix(), 10)
}

//0123456789 select 6 password number
const C_RAND_TMP= "0123456789"
func RandPassword(length int, chars []byte) string {
	newPwd := make([]byte, length)
	random := make([]byte, length+(length/4)) // storage for random bytes.
	charsLength := byte(len(chars))
	max := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, random); err != nil {
			panic(err)
		}
		for _, c := range random {
			if c >= max {
				continue
			}
			newPwd[i] = chars[c%charsLength]
			i++
			if i == length {
				return string(newPwd)
			}
		}
	}
	panic("unreachable")
}