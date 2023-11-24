package storageSQL

import (
	"app/entity"
	"sync"
)

type UserStorage struct {
	data map[uint32]entity.User //Мапа, контейнер данных
	iter uint32                 //Индексатор для добавления новых
	mtx  sync.Mutex             //Блокировщик доступа к контейнеру, чтобы разные потоки не сломали его
}

var USERS UserStorage //Выделяем память под хранилище

func init() {
	USERS.data = make(map[uint32]entity.User) //Инициализация контейнера
}

func AddUser(user entity.User) { //Экв. CreateUser в пакете service
	USERS.mtx.Lock()                 //Блокировка хранилища
	defer USERS.mtx.Unlock()         //Запланируем разблокировку на момент выхода из функции
	user.UserId = uint64(USERS.iter) //Копируем итератор хранилища как id объекта
	USERS.data[USERS.iter] = user    //Добавляем объект в контейнер
	USERS.iter++                     //Инкрементируем индексатор
}

func ChangeUser(id uint32, updts entity.User) { //Экв. UpdateUser в пакете service
	USERS.mtx.Lock()         //Блокировка хранилища
	defer USERS.mtx.Unlock() //Запланируем разблокировку на момент выхода из функции
	USERS.data[id] = updts   //Обновление данных
}

func DeleteUser(id uint32) { //Экв. DeleteUser в пакете service
	USERS.mtx.Lock()         //Блокировка хранилища
	defer USERS.mtx.Unlock() //Запланируем разблокировку на момент выхода из функции
	delete(USERS.data, id)   //Удаляем объект
}

func GetUserById(id uint32) entity.User { //Экв. GetUserById в пакете service
	USERS.mtx.Lock()         //Блокировка хранилища
	defer USERS.mtx.Unlock() //Запланируем разблокировку на момент выхода из функции
	return USERS.data[id]    //Возвращаем найденный объект
}

func GetUserAll() []entity.User { //Экв. GetUserAll в пакете service
	USERS.mtx.Lock()               //Блокировка хранилища
	defer USERS.mtx.Unlock()       //Запланируем разблокировку на момент выхода из функции
	list := []entity.User{}        //Подготавливаем массив, в который положем объекты
	for _, v := range USERS.data { //Переносим объекты
		list = append(list, v) //Копируем значения в массив
	}
	return list //Возвращаем массив
}
