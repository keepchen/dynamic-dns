package client

//Record dns记录
type Record struct {
	Value      string
	TTL        int64
	Remark     string
	DomainName string
	RR         string
	Priority   int64
	RecordId   string
	Status     string
	Locked     bool
	Weight     int
	Line       string
	Type       string
}

//GetDNSRecordsResponse 获取解析列表返回结构体
type GetDNSRecordsResponse struct {
	TotalCount int64
	PageSize   int64
	Records    []Record
	RequestID  string
}

//GetDNSRecordResponse 获取解析记录返回结构体
type GetDNSRecordResponse struct {
	DomainName string
	Line       string
	Locked     bool
	Priority   int64
	RR         string //解析类型
	RecordID   string //解析记录ID
	Status     string //["Enable", "Disable"]
	TTL        int64
	Type       string //记录类型
	Value      string //记录值
}

//CreateDNSRecordResponse 新建解析记录返回结构体
type CreateDNSRecordResponse struct {
	RecordID  string
	RequestID string
}

//UpdateDNSRecordResponse 更新解析记录返回结构体
type UpdateDNSRecordResponse struct {
	RecordID  string
	RequestID string
}

//DeleteDNSRecordResponse 删除解析记录返回结构体
type DeleteDNSRecordResponse struct {
	RequestID  string
	RR         string
	TotalCount string
}
