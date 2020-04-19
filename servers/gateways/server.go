package gateways

import (
	"caixin.app/tokit/client"
	"caixin.app/tokit/config"
	"caixin.app/tokit/constant"
	"caixin.app/tokit/contract"
	"caixin.app/tokit/servers/events"
	"github.com/go-kit/kit/endpoint"
	"net/http"
	"net/http/httputil"
	"time"
	"errors"
	"fmt"
	"context"
)

type Server struct {
	handlers map[string]endpoint.Endpoint
	Logger   contract.ILogger
}

func NewServer() *Server {
	ss := &Server{
		handlers: make(map[string]endpoint.Endpoint),
	}
	return ss
}

func (s *Server) Register(method, path string, endpoint endpoint.Endpoint) {
	key := method + "_" + path
	s.handlers[key] = endpoint
}

func (s *Server) Serve() error {
	config := config.LoadHttpConfig()
	address := config.HttpHost + ":" + config.HttpPort
	s.Logger.Info("Http Server Start ", address)
	handler := s
	return http.ListenAndServe(address, handler)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	var resp contract.Response
	//通过编解码 进行 路由路由处理
	ctx := r.Context()
	req, err := decodeRequest(ctx, r)
	if err != nil {
		resp = contract.ResponseFailed(err)
		_ = encodeResponse(ctx, w, resp)
		return
	}
	key := req.Method + "_" + r.URL.Path
	filter, ok := s.handlers[key]
	if ok && filter != nil {
		// 如果有注册管理,则注册管理处理
		//注意filter的endpoint可以只过滤,不进行service处理,
		// gateway_endpoint负责返回GATEWAY,h或者error
		resp = s.runFilter(filter, ctx, req)
		data, exist := resp.Data.(map[string]interface{})
		if exist && data != nil {
			req.Data = data
		}
	}
	if !ok || req.Data["GATEWAY"] == "GATEWAY" {
		if req.Service == "" {
			resp = contract.ResponseFailed(errors.New("9999:没有响应的服务"))
			_ = encodeResponse(ctx, w, resp)
			return
		}
		var tag, host string
		defer s.fireEvent(time.Now(), &key, &tag, &host)

		//服务发现
		entity, err := client.GetConsulService(req.Service)
		if err != nil {
			resp = contract.ResponseFailed(err)
			_ = encodeResponse(ctx, w, resp)
			return
		}
		tag = entity.Service.Tags[0]
		host = fmt.Sprintf("%s:%d", entity.Service.Address, entity.Service.Port)
		if tag == "http" {
			director := func(dr *http.Request) {
				dr.URL.Scheme = "http"
				dr.URL.Host = host
				dr.URL.Path = req.Dest
				dr.Method = req.Method
			}
			gateway := &httputil.ReverseProxy{Director: director}
			gateway.ServeHTTP(w, r)
			return

		} else if tag == "grpc" && req.Route != "" {
			resp = client.NewGrpcCall(host, req.Route, req.Data)
			_ = encodeResponse(ctx, w, resp)
			return
		}
	}
	_ = encodeResponse(ctx, w, resp)
}

func (s *Server) runFilter(filter endpoint.Endpoint, ctx context.Context, req *contract.GateWayRequest) contract.Response {
	filterResp, err := filter(ctx, contract.Request{
		Id:   req.Id,
		Data: req.Data,
	})
	if err != nil {
		return contract.ResponseFailed(err)
	} else {
		return filterResp.(contract.Response)
	}
}
func (s *Server) fireEvent(begin time.Time, key, tag, host *string) {
	params := make(map[string]interface{})
	params["url"] = key
	params["begin"] = begin.Format(constant.YmdHis)
	params["took"] = time.Since(begin)
	params["tag"] = *tag
	params["host"] = *host
	payload := &contract.Payload{
		Route:  config.EnvString("gateway.event_handler", "GATEWAY_EVENT_HANDLER"),
		Params: params,
	}
	events.Fire(payload)
}
func (s *Server) Close() {

}