package httpbz

import (
	"net/http"

	"github.com/pkg/errors"
)

// Get Get请求
// 针对一些依赖 cookie 的请求, 可以传入 client
func Get(url string, client *http.Client) (data string, statusCode int, err error) {
	if client == nil {
		client = &http.Client{}
	}
	response, err := client.Get(url)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer closeBody(response.Body)
	statusCode = response.StatusCode
	data, err = readBody(response.Body)
	return
}
