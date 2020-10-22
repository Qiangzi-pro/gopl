package parser

import (
	"learngo/gopl.io/crawler/model"
	"learngo/gopl.io/crawler/types"
	"regexp"
	"strconv"
)

const profileRe = `href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)">([^<]+)</a>`
var hukouRe = regexp.MustCompile(`<div .* class="m-btn pink">籍贯:([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div .* class="m-btn purple">月收入:([^<]+)</div>`)


func  ParseProfile(content []byte, name string) types.ParseResult {

	profile := model.Profile{}
	profile.Name = name
	profile.Hokou = extractString(hukouRe, content)
	income, err := strconv.Atoi(extractString(incomeRe, content))
	if err != nil {

	} else {
		profile.Income = income

	}

	result := types.ParseResult{
		Items: []interface{}{profile},
	}
	// 爬取其他profile页面

	return result
}

func extractString(re *regexp.Regexp, content []byte) string {
	match := re.FindSubmatch(content)
	if len(match) > 1 {
		return string(match[1])
	}
	return ""
}
