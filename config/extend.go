package config

var ExtConfig Extend

// Extend 扩展配置
//
//	extend:
//	  mqtt:
//	    host: localhost
//	    port: 1883
//
// 使用方法： config.ExtConfig.Mqtt.Host 即可
type Extend struct {
	AMap    AMap    // 地图配置
	Alipay  Alipay  // 支付宝配置
	Mqtt    Mqtt    // MQTT 配置
	RPC     RPC     // gRPC 配置
	Alert   Alert   // 告警配置（飞书 / 钉钉）
}

type AMap struct {
	Key string
}

type Alipay struct {
	APPID           string
	PrivateKey      string
	AliPayPublicKey string
	NotifyUrl       string
	IsProd          bool
}

type Mqtt struct {
	Host     string
	Port     int
	Username string
	Password string
	ClientId string
}

type RPC struct {
	Client  bool
	Address string
}

// Alert 告警配置
type Alert struct {
	Enabled     bool     `yaml:"enabled"`
	Environment string   `yaml:"environment"`
	ServiceName string   `yaml:"serviceName"`
	Feishu      Feishu   `yaml:"feishu"`
	DingTalk    DingTalk `yaml:"dingtalk"`
}

type Feishu struct {
	Enabled bool   `yaml:"enabled"`
	Webhook string `yaml:"webhook"`
	Secret  string `yaml:"secret"`
}

type DingTalk struct {
	Enabled bool   `yaml:"enabled"`
	Webhook string `yaml:"webhook"`
	Secret  string `yaml:"secret"`
}
