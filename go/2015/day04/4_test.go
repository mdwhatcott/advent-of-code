package advent

import (
	"crypto/md5"
	"encoding/hex"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func Test4_Sanity(t *testing.T) {
	input := []byte("abcdef609043")
	hash := md5.New()
	hash.Write(input)
	should.So(t, hex.EncodeToString(hash.Sum(nil)), should.Equal, "000001dbbfa3a5c83a2d506429c7b00e")
	//t.Log(hex.EncodeToString(hash.Sum(nil)))
}

func Test4_Examples(t *testing.T) {
	should.So(t, smallestWithZeroPrefix("abcdef", 5), should.Equal, 609043)
}

func Test4_Answer(t *testing.T) {
	if testing.Short() {
		//t.Skip("Long-running...")
		return
	}
	should.So(t, smallestWithZeroPrefix("ckczppom", 5), should.Equal, 117946)
	t.Log("Smallest number appended to 'ckczppom' that produces an md5 starting with '00000':", 117946)

	should.So(t, smallestWithZeroPrefix("ckczppom", 6), should.Equal, 3938038)
	t.Log("Smallest number appended to 'ckczppom' that produces an md5 starting with '000000':", 3938038)
}
