package queries

import (
	"context"
	"fmt"
	. "main/models"
	"time"
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

func (repo *OperataionsRepo) GetOperationById(id int) (Operation, error) {
	queryStr := fmt.Sprintf("SELECT * FROM billing_service.operations WHERE id ='%d'", id)
	var operation Operation

	var creatingTime time.Time
	var updatingTime time.Time
	if err := repo.db.QueryRow(context.Background(), queryStr).Scan(&operation.Id, &operation.Cost, &operation.AccountId, &operation.ServiceId, &operation.HealthStatus, &creatingTime, &updatingTime); err != nil {
		fmt.Println("Error OperationById: ", err)
		return operation, ErrNotFound
	}
	operation.CreationDate = creatingTime.String()
	operation.UpdateDate = updatingTime.String()
	fmt.Println(operation)
	return operation, nil
}

func (repo *OperataionsRepo) UpdateExistingOperation(newState Operation) (Operation, error) {
	queryStr := fmt.Sprintf("UPDATE billing_service.operations SET health_status=%d,update_date='%s' WHERE id =%d", newState.HealthStatus, newState.UpdateDate, newState.Id)
	var operation Operation
	_, err := repo.db.Exec(context.Background(), queryStr)
	if err != nil {
		fmt.Println("Error OperationById: ", err)
		return operation, ErrNotFound
	}
	return newState, nil
}

func NewOperationsRepo(pg *Postgres) OperataionsRepo {
	return OperataionsRepo{pg}
}
