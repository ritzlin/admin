package api

import (
	"github.com/gin-gonic/gin"
)


func InitRouter(engine *gin.Engine) {
	initUserGroupRoute(engine)


}
