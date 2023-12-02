package service

import (
	"app/entity"
	"app/logger"
	"app/storageSQL"
)

func init() {
	CRUDS["UserCreate"] = CreateUser
	CRUDS["UserUpdate"] = UpdateUser
	CRUDS["UserDelete"] = DeleteUser
	CRUDS["Users"] = GetUserAll
	CRUDS["User"] = GetUserById
}

func CreateUser(ctx *Context) { //HTTP.POST
	var cart entity.User
	cart, err := ExtractData[entity.User](ctx)
	if err != nil {
		go logger.Log(packname, "Error decoding json for create User", err)
		ctx.SendError(500, "Error decoding json for create User")
	}
	storageSQL.AddUser(cart)
}

func UpdateUser(ctx *Context) { //HTTP.PUT
	var cart entity.User
	cart, err := ExtractData[entity.User](ctx)
	if err != nil {
		go logger.Log(packname, "Error decoding json for update User #", ctx.Id, err)
		ctx.SendError(500, "Error decoding json for update User")
	}
	storageSQL.ChangeUser(ctx.Id, cart)
}

func DeleteUser(ctx *Context) { //HTTP.DELETE
	storageSQL.DeleteUser(ctx.Id)
}

func GetUserById(ctx *Context) { //HTTP.GET
	cart := storageSQL.GetUserById(ctx.Id)
	ctx.SendAnswer(cart)
}

func GetUserAll(ctx *Context) { //HTTP.GET
	carts := storageSQL.GetUserAll()
	ctx.SendAnswer(carts)
}
