package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
)

const (
	dbDSN = "host=localhost port=54322 dbname=auth-db user=auth password=qwerty1 sslmode=disable"
)

func insertAuthToDb(pool *pgxpool.Pool, ctx context.Context) {
	buildInsert := sq.Insert("auth").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "password", "created_at").
		Values(gofakeit.Username(), gofakeit.Password(true, true, true, true, false, 12), gofakeit.Date()).
		Suffix("RETURNING id")

	query, args, err := buildInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var noteId int
	err = pool.QueryRow(ctx, query, args...).Scan(&noteId)
	if err != nil {
		log.Fatalf("failed to insert note: %v", err)
	}

	log.Printf("inserted note with id: %d", noteId)
}

func getAuthValueFromDb(pool *pgxpool.Pool, ctx context.Context) {
	builderSelect := sq.Select("id", "name", "password", "created_at").
		From("auth").
		PlaceholderFormat(sq.Dollar).
		OrderBy("id ASC").
		Limit(10)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to select rows: %v", err)
	}
	defer rows.Close()

	var (
		id        int64
		name      string
		password  string
		createdAt time.Time
	)

	for rows.Next() {
		err = rows.Scan(&id, &name, &password, &createdAt)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		log.Printf("id: %d, name: %s, password: %s, created_at: %v", id, name, password, createdAt)
	}
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	insertAuthToDb(pool, ctx)
	getAuthValueFromDb(pool, ctx)
}
