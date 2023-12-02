package service

import (
	"app/entity"
	"app/logger"
	"app/storageSQL"
)

func init() {
	CRUDS["OrderCreate"] = CreateOrder
	CRUDS["OrderUpdate"] = UpdateOrder
	CRUDS["OrderDelete"] = DeleteOrder
	CRUDS["Orders"] = GetOrderAll
	CRUDS["Order"] = GetOrderById
}

func CreateOrder(ctx *Context) { //HTTP.POST
	var cart entity.Order
	cart, err := ExtractData[entity.Order](ctx)
	if err != nil {
		go logger.Log(packname, "Error decoding json for create Order", err)
		ctx.SendError(500, "Error decoding json for create Order")
	}
	storageSQL.AddOrder(cart)
}

func UpdateOrder(ctx *Context) { //HTTP.PUT
	var cart entity.Order
	cart, err := ExtractData[entity.Order](ctx)
	if err != nil {
		go logger.Log(packname, "Error decoding json for update Order #", ctx.Id, err)
		ctx.SendError(500, "Error decoding json for update Order")
	}
	storageSQL.ChangeOrder(ctx.Id, cart)
}

func DeleteOrder(ctx *Context) { //HTTP.DELETE
	storageSQL.DeleteOrder(ctx.Id)
}

func GetOrderById(ctx *Context) { //HTTP.GET
	cart := storageSQL.GetOrderById(ctx.Id)
	ctx.SendAnswer(cart)
}

func GetOrderAll(ctx *Context) { //HTTP.GET
	carts := storageSQL.GetOrderAll()
	ctx.SendAnswer(carts)
}
