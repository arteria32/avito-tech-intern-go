package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	. "main/models"
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

// HandlerAccounts godoc
// @Summary      AccountsInfo
// @Description  get Account Info by id
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  User
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /Accounts/{id} [get]
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

// HandlerBalance godoc
// @Summary      BalanceInfo
// @Description  get Balance Info by id
// @Tags         Balance
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  User
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /AccountsBalance/{id} [get]
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

// HandlerBalance godoc
// @Summary      Operations
// @Description  Create New Operation
// @Tags         Operations
// @Accept       json
// @Produce      json
// @Success      200  {object}  Operation
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /Operations [post]
func HandlerOperations(wr http.ResponseWriter,
	req *http.Request,
	service services.OperationsService) {
	vars := mux.Vars(req)
	switch req.Method {
	case "POST":
		{
			log.Printf("post")
			idAccount := vars["id"]
			log.Println(idAccount)
			var newOperation Operation
			err := json.NewDecoder(req.Body).Decode(&newOperation)
			if err != nil {
				http.Error(wr, "StatusBadRequest", http.StatusBadRequest)
				return
			}
			res, err := service.AddNewOperation(newOperation)
			if err != nil {
				http.Error(wr, "Bad Request", http.StatusBadRequest)
				return
			}
			jsonResponse, _ := json.Marshal(res)
			wr.WriteHeader(http.StatusOK)
			wr.Write(jsonResponse)
		}
	case "PUT":
		{
			log.Printf("post")
			idAccount := vars["id"]
			log.Println(idAccount)
			var newOperation Operation
			err := json.NewDecoder(req.Body).Decode(&newOperation)
			if err != nil {
				http.Error(wr, "StatusBadRequest", http.StatusBadRequest)
				return
			}
			res, err := service.UpdateOperationStatus(newOperation)
			if err != nil {
				http.Error(wr, "Bad Request", http.StatusBadRequest)
				return
			}
			jsonResponse, _ := json.Marshal(res)
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
