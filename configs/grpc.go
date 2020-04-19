package configs

type GrpcConfig struct {
	GrpcHost string `json:"grpc_host"`
	GrpcPort string `json:"grpc_port"`
}

func LoadGrpcConfig() *GrpcConfig {
	config := &GrpcConfig{
		GrpcHost: EnvString("servers.grpc_host", "127.0.0.1"),
		GrpcPort: EnvString("servers.grpc_port", "9341"),
	}
	return config
}
