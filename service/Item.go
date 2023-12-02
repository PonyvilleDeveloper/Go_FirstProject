package service

import (
	"app/entity"
	"app/logger"
	"app/storageSQL"
)

func init() {
	CRUDS["ItemCreate"] = CreateItem
	CRUDS["ItemUpdate"] = UpdateItem
	CRUDS["ItemDelete"] = DeleteItem
	CRUDS["Items"] = GetItemAll
	CRUDS["Item"] = GetItemById
}

func CreateItem(ctx *Context) { //HTTP.POST
	var cart entity.Item
	cart, err := ExtractData[entity.Item](ctx)
	if err != nil {
		go logger.Log(packname, "Error decoding json for create Item", err)
		ctx.SendError(500, "Error decoding json for create Item")
	}
	storageSQL.AddItem(cart)
}

func UpdateItem(ctx *Context) { //HTTP.PUT
	var cart entity.Item
	cart, err := ExtractData[entity.Item](ctx)
	if err != nil {
		go logger.Log(packname, "Error decoding json for update Item #", ctx.Id, err)
		ctx.SendError(500, "Error decoding json for update Item")
	}
	storageSQL.ChangeItem(ctx.Id, cart)
}

func DeleteItem(ctx *Context) { //HTTP.DELETE
	storageSQL.DeleteItem(ctx.Id)
}

func GetItemById(ctx *Context) { //HTTP.GET
	cart := storageSQL.GetItemById(ctx.Id)
	ctx.SendAnswer(cart)
}

func GetItemAll(ctx *Context) { //HTTP.GET
	carts := storageSQL.GetItemAll()
	ctx.SendAnswer(carts)
}
