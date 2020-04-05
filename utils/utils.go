package utils

import "fmt"

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

// ColorPrintf print message with color to terminal
func ColorPrintf(str string, col int) {
	fmt.Printf("%c[1;10;%dm%s%c[0m", 0x1B, col, str, 0x1B)
}

// ErrOccur check error occur or not
func ErrOccur(err error) bool {
	if err != nil {
		ColorPrintf(err.Error()+"\n", Red)
		return true
	}
	return false
}
