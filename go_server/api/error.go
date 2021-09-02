package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type requestError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func errorOfBadRequest(context *gin.Context, message string) {
	context.JSON(http.StatusBadRequest, requestError{
		Message: message,
		Type:    "error",
	})
}

func errorOfNotFound(context *gin.Context) {
	context.Status(http.StatusNotFound)
}

func errorOfServer(context *gin.Context, err error) {
	context.JSON(http.StatusInternalServerError, err.Error())
}

func errorOfDuplicateField(context *gin.Context, field string) {
	context.JSON(http.StatusBadRequest, requestError{
		Message: field,
		Type:    "duplicate",
	})
}

func errorOfEmptyField(context *gin.Context, field string) {
	context.JSON(http.StatusBadRequest, requestError{
		Message: field,
		Type:    "empty",
	})
}
