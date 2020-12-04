package models

import (
	"context"
	"database/sql"
	"jobs-api/requests"
	"jobs-api/responses"
	"jobs-api/utils/databases"
	"jobs-api/utils/logger"
)

// Review model struct
// swagger:model Review
type Review struct {
	ID      int64   `db:"id"`
	Value   string  `db:"val"`
	Comment *string `db:"comment"`
}

// Create review
func Create(ctx context.Context, log logger.LogEnv, dbEnv databases.DBEnv, request requests.Review) (saved bool, err error) {
	db := dbEnv.PostgresDBWriter()

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		_, err := databases.NewDbEnv()
		if err != nil {
			return false, err
		}
		db = dbEnv.PostgresDBWriter()
	}

	tsql := `INSERT INTO reviews (val, comment)
			 VALUES (@Val, @Comment)`

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return
	}

	result, err := stmt.ExecContext(ctx,
		sql.Named("Val", request.Value),
		sql.Named("Comment", request.Comment))
	if err != nil {
		return
	}

	rowsAmount, err := result.RowsAffected()
	if err != nil {
		return
	}

	if rowsAmount > 0 {
		saved = true
	}

	return
}

// Show review
func Show(ctx context.Context, log logger.LogEnv, dbEnv databases.DBEnv, value string) ([]responses.Show, error) {
	reviews := make([]responses.Show, 0) // todo response validator
	var err error

	db := dbEnv.PostgresDBReader()

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		_, err := databases.NewDbEnv()
		if err != nil {
			return nil, err
		}
		db = dbEnv.PostgresDBReader()
	}

	tsql := `SELECT id FROM reviews WHERE val = @value`

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, sql.Named("value", value))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var review responses.Show
		err = rows.Scan(&review.ID, &review.Value, &review.Comment)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, err
}
