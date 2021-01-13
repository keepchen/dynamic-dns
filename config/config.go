package config

//Config 配置
type Config struct {
	AccessKey    string `yaml:"access_key"`
	AccessSecret string `yaml:"access_secret"`
	RegionID     string `yaml:"region_id"`
	IPAddressURL string `yaml:"ip_address_url"`
}
