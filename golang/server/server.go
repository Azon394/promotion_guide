package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct { // Структура данных пользователя
	Name      string      `json:"name"`
	Daystitle string      `json:"daystitle"`
	Shops_ids []string    `json:"shops_ids"`
	Imagefull map[any]any `json:"imagefull"`
}

type Itemsa struct { // Структура данных пользователя
	Products []Item `json:"products"`
}

type Client struct { // Структура данных пользователя
	Login string `json:"login"`
	Pswd  string `json:"password"`
}

const (
	JWTCODE = "123456789"
	MONGO   = "mongodb://127.0.0.1:27017"
)

func addClient(login, pswd string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO)) // создаём дэфолтного клиента
	if err != nil {                                                  // проверяем ошибку если она есть
		log.Println(err)
	}
	// создаём соединение
	err = client.Connect(context.TODO())
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	// проверяем соединение
	err = client.Ping(context.TODO(), nil)
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	// обращаемся к коллекции clients из базы tg
	collection := client.Database("promotion_guide").Collection("clients")
	// создаём переменную в виде структуры clients
	current_client := Client{login, pswd}
	// добавляем одиночный документ в коллекцию
	insertResult, err := collection.InsertOne(context.TODO(), current_client)
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	// выводим внутренний ID добавленного документа
	log.Println("Inserted a single document: ", insertResult.InsertedID)
} // Функция добавления данных нового клиента по умолчанию в бд

func udateAlc(bsonData bson.M, col string) { //
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO)) // создаём дэфолтного клиента
	if err != nil {                                                  // проверяем ошибку если она есть
		log.Println(err)
	}
	// создаём соединение
	err = client.Connect(context.TODO())
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	// проверяем соединение
	err = client.Ping(context.TODO(), nil)
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	// обращаемся к коллекции clients из базы tg
	collection := client.Database("promotion_guide").Collection(col)

	// Условие для обновления (например, установим поле "status" в "active" для всех документов)
	filter := bson.M{} // Пустой фильтр выбирает все документы

	// Выполнение обновления
	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	insertResult, err := collection.InsertOne(context.TODO(), bsonData)
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	log.Println(insertResult)

}

func json_to_bson(jsonStr string) bson.M {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return bson.M(data)
	}

	// Convert map to bson.M
	bsonData := bson.M(data)
	return bsonData
}

func to_boolean(s string) bool {
	if s == "true" {
		return true
	} else if s == "false" {
		return false
	}
	return false
}

func main() {
	http.HandleFunc("/addalc", addalchandler)
	http.HandleFunc("/addprod", addprodhandler)
	http.HandleFunc("/addcandy", addcandyhandler)
	http.HandleFunc("/addbit", addbithandler)
	http.HandleFunc("/addmeat", addmeathandler)
	http.HandleFunc("/addcof", addcofhandler)
	http.HandleFunc("/addfeed", addfeedhandler)
	http.HandleFunc("/addpowder", addpowderhandler)
	http.HandleFunc("/adddes", adddeshandler)
	http.ListenAndServe(":6969", nil)
}

func adddeshandler(w http.ResponseWriter, r *http.Request) {
	log.Println("adddesert handler started")

	body, err := ioutil.ReadAll(r.Body)
	var data string
	err = json.Unmarshal(body, &data)
	if err == nil {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(data), &obj)
		if err != nil {
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			log.Println(err)
			return
		}
		bsonData := json_to_bson(string(jsonStr))
		//log.Println("\n", bsonData)
		udateAlc(bsonData, "desert")
	} else {
		log.Println(err)
	}
}

func addpowderhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("addpowder handler started")

	body, err := ioutil.ReadAll(r.Body)
	var data string
	err = json.Unmarshal(body, &data)
	if err == nil {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(data), &obj)
		if err != nil {
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			log.Println(err)
			return
		}
		bsonData := json_to_bson(string(jsonStr))
		//log.Println("\n", bsonData)
		udateAlc(bsonData, "powder")
	} else {
		log.Println(err)
	}
}

func addfeedhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("addfeed handler started")

	body, err := ioutil.ReadAll(r.Body)
	var data string
	err = json.Unmarshal(body, &data)
	if err == nil {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(data), &obj)
		if err != nil {
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			log.Println(err)
			return
		}
		bsonData := json_to_bson(string(jsonStr))
		//log.Println("\n", bsonData)
		udateAlc(bsonData, "feed")
	} else {
		log.Println(err)
	}
}

func addcofhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("addcoffee handler started")

	body, err := ioutil.ReadAll(r.Body)
	var data string
	err = json.Unmarshal(body, &data)
	if err == nil {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(data), &obj)
		if err != nil {
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			log.Println(err)
			return
		}
		bsonData := json_to_bson(string(jsonStr))
		//log.Println("\n", bsonData)
		udateAlc(bsonData, "coffee")
	} else {
		log.Println(err)
	}
}

func addmeathandler(w http.ResponseWriter, r *http.Request) {
	log.Println("addmeat handler started")

	body, err := ioutil.ReadAll(r.Body)
	var data string
	err = json.Unmarshal(body, &data)
	if err == nil {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(data), &obj)
		if err != nil {
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			log.Println(err)
			return
		}
		bsonData := json_to_bson(string(jsonStr))
		//log.Println("\n", bsonData)
		udateAlc(bsonData, "meat")
	} else {
		log.Println(err)
	}
}

func addbithandler(w http.ResponseWriter, r *http.Request) {
	log.Println("addbitovuha handler started")

	body, err := ioutil.ReadAll(r.Body)
	var data string
	err = json.Unmarshal(body, &data)
	if err == nil {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(data), &obj)
		if err != nil {
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			log.Println(err)
			return
		}
		bsonData := json_to_bson(string(jsonStr))
		//log.Println("\n", bsonData)
		udateAlc(bsonData, "bitovuha")
	} else {
		log.Println(err)
	}
}

func addcandyhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("addcandy handler started")

	body, err := ioutil.ReadAll(r.Body)
	var data string
	err = json.Unmarshal(body, &data)
	if err == nil {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(data), &obj)
		if err != nil {
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			log.Println(err)
			return
		}
		bsonData := json_to_bson(string(jsonStr))
		//log.Println("\n", bsonData)
		udateAlc(bsonData, "candy")
	} else {
		log.Println(err)
	}
}

func addprodhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("addprod handler started")

	body, err := ioutil.ReadAll(r.Body)
	var data string
	err = json.Unmarshal(body, &data)
	if err == nil {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(data), &obj)
		if err != nil {
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			log.Println(err)
			return
		}
		bsonData := json_to_bson(string(jsonStr))
		//log.Println("\n", bsonData)
		udateAlc(bsonData, "product")
	} else {
		log.Println(err)
	}
}

func addalchandler(w http.ResponseWriter, r *http.Request) {
	log.Println("adddata handler started")

	body, err := ioutil.ReadAll(r.Body)
	var data string
	err = json.Unmarshal(body, &data)
	if err == nil {
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(data), &obj)
		if err != nil {
			log.Println(err)
			return
		}

		jsonStr, err := json.Marshal(obj)
		if err != nil {
			log.Println(err)
			return
		}
		bsonData := json_to_bson(string(jsonStr))
		//log.Println("\n", bsonData)
		udateAlc(bsonData, "alc")
	} else {
		log.Println(err)
	}
}
