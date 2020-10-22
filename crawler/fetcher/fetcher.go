package fetcher

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"learngo/gopl.io/crawler/types"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type DefaultFetcher struct{}

var baseRe = regexp.MustCompile(`((http|https)://)?([w.]{0,4}[^/]+)/.*`)
func getUrlBase(url string) string {
	match := baseRe.FindStringSubmatch(url)
	if match != nil {
		return match[3]
	}
	return ""
}

var cookies = generateCookies()

func generateCookies() (ret []*http.Cookie) {
	const cookieValues = "sid=89f9f5b8-fba1-4b9e-abd0-e79298128a24; ec=lhpvKfOg-1601368795488-cdd2fc0facd5a-115353789; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601368810; notificationPreAuthorizeSwitch=157728; __channelId=905821%2C0; _exid=bQ3XNFXNaDAYg67K4SJS5RLdJ%2BzG7q5BzsH8jJRCTlElxrk%2Bo%2Fhy9Ua1fHQRlE4yZx6IjpEA3j5P2fFg2D2Phg%3D%3D; _efmdata=ScTRp5N13XaQrrDgzGwiG9T2Ds8BhOi4AONN3fJO61C6S1wSw52FCYhdD8EGKWK6kS4S44nZ8ZElJYWMF4qyX59dbj0MJjk240RMKP%2BorcI%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601696107"

	pairs := strings.Split(cookieValues, "; ")
	for _, pair := range pairs {
		arr := strings.Split(pair, "=")
		cookie := &http.Cookie{
			Name: arr[0],
			Value: arr[1],
		}
		ret = append(ret, cookie)
	}
	return ret
}

func (df *DefaultFetcher) Fetch(url string, headers map[string]string) (io.Reader, error) {
	req, _ := http.NewRequest(http.MethodGet,url, nil)
	req.Header.Add(
		"Connection",
		"keep-alive")

	req.Header.Add(
		"Cache-Control",
		"max-age=0")

	req.Header.Add(
		"Upgrade-Insecure-Requests",
		"1")

	req.Header.Add(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")

	req.Header.Add(
		"Accept",
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	req.Header.Add(
		"Accept-Language",
		"zh-CN,zh;q=0.9")

	req.Header.Add(
		"Cookie",
		"sid=89f9f5b8-fba1-4b9e-abd0-e79298128a24; ec=lhpvKfOg-1601368795488-cdd2fc0facd5a-115353789; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601368810; FSSBBIl1UgzbN7NO=5CLeVvQ4I6AfIWoD4CKw0TqyMBHNNN8sW.lqHx2Wn_qYIvcyF9Wt9g_Uhk1H3jNGaGuOoLE5ABEMGv6rPoWCewa; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601702950; FSSBBIl1UgzbN7NP=5U6GWxTegkcVqqqm04ie_mAti70s.MWUP.QZKcl2azI51hLtO7PBdSeNv501s6wiO0lBgZKO3ZshCgfb2na9CrY7aVtpg4xwNJLI9cmrz70qCel0WWS06DraXr3MlJt7f5gb.CU0u7.jGQbvZyXkNZRCHCyaaAeORBuhl.H3gI6ccUbcSvIWoSo732wuixyJHEXWEiIRBNxY2EBAy4Xam18At4sDSniDXMOjxxtASWmSRviJkZe9LPOfOdHKODF.Yohns87TtIqX5OtXJeNTPCE; _efmdata=ScTRp5N13XaQrrDgzGwiG9T2Ds8BhOi4AONN3fJO61C6S1wSw52FCYhdD8EGKWK6d0VmilqfV6BK3nPiXFLV9ZmtFEaIdrLpflyY6LRDBs0%3D; _exid=num6WVoUp5cCavZo6a%2BUfzspeEDnL%2BD89PmFrzQdEQJhzuVMn%2B8vUSwPvKqCq0v%2FK32g4%2FRNt%2BdKgE%2FCKSdzkg%3D%3D")

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	//for _, cookie := range cookies {
	//	req.AddCookie(cookie)
	//}

	fmt.Println(req.Header)

	var resp *http.Response

	client := http.Client{
		//CheckRedirect: func(req *http.Request, via []*http.Request) error {
		//	fmt.Println("Redirect:", req)
		//	return nil
		//},
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	all, _ := ioutil.ReadAll(resp.Body)
	return bytes.NewReader(all), nil
}

var defaultFetcher = DefaultFetcher{}

var rateLimiter = time.Tick(10 * time.Millisecond)

// 根据url 获取utf-8的网页内容
func Fetcher(url string, fetcher types.IGather, h map[string]string) ([]byte, error) {
	<-rateLimiter

	var reader io.Reader
	var err error

	if fetcher != nil {
		reader, err = fetcher.Fetch(url)
	} else {
		reader, err = defaultFetcher.Fetch(url, h)
	}

	if err != nil {
		return nil, err
	}

	bufferReader := bufio.NewReader(reader)

	e := determineEncoding(bufferReader)
	utf8Reader := transform.NewReader(bufferReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	content, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(content, "")
	//log.Println(e, name, certain)
	return e
}
