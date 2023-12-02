package storageSQL

import (
	"app/entity"
)

func AddOrder(order entity.Order) { //Экв. CreateOrder в пакете service
	database.Table("orders").Create(order)
}

func ChangeOrder(id uint32, updts entity.Order) { //Экв. UpdateOrder в пакете service
	database.Table("orders").Where("cartId= ?", id).Updates(updts)
}

func DeleteOrder(id uint32) { //Экв. DeleteOrder в пакете service
	var order entity.Order
	database.Table("orders").Delete(&order, id)
}

func GetOrderById(id uint32) entity.Order { //Экв. GetOrderById в пакете service
	var order entity.Order
	database.Table("orders").Where("id = ?", id).First(&order)
	return order
}

func GetOrderAll() []*entity.Order { //Экв. GetOrderAll в пакете service
	var orders []*entity.Order
	database.Table("orders").Find(&orders)
	return orders
}
