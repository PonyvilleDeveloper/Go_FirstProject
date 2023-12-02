package storageSQL

import (
	"app/entity"
)

func AddUser(user entity.User) { //Экв. CreateUser в пакете service
	database.Table("users").Create(user)
}

func ChangeUser(id uint32, updts entity.User) { //Экв. UpdateUser в пакете service
	database.Table("users").Where("cartId= ?", id).Updates(updts)
}

func DeleteUser(id uint32) { //Экв. DeleteUser в пакете service
	var user entity.User
	database.Table("users").Delete(&user, id)
}

func GetUserById(id uint32) entity.User { //Экв. GetUserById в пакете service
	var user entity.User
	database.Table("users").Where("id = ?", id).First(&user)
	return user
}

func GetUserAll() []*entity.User { //Экв. GetUserAll в пакете service
	var users []*entity.User
	database.Table("users").Find(&users)
	return users
}

func UserExists(login, password string) (res bool, id uint32) {
	rows := database.Table("users").Where("Login = ?, Password = ?", login, password)
	rows.Select("EXISTS").First(&res)
	if res {
		rows.Select("UserId").First(&id)
	}
	return
}
