package service

import (
	"LogServer/dao"
	"LogServer/entity"
	"LogServer/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogServer struct {
	user   *dao.UserDao
	log    *dao.LogDao
	cookie *dao.CookieDao
}

func (server *LogServer) Init(logDao *dao.LogDao, cookieDao *dao.CookieDao, user *dao.UserDao) {
	server.cookie = cookieDao
	server.log = logDao
	server.user = user
}

func (server *LogServer) login(c *gin.Context) {
	user := &entity.UserEntity{}
	err := c.ShouldBind(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusForbidden, gin.H{"code": 1, "msg": "用户名或密码为空"})
		return
	}
	if server.user.Check(user.Username, user.Password) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "登入成功"})
		err := server.cookie.SetCookie(user.Username)
		fmt.Println(err)
		util.SetCookieToClient(c, user.Username)
		return
	}
	c.JSON(http.StatusForbidden, gin.H{"code": -1, "msg": "用户名或密码错误"})
}
func (server LogServer) authcheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, err := c.Cookie("userid")
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"code": -1, "msg": "NOT LOGIN"})
			c.Abort()
			return
		}
		_, err = server.cookie.GetUsername(userid)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"code": -1, "msg": "NOT LOGIN"})
			c.Abort()
			return
		}
	}
}

func (server *LogServer) ListenAndServer() {
	r := gin.Default()
	r.POST("/api/login")
	log := r.Group("/api/log")
	log.Use(server.authcheck())
	err := r.Run("127.0.0.1:9998")
	fmt.Println(err)
}
