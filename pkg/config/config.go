package config

import "net/http"

const (
	EnvProduct = string("product")
	EnvSandBox = string("sand_box")
	Https      = string("Https")
	Http       = string("Http")
)

type contextKey string

func (k contextKey) String() string {
	return "auth " + string(k)
}

type PublicParam string

func (p PublicParam) String() string {
	return string(p)
}

var (
	hostMap = map[string]map[string]string{
		EnvProduct: {
			Http:  "http://gw.api.taobao.com/router/rest",
			Https: "https://eco.taobao.com/router/rest",
		},
		EnvSandBox: {
			Http:  "http://gw.api.tbsandbox.com/router/rest",
			Https: "https://gw.api.tbsandbox.com/router/rest",
		},
	}

	TargetAppKey = PublicParam("target_app_key")
	SignMethod   = PublicParam("sign_method")
	Session      = PublicParam("session")
	Format       = PublicParam("format")
	V            = PublicParam("v")
	PartnerId    = PublicParam("partner_id")
	Simplify     = PublicParam("simplify")

	DefaultPublicParams = map[string]interface{}{
		SignMethod.String(): "md5",
		Format.String():     "json",
		V.String():          "2.0",
		Simplify.String():   false,
	}
)

//GetHost get host by environment and protocol
func GetHost(isSandBox, isHttp bool) (host string) {
	env := EnvProduct
	if isSandBox {
		env = EnvSandBox
	}

	protocol := Https
	if isHttp {
		protocol = Http
	}

	if v, found := hostMap[env]; found {
		host, _ = v[protocol]
	}
	return
}

type Config struct {
	HTTPClient *http.Client
}
