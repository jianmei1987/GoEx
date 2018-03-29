package yunbi

import (
	"testing"
	"net/http"
	. "github.com/jianmei1987/GoEx"
)

var (
	yb = New(http.DefaultClient, "", "")
)

func TestYunBi_GetTicker(t *testing.T) {
	t.Log(yb.GetTicker(BTS_CNY))
	t.Log(yb.GetTicker(SC_CNY))
	t.Log(yb.GetTicker(EOS_CNY))
}
