package parser

import (
	"learngo/gopl.io/crawler/types"
	"regexp"
)

const cityListRe = `href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)">([^<]+)</a>`

func  ParseCityList(content []byte) types.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(content, -1)
	result := types.ParseResult{}

	// 方便测试，限制入队城市数量
	limit := 10
	for _, m := range matches {
		result.Requests = append(result.Requests, types.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
		result.Items = append(result.Items, "City " + string(m[2]) + ", " + string(m[1]))
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
