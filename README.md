# dynamic-dns
域名动态解析

#### 简介
该项目用于将指定域名解析到指定ip地址上。 
WARNING:**Do not use it in production!**

#### 功能列表
- [x] 查询域名解析列表
- [x] 查询域名解析记录
- [x] 新增域名解析
- [x] 修改域名解析
- [x] 删除域名解析
- [x] 查询当前机器IP地址

#### 支持的运营商
- [x] 阿里云
- [] 腾讯云

#### 食用方法
- 查询当前机器IP地址
```golang
# 首先，你需要找到一个可以查询ip的第三方接口
# 然后，自己实现解析这个接口的返回值方法，就像
# 下面的代码一样：
import "github.com/keepchen/dynamic-dns/ipaddress"

//比如某接口的返回值是这样的：
/*
{
ip: "171.221.151.23",
origin: "171.221.151.23:11341",
port: "11341"
}
*/
//那么实现解析方法如下：
parser := func(res []byte) (string, error) {
    type Decoder struct {
        IP string `json:"ip"`
        Origin string `json:"origin"`
        Port string `json:"port"`
    }
    var decode Decoder
    err := json.Unmarshal(res, &decode)
    return decode.IP, err
}
//唤起请求客户端
ipClient := ipaddress.NewClient("第三方接口地址", parser)
//打印查询结果
fmt.Println(ipClient.Query("GET", nil, nil, 5))
```

- 查询域名解析列表
```golang
import "github.com/keepchen/dynamic-dns/client"

//腾讯云
//provider, err = client.NewTencentClient(cfg.AccessKey, cfg.AccessSecret, cfg.RegionID)

//阿里云
provider, err = client.NewAliyunClient(cfg.AccessKey, cfg.AccessSecret, cfg.RegionID)
if err != nil {
    fmt.Printf("实例化SDK客户端失败：\n%s\n", err.Error())
    return
}
cli := client.NewClient(provider)
resp, err := cli.GetDNSRecords(&client.GetDNSRecordsRequest{
    //你的各种参数....
})
```

- 查询域名解析记录
```golang
resp, err := cli.GetDNSRecord(&client.GetDNSRecordRequest{
    //你的各种参数....
})
```

- 新增域名解析记录
```golang
resp, err := cli.CreateDNSRecord(&client.CreateDNSRecordRequest{
    //你的各种参数....
})
```

- 修改域名解析记录
```golang
resp, err := cli.UpdateDNSRecord(&client.UpdateDNSRecordRequest{
    //你的各种参数....
})
```

- 删除域名解析记录
```golang
resp, err := cli.DeleteDNSRecord(&client.DeleteDNSRecordRequest{
    //你的各种参数....
})
```

#### 总结
- 细心的小伙伴可能已经发现了，综合上述的操作，简单的run一个定时任务，就可以将域名映射到家里。
- 目前还在开发阶段，**！！！请不要用于生产环境！！！**