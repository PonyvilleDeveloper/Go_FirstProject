package main

import (
	"app/service"
	"net/http"
	"regexp"
)

func mapping(w http.ResponseWriter, r *http.Request) {
	var ctx service.Context
	ctx.Request = r
	ctx.Response = w
	APIcall := false

	for action, srvc := range cfg.Api {
		correctUrl, _ := regexp.MatchString(srvc.Url, r.URL.Path)
		if r.Method == srvc.Method && correctUrl {
			service.CRUDS[action](&ctx)
			APIcall = true
		}
	}

	if !APIcall {
		http.ServeFile(w, r, r.URL.Path)
	}
}
