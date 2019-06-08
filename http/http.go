package http

import (
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

// CG new CoinGecko http client
var CG *coingecko.Client

// HTTPClient for this app
var HTTPClient *http.Client

func init() {
	HTTPClient = &http.Client{
		Timeout: time.Second * 10,
	}
	CG = coingecko.NewClient(HTTPClient)
}
