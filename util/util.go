package util

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
)

var logger *log.Logger

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

// CheckErr check if error occur
func CheckErr(err error) bool {
	if err != nil {
		ErrorLog(err.Error())
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

// InitLog func
func InitLog() {
	logFile, err := os.OpenFile("serverlog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if CheckErr(err) == false {
		logFile.Close()
		return
	}
	logger = log.New(logFile, "INFO ", log.Ldate|log.Ltime)
}

// InfoLog log info
func InfoLog(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// ErrorLog log error
func ErrorLog(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}
