package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

//Repository services to export from repository module
type Repository interface {
	Close() /// to close database
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

	return &postgresqlRepository{pool}, nil
}

// Close is used to close the db connection
func (r *postgresqlRepository) Close() {
	// close db
	r.db.Close()
}