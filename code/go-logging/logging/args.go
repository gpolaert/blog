package logging

import (
	"github.com/Sirupsen/logrus"
	"fmt"
)

// We define an object compatible with the targeted logger.
// In that case, we use an Entry from the Logrus library
type ArgsLogger struct {
	entry *logrus.Entry
}

// Here is the Event structure. This struct can be defined in an other file
type Event struct {
	id      int
	message string
}
// A simple method to translate the Event into a beautiful log line
func (e Event)toString() string {
	return fmt.Sprintf("%s%05d - %s", loggerNamePrefix, e.id, e.message)
}

// Inspired from the JBoss Logging Library, all loggers have their code
// Make sure the Logger Name Prefix is unique through all your logger classes
const loggerNamePrefix = "ARGS"

// The main part, we define all messages right here.
// The Event struct is pretty simple. We maintain an Id to be sure to
// retrieve simply all messages once they are logged
var (
	invalidArgMessage = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid arg value: %s => %v"}
	missingArgMessage = Event{3, "Missing arg: %s"}
)

// This function instantiate the logger. The bind our instance with a Logrus one.
// The main goal here is to set the global field like the logger name, but we can imagine
// setting some meta like the user name, the cmd name, etc.
func NewArgLogger(l *logrus.Logger) *ArgsLogger {
	e := l.WithField("logger.name", loggerNamePrefix)
	return &ArgsLogger{e}
}

// And here we were, all log events that can be used in our app
func (l *ArgsLogger)InvalidArg(name string) {
	l.entry.Errorf(invalidArgMessage.toString(), name)
}
func (l *ArgsLogger)InvalidArgValue(name string, value interface{}) {
	l.entry.WithField("arg." + name, value).Errorf(invalidArgValueMessage.toString(), name, value)
}
func (l *ArgsLogger)MissingArg(name string) {
	l.entry.Errorf(missingArgMessage.toString(), name)
}



