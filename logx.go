package logx

import (
	"fmt"     // Import the "fmt" package for formatting and printing.
	"io"      // Import the "io" package for I/O operations.
	"log"     // Import the "log" package for creating loggers.
	"os"      // Import the "os" package for file operations.
	"strings" // Import the "strings" package for string manipulation.
	"time"    // Import the "time" package for time-related operations.
)

const (
	RESET = "\x1b[0m" // Define the ANSI escape code for resetting text attributes.

	// Foreground text colors
	FGBLACK   = "\x1b[30m" // Define the ANSI escape code for black text.
	FGRED     = "\x1b[31m" // Define the ANSI escape code for red text.
	FGGREEN   = "\x1b[32m" // Define the ANSI escape code for green text.
	FGYELLOW  = "\x1b[33m" // Define the ANSI escape code for yellow text.
	FGBLUE    = "\x1b[34m" // Define the ANSI escape code for blue text.
	FGMAGENTA = "\x1b[35m" // Define the ANSI escape code for magenta text.
	FGCYAN    = "\x1b[36m" // Define the ANSI escape code for cyan text.
	FGWHITE   = "\x1b[37m" // Define the ANSI escape code for white text.

	// Background colors
	BGBLACK   = "\x1b[40m" // Define the ANSI escape code for a black background.
	BGRED     = "\x1b[41m" // Define the ANSI escape code for a red background.
	BGGREEN   = "\x1b[42m" // Define the ANSI escape code for a green background.
	BGYELLOW  = "\x1b[43m" // Define the ANSI escape code for a yellow background.
	BGBLUE    = "\x1b[44m" // Define the ANSI escape code for a blue background.
	BGMAGENTA = "\x1b[45m" // Define the ANSI escape code for a magenta background.
	BGCYAN    = "\x1b[46m" // Define the ANSI escape code for a cyan background.
	BGWHITE   = "\x1b[47m" // Define the ANSI escape code for a white background.
)

const (
	INFO    = BGBLUE   // Define the background color for informational log messages.
	WARNING = BGYELLOW // Define the background color for warning log messages.
	ERROR   = BGRED    // Define the background color for error log messages.
	SUCCESS = BGGREEN  // Define the background color for success log messages.
)

const (
	// Other text formatting options
	BOLD      = "\x1b[1m" // Define the ANSI escape code for bold text.
	UNDERLINE = "\x1b[4m" // Define the ANSI escape code for underlined text.
	BLINK     = "\x1b[5m" // Define the ANSI escape code for blinking text.
)

var ColoringEnabled = false // Create a variable to control whether coloring is enabled for log messages.

func Log(message string, fgColor string, bgColor string) {
	// Log function for printing log messages with specified text and background colors.
	// If coloring is enabled, it will format the message with the specified colors.
	// If coloring is disabled, it will print the message in plain text.
	if ColoringEnabled {
		if fgColor == "" {
			fgColor = FGWHITE // Set the default text color to white if not specified.
		}
		coloredMessage := fgColor + bgColor + message + RESET // Apply color codes to the message.
		fmt.Println(coloredMessage)                           // Print the colored message.
	} else {
		fmt.Println(message) // Print the message in plain text.
	}
}

func Logf(format string, fgColor string, bgColor string, args ...interface{}) {
	// Logf function for formatting log messages and specifying text and background colors.
	// Similar to fmt.Printf, it formats the message and calls Log to print it.
	message := fmt.Sprintf(format, args...) // Format the message using Printf format and arguments.
	Log(message, fgColor, bgColor)          // Call the Log function to print the formatted message.
}

func LogWithLevel(message string, logLevel string) {
	// LogWithLevel function for displaying log messages with predefined background colors
	// corresponding to log levels such as INFO, WARNING, ERROR, and SUCCESS.
	bgColor := INFO // Set the default background color to INFO.

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

	Log(message, "", bgColor) // Call the Log function to print the message with the selected background color.
}

func LogM(messages []string, fgColor string, bgColor string, formatOptions ...string) {
	// LogM function for logging multiple messages with customized formatting options.
	// It combines color codes and formatting options, joins multiple messages, and applies the formatting.
	// The resulting message is then passed to the Log function for printing.
	formatting := fgColor + bgColor + strings.Join(formatOptions, "") + RESET // Combine color codes and formatting options.
	formattedMessage := formatting + strings.Join(messages, " ") + RESET      // Join messages and apply formatting.
	Log(formattedMessage, "", "")                                             // Call the Log function to print the formatted message.
}

func LogWithTimestamp(message string, fgColor string, bgColor string) {
	// LogWithTimestamp function for adding a timestamp to log messages.
	// It formats the message with a timestamp and passes it to the Log function for printing.
	timestamp := time.Now().Format("2006-01-02 15:04:05") // Format the current time as a timestamp.
	formattedMessage := timestamp + " " + message         // Combine the timestamp and the message.
	Log(formattedMessage, fgColor, bgColor)               // Call the Log function to print the timestamped message.
}

func LogToFile(message string, fgColor string, bgColor string, filename string) {
	// LogToFile function for logging messages to a file in addition to standard output.
	// It opens the specified file for appending and logs the message to both the file and standard output.
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Open the file for appending.
	if err != nil {
		fmt.Println("Error opening log file:", err) // Print an error message if file opening fails.
		return
	}
	defer file.Close() // Close the file when the function returns.

	writer := io.MultiWriter(os.Stdout, file) // Create a writer that writes to both stdout and the file.
	log := log.New(writer, "", 0)             // Create a logger with the multi-writer.
	log.Print(message)                        // Log the message to both stdout and the file.
}
