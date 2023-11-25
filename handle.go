package main

import (
	"app/service"
	"net/http"
	"regexp"
	"strconv"
)

func mapping(w http.ResponseWriter, r *http.Request) {
	var data []byte
	r.Body.Read(data)
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	for action, srvc := range cfg.Api {
		correctUrl, _ := regexp.MatchString(srvc.Url, r.URL.Path)
		if r.Method == srvc.Method && correctUrl {
			w.Write(service.CRUDS[action](service.Unprepared{data, uint32(id)}))
		}
	}
}
