package tokit

import (
	"caixin.app/tokit/args"
	"caixin.app/tokit/contract"
	"caixin.app/tokit/logger"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd/consul"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Application struct {
	Status bool
	//必须初始化
	Service  map[string]contract.IService
	Consul   map[string]*consul.Registrar
	Handlers map[string]endpoint.Endpoint
	Routers  map[string]contract.IRouter
	tomls    map[string]*viper.Viper
}

var App *Application

//初始化成全局变量
func init() {
	App = &Application{
		Status:   true,
		Service:  make(map[string]contract.IService),
		Consul:   make(map[string]*consul.Registrar),
		Handlers: make(map[string]endpoint.Endpoint),
		Routers:  make(map[string]contract.IRouter),
		tomls:    make(map[string]*viper.Viper),
	}
}

func Provider(p contract.IProvider) {
	p.Boot()
	p.Register()
}

func Handler(name string, endpoint ...endpoint.Endpoint) endpoint.Endpoint {
	if endpoint == nil {
		ret, exist := App.Handlers[name]
		if exist {
			return ret
		}
	} else {
		App.Handlers[name] = endpoint[0]
	}
	return nil
}

func Toml(name string, fileName ...string) *viper.Viper {
	if fileName == nil {
		//这里是get
		ret, exist := App.tomls[name]
		if exist {
			return ret
		}
	} else {
		//这里是注册
		c := viper.New()
		c.SetConfigName(fileName[0])
		c.AddConfigPath("./toml/")
		c.AddConfigPath("../toml/")
		c.AddConfigPath("../../toml/")
		c.SetConfigType("toml")
		if err := c.ReadInConfig(); err != nil {
			panic(err)
		}
		App.tomls[name] = c
	}
	return nil
}
func Service(name string, service ...contract.IService) contract.IService {
	if service == nil {
		ret, exist := App.Service[name]
		if exist {
			return ret
		}
	} else {
		App.Service[name] = service[0]
	}
	return nil
}

func Router(name string, server contract.IRouter) {
	server.Boot()
	server.Load()
	server.Register()
	App.Routers[name] = server
}

//启动server
func Start() {
	servers := strings.Split(args.Server, ",")
	routers := make(map[string]contract.IRouter)
	for _, s := range servers {
		if ss, exist := App.Routers[strings.Trim(s, " ")]; exist == true {
			routers[s] = ss
		}
	}
	errChans := make(map[string]chan error)
	for key, router := range routers {
		errChans[key] = make(chan error)
		go func(errChan chan error, server contract.IRouter) {
			errChan <- server.Start()
		}(errChans[key], router)
		go func(errChan chan error) {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			errChan <- fmt.Errorf("%s", <-c)
		}(errChans[key])
	}
	for _, errChan := range errChans {
		logger.GetLog().Info(<-errChan)
	}
	//关闭各种路由服务
	for _, server := range routers {
		server.Close()
	}

}
