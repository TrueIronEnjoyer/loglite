# loglite

`loglite` is a Go logging package that provides multiple log levels (Debug, Info, Warning, Error, Fatal) and logs messages to a file with timestamps and a unique identifier for each entry.

## Features

- **Multiple Log Levels**: Debug, Info, Warning, Error, and Fatal.
- **File Logging**: Logs are written to a specified file with a unique ID and timestamp.

## Installation

To use `loglite`, you can include the package in your Go project by copying the code or integrating it into your Go module.

## Summary of Go File

This Go file implements a logging library called `loglite` with the following key components:

### Types and Constants

- **`loger` struct**: Handles the logger's state, including log level, file path, file handler, and internal counters.
- **`log` struct**: Represents a log entry with a unique ID, log level, message, and timestamp.
- **Log Levels**: Constants `Debug`, `Info`, `Warning`, and `Error` control log severity.

### Methods

- **`NewLogger(lvl int, filePath string) *loger`**: Initializes a new logger, sets the log level, and opens/creates a log file.

#### Logging Methods

- **`Debug(msg string, args ...any)`**: Logs debug messages.
- **`Info(msg string, args ...any)`**: Logs informational messages.
- **`Warning(msg string, args ...any)`**: Logs warning messages.
- **`Error(msg string, args ...any)`**: Logs error messages.
- **`Fatal(msg string, args ...any)`**: Logs a fatal error and stops the logger.

### Internal Functions

- **`newLog(lvl, msg string, args ...any) log`**: Creates a new log entry with a unique ID, log level, and formatted message.
- **`save(lg log) string`**: Writes the log entry to the file.

## Usage

### Importing

```go
import "path/to/loglite"
```

### Starting the Logger

Initialize the logger by specifying the log level and the file path where the logs will be saved:

```go
logger := loglite.NewLogger(loglite.Debug, "app.log")
```

### Logging Messages

Use the provided methods to log messages at different levels:

```go
logger.Debug("This is a debug message")
logger.Info("This is an info message")
logger.Warning("This is a warning message")
logger.Error("This is an error message")
logger.Fatal("This is a fatal message")
```

## Example

Here’s an example of how to use `loglite` in your application:

```go
package main

import "github.com/TrueIronEnjoyer/loglite"

func main() {
	// Initialize the logger
	logger := loglite.NewLogger(loglite.Debug, "app.log")

	// Log messages at different levels
	logger.Debug("Debug message with %d arguments", 1)
	logger.Info("Info message with %d arguments", 1)
	logger.Warning("Warning message with %d arguments", 1)
	logger.Error("Error message with %d arguments", 1)
	logger.Fatal("Fatal message with %d arguments", 1)
}
```

## License

This package is released under the [MIT License](LICENSE).

## Author

TrueIronEnjoyer
