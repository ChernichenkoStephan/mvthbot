package logging

import (
	"sync"

	"go.uber.org/zap"
)

type Options struct {
	LogFileDir    string //Log path
	AppName       string // Filename is the prefix of the file to be written to the log
	ErrorFileName string
	WarnFileName  string
	InfoFileName  string
	DebugFileName string
	MaxSize       int // How many m of a file is greater than this number to start file segmentation
	MaxBackups    int // MaxBackups is the maximum number of old log files to keep
	MaxAge        int // MaxAge is the maximum number of days old log files are retained by date
	zap.Config
}

type Logger struct {
	*zap.SugaredLogger
	sync.RWMutex
	Opts      *Options `json:"opts"`
	zapConfig zap.Config
	inited    bool
}
