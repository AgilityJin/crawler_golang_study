package engine

import (
	"log"

	"../fetcher"
)

// 接收所有请求
func Run(seeds ...Request) {
	var requests []Request
	// 将所有种子加入处理队列
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// 取出第一个需要处理的请求
		r := requests[0]
		// 存下其他需要处理的请求
		requests = requests[1:]
		// 利用读取器取回资源
		body, err := fetcher.Fetch(r.Url)
		// 某一资源获取发生错误时,仅退出当前请求
		if err != nil {
			log.Printf("Fetcher: error "+"fetching url %s: %v", r.Url, err)
			continue
		}

		// 调用预设的 ParserFunc 提取body内需要的信息元素
		parseResult := r.ParserFunc(body)
		// 将得到的所有请求url加入处理队列
		requests = append(requests, parseResult.Requests...)
		// 遍历Items内存储的信息打印log
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
