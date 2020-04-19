package configs

type WebSocketConfig struct {
	WebSocketHost string `json:"web_socket_host"`
	WebSocketPort string `json:"web_socket_port"`
	Path          string
}

func LoadWebSocketConfig() *WebSocketConfig {
	config := &WebSocketConfig{
		Path:          "/ws",
		WebSocketPort: EnvString("servers.websocket_port", "8342"),
		WebSocketHost: EnvString("servers.websocket_host", "127.0.0.1"),
	}
	return config
}
