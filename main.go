package main

import (
	"./engine"
	"./zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		// 指定了该结构体的解析器为 ParseCityList
		ParserFunc: parser.ParseCityList,
	})
}
