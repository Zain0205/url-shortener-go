package repository

import (
	"context"

	"github.com/Zain0205/url-shortener-go/app/db"
	"github.com/Zain0205/url-shortener-go/app/model"
)

func SaveURL(url model.URL) error {
	_, err := db.Pool.Exec(context.Background(),
		"INSERT INTO urls (short_code, original_url) VALUES ($1, $2)",
		url.ShortCode, url.OriginalURL,
	)
	return err
}

func FindByShortCode(code string) (string, error) {
	var original string
	err := db.Pool.QueryRow(context.Background(),
		"SELECT original_url FROM urls WHERE short_code=$1", code).Scan(&original)
	if err != nil {
		return "", err
	}
	return original, nil
}

func IncrementClickCount(code string) error {
	_, err := db.Pool.Exec(context.Background(),
		"UPDATE urls SET click_count = click_count + 1 WHERE short_code=$1", code)
	return err
}
