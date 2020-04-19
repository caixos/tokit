package client

import (
	"caixin.app/caixos/tokit/contract"
	"caixin.app/caixos/tokit/tools/convert"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
)

func NewHttpPostCall(host, service string, params map[string]interface{}) (ret contract.Response) {

	path := "http://" + host + "/" + strings.Replace(service, ".", "/", -1)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm(path, convert.FormEncode(params))
	if err != nil {
		ret = contract.ResponseFailed(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	response := &contract.Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		ret = contract.ResponseFailed(err)
	} else {
		m := response.Data.(map[string]interface{})
		m["call_method"] = "http"
		response.Data = m
		ret = *response
	}
	return
}
