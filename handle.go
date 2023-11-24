package main

import (
	"app/service"
	"net/http"
	"strconv"
	"strings"
)

func HandleGet(args []string) (data []byte) {
	switch args[2] {
	case "item":
		if args[3] == "id" {
			id, _ := strconv.ParseUint(args[4], 10, 64)
			data = service.GetItemById(uint32(id))
		}
		if args[3] == "all" {
			data = service.GetItemAll()
		} else {
			//error.html
		}
	case "order":
		if args[3] == "id" {
			id, _ := strconv.ParseUint(args[4], 10, 64)
			data = service.GetOrderById(uint32(id))
		}
		if args[3] == "all" {
			data = service.GetOrderAll()
		}
	case "cart":
		if args[3] == "id" {
			id, _ := strconv.ParseUint(args[4], 10, 64)
			data = service.GetCartById(uint32(id))
		}
		if args[3] == "all" {
			data = service.GetCartAll()
		}
	case "user":
		if args[3] == "id" {
			id, _ := strconv.ParseUint(args[4], 10, 64)
			data = service.GetUserById(uint32(id))
		}
		if args[3] == "all" {
			data = service.GetUserAll()
		}
	}
	return
}

func HandlePut(args []string, data []byte) {
	id, _ = strconv.ParseUint(args[4], 10, 64)
	switch args[3] {
	case "item":
		service.UpdateItem(uint32(id), data)
	case "order":
	case "user":
	case "cart":

	}
}

func mapping(w http.ResponseWriter, r *http.Request) {

	args := strings.Split(r.URL.Path, "/")

	switch r.Method {
	case "GET":
		w.Write(HandleGet(args))
	case "PUT":
		HandlePut(args, r.)
	case "POST":
	case "DELETE":
	}
}
