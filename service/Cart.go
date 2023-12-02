package service

import (
	"app/entity"
	"app/logger"
	"app/storageSQL"
)

func init() {
	CRUDS["CartCreate"] = CreateCart
	CRUDS["CartUpdate"] = UpdateCart
	CRUDS["CartDelete"] = DeleteCart
	CRUDS["Carts"] = GetCartAll
	CRUDS["Cart"] = GetCartById
}

func CreateCart(ctx *Context) { //HTTP.POST
	var cart entity.Cart
	cart, err := ExtractData[entity.Cart](ctx)
	if err != nil {
		go logger.Log(packname, "Error decoding json for create Cart", err)
		ctx.SendError(500, "Error decoding json for create Cart")
	}
	storageSQL.AddCart(cart)
}

func UpdateCart(ctx *Context) { //HTTP.PUT
	var cart entity.Cart
	cart, err := ExtractData[entity.Cart](ctx)
	if err != nil {
		go logger.Log(packname, "Error decoding json for update Cart #", ctx.Id, err)
		ctx.SendError(500, "Error decoding json for update Cart")
	}
	storageSQL.ChangeCart(ctx.Id, cart)
}

func DeleteCart(ctx *Context) { //HTTP.DELETE
	storageSQL.DeleteCart(ctx.Id)
}

func GetCartById(ctx *Context) { //HTTP.GET
	cart := storageSQL.GetCartById(ctx.Id)
	ctx.SendAnswer(cart)
}

func GetCartAll(ctx *Context) { //HTTP.GET
	carts := storageSQL.GetCartAll()
	ctx.SendAnswer(carts)
}
