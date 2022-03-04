package req

import (
	"context"
	"fmt"
	"log"
	"rProject/config"
	"rProject/models"
	"strings"
	"time"
)

const (
	table          = "req"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Req, error) {

	var reqs []models.Req

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var req models.Req
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&req.ID,
			&req.Email,
			&req.Text,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		req.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		req.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		reqs = append(reqs, req)
	}

	return reqs, nil
}

func Insert(ctx context.Context, post models.Req) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	entries := strings.Split(post.Email, ",")
	for _, email := range entries {
		queryText := fmt.Sprintf("INSERT INTO %v (email, text, created_at, updated_at) values('%v','%v','%v','%v')", table,
			email,
			post.Text,
			time.Now().Format(layoutDateTime),
			time.Now().Format(layoutDateTime))

		_, err = db.ExecContext(ctx, queryText)
	}

	if err != nil {
		return err
	}
	return nil
}
