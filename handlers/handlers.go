package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"main/services"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerGetHelloWorld(wr http.ResponseWriter,
	req *http.Request) {
	fmt.Fprintf(wr, "Hello, World\n")
	log.Println(req.Method) // request method
	log.Println(req.URL)    // request URL
	log.Println(req.Header) // request headers
	log.Println(req.Body)   // request body)
}
func HandlerAccounts(wr http.ResponseWriter,
	req *http.Request,
	service services.AccountsService) {
	vars := mux.Vars(req)
	switch req.Method {
	case "GET":
		{
			idAccount := vars["id"]
			log.Println(idAccount)
			result, err := service.GetAccoundByUserId(idAccount)
			if err != nil {
				http.Error(wr, "Not found", http.StatusNoContent)
				return
			}
			jsonResponse, _ := json.Marshal(result)
			wr.WriteHeader(http.StatusOK)
			wr.Write(jsonResponse)
		}
	default:
		{
			http.Error(wr, "Not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

	log.Println(req.Method) // request method
	log.Println(req.URL)    // request URL
	log.Println(req.Header) // request headers
	log.Println(req.Body)   // request body)
}
func HandlerAccountsBalance(wr http.ResponseWriter,
	req *http.Request,
	service services.AccountsService) {
	vars := mux.Vars(req)
	switch req.Method {
	case "GET":
		{
			idAccount := vars["id"]
			log.Println(idAccount)
			result, err := service.GetAccoundByUserId(idAccount)
			if err != nil {
				http.Error(wr, "Not found", http.StatusNoContent)
				return
			}
			jsonResponse, _ := json.Marshal(result.RealAccount)
			wr.WriteHeader(http.StatusOK)
			wr.Write(jsonResponse)
		}
	case "POST":
		{
			log.Printf("post")
			idAccount := vars["id"]
			log.Println(idAccount)
			var money float64
			err := json.NewDecoder(req.Body).Decode(&money)
			if err != nil {
				http.Error(wr, "StatusBadRequest", http.StatusBadRequest)
				return
			}
			result, err := service.UpdateBalanceAccount(idAccount, money)
			if err != nil {
				http.Error(wr, "Not found", http.StatusNoContent)
				return
			}
			jsonResponse, _ := json.Marshal(result.RealAccount)
			wr.WriteHeader(http.StatusOK)
			wr.Write(jsonResponse)
		}
	default:
		{
			http.Error(wr, "Not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

}
