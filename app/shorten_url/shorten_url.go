package shortenurl

import (
	"context"
	"math/rand"
	"strings"
	"time"
)

const (
	host           = "http://localhost:3000"
	alphabet       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphabetLength = uint64(len(alphabet))
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func EncodeShortenUrl(url string) ShortenUrlModel {
	shortenUrlRepo := NewShortenUrlRepository()
	ctx := context.Background()

	existShortUrl, err := shortenUrlRepo.GetByUrl(ctx, url)
	if err != nil {
		panic(err)
	}

	if existShortUrl == (ShortenUrlModel{}) {
		shortUrl := ShortenUrlModel{
			ShortUrl: generateShortUrl(),
			Url:      url,
		}

		err := shortenUrlRepo.Create(ctx, &shortUrl)
		if err != nil {
			panic(err)
		}

		return shortUrl
	}

	return existShortUrl
}

func DecodeShortenUrl(url string) ShortenUrlModel {
	shortenUrlRepo := NewShortenUrlRepository()
	ctx := context.Background()

	shortUrl, err := shortenUrlRepo.GetByShortUrl(ctx, url)
	if err != nil {
		panic(err)
	}

	if shortUrl == (ShortenUrlModel{}) {
		return ShortenUrlModel{}
	}

	return shortUrl
}

func generateShortUrl() string {
	randString := generateRandomString()

	return host + "/" + randString
}

func generateRandomString() string {
	num := rand.Uint64()

	var stringBuilder strings.Builder

	for ; num > 0; num = num / alphabetLength {
		stringBuilder.WriteByte(alphabet[(num % alphabetLength)])
	}

	return stringBuilder.String()
}
