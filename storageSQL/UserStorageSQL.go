package storageSQL

import (
	"app/db"
	"app/entity"
)

func AddUser(user entity.User) { //Экв. CreateUser в пакете service
	db.DB().Table("users").Create(user)
}

func ChangeUser(id uint32, updts entity.User) { //Экв. UpdateUser в пакете service
	db.DB().Table("users").Where("cartId= ?", id).Updates(updts)
}

func DeleteUser(id uint32) { //Экв. DeleteUser в пакете service
	var user entity.User
	db.DB().Table("users").Delete(&user, id)
}

func GetUserById(id uint32) entity.User { //Экв. GetUserById в пакете service
	var user entity.User

	db.DB().Table("users").Where("id = ?", id).Find(&user)
	return user
}

func GetUserAll() []*entity.User { //Экв. GetUserAll в пакете service
	var users []*entity.User
	db.DB().Table("users").Find(&users)
	return users
}
