package client

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

//AliyunClient 阿里云客户端
type AliyunClient struct {
	accessKey    string
	accessSecret string
	regionID     string
	instance     *alidns.Client
}

//NewAliyunClient 获取实例
func NewAliyunClient(accessKey, accessSecret, regionID string) (*AliyunClient, error) {
	client, err := alidns.NewClientWithAccessKey(regionID, accessKey, accessSecret)
	if err != nil {
		return nil, err
	}
	return &AliyunClient{
		accessKey:    accessKey,
		accessSecret: accessSecret,
		regionID:     regionID,
		instance:     client,
	}, nil
}

//GetDNSRecords 获取解析记录列表
func (cli *AliyunClient) GetDNSRecords(req *GetDNSRecordsRequest) (*GetDNSRecordsResponse, error) {
	var resp = new(GetDNSRecordsResponse)

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.Scheme = "https"
	request.DomainName = req.DomainName

	response, err := cli.instance.DescribeDomainRecords(request)

	if err != nil {
		return nil, err
	}

	resp.PageSize = response.PageSize
	resp.TotalCount = response.TotalCount
	resp.RequestID = response.RequestId
	for _, record := range resp.Records {
		resp.Records = append(resp.Records, Record{
			RecordId:   record.RecordId,
			Value:      record.Value,
			TTL:        record.TTL,
			Remark:     record.Remark,
			DomainName: record.DomainName,
			RR:         record.RR,
			Priority:   record.Priority,
			Locked:     record.Locked,
			Weight:     record.Weight,
			Line:       record.Line,
			Type:       record.Type,
			Status:     record.Status,
		})
	}

	return resp, nil
}

//GetDNSRecord 获取解析记录
func (cli *AliyunClient) GetDNSRecord(req *GetDNSRecordRequest) (*GetDNSRecordResponse, error) {
	var resp = new(GetDNSRecordResponse)

	request := alidns.CreateDescribeDomainRecordInfoRequest()
	request.Scheme = "https"
	request.RecordId = req.RecordID

	response, err := cli.instance.DescribeDomainRecordInfo(request)
	if err != nil {
		return nil, err
	}

	resp.DomainName = response.DomainName
	resp.Line = response.Line
	resp.Locked = response.Locked
	resp.Priority = response.Priority
	resp.RecordID = response.RecordId
	resp.Status = response.Status
	resp.RR = response.RR
	resp.TTL = response.TTL
	resp.Type = response.Type
	resp.Value = response.Value

	return resp, nil
}

//CreateDNSRecord 创建解析记录
func (cli *AliyunClient) CreateDNSRecord(req *CreateDNSRecordRequest) (*CreateDNSRecordResponse, error) {
	var resp = new(CreateDNSRecordResponse)

	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = "https"

	request.DomainName = req.DomainName
	request.RR = req.RR
	request.Type = req.Type
	request.Value = req.Value
	request.Lang = req.Lang
	request.Line = req.Line
	//request.Priority = requests.NewInteger(req.Priority)
	//request.TTL = requests.NewInteger(req.TTL)
	request.UserClientIp = req.UserClientIP

	response, err := cli.instance.AddDomainRecord(request)

	if err != nil {
		return nil, err
	}

	resp.RecordID = response.RecordId
	resp.RequestID = response.RequestId

	return resp, nil
}

//UpdateDNSRecord 更新解析记录
func (cli *AliyunClient) UpdateDNSRecord(req *UpdateDNSRecordRequest) (*UpdateDNSRecordResponse, error) {
	var resp = new(UpdateDNSRecordResponse)

	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"

	request.RecordId = req.RecordID
	request.RR = req.RR
	request.Type = req.Type
	request.Value = req.Value
	request.Lang = req.Lang
	request.Line = req.Line
	request.Priority = requests.NewInteger(req.Priority)
	request.TTL = requests.NewInteger(req.TTL)
	request.UserClientIp = req.UserClientIP

	response, err := cli.instance.UpdateDomainRecord(request)

	if err != nil {
		return nil, err
	}

	resp.RequestID = response.RequestId
	resp.RecordID = response.RecordId

	return resp, nil
}

//DeleteDNSRecord 删除解析记录
func (cli *AliyunClient) DeleteDNSRecord(req *DeleteDNSRecordRequest) (*DeleteDNSRecordResponse, error) {
	var resp = new(DeleteDNSRecordResponse)

	request := alidns.CreateDeleteSubDomainRecordsRequest()
	request.Scheme = "https"

	request.Domain = req.DomainName
	request.RR = req.RR
	request.Type = req.Type
	request.Lang = req.Lang
	request.UserClientIp = req.UserClientIP

	response, err := cli.instance.DeleteSubDomainRecords(request)

	if err != nil {
		return nil, err
	}

	resp.RequestID = response.RequestId
	resp.TotalCount = response.TotalCount
	resp.RR = response.RR

	return resp, nil
}
