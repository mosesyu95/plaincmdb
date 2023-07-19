package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
)

func InitLog() {
	level, err := logrus.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		fmt.Printf("unknown log level for \"%s\"\n"+
			"set Default Warning level ", viper.GetString("log.level"))
		level = logrus.WarnLevel
	}
	writers := []io.Writer{os.Stdout}
	file, err := os.OpenFile(viper.GetString("log.path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		writers = append(writers, file)
	} else {
		fmt.Printf("Failed to log to file, using default stdout")
	}
	logrus.SetLevel(level)
	logrus.SetOutput(io.MultiWriter(writers...))
	logrus.SetFormatter(&logrus.TextFormatter{})
}
