package sdk

import (
	"context"
	"github.com/truanter/top-go/pkg/api"
	"github.com/truanter/top-go/pkg/config"
	"net/http"
)

type SDKClient struct {
	http.RoundTripper
	Config         *config.SDKConfig
	API            *api.APIClient
	Ctx            *context.Context
	middlewareList []Middleware
	AppKey         string
	Secret         string
	PublicParams   map[string]interface{}
}

func NewSdkClient(cfg *config.SDKConfig) *SDKClient {
	ctx := context.Background()
	client := api.NewAPIClient(cfg)
	sdkClient := &SDKClient{
		RoundTripper: http.DefaultTransport,
		Config:       cfg,
		API:          client,
		Ctx:          &ctx,
		PublicParams: config.DefaultPublicParams,
	}
	sdkClient.API.Cfg.HTTPClient.Transport = sdkClient
	sdkClient.middlewareList = []Middleware{
		&AuthMiddleware{sdkClient},
		&LogMiddleware{sdkClient},
	}
	return sdkClient
}

func (c *SDKClient) UseSandBox() {
	c.Config.IsDebug = true
}

func (c *SDKClient) UseProduction() {
	c.Config.IsDebug = false
}

func (c *SDKClient) UseHttp() {
	c.Config.IsHttp = true
}

func (c *SDKClient) UseHttps() {
	c.Config.IsHttp = false
}

func (c *SDKClient) SetDebug(debug bool) {
	c.Config.IsDebug = debug
}

func (c *SDKClient) IsDebug() bool {
	return c.Config.IsDebug
}

func (c *SDKClient) GetDebugFile() string {
	return c.Config.DebugFilePath
}

func (c *SDKClient) SetDebugFile(debugFilePath string) {
	c.Config.DebugFilePath = debugFilePath
}

func (c *SDKClient) SetKeyAndSecret(appKey, secret string) {
	c.AppKey = appKey
	c.Secret = secret
}

func (c *SDKClient) SetTargetAppKey(targetAppKey string) {
	c.PublicParams[config.TargetAppKey.String()] = targetAppKey
}

func (c *SDKClient) SetSession(session string) {
	c.PublicParams[config.Session.String()] = session
}

func (c *SDKClient) SetPartnerID(partnerID string) {
	c.PublicParams[config.PartnerId.String()] = partnerID
}

func (c *SDKClient) SetSimplify(simplify bool) {
	c.PublicParams[config.Simplify.String()] = simplify
}

func (c *SDKClient) resetPublicParams() {
	c.PublicParams = config.DefaultPublicParams
}

func (c *SDKClient) RoundTrip(req *http.Request) (rsp *http.Response, err error) {
	beforeFunc := func(req *http.Request) (rsp *http.Response, err error) {
		return c.RoundTripper.RoundTrip(req)
	}
	middlewareCount := len(c.middlewareList)
	for i := middlewareCount - 1; i >= 0; i-- {
		beforeFunc = c.genMiddlewareHandleFunc(c.middlewareList[i], beforeFunc)
	}
	return beforeFunc(req)
}

func (c *SDKClient) genMiddlewareHandleFunc(
	middleware Middleware,
	beforeFunc func(req *http.Request) (rsp *http.Response, err error),
) func(req *http.Request) (rsp *http.Response, err error) {
	return func(req *http.Request) (rsp *http.Response, err error) {
		return middleware.Handle(req, beforeFunc)
	}
}
