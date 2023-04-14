# Core - Logging

Sets up global logging facilities using `sirupsen/logrus`. Provides global formatting functions to speed up logging messages. Supports the levels `DEBUG`, `INFO`,  and `ERROR`. The logging level dictates that only messages of that level or higher (error being highest) will be logged. Note that no warning level is provided. It is under the assumption that anything that would be a warning should actually be an error and reported as one. The general usage of the logging is to output to a file. The `natefinch/lumberjack` package is used for file rotation, but it should be noted I have not observed this to actually rotate any files. 

In development environments (as provided by the `config` package), the log file will be removed when applications starts.

## Usage

Like most other packages, it must be initialized with the given `config` options using the `log.Initialize()` function. From here, formatting functions are provided. For each log level a variety of `[Level]()`, and `[Level]f()` are available. The difference being like the `fmt` package, the one suffixed with f is for use with formatting strings.

- `Debug(...)` logs a message as debug level
- `Debugf(fmt, ...)` logs a formatted message as debug, using the format string provided
- `Info(...)` logs a message as info level
- `Infof(fmt, ...)` logs a formatted message as info, using the format string provided
- `Error(...)` logs a message as error level
- `Errorf(fmt, ...)` logs a formatted message as error, using the format string provided
- `Err(msg, error)` convenience function that logs an error level message with the msg provided. Applies the error provided as a "field" within the message. **Use this instead of the Error\* functions where possible**