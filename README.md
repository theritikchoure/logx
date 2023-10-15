# logx: Colorful Logging for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/theritikchoure/logx)](https://goreportcard.com/report/github.com/theritikchoure/logx)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/theritikchoure/logx/blob/main/LICENSE)
[![GitHub Release](https://img.shields.io/github/v/release/theritikchoure/logx)](https://github.com/theritikchoure/logx/releases)
[![Build Status](https://travis-ci.org/theritikchoure/logx.svg?branch=main)](https://github.com/theritikchoure/logx/actions)
[![GoDoc](https://pkg.go.dev/badge/github.com/theritikchoure/logx)](https://pkg.go.dev/github.com/theritikchoure/logx)
[![Coverage Status](https://coveralls.io/repos/github/theritikchoure/logx/badge.svg?branch=main)](https://coveralls.io/github/theritikchoure/logx?branch=main)


`logx` is a Go package that allows you to add color and formatting to your console log messages, making it easier to distinguish different types of log entries or to add emphasis to specific messages. It provides a simple way to enhance your application's log output.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Enabling or Disabling Color](#enabling-or-disabling-color)
  - [Log](#log)
  - [Logf](#logf)
  - [LogWithLevel](#logwithlevel)
  - [LogM](#logm)
  - [LogWithTimestamp](#logwithtimestamp)
  - [LogToFile](#logtofile)
- [Guide](#guide)
- [Contributing](#contributing)
- [License](#license)

## Installation

To use `logx`, you need to import it in your Go code:

```go
import "github.com/theritikchoure/logx"
```

Then, install it with `go get`:

```go
go get github.com/theritikchoure/logx
```

## Usage

### Enabling or Disabling Color
By default, colorized logging is disabled. You can enable or disable it by modifying the `ColoringEnabled` variable in your code. Setting it to `true` will enable colored output, and setting it to `false` will disable it.

```go
logx.ColoringEnabled = true // Enable colorized logging
logx.ColoringEnabled = false // Disable colorized logging (default)
```

### Log
The `Log` function allows you to print log messages with specified foreground and background colors. If coloring is disabled, it will print plain text.

```go
logx.Log("This is a log message", logx.FGRED, logx.BGGREEN)
```

### Logf
Use `Logf` to format log messages and specify colors. It is similar to `fmt.Printf`.

```go
logx.Logf("User: %s logged in.", logx.FGBLUE, logx.BGWHITE, "JohnDoe")
```

### LogWithLevel
`LogWithLevel` is helpful for displaying log messages with predefined background colors corresponding to log levels such as INFO, WARNING, ERROR, and SUCCESS.

```go
logx.LogWithLevel("An error occurred", "ERROR")
```

### LogM
With `LogM`, you can log multiple messages with customized formatting options.

```go
logx.LogM([]string{"Applying", "updates..."}, logx.FGBLUE, logx.BGYELLOW, logx.BOLD, logx.UNDERLINE)
```

### LogWithTimestamp
`LogWithTimestamp` adds a timestamp to your log message. It's particularly useful for recording when an event occurred.
```go
logx.LogWithTimestamp("System restarted", logx.FGCYAN, logx.BGWHITE)
```

### LogToFile
`LogToFile` logs messages to a file in addition to standard output.

```go
logx.LogToFile("Log entry written to file", logx.FGRED, logx.BGGREEN, "log.txt")
```

## Guide

- [LogX — Enhance Your GoLang Projects with Colorful Logging](https://ritikchourasiya.medium.com/logx-enhance-your-golang-projects-with-colorful-logging-1f93df825cd8)

## Contributing
We welcome contributions from the community! If you'd like to contribute to logx, follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and add tests if applicable.
4. Run tests and ensure they pass.
5. Submit a pull request with a clear description of your changes.

We appreciate your help in improving logx.

## License
This project is licensed under the MIT License. See the [LICENSE](https://github.com/theritikchoure/logx/blob/main/LICENSE) file for details.