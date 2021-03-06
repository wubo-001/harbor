package logger

import "strings"

const (
	// LoggerNameFile is unique name of the file logger.
	LoggerNameFile = "FILE"
	// LoggerNameStdOutput is the unique name of the std logger.
	LoggerNameStdOutput = "STD_OUTPUT"
)

// Declaration is used to declare a supported logger.
// Use this declaration to indicate what logger and sweeper will be provided.
type Declaration struct {
	Logger  Factory
	Sweeper SweeperFactory
	Getter  GetterFactory
	// Indicate if the logger is a singleton logger
	Singleton bool
}

// knownLoggers is a static logger registry.
// All the implemented loggers (w/ sweeper) should be registered
// with an unique name in this registry. Then they can be used to
// log info.
var knownLoggers = map[string]*Declaration{
	// File logger
	LoggerNameFile: {FileFactory, FileSweeperFactory, FileGetterFactory, false},
	// STD output(both stdout and stderr) logger
	LoggerNameStdOutput: {StdFactory, nil, nil, true},
}

// IsKnownLogger checks if the logger is supported with name.
func IsKnownLogger(name string) bool {
	_, ok := knownLoggers[name]

	return ok
}

// HasSweeper checks if the logger with the name provides a sweeper.
func HasSweeper(name string) bool {
	d, ok := knownLoggers[name]

	return ok && d.Sweeper != nil
}

// HasGetter checks if the logger with the name provides a log data getter.
func HasGetter(name string) bool {
	d, ok := knownLoggers[name]

	return ok && d.Getter != nil
}

// KnownLoggers return the declaration by the name
func KnownLoggers(name string) *Declaration {
	return knownLoggers[name]
}

// All known levels which are supported.
var debugLevels = []string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"FATAL",
}

// IsKnownLevel is used to check if the logger level is supported.
func IsKnownLevel(level string) bool {
	if len(level) == 0 {
		return false
	}

	for _, lvl := range debugLevels {
		if lvl == strings.ToUpper(level) {
			return true
		}
	}

	return false
}
