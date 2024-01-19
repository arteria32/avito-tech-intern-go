package queries

import (
	"context"
	"fmt"
	. "main/models"
)

type AccountsRepo struct {
	*Postgres
}

func (repo *AccountsRepo) GetUserById(name string) (User, error) {
	queryStr := fmt.Sprintf("SELECT * FROM billing_service.accounts WHERE user_id ='%s'", name)
	var user User
	// Executing query for single row
	if err := repo.db.QueryRow(context.Background(), queryStr).Scan(&user.Id, &user.UserId, &user.RealAccount, &user.ReservingAccount); err != nil {
		fmt.Println("Error occur while finding user: ", err)
		return user, ErrNotFound
	}
	fmt.Println(user)
	return user, nil
}

func NewAccountsRepo(pg *Postgres) AccountsRepo {
	return AccountsRepo{pg}
}
