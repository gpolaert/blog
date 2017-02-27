package main

import (
	"github.com/logmatic/walter/logging"
	"github.com/Sirupsen/logrus"
)

func main() {

	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{}

	// a new instance of my Logger bind to the "l" Logrus logger
	argLogger := logging.NewArgLogger(l)

	//opts := os.Args
	// checking your opts ...

	// And logs all things that go wrong :)
	// an opt is missing ...
	argLogger.MissingArg("fizz")

	// here, the value is not valid ...
	argLogger.InvalidArgValue("buzz", 12)


	/// OUTPUT
	// output - plain text
	//time="2017-02-24T23:12:31+01:00" level=error msg="ARGS00003 - Missing arg: fizz" arg.fizz=<nil> logger.name=ARGS
	//time="2017-02-24T23:12:31+01:00" level=error msg="ARGS00002 - Invalid arg value: buzz => 12" arg.buzz=12 logger.name=ARGS

	// output - as json
	//{"arg.fizz":null,"level":"error","logger.name":"ARGS","msg":"ARGS00003 - Missing arg: fizz","time":"2017-02-24T23:14:28+01:00"}
	//{"arg.buzz":12,"level":"error","logger.name":"ARGS","msg":"ARGS00002 - Invalid arg value: buzz =\u003e 12","time":"2017-02-24T23:14:28+01:00"}

}
