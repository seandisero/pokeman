package pokapi

import (
	"net/http"
	"time"

	"pokeman/internal/pokcache"
)

// Client -
type Client struct {
	httpClient http.Client
	cache      pokcache.Cache
}

// NewClient -
func NewClient(timeout, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokcache.NewCache(interval),
	}
}
