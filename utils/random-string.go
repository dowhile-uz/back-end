package utils

import (
	"crypto/rand"
	"encoding/hex"
	"math"
)

// https://stackoverflow.com/a/55860599
func RandomBase16String(l int) string {
	buff := make([]byte, int(math.Ceil(float64(l)/2)))
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	return str[:l]
}
