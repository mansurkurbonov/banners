package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenRandNum generates random number 5 digits
func GenRandNum() string {
	return strconv.Itoa(rand.Intn(99999-10000) + 10000)
}

// GenRand8Nums -
func GenRand8Nums() string {
	return strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)
}

// StrEmpty best way to check if string is empty.
func StrEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// ReadJSON decodes json file.
func ReadJSON(path string, conf interface{}) (err error) {
	var (
		file *os.File
	)
	file, err = os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(conf)
	return
}

//RmBadChars remove chars
func RmBadChars(s string) string {
	return strings.Map(
		func(r rune) rune {
			if unicode.IsDigit(r) || unicode.IsLetter(r) {
				return r
			}
			return -1
		},
		s,
	)
}

// Hs256 SHA-256 HMAC hex hash
func Hs256(key, val []byte) string {
	hashhmac := hmac.New(sha256.New, key)
	hashhmac.Write(val)
	return hex.EncodeToString(hashhmac.Sum(nil))
}
