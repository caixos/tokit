package mqtts

import (
	"caixin.app/caixos/tokit/config"
	"caixin.app/caixos/tokit/constant"
	"encoding/json"
	"errors"
)

func Publish(topic string, payload interface{}) error {
	if GetIns() == nil {
		return errors.New(constant.ErrMQTTConnect)
	}
	param, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	config := config.LoadMqttConfig()
	token := GetIns().Publish(topic, config.PublishQos, false, param)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
