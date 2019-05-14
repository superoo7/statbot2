package command

import (
	"time"

	"github.com/superoo7/go-gecko/v3/types"
	"github.com/superoo7/statbot2/coingecko"
)

var Coinlist *types.CoinList
var LastSavedCoinList time.Time

func LoadCoinList() {
	now := time.Now()
	diff := now.Sub(LastSavedCoinList)
	days := int(diff.Hours() / 24)

	if days > 1 {
		cl, _ := coingecko.CG.CoinsList()
		Coinlist = cl
		LastSavedCoinList = now
	}
}
