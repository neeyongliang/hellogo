package parser

import (
	"fmt"
	"regexp"

	"gitee.com/wikinee/gogogo/imooc/crawler/engine"
)

const CityListRE = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]+>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(CityListRE)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		// index 0 is full string, index 1 is URL, index 2 is City
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	fmt.Printf("result len: %d", len(result.Items))
	return result
}
