package ipaddress

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//Client 查询客户端
type Client struct {
	url    string
	parser func([]byte) (string, error)
}

//NewClient 实例化客户端
func NewClient(url string, parser func([]byte) (string, error)) *Client {
	return &Client{
		url:    url,
		parser: parser,
	}
}

//Query 查询
func (cli *Client) Query(method string, headers map[string]string,
	params map[string]interface{}, timeout int) (string, error) {
	var (
		httpMethod string
		payload    string
	)
	if strings.ToUpper(method) == "POST" {
		httpMethod = "POST"
		b, err := json.Marshal(params)
		if err != nil {
			payload = string(b)
		}
	} else {
		httpMethod = "GET"
	}
	request, err := http.NewRequest(httpMethod, cli.url, strings.NewReader(payload))
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		request.Header[k] = []string{v}
	}
	client := http.Client{Timeout: time.Second * time.Duration(timeout)}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer func() { _ = response.Body.Close() }()
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return cli.parser(responseByte)
}
