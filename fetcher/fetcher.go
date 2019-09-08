package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Fetcher(url string) ([]byte, error) {
	//防止太频繁被屏蔽
	time.Sleep(1 * time.Second)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong request response code:%d", resp.StatusCode)
	}
	e := determineEncoding(resp.Body)
	utf8reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8reader)
}

/**
*
 */
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
