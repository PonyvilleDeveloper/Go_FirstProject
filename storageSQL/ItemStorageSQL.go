package storageSQL

import (
	"app/db"
	"app/entity"
)

func AddItem(item entity.Item) { //Экв. CreateItem в пакете service
	db.DB().Table("items").Create(item)
}

func ChangeItem(id uint32, updts entity.Item) { //Экв. UpdateItem в пакете service
	db.DB().Table("items").Where("cartId= ?", id).Updates(updts)
}

func DeleteItem(id uint32) { //Экв. DeleteItem в пакете service
	var item entity.Item
	db.DB().Table("items").Delete(&item, id)
}

func GetItemById(id uint32) entity.Item { //Экв. GetItemById в пакете service
	var item entity.Item

	db.DB().Table("items").Where("id = ?", id).Find(&item)
	return item
}

func GetItemAll() []*entity.Item { //Экв. GetItemAll в пакете service
	var items []*entity.Item
	db.DB().Table("items").Find(&items)
	return items
}
