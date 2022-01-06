package calendar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/calendar/v3/controller"
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
