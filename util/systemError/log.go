package systemError

import (
	"github.com/sirupsen/logrus"
)

func Log(error string) {
	logrus.Errorln(error)
}
