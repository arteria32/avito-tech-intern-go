package services

import (
	"errors"
	"fmt"
	. "main/models"
	"main/queries"
	"time"
)

var (
	ErrorNotEnoughMoney = errors.New("not enough money on account")
)

type OperationsService struct {
	accountsRepo   queries.AccountsRepo
	operationsRepo queries.OperataionsRepo
}

func (service *OperationsService) AddNewOperation(newBody Operation) (User, error) {
	fmt.Println("newBody", newBody)
	userId := newBody.AccountId
	curAccountState, errAcc := service.accountsRepo.GetUserById(userId)
	fmt.Println(curAccountState)
	if errAcc != nil {
		return curAccountState, errAcc
	}
	curAccount := curAccountState.RealAccount
	fmt.Println("compare", curAccount, newBody.Cost)

	if curAccount < newBody.Cost {
		return curAccountState, ErrorNotEnoughMoney
	}
	curAccountState.RealAccount = curAccountState.RealAccount - newBody.Cost
	curAccountState.ReservingAccount = curAccountState.ReservingAccount + newBody.Cost
	newAccountState, updatingError := service.accountsRepo.UpdateExistingUser(curAccountState)
	if updatingError != nil {
		return newAccountState, errAcc
	}
	newBody.HealthStatus = PendingStatus
	creatingDate := time.Now().UTC().Format(time.RFC3339Nano)
	newBody.CreationDate = creatingDate
	newBody.UpdateDate = creatingDate
	_, addingError := service.operationsRepo.AddNewOperation(newBody)
	if addingError != nil {
		return newAccountState, addingError
	}
	return newAccountState, nil
}
func NewOperationService(dbAccounts queries.AccountsRepo, dbOperations queries.OperataionsRepo) OperationsService {
	newOperService := OperationsService{
		operationsRepo: dbOperations,
		accountsRepo:   dbAccounts,
	}
	return newOperService
}
