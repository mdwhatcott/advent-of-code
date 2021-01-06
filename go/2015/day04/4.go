package advent

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func smallestWithZeroPrefix(key string, zeros int) int {
	hash := md5.New()
	for x := 0; ; x++ {
		hash.Write([]byte(key))
		salt := strconv.Itoa(x)
		hash.Write([]byte(salt))
		sum := hash.Sum(nil)
		digest := hex.EncodeToString(sum)
		if strings.HasPrefix(digest, strings.Repeat("0", zeros)) {
			return x
		}
		hash.Reset()
	}
	return -1
}
