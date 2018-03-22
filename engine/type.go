package engine

type RequestTest struct{
	UrlTest string
	ParserFuncTest func([]byte) ParseResultTest
}

type ParseResultTest struct{
	RequestTest []RequestTest
	ItemsTest []interface{}
}

type Request struct{
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct{
	Requests []Request
	Items []interface{}
}