package log4

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/idoall/TokenExchangeCommon/commonutils"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type apiFileLogger struct {
	LogFileWrite *logrus.Logger
	LogOutPut    *logrus.Logger
}

// NewFileLogger New Creates a new logger with a "file" writer to send
// log messages at or above lvl to standard output.
func NewFileLogger(fileName string) Logger4 {
	var log = logrus.New()

	// -------- 判断目录是否存在 -------- Begin
	// 判断当前应用程序目录,在用 go run 命令执行时，一般在一个临时目录
	// if AppPath, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("AppPath", AppPath)
	// }

	// 获取当前工作目录
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 设置 logs 输出文件目录
	logFolderPath := filepath.Join(workPath, "logs")

	// 如果目录不存在，则创建
	if !commonutils.PathExists(logFolderPath) {
		if err = os.Mkdir(logFolderPath, os.ModePerm); err != nil {
			panic(err)
		}
	}
	// -------- 判断目录是否存在 -------- End

	baseLogPath := fmt.Sprintf("logs/%s", fileName)
	newLogName := baseLogPath[:strings.LastIndex(baseLogPath, ".")]
	fileExe := baseLogPath[strings.LastIndex(baseLogPath, ".")+1:]
	writer, _ := rotatelogs.New(
		newLogName+".%F."+fileExe,
		rotatelogs.WithClock(rotatelogs.Local),    // 使用本地时区
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	// 输出到文本的格式
	logFormat := &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true, DisableTimestamp: false, DisableColors: false, ForceColors: true, DisableSorting: false}
	// logFormat := &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05", DisableTimestamp: false}

	// 参考:http://xiaorui.cc/2018/01/11/golang-logrus%E7%9A%84%E9%AB%98%E7%BA%A7%E9%85%8D%E7%BD%AEhook-logrotate/
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, logFormat)
	log.AddHook(lfHook)

	// 输出到屏幕的格式
	log.Formatter = &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true, DisableTimestamp: false, DisableColors: false, ForceColors: true, DisableSorting: false}
	log.SetLevel(logrus.DebugLevel)
	log.Out = os.Stdout
	// log.SetFormatter(logFormat)

	return &apiFileLogger{
		LogFileWrite: log,
		LogOutPut:    nil,
	}

}

func (e *apiFileLogger) Close() {
	e.LogFileWrite.SetNoLock()
	e.LogOutPut.SetNoLock()
}

func (e *apiFileLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	return e.LogFileWrite.WithFields(fields)
}

func (e *apiFileLogger) WithError(err error) *logrus.Entry {
	return e.LogFileWrite.WithError(err)
}

func (e *apiFileLogger) Fatal(args ...interface{}) {
	e.LogFileWrite.Fatal(args)
}
func (e *apiFileLogger) Fatalf(format string, args ...interface{}) {
	e.LogFileWrite.Fatalf(format, args)
}

func (e *apiFileLogger) Debug(args ...interface{}) {
	e.LogFileWrite.Debug(args)
}
func (e *apiFileLogger) Debugf(format string, args ...interface{}) {
	e.LogFileWrite.Debugf(format, args)
}

func (e *apiFileLogger) Warning(args ...interface{}) {
	e.LogFileWrite.Warning(args)
}
func (e *apiFileLogger) Warningf(format string, args ...interface{}) {
	e.LogFileWrite.Warningf(format, args)
}

func (e *apiFileLogger) Info(args ...interface{}) {
	e.LogFileWrite.Info(args)
}
func (e *apiFileLogger) Infof(format string, args ...interface{}) {
	e.LogFileWrite.Infof(format, args)
}

func (e *apiFileLogger) Error(args ...interface{}) {
	e.LogFileWrite.Error(args)
}
func (e *apiFileLogger) Errorf(format string, args ...interface{}) {
	e.LogFileWrite.Errorf(format, args)
}
