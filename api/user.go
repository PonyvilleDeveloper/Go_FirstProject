package api

import (
	"app/db"
	"app/engine"
	"app/entity"
	"github.com/google/uuid"
	"time"
)

func (a *Api) UserAuth(ctx *engine.Context) {
	user, err := engine.ToStruct[entity.UserAuth](ctx)
	if err != nil {
		ctx.Error(400, "Bad user data")
	}
	var usr entity.User
	db.DB().Table("users").Where("Login = ? and Password = ?", user.Login, user.Password).Find(&usr)
	if usr.UserId == 0 {
		ctx.Error(401, "Bad Auth")
		return
	}
	token := entity.Token{
		UserId:  usr.UserId,
		Token:   uuid.NewString(),
		Expired: time.Now().Add(1 * time.Hour),
	}
	db.DB().Table("users").Update("Token", &token)
	ctx.Print(token)
}
