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
			service.GetAccoundByUserId(idAccount)
			jsonResponse, _ := json.Marshal(nil)
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
