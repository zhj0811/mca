package config

import (
	"io"
	"os"
)

type GlobalConfig struct{
	writer io.Writer
}

var globalConf GlobalConfig

func init() {
	globalConf.writer = os.Stdout
}

func GetDefaultLogWriter() io.Writer {
	return globalConf.writer
}