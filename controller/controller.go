package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HappyDayCountdown(ctx *gin.Context) {
	content := Happy{
		Title:    "Happy",
		Contents: HappyDayJson(),
		Tail:     Gift(),
	}
	IPComing.Println(GetRequestIP(ctx))
	ctx.JSON(http.StatusOK, content)
}
func SadDayCountdown(ctx *gin.Context) {
	content := Sad{
		Title:    "Sad",
		Contents: SadDayJson(),
	}
	IPComing.Println(GetRequestIP(ctx))
	ctx.JSON(http.StatusOK, content)
}

func StandRequest(ctx *gin.Context) {
	content := Stand{
		Title: StandHeadJson(),
		Count: HappyDayJson(),
		Tail:  StandTailJson(),
	}
	IPComing.Println(GetRequestIP(ctx))
	ctx.JSON(http.StatusOK, content)
}
