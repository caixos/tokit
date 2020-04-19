package provider

import (
	"caixin.app/caixos/tokit"
	"caixin.app/caixos/tokit/client"
	"caixin.app/caixos/tokit/config"
	"caixin.app/caixos/tokit/args"

	"strings"

)

type ConsulRegistyProvider struct {}

func (s *ConsulRegistyProvider) Boot() {}

func (s *ConsulRegistyProvider) Register() {
	if args.Registy != "" {
		s.consulHttpRegister()
		s.consulGrpcRegister()
	}
}

func (s *ConsulRegistyProvider) consulHttpRegister() {
	if strings.Contains(args.Server, "http") || strings.Contains(args.Server, "gateway") {
		httpConfig := config.LoadHttpConfig()
		tokit.App.Consul["http"] = client.NewConsulHttpRegister(
			args.Name,
			httpConfig.HttpHost,
			httpConfig.HttpPort,
		)
	}
}

func (s *ConsulRegistyProvider) consulGrpcRegister() {
	if strings.Contains(args.Server, "grpc") {
		grpcConfig := config.LoadGrpcConfig()
		tokit.App.Consul["grpc"] = client.NewConsulGrpcRegister(
			args.Name,
			grpcConfig.GrpcHost,
			grpcConfig.GrpcPort,
		)
	}
}
