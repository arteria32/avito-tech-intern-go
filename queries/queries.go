package queries

import (
	"context"
	"fmt"
	. "main/models"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPG(connString string) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(context.Background(), connString)
		if err != nil {
			fmt.Errorf("unable to create connection pool: %w", err)
			return
		}

		pgInstance = &Postgres{db}
	})

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}

func (pg *Postgres) FindRowByProrFromTable(propName, propValue, tableName string) (User, error) {
	queryStr := fmt.Sprintf("SELECT * FROM billing_service.%s WHERE %s ='%s'", tableName, propName, propValue)
	var user User
	// Executing query for single row
	if err := pg.db.QueryRow(context.Background(), queryStr).Scan(&user.Id, &user.UserId, &user.RealAccount, &user.ReservingAccount); err != nil {
		fmt.Println("Error occur while finding user: ", err)
		return user, err
	}
	fmt.Println(user)
	return user, nil
}
