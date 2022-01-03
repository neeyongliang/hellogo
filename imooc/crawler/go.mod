module crawler

replace (
	gitee.com/neeyongliang/gogogo/imooc/crawler/engine => ./engine
	gitee.com/neeyongliang/gogogo/imooc/crawler/fetcher => ./fetcher
	gitee.com/neeyongliang/gogogo/imooc/crawler/model => ./model
	gitee.com/neeyongliang/gogogo/imooc/crawler/zhenai/parser => ./zhenai/parser
)

go 1.13
