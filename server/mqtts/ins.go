package mqtts

import (
	"caixin.app/caixos/tokit/config"
	"caixin.app/caixos/tokit/tools/idwork"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

var ins mqtt.Client
var once sync.Once

func GetIns() mqtt.Client {
	once.Do(func() {
		ins = init_mc()
	})
	return ins
}
func init_mc() mqtt.Client {
	config := config.LoadMqttConfig()
	opts := mqtt.NewClientOptions().AddBroker(config.Host)
	opts.SetUsername(config.UserName)
	opts.SetPassword(config.PassWord)
	opts.SetClientID(idwork.ID())
	mc := mqtt.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		return nil
	}
	return mc
}
