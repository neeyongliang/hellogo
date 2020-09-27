module crawler

replace (
	gitee.com/wikinee/gogogo/imooc/crawler/engine => ./engine
	gitee.com/wikinee/gogogo/imooc/crawler/fetcher => ./fetcher
	gitee.com/wikinee/gogogo/imooc/crawler/model => ./model
	gitee.com/wikinee/gogogo/imooc/crawler/zhenai/parser => ./zhenai/parser
)

go 1.13
