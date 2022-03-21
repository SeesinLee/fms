package errorLog

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func InitErrorLog() {
	logrus.SetReportCaller(true)
	write1 :=os.Stdout
	write2,err  := os.OpenFile("errorLog.txt",os.O_WRONLY|os.O_CREATE,0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	logrus.SetOutput(io.MultiWriter(write1,write2))
}