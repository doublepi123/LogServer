package service

import (
	"LogServer/dao"
	"LogServer/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogServer struct {
	log    *dao.LogDao
	cookie *dao.CookieDao
}

func (server *LogServer) Init(logDao *dao.LogDao, cookieDao *dao.CookieDao) {
	server.cookie = cookieDao
	server.log = logDao
}

func (server *LogServer) login(c *gin.Context) {
	user := &entity.UserEntity{}
	err := c.ShouldBind(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

}

func (server *LogServer) ListenAndServer() {
	r := gin.Default()
	r.POST("/api/login")
}
