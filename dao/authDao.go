package dao

import "go-blog/entity"

func CheckAuth(username, password string) bool {
	var auth entity.Auth
	db.Select("id").Where(entity.Auth{Username : username, Password : password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}