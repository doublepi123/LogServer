package service

import (
	"LogServer/dao"
	"LogServer/entity"
	"LogServer/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LogServer struct {
	user    *dao.UserDao
	log     *dao.LogDao
	cookie  *dao.CookieDao
	message *dao.MessageDao
}

func (server *LogServer) Init(logDao *dao.LogDao, cookieDao *dao.CookieDao, user *dao.UserDao, messageDao *dao.MessageDao) {
	server.cookie = cookieDao
	server.log = logDao
	server.user = user
	server.message = messageDao
}

func (server *LogServer) INFO(message string) {
	server.message.LogAdd("INFO", message)
}

func (server *LogServer) ERROR(message string) {
	server.message.LogAdd("ERROR", message)
}

func (server *LogServer) login(c *gin.Context) {
	user := &entity.UserEntity{}
	err := c.ShouldBind(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusForbidden, gin.H{"code": 1, "msg": "用户名或密码为空"})
		server.INFO(c.Request.Host + time.Now().String() + "fail")
		return
	}
	if server.user.Check(user.Username, user.Password) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "登入成功"})
		server.INFO(c.Request.Host + time.Now().String() + "success")
		err := server.cookie.SetCookie(user.Username)
		if err != nil {
			server.ERROR(c.Request.Host + time.Now().String() + fmt.Sprint(err))
			fmt.Println(err)
		}
		util.SetCookieToClient(c, user.Username)
		return
	}
	c.JSON(http.StatusForbidden, gin.H{"code": -1, "msg": "用户名或密码错误"})
}
func (server LogServer) authcheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, err := c.Cookie("userid")
		if err != nil {
			server.INFO(c.Request.Host + time.Now().String() + "###not auth###" + c.Request.RequestURI)
			c.JSON(http.StatusForbidden, gin.H{"code": -1, "msg": "NOT LOGIN"})
			c.Abort()
			return
		}
		_, err = server.cookie.GetUsername(userid)
		if err != nil {
			server.INFO(c.Request.Host + time.Now().String() + "###not auth###" + c.Request.RequestURI)
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
