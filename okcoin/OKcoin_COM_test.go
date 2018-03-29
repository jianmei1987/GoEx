package okcoin

import (
	"github.com/jianmei1987/GoEx"
	"net/http"
	"testing"
)

var okcom = NewCOM(http.DefaultClient, "", "")

func TestOKCoinCOM_API_GetTicker(t *testing.T) {
	ticker, _ := okcom.GetTicker(goex.BTC_CNY)
	t.Log(ticker)
}
