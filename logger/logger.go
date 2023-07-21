package logger

import (
	"fmt"
	"os"
	"todo/cfg"
)

var (
	LOGGER = SimpleLogger{logPath: cfg.LogPath}
)

type InfoLog struct {
}

type Logger interface {
	Log(err error)
}

type SimpleLogger struct {
	logPath string
}

func (sl SimpleLogger) Error(err error) {
	sl.Log("[ERROR] " + err.Error())
}

func (sl SimpleLogger) Info(message string) {
	sl.Log("[INFO] " + message)
}

func (sl SimpleLogger) Log(message string) {

	file, fErr := os.OpenFile(sl.logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if fErr != nil {
		fmt.Printf("cant not open file %s, err: %v\n", sl.logPath, fErr)
		return
	}

	if _, wErr := file.Write(append([]byte(message), []byte("\n")...)); wErr != nil {
		fmt.Printf("cant not write in file %s, err: %v\n", sl.logPath, wErr)
		return
	}
}
