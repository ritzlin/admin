package dto

import "server/model"

type UserGroupResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewUserGroupResponse(userGroup model.UserGroup) UserGroupResponse {
	return UserGroupResponse{
		Id: userGroup.Id,
		Name: userGroup.Name,
	}
}


type CreateUserGroupRequest struct {
	Name string `json:"name" binding:"required"`
}

type DeleteUserGroupRequest struct {
	Id string `json:"id" binding:"required,uuid"`
}