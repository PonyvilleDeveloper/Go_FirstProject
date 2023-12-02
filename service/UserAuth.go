package service

import (
	"app/entity"
	"app/logger"
	"app/storageSQL"
	"net/http"
)

func init() {
	CRUDS["UserAuth"] = UserAuth
}

func UserAuth(ctx *Context) {
	err := ctx.Request.ParseForm()
	if err != nil {
		logger.Log(packname, "Error to parse auth form", err)
		ctx.SendError(400, "Some error in auth form")
	}
	login := ctx.Request.PostForm.Get("login")
	password := ctx.Request.PostForm.Get("password")
	exists, id := storageSQL.UserExists(login, password)
	if exists {
		token := entity.NewToken(id)
		storageSQL.AddToken(token)
		http.SetCookie(ctx.Response, &http.Cookie{Value: token.Token, Name: "Token", Expires: token.Expires})
	} else {
		ctx.SendError(400, "Unregistred user.")
	}
}
