package benchmarks

import (
	"github.com/dyweb/gommon/log"
	"io/ioutil"
)

func newDisabledGommon() *log.Entry {
	logger := log.NewLogger()
	logger.Out = ioutil.Discard
	logger.Level = log.ErrorLevel
	entry := logger.RegisterPkg()
	return entry
}

func newGommon() *log.Entry {
	logger := log.NewLogger()
	logger.Out = ioutil.Discard
	logger.Level = log.DebugLevel
	entry := logger.RegisterPkg()
	return entry
}