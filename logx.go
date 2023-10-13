package logx

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	RESET = "\x1b[0m"

	// Foreground text colors
	FGBLACK   = "\x1b[30m"
	FGRED     = "\x1b[31m"
	FGGREEN   = "\x1b[32m"
	FGYELLOW  = "\x1b[33m"
	FGBLUE    = "\x1b[34m"
	FGMAGENTA = "\x1b[35m"
	FGCYAN    = "\x1b[36m"
	FGWHITE   = "\x1b[37m"

	// Background colors
	BGBLACK   = "\x1b[40m"
	BGRED     = "\x1b[41m"
	BGGREEN   = "\x1b[42m"
	BGYELLOW  = "\x1b[43m"
	BGBLUE    = "\x1b[44m"
	BGMAGENTA = "\x1b[45m"
	BGCYAN    = "\x1b[46m"
	BGWHITE   = "\x1b[47m"
)

const (
	INFO    = BGBLUE
	WARNING = BGYELLOW
	ERROR   = BGRED
	SUCCESS = BGGREEN
)

const (
	// Other text formatting options
	BOLD      = "\x1b[1m"
	UNDERLINE = "\x1b[4m"
	BLINK     = "\x1b[5m"
)

var ColoringEnabled = false

func Log(message string, fgColor string, bgColor string) {

	if ColoringEnabled {
		if fgColor == "" {
			fgColor = FGWHITE
		}

		coloredMessage := fgColor + bgColor + message + RESET
		fmt.Println(coloredMessage)
	} else {
		fmt.Println(message)
	}
}

func Logf(format string, fgColor string, bgColor string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)

	Log(message, fgColor, bgColor)
}

func LogWithLevel(message string, logLevel string) {
	bgColor := INFO

	switch logLevel {
	case "INFO":
		bgColor = INFO
	case "WARNING":
		bgColor = WARNING
	case "ERROR":
		bgColor = ERROR
	case "SUCCESS":
		bgColor = SUCCESS
	}

	Log(message, "", bgColor)
}

func LogM(messages []string, fgColor string, bgColor string, formatOptions ...string) {
	// Combine color codes and formatting options
	formatting := fgColor + bgColor + strings.Join(formatOptions, "") + RESET

	// Join multiple messages and apply formatting
	formattedMessage := formatting + strings.Join(messages, " ") + RESET

	Log(formattedMessage, "", "")
}

func LogWithTimestamp(message string, fgColor string, bgColor string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formattedMessage := timestamp + " " + message
	Log(formattedMessage, fgColor, bgColor)
}

func LogToFile(message string, fgColor string, bgColor string, filename string) {
	// Open the file for appending
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	// Log to both the file and standard output
	writer := io.MultiWriter(os.Stdout, file)
	log := log.New(writer, "", 0)
	log.Print(message)
}
