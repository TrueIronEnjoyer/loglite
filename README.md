# loglite

`loglite` is a flexible logging package for Go, providing multiple log levels (Debug, Info, Warning, Error, Fatal) with configuration to enable or disable each level. Logs are saved to a file with timestamps and a unique identifier.

## Installation

To use `loglite`, include it in your Go project. You can copy the package into your project or use it as part of a Go module.

## Usage

### Importing

```go
import "TrueIronEnjoyer/loglite"
```

### Starting the Logger

Initialize the logger and specify the log file name. You can also configure which log levels are enabled.

```go
loglite.Loger.Mods = loglite.mods{
    Debug:   true,
    Info:    true,
    Warning: true,
    Error:   true,
    Fatal:   true,
}
loglite.Start("path/to/your/logfile.log")
```

### Logging Messages

Use the provided functions to log messages at different levels. Logging will only occur if the corresponding log level is enabled.

```go
loglite.Debug("This is a debug message")
loglite.Info("This is an info message")
loglite.Warning("This is a warning message")
loglite.Error("This is an error message")
loglite.Fatal("This is a fatal message")
```

### Example

Here’s a simple example of how to use `loglite` in your application:

```go
package main

import (
	"loglite"
)

func main() {
	// Configure log levels
	loglite.Loger.Mods = loglite.Mods{
		Debug:   true,
		Info:    true,
		Warning: true,
		Error:   true,
		Fatal:   true,
	}

	// Initialize the logger
	loglite.Start("app.log")

	// Log messages at different levels
	loglite.Debug("Debug message with %d arguments", 1)
	loglite.Info("Info message with %d arguments", 1)
	loglite.Warning("Warning message with %d arguments", 1)
	loglite.Error("Error message with %d arguments", 1)
	loglite.Fatal("Fatal message with %d arguments", 1)
}
```

## Features

- **Configurable Log Levels**: Enable or disable log levels (Debug, Info, Warning, Error, Fatal).
- **File Logging**: Logs are written to a specified file with a unique ID and timestamp.

## Contributing

Contributions are welcome! If you find a bug or have suggestions for improvements, please open an issue or submit a pull request.

## License

This package is released under the [MIT License](LICENSE).

## Author

TrueIronEnjoyer