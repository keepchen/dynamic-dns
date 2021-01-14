package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/keepchen/dynamic-dns/client"
	"github.com/keepchen/dynamic-dns/config"
	"github.com/keepchen/dynamic-dns/ipaddress"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	var (
		isp      *string
		cfgFile  *string
		provider client.IClient
		err      error
	)

	isp = flag.String("isp", "", "运营商(阿里云:aliyun,腾讯云:tencent)")
	cfgFile = flag.String("config", "", "配置文件")
	flag.Parse()
	fmt.Printf("|--isp: %s\n|--config: %s\n", *isp, *cfgFile)

	bt, err := ioutil.ReadFile(*cfgFile)
	if err != nil {
		fmt.Printf("打开配置文件失败:\n%s\n", err.Error())
		return
	}

	var cfg config.Config
	err = yaml.Unmarshal(bt, &cfg)
	if err != nil {
		fmt.Printf("配置文件解析失败:\n%s\n", err.Error())
		return
	}
	fmt.Println(cfg)

	if *isp == "aliyun" {
		provider, err = client.NewAliyunClient(cfg.AccessKey, cfg.AccessSecret, cfg.RegionID)
	} else {
		provider, err = client.NewTencentClient(cfg.AccessKey, cfg.AccessSecret, cfg.RegionID)
	}

	if err != nil {
		fmt.Printf("实例化SDK客户端失败：\n%s\n", err.Error())
		return
	}

	cli := client.NewClient(provider)

	fmt.Println(cli.GetDNSRecords(&client.GetDNSRecordsRequest{}))

	parser := func(res []byte) (string, error) {
		type Decoder struct {
			IP     string `json:"ip"`
			Origin string `json:"origin"`
			Port   string `json:"port"`
		}
		var decode Decoder
		err := json.Unmarshal(res, &decode)
		return decode.IP, err
	}
	ipClient := ipaddress.NewClient(cfg.IPAddressURL, parser)

	fmt.Println(ipClient.Query("GET", nil, nil, 5))
}
