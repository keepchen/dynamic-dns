package client

//IClient 公共方法接口
type IClient interface {
	GetDNSRecords(r *GetDNSRecordsRequest) (*GetDNSRecordsResponse, error)
	GetDNSRecord(r *GetDNSRecordRequest) (*GetDNSRecordResponse, error)
	CreateDNSRecord(r *CreateDNSRecordRequest) (*CreateDNSRecordResponse, error)
	UpdateDNSRecord(r *UpdateDNSRecordRequest) (*UpdateDNSRecordResponse, error)
	DeleteDNSRecord(r *DeleteDNSRecordRequest) (*DeleteDNSRecordResponse, error)
}

//Client 客户端接口
type Client struct {
	IClient
}
