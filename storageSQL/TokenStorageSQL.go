package storageSQL

import "app/entity"

func AddToken(token entity.Token) {
	database.Table("tokens").Create(token)
}

func ExistToken(user_id uint32) (res bool) {
	database.Table("tokens").Exec("SELECT EXISTS(SELECT * WHERE UserId = ?) AS value", user_id).First(&res)
	return
}

func UpdateToken(token entity.Token) {
	database.Table("tokens").Where("UserId = ?", token.UserId).Updates(token)
}
