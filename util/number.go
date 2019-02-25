package util

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomNumber(maxLength int) string {
	rand.Seed(time.Now().UnixNano())

	var buffer bytes.Buffer
	for i := 0; i < maxLength; i++ {
		str := strconv.Itoa(rand.Intn(10))
		buffer.WriteString(str)
	}

	return buffer.String()
}
