package client

//GetDNSRecordsRequest 获取解析列表请求结构体
type GetDNSRecordsRequest struct {
	DomainName string
	PageNumber int
	PageSize   int
	OrderBy    string
	Status     string
}

//GetDNSRecordsRequest 获取解析记录请求结构体
type GetDNSRecordRequest struct {
	RecordID     string
	Lang         string
	UserClientIP string
}

//CreateDNSRecordRequest 新建解析记录请求结构体
type CreateDNSRecordRequest struct {
	DomainName   string //域名
	RR           string //主机记录值，如a.test.com中的a
	Type         string //解析类型
	Value        string //记录值，如A记录的123.123.123.123
	Lang         string //语言
	Line         string //解析线路
	Priority     int    //MX记录的优先级[1,10]
	TTL          int    //解析生效时间(秒)，默认600秒
	UserClientIP string //用户端IP
}

//UpdateDNSRecordRequest 更新解析记录请求结构体
type UpdateDNSRecordRequest struct {
	RecordID     string //记录编号
	RR           string //主机记录值，如a.test.com中的a
	Type         string //解析类型
	Value        string //记录值，如A记录的123.123.123.123
	Lang         string //语言
	Line         string //解析线路
	Priority     int    //MX记录的优先级[1,10]
	TTL          int    //解析生效时间(秒)，默认600秒
	UserClientIP string //用户端IP
}

//DeleteDNSRecordRequest 删除解析记录请求结构体
type DeleteDNSRecordRequest struct {
	DomainName   string
	RR           string
	Lang         string
	Type         string
	UserClientIP string
}
