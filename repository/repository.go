package repository

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

//Repository services to export from repository module
type Repository interface {
	Close()                                                                                   /// to close database
	Insert(ctx context.Context, statement string, args ...interface{}) (pgtype.Record, error) /// to close database
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

// Close is used to insert into the db
func (r *postgresqlRepository) Insert(ctx context.Context, statement string, args ...interface{}) (pgtype.Record, error) {
	var row pgtype.Record
	err := r.db.QueryRow(ctx, statement, args).Scan(&row)
	return row, err
}
