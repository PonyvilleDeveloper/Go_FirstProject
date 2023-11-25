package service

import (
	"app/entity"
	"app/storage"
	"encoding/json"
)

func init() {
	CRUDS["UserCreate"] = CreateUser
	CRUDS["UserUpdate"] = UpdateUser
	CRUDS["UserDelete"] = DeleteUser
	CRUDS["Users"] = GetUserAll
	CRUDS["User"] = GetUserById
}

func CreateUser(unprepared Unprepared) []byte { //HTTP.POST
	var user entity.User
	err := json.Unmarshal(unprepared.Data, &user)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json user for adding error: %v\n", err)
	}
	storage.AddUser(user)
	return nil
}

func UpdateUser(unprepared Unprepared) []byte { //HTTP.PUT
	var user entity.User
	err := json.Unmarshal(unprepared.Data, &user)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json user for changing error: %v\n", err)
	}
	storage.ChangeUser(unprepared.Id, user)
	return nil
}

func DeleteUser(unprepared Unprepared) []byte { //HTTP.DELETE
	storage.DeleteUser(unprepared.Id)
	return nil
}

func GetUserById(unprepared Unprepared) []byte { //HTTP.GET
	user := storage.GetUserById(unprepared.Id)
	json, err := json.Marshal(user)
	if err != nil {
		log("\t\t[SERVICE]: Encoding current json user error: %v\n", err)
	}
	return json
}

func GetUserAll(unprepared Unprepared) []byte { //HTTP.GET
	users := storage.GetUserAll()
	json, err := json.Marshal(users)
	if err != nil {
		log("\t\t[SERVICE]: Encoding json user list error: %v\n", err)
	}
	return json
}
