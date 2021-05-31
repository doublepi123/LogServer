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
	server.message.LogAdd("INFO", "LogServer:"+util.GetNowTimeFormat()+" "+message)
}

func (server *LogServer) ERROR(message string) {
	server.message.LogAdd("ERROR", "LogServer:"+util.GetNowTimeFormat()+" "+message)
}

func (server *LogServer) login(c *gin.Context) {
	user := &entity.UserEntity{}
	err := c.ShouldBind(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusForbidden, gin.H{"code": 1, "msg": "用户名或密码为空"})
		server.INFO(c.ClientIP() + " " + "fail")
		return
	}
	if server.user.Check(user.Username, user.Password) {
		server.INFO(c.Request.Host + " " + "success")
		userid, err := server.cookie.SetCookie(user.Username)
		fmt.Println(userid)
		if err != nil {
			server.ERROR(c.ClientIP() + " " + fmt.Sprint(err))
			fmt.Println(err)
		}
		util.SetCookieToClient(c, userid)
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "登入成功"})
		return
	}
	c.JSON(http.StatusForbidden, gin.H{"code": -1, "msg": "用户名或密码错误"})
}
func (server LogServer) authcheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, err := c.Cookie("userid")
		if err != nil {
			server.INFO(c.ClientIP() + "###not auth###" + c.Request.RequestURI)
			c.JSON(http.StatusForbidden, gin.H{"code": -1, "msg": "NOT LOGIN"})
			c.Abort()
			return
		}
		_, err = server.cookie.GetUsername(userid)
		if err != nil {
			server.INFO(c.ClientIP() + "###not auth###" + c.Request.RequestURI)
			c.JSON(http.StatusForbidden, gin.H{"code": -1, "msg": "NOT LOGIN"})
			c.Abort()
			return
		}
		server.cookie.UpdateCookie(userid)
		util.SetCookieToClient(c, userid)
		c.Next()
	}
}

func (server *LogServer) connectlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		server.INFO(c.ClientIP() + " " + c.Request.Host + " " + c.Request.RequestURI + " " + c.Request.Method)
		c.Next()
	}
}

func (server *LogServer) getItem(c *gin.Context) {
	items, err := server.log.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "无法获取items列表",
		})
		return
	}
	c.JSON(http.StatusOK, items)

}

func (server *LogServer) findlog(c *gin.Context) {
	m := &entity.ItemReturn{}
	err := c.ShouldBind(m)
	ans, err := server.log.GetLog(m.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, ans)
	return
}

func (server *LogServer) ListenAndServer() {
	r := gin.Default()
	server.user.Add("root", "toor")
	root := r.Group("/api")
	{
		root.Use(server.connectlog())
		root.POST("/login", server.login)
		log := root.Group("/log")
		{
			log.Use(server.authcheck())
			log.GET("/item", server.getItem)
			log.POST("/find", server.findlog)
			log.POST("/count", func(c *gin.Context) {
				m := struct {
					Name string
				}{}
				err := c.ShouldBind(&m)
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusBadRequest, gin.H{
						"msg": err,
					})
					return
				}
				count, err := server.log.Count(m.Name)
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusBadRequest, gin.H{
						"msg": err,
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"count": count,
				})

			})
			log.POST("/select", func(c *gin.Context) {
				m := struct {
					Name string
					From int
					To   int
				}{}
				err := c.ShouldBind(&m)
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusBadRequest, gin.H{
						"msg": err,
					})
					return
				}
				ans, err := server.log.Select(m.Name, m.From, m.To)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"msg": err,
					})
					return
				}
				c.JSON(http.StatusOK, ans)

			})
		}
	}
	err := r.Run("0.0.0.0:39998")
	if err != nil {
		fmt.Println(err)
		server.ERROR(fmt.Sprint(err))
	}

}
