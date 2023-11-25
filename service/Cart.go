package service

import (
	"app/entity"
	"app/storage"
	"encoding/json"
)

func init() {
	CRUDS["CartCreate"] = CreateCart
	CRUDS["CartUpdate"] = UpdateCart
	CRUDS["CartDelete"] = DeleteCart
	CRUDS["Carts"] = GetCartAll
	CRUDS["Cart"] = GetCartById
}

func CreateCart(unprepared Unprepared) []byte { //HTTP.POST
	var cart entity.Cart
	err := json.Unmarshal(unprepared.Data, &cart)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json cart for adding error: %v\n", err)
	}
	storage.AddCart(cart)
	return nil
}

func UpdateCart(unprepared Unprepared) []byte { //HTTP.PUT
	var cart entity.Cart
	err := json.Unmarshal(unprepared.Data, &cart)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json cart for changing error: %v\n", err)
	}
	storage.ChangeCart(unprepared.Id, cart)
	return nil
}

func DeleteCart(unprepared Unprepared) []byte { //HTTP.DELETE
	storage.DeleteCart(unprepared.Id)
	return nil
}

func GetCartById(unprepared Unprepared) []byte { //HTTP.GET
	cart := storage.GetCartById(unprepared.Id)
	json, err := json.Marshal(cart)
	if err != nil {
		log("\t\t[SERVICE]: Encoding current json cart error: %v\n", err)
	}
	return json
}

func GetCartAll(unprepared Unprepared) []byte { //HTTP.GET
	carts := storage.GetCartAll()
	json, err := json.Marshal(carts)
	if err != nil {
		log("\t\t[SERVICE]: Encoding json cart list error: %v\n", err)
	}
	return json
}
