package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type dataResponse struct {
	Data interface{} `json:"data"`
}

func responseOK(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, dataResponse{
		Data: data,
	})
}

func responseOkNoData(context *gin.Context) {
	context.Status(http.StatusNoContent)
}