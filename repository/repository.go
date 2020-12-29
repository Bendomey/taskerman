package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

//Repository services to export from repository module
type Repository interface {
	Close()                                                                              /// to close database
	Insert(ctx context.Context, statement string, args ...interface{}) (bool, error)     /// to insert and return boolean database
	GetSingle(ctx context.Context, statement string, args ...interface{}) pgx.Row        /// return single row database
	GetAll(ctx context.Context, statement string, args ...interface{}) (pgx.Rows, error) /// return single row database
	DeleteSingle(ctx context.Context, statement string, args ...interface{}) error       /// to insert and return boolean database

}

//type to hold postgres db
type postgresqlRepository struct {
	db *pgxpool.Pool
}

// NewPostgresqlRepository takes in the url of db and returns a repository or an error
func NewPostgresqlRepository(url string) (Repository, error) {
	var err error
	var pool *pgxpool.Pool
	pool, err = pgxpool.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	log.Println("Taskerman connected to db successfully")

	return &postgresqlRepository{pool}, nil
}

// Close is used to close the db connection
func (r *postgresqlRepository) Close() {
	// close db
	r.db.Close()
}

// InsertAndReturn is used to insert into the db
func (r *postgresqlRepository) GetSingle(ctx context.Context, statement string, args ...interface{}) pgx.Row {
	// var row pgx.Row
	row := r.db.QueryRow(ctx, statement, args...)
	return row
}

// Insert is used to insert into the db
func (r *postgresqlRepository) Insert(ctx context.Context, statement string, args ...interface{}) (bool, error) {
	_, err := r.db.Exec(ctx, statement, args...)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteSingle is used to delete from db
func (r *postgresqlRepository) DeleteSingle(ctx context.Context, statement string, args ...interface{}) error {
	_, err := r.db.Exec(ctx, statement, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresqlRepository) GetAll(ctx context.Context, statement string, args ...interface{}) (pgx.Rows, error) {
	rows, err := r.db.Query(ctx, statement, args...)
	return rows, err
}
