package service

import (
	"app/entity"
	"app/storage"
	"encoding/json"
)

func init() {
	CRUDS["ItemCreate"] = CreateItem
	CRUDS["ItemUpdate"] = UpdateItem
	CRUDS["ItemDelete"] = DeleteItem
	CRUDS["Items"] = GetItemAll
	CRUDS["Item"] = GetItemById
}

func CreateItem(unprepared Unprepared) []byte { //HTTP.POST
	var item entity.Item
	err := json.Unmarshal(unprepared.Data, &item)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json item for adding error: %v\n", err)
	}
	storage.AddItem(item)
	return nil
}

func UpdateItem(unprepared Unprepared) []byte { //HTTP.PUT
	var item entity.Item
	err := json.Unmarshal(unprepared.Data, &item)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json item for changing error: %v\n", err)
	}
	storage.ChangeItem(unprepared.Id, item)
	return nil
}

func DeleteItem(unprepared Unprepared) []byte { //HTTP.DELETE
	storage.DeleteItem(unprepared.Id)
	return nil
}

func GetItemById(unprepared Unprepared) []byte { //HTTP.GET
	item := storage.GetItemById(unprepared.Id)
	json, err := json.Marshal(item)
	if err != nil {
		log("\t\t[SERVICE]: Encoding current json item error: %v\n", err)
	}
	return json
}

func GetItemAll(unprepared Unprepared) []byte { //HTTP.GET
	items := storage.GetItemAll()
	json, err := json.Marshal(items)
	if err != nil {
		log("\t\t[SERVICE]: Encoding json item list error: %v\n", err)
	}
	return json
}
