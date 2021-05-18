package util

import (
	"github.com/gin-gonic/gin"
	"time"
)

func PauseForRun() {
	for {
		time.Sleep(time.Hour)
	}
}

func SetCookieToClient(c *gin.Context, userid string) {
	c.SetCookie("userid", userid, int(time.Minute*15), "/", c.Request.Host, false, true)
}

func GetNowTimeFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
