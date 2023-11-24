package service

import (
	"encoding/json"
	"app/entity"
	"app/storage"
	"fmt"
)

func CreateItem(data []byte) {	//HTTP.POST
	var item entity.Item
	err := json.Unmarshal(data, &item)
	if err != nil {
		service.log("\t\t[SERVICE]: Decoding json item for adding error: %v\n", err)
	}
	storage.AddItem(item)
}

func UpdateItem(id uint32, data []byte) {	//HTTP.PUT
	var item entity.Item
	err := json.Unmarshal(data, &item)
	if err != nil {
		service.log("\t\t[SERVICE]: Decoding json item for changing error: %v\n", err)
	}
	storage.ChangeItem(id, item)
}

func DeleteItem(id uint32) {	//HTTP.DELETE
	storage.DeleteItem(id)
}

func GetItemById(id uint32) []byte {	//HTTP.GET
	item := storage.GetItemById(id)
	json, err := json.Marshal(item)
	if err != nil {
		service.log("\t\t[SERVICE]: Encoding current json item error: %v\n", err)
	}
	return []byte(json)
}

func GetItemAll() []byte {	//HTTP.GET
	items := storage.GetItemAll()
	json, err := json.Marshal(items)
	if err != nil {
		service.log("\t\t[SERVICE]: Encoding json item list error: %v\n", err)
	}
	return []byte(json)
}
