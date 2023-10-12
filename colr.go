package colr

import "fmt"

const (
	RESET   = `\x1b[0m`
	BLACK   = `\x1b[30m`
	RED     = `\x1b[31m`
	GREEN   = `\x1b[32m`
	YELLOW  = `\x1b[33m`
	BLUE    = `\x1b[34m`
	MAGENTA = `\x1b[35m`
	CYAN    = `\x1b[36m`
	WHITE   = `\x1b[37m`
)

func Log(message string) {
	fmt.Println("\x1b[31mRed Text\x1b[0m")
}
