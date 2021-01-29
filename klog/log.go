/**
 *创建人：kun
 *日期：2020/3/25
 */
package klog

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var logger *log.Entry

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
	//log.SetReportCaller(true)

	logger = log.WithFields(log.Fields{
		//"fields": "service-platform",
	})
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
