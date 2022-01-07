package controller

import (
	"io"
	"log"
	"os"
)

var IPComing *log.Logger

func init() {
	//log.SetFlags(log.Lshortfile)
	fInfo, err := os.OpenFile("info.log", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("打开日志文件错误")
	}
	IPComing = log.New(io.MultiWriter(fInfo, os.Stdout), "访问来自:", log.Lmsgprefix)
}
