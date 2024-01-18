package services

import (
	"fmt"
	"main/queries"
)

type AccountsService struct {
	dbDeep *queries.Postgres
}

func (service *AccountsService) GetAccoundByUserId(userId string) {
	fmt.Println("userId", userId)
	service.dbDeep.FindRowByProrFromTable("user_id", userId, "accounts")
}
func NewAccountService(dbDeep *queries.Postgres) AccountsService {
	return AccountsService{dbDeep}
}
