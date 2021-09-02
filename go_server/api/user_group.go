package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"server/api/dto"
	"server/db"
	"server/model"
	"server/utils"
)

const urlPrefix = "/user-group"

func initUserGroupRoute(engine *gin.Engine) {
	engine.GET(urlPrefix, getUserGroups)
	engine.POST(urlPrefix, createUserGroup)
	engine.DELETE(urlPrefix+"/:id", deleteUserGroup)
}

func getUserGroups(context *gin.Context) {
	groups, err := db.GetUserGroups()
	if err != nil {
		errorOfServer(context, err)
		return
	}
	response := make([]dto.UserGroupResponse, 0, len(groups))
	for _, group := range groups {
		response = append(response, dto.NewUserGroupResponse(group))
	}
	responseOK(context, response)
}

func createUserGroup(context *gin.Context) {
	var err error
	var request dto.CreateUserGroupRequest
	if err = context.ShouldBindJSON(&request); err != nil {
		errorOfNotFound(context)
		return
	}
	// check empty
	if utils.IsWhiteSpace(request.Name) {
		errorOfEmptyField(context, "Name")
		return
	}
	// check duplicate
	var isFieldExists bool
	if isFieldExists, err = db.IsUserGroupNameExist(request.Name); err != nil {
		errorOfServer(context, err)
		return
	}
	if isFieldExists {
		errorOfDuplicateField(context, "Name")
		return
	}
	userGroup := model.UserGroup{
		Id:   uuid.New().String(),
		Name: request.Name,
	}
	err = db.CreateUserGroup(userGroup)
	if err != nil {
		errorOfServer(context, err)
	} else {
		responseOkNoData(context)
	}
}

func deleteUserGroup(context *gin.Context) {
	var err error
	var request dto.DeleteUserGroupRequest
	if err = context.ShouldBindUri(&request); err != nil {
		errorOfNotFound(context)
		return
	}
	err = db.DeleteUserGroup(request.Id)
	if err != nil {
		errorOfServer(context, err)
	} else {
		responseOkNoData(context)
	}
}
