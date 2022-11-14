package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/wezhai/kubesphere-webhook-proxy-go/config"
)

var Logger = logrus.New()

var Entry = logrus.NewEntry(Logger)

var Error = Entry.Error
var Errorf = Entry.Errorf
var Errorln = Entry.Errorln

var Info = Entry.Info
var Infof = Entry.Infof

var Print = Entry.Info
var Printf = Entry.Infof
var Println = Entry.Println

var Debug = Entry.Debug
var Debugf = Entry.Debugf
var Debugln = Entry.Debugln

var Panicf = Entry.Panicf
var Panic = Entry.Panic

var Trace = Entry.Trace
var Tracef = Entry.Tracef

var Warn = Entry.Warn
var Warnf = Entry.Warnf

var Fatal = Entry.Fatal
var Fatalf = Entry.Fatalf

func init() {
	if config.Config.DebugMode {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}
	Logger.SetReportCaller(true) // 显示调用者
	Logger.SetFormatter(&nested.Formatter{
		// TimestampFormat: time.RFC3339,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerFirst:     true, // 调用显示在第一列
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			funcInfo := runtime.FuncForPC(frame.PC)
			if funcInfo == nil {
				return "error during runtime.FuncForPC"
			}
			fullPath, line := funcInfo.FileLine(frame.PC)
			return fmt.Sprintf(" [%v:%v]", filepath.Base(fullPath), line)
		},
	})
	// Logger.SetFormatter(&logrus.JSONFormatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })

	fileName := "app.log"
	stdoutWriter := os.Stdout
	fileWriter, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Create file %v failed: %v", fileName, err)
		os.Exit(1)
	}
	Logger.SetOutput(io.MultiWriter(stdoutWriter, fileWriter))
}
