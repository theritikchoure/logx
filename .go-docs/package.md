package logx // import "github.com/theritikchoure/logx"


CONSTANTS

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
const (
	DEBUG = BGCYAN // Define the background color for debug log messages.
	FATAL = BGRED  // Define the background color for fatal log messages.
)

VARIABLES

var ColoringEnabled = true // coloring is enabled for log messages by default.

FUNCTIONS

func Debug(format string, args ...interface{})
func Fatal(format string, args ...interface{})
func Log(message string, fgColor string, bgColor string)
func LogM(messages []string, fgColor string, bgColor string, formatOptions ...string)
func LogToFile(message string, fgColor string, bgColor string, filename string)
func LogWithLevel(message string, logLevel string)
func LogWithLevelAndTimestamp(message string, logLevel string, fgColor string, bgColor string)
func LogWithTimestamp(message string, fgColor string, bgColor string)
func Logf(format string, fgColor string, bgColor string, args ...interface{})
