package pkg

import (
	"io"
	"log"
	"os"
	"time"
)

const timeFormat  =  "20060102"

var Info,Error,Warning *log.Logger

func init()  {
	filepath:="logs/"+time.Now().Format(timeFormat)+".log"

	file,err := os.OpenFile(filepath,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil {
		log.Fatalln("failed to open or create log file",err)
	}
	Info = log.New(os.Stdout,"INFO:",log.Ldate|log.Ltime|log.Llongfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	Warning = log.New(os.Stdout,"WARNING:",log.Ldate|log.Ltime|log.Llongfile)
}