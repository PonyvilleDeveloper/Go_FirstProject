package main

/* ИДЕЯ URL'ОВ
/page - запрос страницы
/item/7000 - запрос item'а с id 7000
/cart/all - запрос всех cart'ов
/img/Image.jpg - запрос картинки (других ресурсов - аналогично)
*/
import (
	"app/service"
	"net/http"
	"strconv"
	"strings"
)

func getById(ent_name string, id uint32) (data []byte) { //Обработка HTTP.GET
	switch ent_name { //По переданному имени возвращаем
	case "item": //конкретный экземпляр нужной сущности
		data = service.GetItemById(id)
	case "order":
		data = service.GetOrderById(id)
	case "user":
		data = service.GetUserById(id)
	case "cart":
		data = service.GetCartById(id)
	}
	return
}

func getAll(ent_name string) (data []byte) { //Обработка HTTP.GET этап 2
	switch ent_name { //По переданному имени возвращаем
	case "item": //все экземпляры нужной сущности
		data = service.GetItemAll()
	case "order":
		data = service.GetOrderAll()
	case "user":
		data = service.GetUserAll()
	case "cart":
		data = service.GetCartAll()
	}
	return
}

func handleGet(ent_name string, id uint32) []byte { //Обработка HTTP.GET 1 этап
	if id != 0 { //Если id != 0, служебное значение
		return getById(ent_name, id) //Мы можем вернуть сущность по нему
	} else { //Иначе
		return getAll(ent_name) //Мы возвращаем все
	}
}

func handlePut(ent_name string, id uint32, data []byte) { //Обработка HTTP.PUT
	switch ent_name { //В зависимости от имени сущности
	case "item": //Обновляем нужный экземпляр
		service.UpdateItem(id, data)
	case "order":
		service.UpdateOrder(id, data)
	case "user":
		service.UpdateUser(id, data)
	case "cart":
		service.UpdateCart(id, data)
	}
}

func handlePost(ent_name string, data []byte) { //Обработка HTTP. POST
	switch ent_name { //По имени сущности
	case "item": //Создаём новый экземпляр с данными
		service.CreateItem(data)
	case "order":
		service.CreateOrder(data)
	case "user":
		service.CreateUser(data)
	case "cart":
		service.CreateCart(data)
	}
}

func handleDel(ent_name string, id uint32) { //Обработка HTTP.DELETE
	switch ent_name { //В зависимости от имени сущности
	case "item": //Удаляем заданный экземпляр
		service.DeleteItem(id)
	case "order":
		service.DeleteOrder(id)
	case "user":
		service.DeleteUser(id)
	case "cart":
		service.DeleteCart(id)
	}
}

func mapping(w http.ResponseWriter, r *http.Request) { //Маппинг
	args := strings.Split(r.URL.Path, "/")        //Разделяем URL
	ent_name := args[1]                           //Получаем имя сущности, с которой работаем
	id, err := strconv.ParseUint(args[2], 10, 64) //Пытаемся получить id
	if err != nil {                               //если его нет
		id = 0 //Он будет = 0, служебное знаечение
	}
	var input []byte            //Память под данные с клиента
	_, err = r.Body.Read(input) //Читаем их из запроса
	if err != nil {             //Если ошибка
		input = make([]byte, 0) //Данные пусты
	}
	switch ent_name { //Далее всё зависит от имени сущности
	case "item", "cart", "order", "user": //Если запрос именно сущности
		switch r.Method { //Определяем метод и вызываем соответсвующий обработчик
		case "GET":
			w.Write(handleGet(ent_name, uint32(id)))
		case "PUT":
			handlePut(ent_name, uint32(id), input)
		case "POST":
			handlePost(ent_name, input)
		case "DELETE":
			handleDel(ent_name, uint32(id))
		}
	case "js", "css", "img": //Если запрошенная сущность - какой-то ресурс
		http.ServeFile(w, r, r.URL.Path) //Просто выдаём соответствующий файл
	default: //В любом ином случае, мы запрашиваем страницу
		http.ServeFile(w, r, "pages"+r.URL.Path+".html") //И выдаём её
	}
}
