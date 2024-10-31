package util

import (
	"github.com/xiaoyuer11223344/wangsu/common/model"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func Call(requestMsg model.HttpRequestMsg) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(requestMsg.Method, requestMsg.Url, strings.NewReader(requestMsg.Body))
	if err != nil {
		return "", err
	}

	for key := range requestMsg.Headers {
		req.Header.Set(key, requestMsg.Headers[key])
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
