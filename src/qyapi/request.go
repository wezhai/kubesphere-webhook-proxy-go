package qyapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/wezhai/kubesphere-webhook-proxy-go/logger"
)

func request(url string, jsonStr []byte) (body []byte, err error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	content := bytes.NewBuffer(jsonStr)
	logger.Debugf("RequestUrl: %s RequestPayload: %s", url, content)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Body.Close()

	body, _ = ioutil.ReadAll(resp.Body)
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		logger.Error("Unmarshal body error")
	}

	errcode := string(*objmap["errcode"])
	if errcode != "0" {
		errmsg := "请求失败 errmsg: " + string(*objmap["errmsg"])
		logger.Errorf(errmsg)
		err = errors.New(errmsg)
	}

	return body, err
}
