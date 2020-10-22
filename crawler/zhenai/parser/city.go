package parser

import (
	"learngo/gopl.io/crawler/types"
	"regexp"
)


var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func  ParseCity(content []byte) types.ParseResult {
	matches := cityRe.FindAllSubmatch(content, -1)
	result := types.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User " + name + ", " + string(m[1]))

		result.Requests = append(result.Requests, types.Request{
			Url:       string(m[1]),
			ParseFunc: func(bytes []byte) types.ParseResult {
				return ParseProfile(content, name)
			},
		})
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	return result
}