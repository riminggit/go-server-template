package logging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go-server-template/config"
	"os"
	"path"
	"time"
)

var Logger *logrus.Logger

//初始化log
func InitLogging() *logrus.Logger {
	config := projectConfig.AppConfig

	// 日志文件
	fileName := path.Join(config.LogConfig.LOG_SAVE_URL, config.LogConfig.LOG_SAVE_NAME)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	Logger = logrus.New()

	// Logger.SetReportCaller(true) // 显示行号等信息

	//设置日志级别
	Logger.SetLevel(logrus.DebugLevel)
	//设置输出
	Logger.Out = src

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return Logger
}

func WriteMiddlewareLogs(c *gin.Context) {
	//开始时间
	startTime := time.Now()
	//处理请求
	c.Next()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)
	//请求方式
	reqMethod := c.Request.Method
	//请求路由
	reqUrl := c.Request.RequestURI
	//状态码
	statusCode := c.Writer.Status()
	//请求ip
	clientIP := c.ClientIP()

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"status_code":  statusCode,
		"latency_time": latencyTime,
		"client_ip":    clientIP,
		"req_method":   reqMethod,
		"req_uri":      reqUrl,
	}).Info()
}

// 程序运行正常的日志类型
// ALL: 最低等级的，用于打开所有日志记录。
// DEBUG: 主要用于开发过程中打印一些运行信息,指出细粒度信息事件对调试应用程序是非常有帮助的。
// INFO: 通常用于打印一些用户感兴趣的信息，这个可以用于生产环境中输出程序运行的一些重要信息，但是不能滥用，避免打印过多的日志。

// 代表程序运行错误的日志类型
// WARN: 表明会出现潜在错误的情形，有些信息不是错误信息，但是也要给程序员的一些提示。
// ERROR: 指出虽然发生错误事件，但仍然不影响系统的继续运行。打印错误和异常信息，如果不想输出太多的日志，可以使用这个级别。
// FATAL: 指出每个严重的错误事件将会导致应用程序的退出。这个级别比较高了。重大错误，这种级别你可以直接停止程序了。

// Debug output logs at debug level
func Debug(v ...interface{}) {
	//开始时间
	startTime := time.Now()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Debug(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	//开始时间
	startTime := time.Now()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Info(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	//开始时间
	startTime := time.Now()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Warn(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	//开始时间
	startTime := time.Now()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Error(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	//开始时间
	startTime := time.Now()
	//结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 日志格式
	Logger.WithFields(logrus.Fields{
		"latency_time": latencyTime,
	}).Fatal(v)
}
