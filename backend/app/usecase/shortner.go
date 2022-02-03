package usecase

import (
	"context"
	"github.com/Z00mZE/url-shortner/ent/service"
	"net/url"
	"time"
)

type converter interface {
	Encode(int64) string
	Decode(string) int64
}

type Shortener struct {
	userORMClient    *service.ShortUrlClient
	decimalConverter converter
}

func NewShortener(orm *service.ShortUrlClient, converter converter) *Shortener {
	return &Shortener{userORMClient: orm, decimalConverter: converter}
}

func (s *Shortener) SaveLink(ctx context.Context, userURL *url.URL) (*service.ShortUrl, error) {
	entity, saveError := s.
		userORMClient.
		Create().
		SetURL(userURL.String()).
		SetExpiredAt(time.Now().AddDate(0, 1, 0)).
		Save(ctx)
	if saveError != nil {
		return nil, saveError
	}
	entity.HashID = s.decimalConverter.Encode(int64(entity.ID))
	return entity, nil
}

func (s *Shortener) FindLink(ctx context.Context, idHash string) (*service.ShortUrl, error) {
	shortUrlID := s.decimalConverter.Decode(idHash)

	entity, findError := s.
		userORMClient.
		Get(ctx, int(shortUrlID))

	if findError != nil {
		return nil, findError
	}
	entity.HashID = idHash
	return entity, nil
}
