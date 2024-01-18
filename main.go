package main

import (
	"fmt"
	"log"
	"main/handlers"
	"main/queries"
	"main/services"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwerty32"
	dbname   = "postgres"
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
	/* Иницилаизация базы */
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	pg, pgErr := queries.NewPG(connStr)
	if pgErr != nil {
		log.Fatalln(fmt.Errorf("app - Run - pgdb.NewServices: %w", pgErr))
	}
	defer pg.Close()
	log.Println(pg)
	/* Сервисы */
	accountService := services.NewAccountService(pg)
	/*  */
	/* Ручки */
	/* Проверка работоспособности */
	router.HandleFunc("/", handlers.HandlerGetHelloWorld)
	/* 1. Начисление денег на счет */
	router.HandleFunc("/AddMoneyToAccount", handlers.HandlerGetHelloWorld)
	/* 2. Резервирование средств POST */
	router.HandleFunc("/Operations", handlers.HandlerGetHelloWorld)
	/* 3. Cписание  средств PUT */
	router.HandleFunc("/", handlers.HandlerGetHelloWorld)
	/* 4.  Получение баланса GET */
	router.HandleFunc("/Accounts/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerAccounts(w, r, accountService)
	})
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Couldnt ListenAndServe()", err)
	}
}
