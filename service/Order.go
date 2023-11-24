package service

import (
	"app/entity"
	"app/storage"
	"encoding/json"
)

func CreateOrder(data []byte) {
	var order entity.Order
	err := json.Unmarshal(data, &order)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json order for adding error: %v\n", err)
	}
	storage.AddOrder(order)
}

func UpdateOrder(id uint32, data []byte) {
	var order entity.Order
	err := json.Unmarshal(data, &order)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json order for changing error: %v\n", err)
	}
	storage.ChangeOrder(id, order)
}

func DeleteOrder(id uint32) {
	storage.DeleteOrder(id)
}

func GetOrderById(id uint32) []byte {
	order := storage.GetOrderById(id)
	json, err := json.Marshal(order)
	if err != nil {
		log("\t\t[SERVICE]: Encoding current json order error: %v\n", err)
	}
	return []byte(json)
}

func GetOrderAll() []byte {
	orders := storage.GetOrderAll()
	json, err := json.Marshal(orders)
	if err != nil {
		log("\t\t[SERVICE]: Encoding json order list error: %v\n", err)
	}
	return []byte(json)
}
