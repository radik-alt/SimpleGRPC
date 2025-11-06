package main

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
)

const (
	dbDSN = "host=localhost port=54323 dbname=chat-server-db user=chat_server password=qwerty1 sslmode=disable"
)

func insertAuthToDb(pool *pgxpool.Pool, ctx context.Context) {
	buildInsert := sq.Insert("chat").
		PlaceholderFormat(sq.Dollar).
		Columns("message", "created_at", "updated_at").
		Values(gofakeit.DomainName(), gofakeit.Date(), nil).
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
	builderSelect := sq.Select("id", "message", "created_at", "updated_at").
		From("chat").
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
		id         int64
		message    string
		createdAt  time.Time
		updated_at sql.NullTime
	)

	for rows.Next() {
		err = rows.Scan(&id, &message, &createdAt, &updated_at)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		log.Printf("id: %d, message: %s, createdAt: %s, updated_at: %v", id, message, createdAt, updated_at)
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
