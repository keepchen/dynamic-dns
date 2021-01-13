package client

//Response 返回结构体
type Response struct {
	ErrCode string                 `json:"err_code"`
	ErrMsg  string                 `json:"err_msg"`
	Content map[string]interface{} `json:"content"`
}

//NewClient 实例化
func NewClient(cli IClient) *Client {
	return &Client{IClient: cli}
}
