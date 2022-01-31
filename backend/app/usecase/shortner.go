package usecase

import (
	"context"
	"net/url"
)

type Shortner struct {
}

func (s *Shortner) Index() error {
	//TODO implement me
	panic("implement me")
}

func (s *Shortner) SaveLink(ctx context.Context, url url.URL) error {
	//TODO implement me
	panic("implement me")
}

func (s *Shortner) ViewLink(ctx context.Context, hash string) {
	//TODO implement me
	panic("implement me")
}

func (s *Shortner) Redirect(ctx context.Context, url string) error {
	//TODO implement me
	panic("implement me")
}
