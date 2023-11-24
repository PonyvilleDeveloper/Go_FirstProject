package storageSQL

import (
	"app/entity"
	"sync"
)

type OrderStorage struct {
	data map[uint32]entity.Order //Мапа, контейнер данных
	iter uint32                  //Индексатор для добавления новых
	mtx  sync.Mutex              //Блокировщик доступа к контейнеру, чтобы разные потоки не сломали его
}

var ORDERS OrderStorage //Выделяем память под хранилище

func init() {
	ORDERS.data = make(map[uint32]entity.Order) //Инициализация контейнера
}

func AddOrder(order entity.Order) { //Экв. CreateOrder в пакете service
	ORDERS.mtx.Lock()                //Блокировка хранилища
	defer ORDERS.mtx.Unlock()        //Запланируем разблокировку на момент выхода из функции
	order.OrderId = ORDERS.iter      //Копируем итератор хранилища как id объекта
	ORDERS.data[ORDERS.iter] = order //Добавляем объект в контейнер
	ORDERS.iter++                    //Инкрементируем индексатор
}

func ChangeOrder(id uint32, updts entity.Order) { //Экв. UpdateOrder в пакете service
	ORDERS.mtx.Lock()         //Блокировка хранилища
	defer ORDERS.mtx.Unlock() //Запланируем разблокировку на момент выхода из функции
	ORDERS.data[id] = updts   //Обновление данных
}

func DeleteOrder(id uint32) { //Экв. DeleteOrder в пакете service
	ORDERS.mtx.Lock()         //Блокировка хранилища
	defer ORDERS.mtx.Unlock() //Запланируем разблокировку на момент выхода из функции
	delete(ORDERS.data, id)   //Удаляем объект
}

func GetOrderById(id uint32) entity.Order { //Экв. GetOrderById в пакете service
	ORDERS.mtx.Lock()         //Блокировка хранилища
	defer ORDERS.mtx.Unlock() //Запланируем разблокировку на момент выхода из функции
	return ORDERS.data[id]    //Возвращаем найденный объект
}

func GetOrderAll() []entity.Order { //Экв. GetOrderAll в пакете service
	ORDERS.mtx.Lock()               //Блокировка хранилища
	defer ORDERS.mtx.Unlock()       //Запланируем разблокировку на момент выхода из функции
	list := []entity.Order{}        //Подготавливаем массив, в который положем объекты
	for _, v := range ORDERS.data { //Переносим объекты
		list = append(list, v) //Копируем значения в массив
	}
	return list //Возвращаем массив
}
