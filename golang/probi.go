package main

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func main() {
	http.HandleFunc("/reg", reghandler)
	http.HandleFunc("/auth", authhandle)
	http.ListenAndServe(":8000", nil)
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

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6IjEyMzQ1Njc4OTAiLCJwYXNzd29yZCI6IkpvaG4gRG9lIn0.qupq7x2Xp32sZbOk9wH49EefoYyiuyYEd8sG1rs_lfA
