package queries

import (
	"context"
	"fmt"
	. "main/models"
)

type AccountsRepo struct {
	*Postgres
}

func (repo *AccountsRepo) GetUserByName(name string) (User, error) {
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
func (repo *AccountsRepo) UpdateExistingUser(newInstance User) (User, error) {
	queryStr := fmt.Sprintf("UPDATE billing_service.accounts SET real_account=%f,reserving_account=%f WHERE id =%d", newInstance.RealAccount, newInstance.ReservingAccount, newInstance.Id)
	var user User
	// Executing query for single row
	_, err := repo.db.Exec(context.Background(), queryStr)
	if err != nil {
		fmt.Println("Error occur while updating user: ", err)
		return user, ErrNotFound
	}
	return repo.GetUserByName(newInstance.UserId)
}

func NewAccountsRepo(pg *Postgres) AccountsRepo {
	return AccountsRepo{pg}
}
