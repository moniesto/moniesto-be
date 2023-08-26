package systemError

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Log(args ...interface{}) {
	args = append([]interface{}{"SYSTEM-ERROR:"}, args...)

	logrus.Errorln(args...)
}

func LogBody(request string, c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	fmt.Println("REQUEST-BODY-LOG", request, string(body))

	c.Request.Body = io.NopCloser(bytes.NewReader(body))
}
