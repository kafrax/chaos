package chaos

import (
	"strconv"
	"io"
	"crypto/rand"

	"github.com/google/uuid"
	"math/big"
)

//uuid+unix time
func RandId() string {
	return strconv.FormatInt(int64(uuid.New().Time()/10000000000)*10000000000+Int64Range(100000000,10000000000),10)
}

func RandIdInt64() int64 {
	return int64(uuid.New().Time()/10000000000)*10000000000 +Int64Range(100000000,10000000000)
}

//0123456789 select 6 password number
const C_RAND_TMP = "0123456789"

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

func Int64Range(min, max int64) int64 {
	var result int64
	maxRand := max - min
	b, err := rand.Int(rand.Reader, big.NewInt(int64(maxRand)))
	if err != nil {
		return max
	}
	result = min + b.Int64()
	return result
}
