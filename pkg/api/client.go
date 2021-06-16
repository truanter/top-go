package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/truanter/top-go/pkg/config"
	"io/ioutil"
	"net/http"
)

type API interface {
	Methods() string
}

type APIClient struct {
	Cfg       *config.Config
	SDKConfig *config.SDKConfig
	common    service

	TbkDgOptimusMaterial  *TbkDgOptimusMaterialService
	TbkDgMaterialOptional *TbkDgMaterialOptionalService
}

type service struct {
	client *APIClient
	Method string
}

func NewAPIClient(sdkConfig *config.SDKConfig) *APIClient {
	cfg := sdkConfig.Config
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{}
	}

	c := &APIClient{}
	c.SDKConfig = sdkConfig
	c.Cfg = &cfg
	c.common.client = c

	c.TbkDgOptimusMaterial = (*TbkDgOptimusMaterialService)(&c.common)
	c.TbkDgMaterialOptional = (*TbkDgMaterialOptionalService)(&c.common)
	return c
}

func (c *APIClient) prepareRequest(method string, requestData interface{}) (r *http.Request, err error) {
	topUrl := config.GetHost(c.SDKConfig.IsSandBox, c.SDKConfig.IsHttp)

	body := &bytes.Buffer{}
	if err = json.NewEncoder(body).Encode(requestData); err != nil {
		return
	}

	headers := http.Header{
		"Content-Type": []string{"application/json"},
	}
	if r, err = http.NewRequest(http.MethodPost, topUrl, body); err != nil {
		return
	}
	r.Header = headers
	q := r.URL.Query()
	q.Set("method", method)
	r.URL.RawQuery = q.Encode()
	return
}

func (c *APIClient) callAPI(req *http.Request) (*http.Response, error) {
	return c.Cfg.HTTPClient.Do(req)
}

func (c *APIClient) do(api API, reqData, respData interface{}) error {
	method := api.Methods()
	req, err := c.prepareRequest(method, reqData)
	if err != nil {
		return err
	}

	httpResp, err := c.callAPI(req)
	if err != nil || httpResp == nil {
		return err
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if httpResp.StatusCode == 200 {
		dec := json.NewDecoder(bytes.NewBuffer(body))
		dec.UseNumber()
		if err := dec.Decode(respData); err != nil {
			return errors.New(fmt.Sprintf("解析数据失败：%s", err.Error()))
		}
		return nil
	} else {
		return errors.New(fmt.Sprintf("http request error. status code: %d", httpResp.StatusCode))
	}

}
