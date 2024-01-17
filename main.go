package main

import (
	"log"
	"main/handlers"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starts")
	router := mux.NewRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         ":3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	/* Проверка работоспособности */
	router.HandleFunc("/", handlers.HandlerGetHelloWorld)
	/* 1. Начисление денег на счет */
	router.HandleFunc("/AddMoneyToAccount", handlers.HandlerGetHelloWorld)
	/* 2. Резервирование средств POST */
	router.HandleFunc("/Operations", handlers.HandlerGetHelloWorld)
	/* 3. Cписание  средств PUT */
	router.HandleFunc("/", handlers.HandlerGetHelloWorld)
	/* 4.  Получение баланса GET */
	router.HandleFunc("/Accounts", handlers.HandlerGetHelloWorld)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Couldnt ListenAndServe()", err)
	}
}
