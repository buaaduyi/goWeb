package util

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

// CheckErr check if error occur
func CheckErr(err error) bool {
	if err != nil {
		ColorPrintf(err.Error()+"\n", Red)
		return false
	}
	return true
}

// MD5Code encrypto
func MD5Code(message string) (key string) {
	b := md5.Sum([]byte(message))
	for i := 0; i < 16; i++ {
		key += strconv.Itoa(int(b[i]))
	}
	return
}

const (
	//Red color
	Red = 31
	//Green color
	Green = 32
	//Yellow color
	Yellow = 33
	//Blue color
	Blue = 34
)

// ColorPrintf print message with color
func ColorPrintf(message string, color int) {
	fmt.Printf("%c[1;10;%dm%s%c[0m", 0x1B, color, message, 0x1B)
}
