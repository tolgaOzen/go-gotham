package helpers

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math"
	"math/rand"
	"reflect"
	"regexp"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func NextPageCal(page int, totalPage int) int {
	if page == totalPage {
		return page
	}
	return page + 1
}

func PrevPageCal(page int) int {
	if page > 1 {
		return page - 1
	}
	return page
}

func TotalPage(count int64, limit int) int {
	return int(math.Ceil(float64(count) / float64(limit)))
}

func OffsetCal(page int, limit int) int {
	return (page - 1) * limit
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}
	return
}

func RemoveDuplicateValues(intSlice []uint) []uint {
	keys := make(map[uint]bool)
	list := []uint{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ComputeHmacSha1(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func ClearNonAlphanumericalCharacters(val string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(val, ""), nil
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func ByteCountDecimal(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}