package common

import (
	"github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	l := logrus.New()
	l.SetFormatter(&formatter.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel:   true,
	})
	l.SetOutput(os.Stdout)
	Log = l
}
