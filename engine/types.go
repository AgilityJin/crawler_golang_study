package engine

// 定义请求数据结构体
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// 定义请求结果结构体
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
