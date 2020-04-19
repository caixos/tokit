package configs

type HttpConfig struct {
	HttpHost string `json:"http_host"`
	HttpPort string `json:"http_port"`
}

func LoadHttpConfig() *HttpConfig {
	config := &HttpConfig{
		HttpHost: EnvString("servers.http_host", "127.0.0.1"),
		HttpPort: EnvString("servers.http_port", "8341"),
	}
	return config
}
