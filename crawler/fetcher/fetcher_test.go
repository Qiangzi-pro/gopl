package fetcher

import (
	"encoding/json"
	"testing"
)

func TestFetcher(t *testing.T) {
	url := "https://album.zhenai.com/u/113640880"
	h := map[string]string {
		"authority": getUrlBase(url),
		"cache-control": "max-age=0",
		"upgrade-insecure-requests": "1",
		"sec-fetch-site": "cross-site",
		"sec-fetch-mode": "navigate",
		"sec-fetch-user": "?1",
		"sec-fetch-dest": "document",
		"referer": "http://www.zhenai.com/",
		"cookie": "sid=89f9f5b8-fba1-4b9e-abd0-e79298128a24; ec=lhpvKfOg-1601368795488-cdd2fc0facd5a-115353789; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601368810; FSSBBIl1UgzbN7NO=5CLeVvQ4I6AfIWoD4CKw0TqyMBHNNN8sW.lqHx2Wn_qYIvcyF9Wt9g_Uhk1H3jNGaGuOoLE5ABEMGv6rPoWCewa; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601702950; FSSBBIl1UgzbN7NP=5U6GWxTegkcVqqqm04ie_mAti70s.MWUP.QZKcl2azI51hLtO7PBdSeNv501s6wiO0lBgZKO3ZshCgfb2na9CrY7aVtpg4xwNJLI9cmrz70qCel0WWS06DraXr3MlJt7f5gb.CU0u7.jGQbvZyXkNZRCHCyaaAeORBuhl.H3gI6ccUbcSvIWoSo732wuixyJHEXWEiIRBNxY2EBAy4Xam18At4sDSniDXMOjxxtASWmSRviJkZe9LPOfOdHKODF.Yohns87TtIqX5OtXJeNTPCE; _efmdata=ScTRp5N13XaQrrDgzGwiG9T2Ds8BhOi4AONN3fJO61C6S1wSw52FCYhdD8EGKWK6d0VmilqfV6BK3nPiXFLV9ZmtFEaIdrLpflyY6LRDBs0%3D; _exid=num6WVoUp5cCavZo6a%2BUfzspeEDnL%2BD89PmFrzQdEQJhzuVMn%2B8vUSwPvKqCq0v%2FK32g4%2FRNt%2BdKgE%2FCKSdzkg%3D%3D",
	}

	content, err := Fetcher(url, nil, h)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%s\n", content)
	//req, err := http.Get(url)
	//if err != nil {
	//	t.Log(req, err)
	//}
	//t.Log(req)
	//
	//t.Log(req.Request.Header)

	//fmt.Println(getUrlBase("http://www.zhenai.com/zhenghun/henan"))
	json.Marshal()
}
