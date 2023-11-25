package storage

import (
	"app/entity"
	"sync"
)

type CartStorage struct {
	data map[uint32]entity.Cart //Мапа, контейнер данных
	iter uint32                 //Индексатор для добавления новых
	mtx  sync.Mutex             //Блокировщик доступа к контейнеру, чтобы разные потоки не сломали его
}

var CARTS CartStorage //Выделяем память под хранилище

func init() {
	CARTS.data = make(map[uint32]entity.Cart) //Инициализация контейнера
}

func AddCart(item entity.Cart) { //Экв. CreateCart в пакете service
	CARTS.mtx.Lock()              //Блокировка хранилища
	defer CARTS.mtx.Unlock()      //Запланируем разблокировку на момент выхода из функции
	item.CartId = CARTS.iter      //Копируем итератор хранилища как id объекта
	CARTS.data[CARTS.iter] = item //Добавляем объект в контейнер
	CARTS.iter++                  //Инкрементируем индексатор
}

func ChangeCart(id uint32, updts entity.Cart) { //Экв. UpdateCart в пакете service
	CARTS.mtx.Lock()         //Блокировка хранилища
	defer CARTS.mtx.Unlock() //Запланируем разблокировку на момент выхода из функции
	CARTS.data[id] = updts   //Обновление данных
}

func DeleteCart(id uint32) { //Экв. DeleteCart в пакете service
	CARTS.mtx.Lock()         //Блокировка хранилища
	defer CARTS.mtx.Unlock() //Запланируем разблокировку на момент выхода из функции
	delete(CARTS.data, id)   //Удаляем объект
}

func GetCartById(id uint32) entity.Cart { //Экв. GetCartById в пакете service
	CARTS.mtx.Lock()         //Блокировка хранилища
	defer CARTS.mtx.Unlock() //Запланируем разблокировку на момент выхода из функции
	return CARTS.data[id]    //Возвращаем найденный объект
}

func GetCartAll() []entity.Cart { //Экв. GetCartAll в пакете service
	CARTS.mtx.Lock()               //Блокировка хранилища
	defer CARTS.mtx.Unlock()       //Запланируем разблокировку на момент выхода из функции
	list := []entity.Cart{}        //Подготавливаем массив, в который положем объекты
	for _, v := range CARTS.data { //Переносим объекты
		list = append(list, v) //Копируем значения в массив
	}
	return list //Возвращаем массив
}
