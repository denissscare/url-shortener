package random

import (
	"math/rand"
	"time"
)

func NewRandomString(aliasLength int) string {
	var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")
	var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	var alias = make([]rune, aliasLength)

	for i := range alias {
		alias[i] = chars[rnd.Intn(len(chars))]
	}

	return string(alias)
}
