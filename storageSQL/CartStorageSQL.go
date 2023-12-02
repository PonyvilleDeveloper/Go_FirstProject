package storageSQL

import (
	"app/entity"
)

func AddCart(cart entity.Cart) { //Экв. Createcart в пакете service
	database.Table("carts").Create(cart)
}

func ChangeCart(id uint32, updts entity.Cart) { //Экв. Updatecart в пакете service
	database.Table("carts").Where("cartId= ?", id).Updates(updts)
}

func DeleteCart(id uint32) { //Экв. Deletecart в пакете service
	var cart entity.Cart
	database.Table("carts").Delete(&cart, id)
}

func GetCartById(id uint32) entity.Cart { //Экв. GetcartById в пакете service
	var cart entity.Cart
	database.Table("carts").Where("id = ?", id).First(&cart)
	return cart
}

func GetCartAll() []*entity.Cart { //Экв. GetcartAll в пакете service
	var carts []*entity.Cart
	database.Table("carts").Find(&carts)
	return carts
}
