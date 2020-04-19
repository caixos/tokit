package configs

type HttpConfig struct {
	HttpHost string `json:"http_host"`
	HttpPort string `json:"http_port"`
}

func LoadHttpConfig() *HttpConfig {
	config := &HttpConfig{
		HttpHost: EnvString("servers.http_host", "0.0.0.0"),
		HttpPort: EnvString("servers.http_port", "8341"),
	}
	return config
}
