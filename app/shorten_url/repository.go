package shortenurl

import (
	"context"
	database "shorten-url/config"
	"time"

	"github.com/uptrace/bun"
)

type ShortenUrlModel struct {
	bun.BaseModel `bun:"shorten_url"`
	ID            int64
	ShortUrl      string
	Url           string
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type ShortenUrlRepository interface {
	Create(context.Context, *ShortenUrlModel) error
	GetByUrl(context.Context, string) (ShortenUrlModel, error)
	GetByShortUrl(context.Context, string) (ShortenUrlModel, error)
}

func NewShortenUrlRepository() ShortenUrlRepository {
	return new(repo)
}

type repo struct {
}

func (repo *repo) Create(ctx context.Context, shortenUrl *ShortenUrlModel) error {
	_, err := database.PGConnection.NewInsert().Model(shortenUrl).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) GetByUrl(ctx context.Context, url string) (ShortenUrlModel, error) {
	shortUrl := new(ShortenUrlModel)

	err := database.PGConnection.NewSelect().Model(shortUrl).Where("url = ?", url).Scan(ctx, shortUrl)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return *shortUrl, err
	}

	return *shortUrl, nil
}

func (repo *repo) GetByShortUrl(ctx context.Context, url string) (ShortenUrlModel, error) {
	shortUrl := new(ShortenUrlModel)

	err := database.PGConnection.NewSelect().Model(shortUrl).Where("short_url = ?", url).Scan(ctx, shortUrl)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return *shortUrl, err
	}

	return *shortUrl, nil
}
