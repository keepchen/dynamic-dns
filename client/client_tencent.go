package client

import (
	"errors"
)

//TencentClient 腾讯云客户端
type TencentClient struct {
	accessKey    string
	accessSecret string
	regionID     string
}

//NewTencentClient 获取实例
func NewTencentClient(accessKey, accessSecret, regionID string) (*TencentClient, error) {
	return &TencentClient{
		accessKey:    accessKey,
		accessSecret: accessSecret,
		regionID:     regionID,
	}, nil
}

//GetDNSRecords 获取解析记录列表
func (*TencentClient) GetDNSRecords(r *GetDNSRecordsRequest) (*GetDNSRecordsResponse, error) {
	var resp *GetDNSRecordsResponse

	return resp, errors.New("not implement")
}

//GetDNSRecord 获取解析记录
func (*TencentClient) GetDNSRecord(r *GetDNSRecordRequest) (*GetDNSRecordResponse, error) {
	var resp *GetDNSRecordResponse

	return resp, errors.New("not implement")
}

//CreateDNSRecord 创建解析记录
func (*TencentClient) CreateDNSRecord(req *CreateDNSRecordRequest) (*CreateDNSRecordResponse, error) {
	var resp *CreateDNSRecordResponse

	return resp, errors.New("not implement")
}

//UpdateDNSRecord 更新解析记录
func (*TencentClient) UpdateDNSRecord(req *UpdateDNSRecordRequest) (*UpdateDNSRecordResponse, error) {
	var resp *UpdateDNSRecordResponse

	return resp, errors.New("not implement")
}

//DeleteDNSRecord 删除解析记录
func (*TencentClient) DeleteDNSRecord(req *DeleteDNSRecordRequest) (*DeleteDNSRecordResponse, error) {
	var resp *DeleteDNSRecordResponse

	return resp, errors.New("not implement")
}
