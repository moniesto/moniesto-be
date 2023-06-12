package systemError

import (
	"github.com/sirupsen/logrus"
)

func Log(args ...interface{}) {
	args = append([]interface{}{"SYSTEM-ERROR:"}, args...)

	logrus.Errorln(args...)
}
