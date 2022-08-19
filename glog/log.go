package glog

import (
	"sync"

	"github.com/sirupsen/logrus"
)

func init() {

}

var (
	logger *logrus.Logger
	level  LEVEL
	mu     sync.Mutex
)

type LEVEL int

const (
	DEBUG LEVEL = iota
	INFO
	WARN
	ERROR
	NONE
)

func SetLogger(l *logrus.Logger) {
	logger = l
}

func SetLevel(l LEVEL) {
	level = l
}

func checkLogger() {
	if logger == nil && level != NONE {
		panic("logger not inited")
	}
}

func Debug(v ...interface{}) {
	checkLogger()
	if level <= DEBUG {
		mu.Lock()
		defer mu.Unlock()
		//logger.SetPrefix("[DEBUG]")
		//logger.Output(2, fmt.Sprintln(v...))
		logger.Debug(v...)
	}
}
func Debugf(f string, v ...interface{}) {
	checkLogger()
	if level <= DEBUG {
		mu.Lock()
		defer mu.Unlock()
		//logger.SetPrefix("[DEBUG]")
		//logger.Output(2, fmt.Sprintln(fmt.Sprintf(f, v...)))
		logger.Debug(v...)
	}
}
func Info(v ...interface{}) {
	checkLogger()
	if level <= INFO {
		mu.Lock()
		defer mu.Unlock()
		//logger.SetPrefix("[INFO]")
		//logger.Output(2, fmt.Sprintln(v...))
		logger.Info(v...)
	}
}
func Infof(f string, v ...interface{}) {
	checkLogger()
	if level <= INFO {
		mu.Lock()
		defer mu.Unlock()
		//logger.SetPrefix("[INFO]")
		//logger.Output(2, fmt.Sprintln(fmt.Sprintf(f, v...)))
		logger.Info(v...)
	}
}
func Warn(v ...interface{}) {
	checkLogger()
	if level <= WARN {
		mu.Lock()
		defer mu.Unlock()
		//logger.SetPrefix("[WARN]")
		//logger.Output(2, fmt.Sprintln(v...))
		logger.Warn(v...)
	}
}

func Error(v ...interface{}) {
	checkLogger()
	if level <= ERROR {
		mu.Lock()
		defer mu.Unlock()
		//logger.SetPrefix("[ERROR]")
		//logger.Output(2, fmt.Sprintln(v...))
		logger.Error(v...)
	}
}
func Errorf(f string, v ...interface{}) {
	checkLogger()
	if level <= ERROR {
		mu.Lock()
		defer mu.Unlock()
		//logger.SetPrefix("[ERROR]")
		//logger.Output(2, fmt.Sprintln(fmt.Sprintf(f, v...)))
		logger.Error(v...)
	}
}
