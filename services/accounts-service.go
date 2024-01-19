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
	return service.repo.GetUserByName(userId)
}
func (service *AccountsService) UpdateBalanceAccount(userId string, adding float64) (User, error) {
	user, errGet := service.repo.GetUserByName(userId)
	if errGet != nil {
		return user, errGet
	}

	newBalance := user.RealAccount + adding
	user.RealAccount = newBalance
	return service.repo.UpdateExistingUser(user)
}

func NewAccountService(dbDeep queries.AccountsRepo) AccountsService {
	return AccountsService{dbDeep}
}
