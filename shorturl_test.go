package shorturl

import (
	"testing"
)

func TestShorturl(t *testing.T) {
	url := "http://www.example.com?a=1&b=2&c=3&d=4"

	cb := func(url, keyword string) bool {
		// todo 查db或缓存判断keyword是否重复
		return true
	}

	domain := "http://shorturl.cn"
	keyword := Generator(CHARSET_ALPHANUMERIC, domain, url, cb)
	if keyword == "" {
		t.Fatalf("Failed: generator shorturl, url[%s]", url)
	}
}
