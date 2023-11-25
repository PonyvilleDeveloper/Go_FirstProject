package service

import (
	"app/entity"
	"app/storage"
	"encoding/json"
)

func init() {
	CRUDS["OrderCreate"] = CreateOrder
	CRUDS["OrderUpdate"] = UpdateOrder
	CRUDS["OrderDelete"] = DeleteOrder
	CRUDS["Orders"] = GetOrderAll
	CRUDS["Order"] = GetOrderById
}

func CreateOrder(unprepared Unprepared) []byte { //HTTP.POST
	var order entity.Order
	err := json.Unmarshal(unprepared.Data, &order)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json order for adding error: %v\n", err)
	}
	storage.AddOrder(order)
	return nil
}

func UpdateOrder(unprepared Unprepared) []byte { //HTTP.PUT
	var order entity.Order
	err := json.Unmarshal(unprepared.Data, &order)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json order for changing error: %v\n", err)
	}
	storage.ChangeOrder(unprepared.Id, order)
	return nil
}

func DeleteOrder(unprepared Unprepared) []byte { //HTTP.DELETE
	storage.DeleteOrder(unprepared.Id)
	return nil
}

func GetOrderById(unprepared Unprepared) []byte { //HTTP.GET
	order := storage.GetOrderById(unprepared.Id)
	json, err := json.Marshal(order)
	if err != nil {
		log("\t\t[SERVICE]: Encoding current json order error: %v\n", err)
	}
	return json
}

func GetOrderAll(unprepared Unprepared) []byte { //HTTP.GET
	orders := storage.GetOrderAll()
	json, err := json.Marshal(orders)
	if err != nil {
		log("\t\t[SERVICE]: Encoding json order list error: %v\n", err)
	}
	return json
}
