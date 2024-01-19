package services

import (
	"fmt"
	. "main/models"
	"main/queries"
)

type AccountsService struct {
	repo queries.AccountsRepo
}

func (service *AccountsService) GetAccoundByUserId(userId string) (User, error) {
	fmt.Println("userId", userId)
	return service.repo.GetUserById(userId)
}
func NewAccountService(dbDeep queries.AccountsRepo) AccountsService {
	return AccountsService{dbDeep}
}
