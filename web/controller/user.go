package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/utils"
)

func GetSession(c *gin.Context) {
	//初始化返回的  map
	resp := make(map[string]string)
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	c.JSONP(http.StatusOK, resp)
}
