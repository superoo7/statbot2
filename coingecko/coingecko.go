package coingecko

import (
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

// CG new CoinGecko http client
var CG *coingecko.Client

func init() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG = coingecko.NewClient(httpClient)
}
