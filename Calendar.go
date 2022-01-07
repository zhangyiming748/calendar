package calendar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/calendar/v3/controller"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func ShowWeb(port int) {
	if port < 1024 || port > 65535 {
		panic("端口号取值范围1024-65535")
	}
	//1.创建路由
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	//gin.Context，封装了request和response

	r.GET("/HappyCount", controller.HappyDayCountdown)
	r.GET("/SadCount", controller.SadDayCountdown)
	r.GET("/", controller.StandRequest)
	/*
		r.POST("/upload", func(c *gin.Context) {
			file, err := c.FormFile("file")
			if err != nil {
				c.String(500, "上传图片出错")
			}
			// c.JSON(200, gin.H{"message": file.Header.Context})
			c.SaveUploadedFile(file, file.Filename)
			c.String(http.StatusOK, file.Filename)
		})
	*/
	r.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		//headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}
		//headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "text/plain" {
			fmt.Println("只允许上传txt文本文件")
			return
		}
		//c.SaveUploadedFile(headers, "./video/"+headers.Filename)
		c.SaveUploadedFile(headers, "./users.txt")
		c.String(http.StatusOK, headers.Filename)
	})
	//3.监听端口，默认在8080
	//Run("里面不指定端口号默认为8080")
	if err := r.SetTrustedProxies([]string{"0.0.0.0"}); err != nil {
		return
	}

	addr := strings.Join([]string{":", strconv.Itoa(port)}, "")
	err := r.Run(addr)
	if err != nil {
		fmt.Println(err)
		return
	}
}
