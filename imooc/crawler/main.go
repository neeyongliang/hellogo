package main

import (
	"gitee.com/neeyongliang/gogogo/imooc/crawler/engine"
	"gitee.com/neeyongliang/gogogo/imooc/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		//ParserFunc: parser.ParserCityList,
		ParserFunc: parser.ParseProfile,
	})
}
