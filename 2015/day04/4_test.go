package advent

import (
	"crypto/md5"
	"encoding/hex"
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func Test4_Sanity(t *testing.T) {
	input := []byte("abcdef609043")
	hash := md5.New()
	hash.Write(input)
	assertions.New(t).So(hex.EncodeToString(hash.Sum(nil)), should.Equal, "000001dbbfa3a5c83a2d506429c7b00e")
	//t.Log(hex.EncodeToString(hash.Sum(nil)))
}

func Test4_Examples(t *testing.T) {
	assert := assertions.New(t)
	assert.So(smallestWithZeroPrefix("abcdef", 5), should.Equal, 609043)
}

func Test4_Answer(t *testing.T) {
	if testing.Short() {
		//t.Skip("Long-running...")
		return
	}
	assert := assertions.New(t)
	assert.So(smallestWithZeroPrefix("ckczppom", 5), should.Equal, 117946)
	t.Log("Smallest number appended to 'ckczppom' that produces an md5 starting with '00000':", 117946)

	assert.So(smallestWithZeroPrefix("ckczppom", 6), should.Equal, 3938038)
	t.Log("Smallest number appended to 'ckczppom' that produces an md5 starting with '000000':", 3938038)
}
