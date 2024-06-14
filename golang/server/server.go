package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct { // Структура данных пользователя
	Login string `json:"login"`
	Pswd  string `json:"password"`
}

const (
	JWTCODE = "123456789"
	MONGO   = "mongodb://127.0.0.1:27017"
)

func TrimFirstAndLast(s string) string {
	if len(s) > 44 {
		s = s[43 : len(s)-1]
	}
	return s
}

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

func findClient(login string) Client { // возвращает данные о пользоваете в виде структуры
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO)) // создаём дэфолтного клиента
	if err != nil {                                                  // проверяем ошибку если она есть
		log.Println(err)
		return Client{"", ""}
	}
	// создаём соединение
	err = client.Connect(context.TODO())
	if err != nil { // проверяем ошибку если она есть
		log.Println("findClient error to connect to client: ", err)
		return Client{"", ""}
	}
	// проверяем соединение
	err = client.Ping(context.TODO(), nil)
	if err != nil { // проверяем ошибку если она есть
		log.Println("findClient error to ping: ", err)
		return Client{"", ""}
	}
	// обращаемся к коллекции clients из базы tg
	collection := client.Database("promotion_guide").Collection("clients")
	// создаём фильтр по которму мы будем искать клиента. был взят именно ID потому что они не повторяются
	filter := bson.D{{"login", login}}
	// создаём переменную в которую будем записывать полученного клиента в результате поиска
	var result Client
	// собственно ищем
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil { // проверяем ошибку если она есть то возвращаем пустую структуру вида clients
		log.Println("findClient error to find: ", err)
		return Client{"", ""}
	}
	log.Println("Client was found")
	return result // возвращаем в виде структуры clients
}

func is_in_data(login string) bool { // Проверка существует пользоваетль с данным ID в бд
	curUser := findClient(login)
	if curUser.Login == "" {
		return false
	}
	return true
}

func to_boolean(s string) bool {
	if s == "true" {
		return true
	} else if s == "false" {
		return false
	}
	return false
}

func getData(col string) primitive.D { //
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
	filter := bson.M{}
	var result primitive.D
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	log.Println("gave successfuly")
	return result
}

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

func main() {
	http.HandleFunc("/reg", reghandler)
	http.HandleFunc("/auth", authhandle)
	http.HandleFunc("/getstr", getstrhandler)
	http.HandleFunc("/getall", gethandler)
	http.HandleFunc("/addalc", addalchandler)
	http.HandleFunc("/addprod", addprodhandler)
	http.HandleFunc("/addcandy", addcandyhandler)
	http.HandleFunc("/addbit", addbithandler)
	http.HandleFunc("/addmeat", addmeathandler)
	http.HandleFunc("/addcof", addcofhandler)
	http.HandleFunc("/addfeed", addfeedhandler)
	http.HandleFunc("/addpowder", addpowderhandler)
	http.HandleFunc("/adddes", adddeshandler)
	http.ListenAndServe(":8000", nil)
}

func getstrhandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getstr handler started")
	col := r.URL.Query().Get("type")
	log.Println(col)
	data := getData(col)
	fmt.Fprint(w, data)
}

func gethandler(w http.ResponseWriter, r *http.Request) {
	log.Println("get handler started")
	col := r.URL.Query().Get("type")
	log.Println(col)
	url := "http://localhost:8000/getstr?type=" + col
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	str := TrimFirstAndLast(sb)
	log.Println(str)

	fmt.Fprintf(w, str)
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

func reghandler(w http.ResponseWriter, r *http.Request) {
	log.Println("register handler started")

	token := r.URL.Query().Get("token")
	data := decodeValid(token)
	var client Client
	client.Login = data["login"].(string)
	client.Pswd = data["password"].(string)
	log.Println("before checking client")
	if is_in_data(client.Login) {
		fmt.Fprint(w, "такой пользователь уже существует")
		log.Println("такой пользователь уже существует", client.Login)
	} else {
		addClient(client.Login, client.Pswd)
		fmt.Fprint(w, "успешная регистрация")
		log.Println("успешная регистрация", client.Login)
	}
}

func authhandle(w http.ResponseWriter, r *http.Request) {
	log.Println("authrization handler started")
	token := r.URL.Query().Get("token")
	data := decodeValid(token)
	var client Client
	client.Login = data["login"].(string)
	client.Pswd = data["password"].(string)
	if is_in_data(client.Login) {
		if client.Pswd == findClient(client.Login).Pswd {
			fmt.Fprint(w, "успешная авторизация")
			log.Println("успешная авторизация", client.Login)
		} else {
			fmt.Fprint(w, "неверный пароль")
			log.Println("неверный пароль", client.Login)
		}
	} else {
		fmt.Fprint(w, "нет такого пользователя")
		log.Println("нет такого пользователя", client.Login)
	}
}

func decodeValid(tokenString string) jwt.MapClaims {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTCODE), nil
	})
	if err != nil { // проверяем ошибку если она есть
		log.Println(err)
	}
	return claims
}
