package shell

import (
	"os"

	"github.com/zerodha/logf"
)

var logger logf.Logger
var debug bool

func init() {

	debug = os.Getenv("DEBUG") != ""

	level := logf.InfoLevel
	if debug {
		level = logf.DebugLevel
	}
	logger = logf.New(logf.Opts{
		EnableColor: true,
		Level:       level,
	})

}
