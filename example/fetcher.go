package example

import (
	"context"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type HTTPFetcher struct {
	Client *http.Client
}

func (f *HTTPFetcher) FetchByID(ctx context.Context, ID string) (*User, error) {
	return nil, nil
}
