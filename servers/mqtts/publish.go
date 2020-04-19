package mqtts

import (
	"caixin.app/caixos/tokit/configs"
	"caixin.app/caixos/tokit/constants"
	"encoding/json"
	"errors"
)

func Publish(topic string, payload interface{}) error {
	if GetIns() == nil {
		return errors.New(constants.ErrMQTTConnect)
	}
	param, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	config := configs.LoadMqttConfig()
	token := GetIns().Publish(topic, config.PublishQos, false, param)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
