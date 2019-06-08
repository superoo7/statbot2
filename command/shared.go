package command

import (
	"time"

	"github.com/superoo7/go-gecko/v3/types"
	"github.com/superoo7/statbot2/http"
)

var Coinlist *types.CoinList
var LastSavedCoinList time.Time

func LoadCoinList() error {
	now := time.Now()
	diff := now.Sub(LastSavedCoinList)
	days := int(diff.Hours() / 24)

	if days > 1 {
		cl, err := http.CG.CoinsList()
		if err != nil {
			return err
		}
		Coinlist = cl
		LastSavedCoinList = now
	}

	return nil
}

func IsCoinInList(coin string) (bool, types.CoinsListItem) {
	inList := true
	var cc types.CoinsListItem

	for _, c := range *Coinlist {
		if coin == c.ID || coin == c.Name || coin == c.Symbol {
			inList = false
			cc = c
			break
		}
	}
	return inList, cc
}
