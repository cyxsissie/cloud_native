package main

type HTTPReqInfo struct {
	// GET 等方法
	method string
	uri    string
	ipaddr string
	// 响应状态码，如 200，204
	code string
}
