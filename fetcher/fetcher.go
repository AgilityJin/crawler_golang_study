package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 资源查询器
func Fetch(url string) ([]byte, error) {
	// 获取目标网页信息
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// 关闭
	defer resp.Body.Close()

	// 如果返回状态不成功
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求错误码: %d", resp.StatusCode)
	}
	// fmt.Println(reflect.TypeOf(resp.Body))
	e := determineEncoding(resp.Body)
	// 转utf-8编码
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	// 返回获取并转码好的内容
	return ioutil.ReadAll(utf8Reader)
}

// 返回网页编码格式
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	// 如果没有读取成功,返回UTF-8格式
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
