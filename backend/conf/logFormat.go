package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

//[GIN] 2020/06/29 - 10:25:13 | 200 | 19.726817139s |       127.0.0.1 | POST     "/Login"
type MyFormatter struct{}

func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 - 15:04:05")
	start := strings.Index(entry.Caller.File, "backend")
	filePath := entry.Caller.File[start:]
	msg := fmt.Sprintf("[LOG] %s | %s | %s:%d | %s\n", timestamp, strings.ToUpper(entry.Level.String()), filePath, entry.Caller.Line, entry.Message)
	return []byte(msg), nil
}
