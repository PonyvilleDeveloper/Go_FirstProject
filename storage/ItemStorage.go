package storage

import (
	"app/entity"
	"sync"
)

type ItemStorage struct {
	data map[uint32]entity.Item						//Мапа, контейнер данных
	iter uint32										//Индексатор для добавления новых
	mtx  sync.Mutex									//Блокировщик доступа к контейнеру, чтобы разные потоки не сломали его
}

var ITEMS ItemStorage								//Выделяем память под хранилище

func init() {
	ITEMS.data = make(map[uint32]entity.Item)		//Инициализация контейнера
}

func AddItem(item entity.Item) {					//Экв. CreateItem в пакете service
	ITEMS.mtx.Lock()								//Блокировка хранилища
	defer ITEMS.mtx.Unlock()						//Запланируем разблокировку на момент выхода из функции
	item.ItemId = ITEMS.iter						//Копируем итератор хранилища как id объекта
	ITEMS.data[ITEMS.iter] = item					//Добавляем объект в контейнер
	ITEMS.iter++									//Инкрементируем индексатор
}

func ChangeItem(id uint32, updts entity.Item) {		//Экв. UpdateItem в пакете service
	ITEMS.mtx.Lock()								//Блокировка хранилища
	defer ITEMS.mtx.Unlock()						//Запланируем разблокировку на момент выхода из функции
	ITEMS.data[id] = updts							//Обновление данных
}

func DeleteItem(id uint32) {						//Экв. DeleteItem в пакете service
	ITEMS.mtx.Lock()								//Блокировка хранилища
	defer ITEMS.mtx.Unlock()						//Запланируем разблокировку на момент выхода из функции
	delete(ITEMS.data, id)							//Удаляем объект
}

func GetItemById(id uint32) entity.Item {			//Экв. GetItemById в пакете service
	ITEMS.mtx.Lock()								//Блокировка хранилища
	defer ITEMS.mtx.Unlock()						//Запланируем разблокировку на момент выхода из функции
	return ITEMS.data[id]							//Возвращаем найденный объект
}

func GetItemAll() []entity.Item {					//Экв. GetItemAll в пакете service
	ITEMS.mtx.Lock()								//Блокировка хранилища
	defer ITEMS.mtx.Unlock()						//Запланируем разблокировку на момент выхода из функции
	list := []entity.Item{}							//Подготавливаем массив, в который положем объекты
	for _, v := range ITEMS.data {					//Переносим объекты
		list = append(list, v)						//Копируем значения в массив
	}
	return list										//Возвращаем массив
}