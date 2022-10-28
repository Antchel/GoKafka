package main

import (
	"context"
	"fmt"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	urlExample := "postgres://test:test@localhost:5432/test"
	// conn, err := pgx.Connect(context.Background(), urlExample)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }

	sql := sq.Select("*").From("t4")
	sql = sql.Where(sq.Eq{"A": 1})
	fmt.Println(sql.ToSql())
	ctx := context.Background()
	db, err := pgx.Connect(ctx, urlExample)
	// defer conn.Close(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	var version string
	_ = db.QueryRow(ctx, "select version()").Scan(&version)

	fmt.Println(version)
	// rows, err := conn.Query(context.Background(), "select * from node")

	// rows.Next()
	// fmt.Println(rows.Values())
}
