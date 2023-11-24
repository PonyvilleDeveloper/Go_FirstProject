package service

import (
	"encoding/json"
	"app/entity"
	"app/storage"
	"fmt"
)

func CreateCart(data []byte) {
	var cart entity.Cart
	err := json.Unmarshal(data, &cart)
	if err != nil {
		service.log("\t\t[SERVICE]: Decoding json cart for adding error: %v\n", err)
	}
	storage.AddCart(cart)
}

func UpdateCart(id uint32, data []byte) {
	var cart entity.Cart
	err := json.Unmarshal(data, &cart)
	if err != nil {
		service.log("\t\t[SERVICE]: Decoding json cart for changing error: %v\n", err)
	}
	storage.ChangeCart(id, cart)
}

func DeleteCart(id uint32) {
	storage.DeleteCart(id)
}

func GetCartById(id uint32) []byte {
	cart := storage.GetCartById(id)
	json, err := json.Marshal(cart)
	if err != nil {
		service.log("\t\t[SERVICE]: Encoding current json cart error: %v\n", err)
	}
	return []byte(json)
}

func GetCartAll() []byte {
	carts := storage.GetCartAll()
	json, err := json.Marshal(carts)
	if err != nil {
		service.log("\t\t[SERVICE]: Encoding json cart list error: %v\n", err)
	}
	return []byte(json)
}
