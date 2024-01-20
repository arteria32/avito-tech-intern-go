package queries

import (
	"context"
	"fmt"
	. "main/models"
)

type OperataionsRepo struct {
	*Postgres
}

func (repo *OperataionsRepo) AddNewOperation(newInstance Operation) (Operation, error) {
	queryStr := fmt.Sprintf("INSERT INTO billing_service.operations (cost, account_id, service_id, health_status,creation_date,update_date) VALUES(%f,%d,%d,%d,'%s','%s')", newInstance.Cost, newInstance.AccountId, newInstance.ServiceId, newInstance.HealthStatus, newInstance.CreationDate, newInstance.UpdateDate)
	fmt.Print("queryStr", queryStr)
	// Executing query for single row

	res, err := repo.db.Exec(context.Background(), queryStr)
	if err != nil {
		fmt.Println("Error occur while updating user: ", err)
		return newInstance, ErrAlreadyExists
	}
	// Используем метод LastInsertId(), чтобы получить последний ID
	// созданной записи из таблицу snippets.
	fmt.Print(res)
	return newInstance, nil
}
func NewOperationsRepo(pg *Postgres) OperataionsRepo {
	return OperataionsRepo{pg}
}
