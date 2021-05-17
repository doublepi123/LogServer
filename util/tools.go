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
	c.SetCookie("userid", userid, int(time.Minute*15), "/", c.Request.Host, false, false)
}
