package logging

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"github.com/Sirupsen/logrus"
)

var mockIO bytes.Buffer

type MockFormatter struct {
}

func (f *MockFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte("message=" + entry.Message), nil
}

var log *ArgsLogger

func init() {
	// Mock the default writer
	l := logrus.New()
	l.Formatter = &MockFormatter{}
	l.Out = &mockIO

	log = NewArgLogger(l)
}

func Test_MessageLogged(t *testing.T) {

	log.InvalidArg("argname")
	assert.Equal(t, "message=ARGS00001 - Invalid arg: argname", mockIO.String())
	mockIO.Reset()

	log.InvalidArgValue("argname", "argvalue")
	assert.Equal(t, "message=ARGS00002 - Invalid arg value: argname => argvalue", mockIO.String())
	mockIO.Reset()

	log.MissingArg("argname")
	assert.Equal(t, "message=ARGS00003 - Missing arg: argname", mockIO.String())
	mockIO.Reset()
}