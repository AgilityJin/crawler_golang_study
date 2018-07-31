package parser

import (
	"regexp"

	"../../engine/"
)

// 提取正则
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 城市解析器
func ParseCityList(contents []byte) engine.ParseResult {
	// 设置正则,
	re := regexp.MustCompile(cityListRe)
	// 提取所有匹配项, -1表示拿全部
	matches := re.FindAllSubmatch(contents, -1)

	// 初始化返回结构
	result := engine.ParseResult{}

	// 处理返回数据信息
	for _, m := range matches {
		// 添加城市名称至 Items
		result.Items = append(result.Items, string(m[2]))
		// 添加请求处理的url及指定解析器,
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
