package main

import (
	"fmt"
	"log"
	. "net/http"
)

func main() {
	var myPort string
	fmt.Print("Введите порт: ")
	fmt.Scan(&myPort)
	// Обработчик запросов
	HandleFunc("/", func(w ResponseWriter, r *Request) {
		log.Println("Получен запрос на /")
		ServeFile(w, r, "./index.html")
	})
	HandleFunc("/home", func(w ResponseWriter, r *Request) {
		log.Println("Получен запрос на /home")
		ServeFile(w, r, "./home.html")
	})
	// Настройка и запуск сервера
	log.Println("Сервер запущен на http://localhost:" + myPort)
	if err := ListenAndServe(":"+myPort, nil); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}
