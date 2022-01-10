package log

import (
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger //只打印
	Debug *log.Logger //打印并写入文档
)
func init() {

	info, err2 := os.OpenFile("info.log", os.O_WRONLY|os.O_CREATE, 0666)
	if err2 != nil {
		log.Println("打开日志文件错误")
	}

	Info = log.New(os.Stdout, "", log.Lmsgprefix)
	Debug = log.New(io.MultiWriter(info, os.Stdout), "",log.Lmsgprefix)
}