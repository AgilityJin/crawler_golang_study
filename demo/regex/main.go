package main

import (
	"fmt"
	"regexp"
)

const text = `my email is ccmouse@gmail.com@qq.com
my email is test@outlook.com@qq.com.cn
my email is 17621667884@163.com@qq.com`

func main()  {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9]+)`)
	// -1 代表提取所有匹配
	match := re.FindAllStringSubmatch(text, -1)

	for _,m := range match {
		fmt.Println(m)
	}
}