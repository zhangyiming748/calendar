package calendar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/calendar/v3/controller"
)

func ShowWeb() {

	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	//gin.Context，封装了request和response

	r.GET("/HappyCount", controller.HappyDayCountdown)
	r.GET("/SadCount", controller.SadDayCountdown)
	r.GET("/", controller.StandRequest)
	//3.监听端口，默认在8080
	//Run("里面不指定端口号默认为8080")
	if err := r.SetTrustedProxies([]string{"0.0.0.0"}); err != nil {
		return
	}
	//gin.SetMode(gin.ReleaseMode)

	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
