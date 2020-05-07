package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"red-east/config"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
)

var (
	// FileLogger *rotatelogs.RotateLogs
	Writers []io.Writer
)

type NLogger struct {
	info   *log.Logger
	errors *log.Logger
	debug  *log.Logger
	waring *log.Logger
}

func InitLogger() (NLogger, error) {
	//获取配置，因为在同一个utils，不能直接引用
	conf, err := config.InitConfig()
	if err != nil {
		return NLogger{}, err
	}
	logConfig := conf.Logging
	flags := log.Ldate | log.LstdFlags | log.Lshortfile
	//日志需要输出到哪里，标准输出是要的，输出到文件则看配置
	err = InitWriters(logConfig)
	if err != nil {
		return NLogger{}, err
	}
	logger := NLogger{
		log.New(io.MultiWriter(Writers...), "[INFO]", flags),
		log.New(io.MultiWriter(Writers...), "[ERROR]", flags),
		log.New(io.MultiWriter(Writers...), "[DEBUG]", flags),
		log.New(io.MultiWriter(Writers...), "[WARING]", flags),
	}
	return logger, nil
}

func InitWriters(logConfig config.LoggingConfig) error {
	Writers = append(Writers, os.Stdout)
	fileLogger, err := initFileLogger(logConfig)
	if err != nil {
		log.Println("create file logger fail:", err.Error())
		return err
	}
	if fileLogger != nil {
		Writers = append(Writers, fileLogger)
	}
	return nil
}

func initFileLogger(logConfig config.LoggingConfig) (fileLogger *rotatelogs.RotateLogs, err error) {
	if logConfig.FileWrite {
		path := logConfig.FilePath
		fileLogger, err = rotatelogs.New(
			path+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(path),                                 // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(logConfig.FileMaxAge*time.Hour),         // 文件最大保存时间
			rotatelogs.WithRotationTime(logConfig.RotationTime*time.Hour), // 日志切割时间间隔
		)
		if err != nil {
			return nil, err
		}
		return fileLogger, nil
	}
	return nil, nil
}

func (l *NLogger) Info(v ...interface{}) {
	l.info.Output(2, fmt.Sprintln(v...))
}

func (l *NLogger) Infof(format string, v ...interface{}) {
	l.info.Output(2, fmt.Sprintf(format, v...))
}

func (l *NLogger) Error(v ...interface{}) {
	l.errors.Output(2, fmt.Sprintln(v...))
}

func (l *NLogger) Debug(v ...interface{}) {
	l.debug.Output(2, fmt.Sprintln(v...))
}

func (l *NLogger) Debugf(format string, v ...interface{}) {
	l.info.Output(2, fmt.Sprintf(format, v...))
}

func (l *NLogger) Waring(v ...interface{}) {
	l.waring.Output(2, fmt.Sprintln(v...))
}

func (l *NLogger) GetWriter() []io.Writer {
	return Writers
}
