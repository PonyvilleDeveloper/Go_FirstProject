package storageSQL

import (
	"app/db"
	"app/entity"
)

func AddOrder(order entity.Order) { //Экв. CreateOrder в пакете service
	db.DB().Table("orders").Create(order)
}

func ChangeOrder(id uint32, updts entity.Order) { //Экв. UpdateOrder в пакете service
	db.DB().Table("orders").Where("cartId= ?", id).Updates(updts)
}

func DeleteOrder(id uint32) { //Экв. DeleteOrder в пакете service
	var order entity.Order
	db.DB().Table("orders").Delete(&order, id)
}

func GetOrderById(id uint32) entity.Order { //Экв. GetOrderById в пакете service
	var order entity.Order

	db.DB().Table("orders").Where("id = ?", id).Find(&order)
	return order
}

func GetOrderAll() []*entity.Order { //Экв. GetOrderAll в пакете service
	var orders []*entity.Order
	db.DB().Table("orders").Find(&orders)
	return orders
}
