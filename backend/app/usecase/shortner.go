package usecase

import (
	"context"
	"github.com/Z00mZE/url-shortner/ent"
	"net/url"
	"time"
)

type converter interface {
	Encode(int64) string
	Decode(string) int64
}

type Shortner struct {
	userORMClient    *ent.UrlsClient
	decimalConverter converter
}

func NewShortner(orm *ent.UrlsClient, converter converter) *Shortner {
	return &Shortner{userORMClient: orm, decimalConverter: converter}
}

func (s *Shortner) SaveLink(ctx context.Context, userURL *url.URL) (ent.Urls, error) {
	//TODO implement me
	//panic("implement me")
	s.
		userORMClient.
		Create().
		SetURL(userURL.String()).
		SetExpiredAt(time.Now().AddDate(0, 1, 0)).
	return ent.Urls{
		ID:        0,
		URL:       "",
		ExpiredAt: time.Time{},
	}, nil
}

func (s *Shortner) ViewLink(ctx context.Context, s2 string) (ent.Urls, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Shortner) Redirect(ctx context.Context, s2 string) (ent.Urls, error) {
	//TODO implement me
	panic("implement me")
}
