package service

var (
	CRUDS    map[string]func(ctx *Context)
	packname string
)

func init() {
	CRUDS = make(map[string]func(ctx *Context))
	packname = "SERVICE"
}
