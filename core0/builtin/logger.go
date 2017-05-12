package builtin

import (
	"encoding/json"
	"fmt"

	"github.com/g8os/core0/base/pm"
	"github.com/g8os/core0/base/pm/core"
	"github.com/g8os/core0/base/pm/process"
	"github.com/g8os/core0/core0/logger"
	logging "github.com/op/go-logging"
)

type logMgr struct{}

func init() {
	l := (*logMgr)(nil)
	pm.CmdMap["logger.set_level"] = process.NewInternalProcessFactory(l.setLevel)
	pm.CmdMap["logger.reopen"] = process.NewInternalProcessFactory(l.reopen)
}

type LogLevel struct {
	Level string `json:"level"`
}

func (l *logMgr) setLevel(cmd *core.Command) (interface{}, error) {
	var args LogLevel

	if err := json.Unmarshal(*cmd.Arguments, &args); err != nil {
		return nil, err
	}

	level, err := logging.LogLevel(args.Level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %s", args.Level)
	}

	logging.SetLevel(level, "")

	return nil, nil

}

func (l *logMgr) reopen(cmd *core.Command) (interface{}, error) {
	logging.Reset()
	logger.SetupLogging()

	return nil, nil

}
