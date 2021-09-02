package db

import (
	"server/global"
	"server/model"
)


func GetUserGroups() ( []model.UserGroup, error) {
	var groups []model.UserGroup
	query := `SELECT "Id", "Name" FROM "UserGroup"`
	err := global.Db.Select(&groups, query)
	return groups, err
}

func CreateUserGroup(userGroup model.UserGroup) error {
	query := `
		INSERT INTO "UserGroup"
		("Id", "Name")
		VALUES
		($1, $2)`
	_, err := global.Db.Exec(query, userGroup.Id, userGroup.Name)
	return err
}

func DeleteUserGroup(id string) error {
	query := `
		DELETE FROM "UserGroup"
		WHERE "Id" = $1`
	_, err := global.Db.Exec(query, id)
	return err
}

func IsUserGroupNameExist(name string) (bool, error) {
	query :=
		`SELECT EXISTS (
			SELECT 1 FROM "UserGroup"
			WHERE "Name" = $1)`
	var isExists bool
	err := global.Db.Get(&isExists, query, name)
	return isExists, err
}