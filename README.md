# shorturl
	短链接生成算法

## Installation

	$ go get github.com/kaimixu/shorturl
	
## Usage

	package main

	import (
		"fmt"			
		"github.com/kaimixu/shorturl"
	)

	func main() {
		url := "http://www.example.com?a=1&b=2&c=3&d=4"

		cb := func(url, keyword string) bool {
			// todo 查db或缓存判断keyword是否重复
		return true
		}

		domain := "http://shorturl.cn"
		keyword := shorturl.Generator(shorturl.CHARSET_ALPHANUMERIC, domain, url, cb)
		fmt.Println(keyword)
	}
