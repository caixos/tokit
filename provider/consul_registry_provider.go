package provider

import "caixin.app/tokit/args"

type ConsulRegistyProvider struct {
}

func (s *ConsulRegistyProvider) Boot() {

}

func (s *ConsulRegistyProvider) Register() {
	if args.Registy != "" {
		s.consulHttpRegister()
		s.consulGrpcRegister()
	}
}
func (s *ConsulRegistyProvider) consulHttpRegister() {
	if strings.Contains(args.Server, "http") || strings.Contains(args.Server, "gateway") {
		httpConfig := configs.LoadHttpConfig()
		wego.App.Consul["http"] = clients.NewConsulHttpRegister(
			args.Name,
			httpConfig.HttpHost,
			httpConfig.HttpPort,
		)
	}
}
func (s *ConsulRegistyProvider) consulGrpcRegister() {
	if strings.Contains(args.Server, "grpc") {
		grpcConfig := configs.LoadGrpcConfig()
		wego.App.Consul["grpc"] = clients.NewConsulGrpcRegister(
			args.Name,
			grpcConfig.GrpcHost,
			grpcConfig.GrpcPort,
		)
	}
}
