package storageSQL

import (
	"app/db"
	"app/entity"
)

func AddCart(cart entity.Cart) { //Экв. Createcart в пакете service
	db.DB().Table("carts").Create(cart)
}

func ChangeCart(id uint32, updts entity.Cart) { //Экв. Updatecart в пакете service
	db.DB().Table("carts").Where("cartId= ?", id).Updates(updts)
}

func DeleteCart(id uint32) { //Экв. Deletecart в пакете service
	var cart entity.Cart
	db.DB().Table("carts").Delete(&cart, id)
}

func GetCartById(id uint32) entity.Cart { //Экв. GetcartById в пакете service
	var cart entity.Cart

	db.DB().Table("carts").Where("id = ?", id).Find(&cart)
	return cart
}

func GetCartAll() []*entity.Cart { //Экв. GetcartAll в пакете service
	var carts []*entity.Cart
	db.DB().Table("carts").Find(&carts)
	return carts
}
