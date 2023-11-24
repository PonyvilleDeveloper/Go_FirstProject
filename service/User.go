package service

import (
	"app/entity"
	"app/storage"
	"encoding/json"
)

func CreateUser(data []byte) {
	var user entity.User
	err := json.Unmarshal(data, &user)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json user for adding error: %v\n", err)
	}
	storage.AddUser(user)
}

func UpdateUser(id uint32, data []byte) {
	var user entity.User
	err := json.Unmarshal(data, &user)
	if err != nil {
		log("\t\t[SERVICE]: Decoding json user for changing error: %v\n", err)
	}
	storage.ChangeUser(id, user)
}

func DeleteUser(id uint32) {
	storage.DeleteUser(id)
}

func GetUserById(id uint32) []byte {
	user := storage.GetUserById(id)
	json, err := json.Marshal(user)
	if err != nil {
		log("\t\t[SERVICE]: Encoding current json user error: %v\n", err)
	}
	return []byte(json)
}

func GetUserAll() []byte {
	users := storage.GetUserAll()
	json, err := json.Marshal(users)
	if err != nil {
		log("\t\t[SERVICE]: Encoding json user list error: %v\n", err)
	}
	return []byte(json)
}
