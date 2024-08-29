package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(strings.ReplaceAll(err.Error(), ":", "\n\t:"))
		os.Exit(1)
	}
}

//goland:noinspection SqlResolve
func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return fmt.Errorf("run: %w", err)
	}
	defer db.Close()

	newRows := func() *sqlmock.Rows {
		return mock.NewRows([]string{"customer_id", "customer_name", "country"}).
			AddRow(1, "Alfreds Futterkiste", "Germany").
			AddRow(2, "Ana Trujillo Emparedados y helados", "Mexico").
			AddRow(3, "Antonio Moreno Taquería", "Mexico").
			AddRow(4, "Around the Horn", "UK").
			AddRow(5, "Berglunds snabbköp", "Sweden")
	}

	mock.ExpectQuery("SELECT customer_id, customer_name, country FROM customers").WillReturnRows(newRows())
	mock.ExpectQuery("SELECT customer_id, customer_name, country FROM customers").
		WillReturnRows(newRows().RowError(3, context.DeadlineExceeded))

	if err := demo(ctx, db); err != nil {
		return fmt.Errorf("run: %w", err)
	}

	return nil
}

// START DEMO // OMIT
func demo(ctx context.Context, db *sql.DB) error {
	fmt.Println("1st attempt:")
	if err := printCustomers(ctx, db); err != nil {
		return fmt.Errorf("demo: #1: %w", err)
	}
	fmt.Println("2nd attempt:")
	if err := printCustomers(ctx, db); err != nil {
		return fmt.Errorf("demo: #2: %w", err)
	}
	return nil
}

func printCustomers(ctx context.Context, db *sql.DB) error {
	cc, err := fetchCustomers(ctx, db)
	if err != nil {
		return fmt.Errorf("print customers: %w", err)
	}

	for _, c := range cc {
		fmt.Printf("\t%d: %s (%s)\n", c.ID, c.Name, c.Country)
	}

	return nil
}

// END DEMO // OMIT

//goland:noinspection SqlResolve // START FETCH // OMIT
func fetchCustomers(ctx context.Context, db *sql.DB) ([]Customer, error) {
	qry := "SELECT customer_id, customer_name, country FROM customers"
	rows, err := db.QueryContext(ctx, qry)
	if err != nil {
		return nil, fmt.Errorf("fetch customers: execute query: %w", err)
	}

	defer rows.Close()

	var cc []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.Country); err != nil {
			return nil, fmt.Errorf("fetch customers: scan rows: %w", err)
		}

		cc = append(cc, c)
	}

	return cc, nil
}

// END FETCH // OMIT

type Customer struct {
	ID      int
	Name    string
	Country string
}
