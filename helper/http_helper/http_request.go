package http_helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-echo/model/base"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"moul.io/http2curl"
)

type HttpRequestStruct struct {
	Request      *http.Request
	Response     *http.Response
	Db           *gorm.DB
	Logger       *zap.Logger
	ErrorMessage string
}

type HttpRequest interface {
	PerformRequest(
		method string,
		url string,
		body []byte,
		bodyQuery []byte,
		headers map[string]string,
	) (int, []byte, map[string]interface{}, error)
}

func NewHttpRequest() HttpRequest {
	return &HttpRequestStruct{}
}

func (hr *HttpRequestStruct) PerformRequest(
	method string,
	url string,
	body []byte,
	bodyQuery []byte,
	headers map[string]string,
) (int, []byte, map[string]interface{}, error) {
	errorLog := "HttpRequest.PerformRequest.error"
	infoLog := "HttpRequest.PerformRequest.info"
	client := NewClient()

	var dataResponse = make(map[string]interface{})
	dataResponse = map[string]interface{}{
		"URL":            url,
		"RequestHeader":  "",
		"Request":        "",
		"ResponseHeader": "",
		"Response":       "",
	}

	var req *Request
	var err error
	if method == "GET" {
		req, err = NewRequest(method, url, bytes.NewReader(body))
		if err != nil {
			base.LoggerHttpClient(errorLog, fmt.Sprintf("%v", err))
			return 0, nil, dataResponse, err
		}

		if bodyQuery != nil {
			var params map[string]string
			_ = json.Unmarshal(bodyQuery, &params)
			strQuery := urlEncode(params)

			req.URL.RawQuery = strQuery
		}
	} else {
		req, err = NewRequest(method, url, bytes.NewReader(body))
		if err != nil {
			base.LoggerHttpClient(errorLog, fmt.Sprintf("%v", err))
			return 0, nil, dataResponse, err
		}
	}

	for key, val := range headers {
		req.Header.Add(key, val)
	}
	jsonHeaders, _ := json.Marshal(headers)

	dataResponse["RequestHeader"] = string(jsonHeaders)
	dataResponse["Request"] = string(body)

	//Log CURL
	command, _ := http2curl.GetCurlCommand(req.Request)
	base.LoggerHttpClient(infoLog, fmt.Sprintf("%v", command))

	resp, err := client.Do(req)
	if err != nil {
		base.LoggerHttpClient(errorLog, fmt.Sprintf("%v", err))
		return 0, nil, dataResponse, err
	}

	dataRsh, _ := json.Marshal(resp.Header)
	dataRs, _ := io.ReadAll(resp.Body)
	dataResponse["ResponseHeader"] = string(dataRsh)
	dataResponse["Response"] = string(dataRs)

	defer resp.Body.Close()

	return resp.StatusCode, dataRs, dataResponse, nil
}

func urlEncode(data map[string]string) string {
	params := url.Values{}
	for k, v := range data {
		params.Add(k, v)
	}
	return params.Encode()
}
