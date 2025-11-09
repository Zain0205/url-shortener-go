package service

import (
	"github.com/Zain0205/url-shortener-go/app/db"
	"github.com/Zain0205/url-shortener-go/app/model"
	"github.com/Zain0205/url-shortener-go/app/repository"
	"github.com/Zain0205/url-shortener-go/app/utils"
)

func ShortenURL(original string) (string, error) {
	code := utils.GenerateShortCode(6)

	url := model.URL{
		ShortCode:   code,
		OriginalURL: original,
	}

	err := repository.SaveURL(url)
	if err != nil {
		return "", err
	}

	db.Redis.Set(db.Ctx, code, original, 0)
	return code, nil
}

func ResolveURL(code string) (string, error) {
	val, err := db.Redis.Get(db.Ctx, code).Result()
	if err == nil {
		repository.IncrementClickCount(code)
		return val, nil
	}

	original, err := repository.FindByShortCode(code)
	if err != nil {
		return "", err
	}

	db.Redis.Set(db.Ctx, code, original, 0)
	repository.IncrementClickCount(code)
	return original, nil
}
