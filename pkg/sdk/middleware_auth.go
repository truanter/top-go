package sdk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type AuthMiddleware struct {
	top *SDKClient
}

func (a *AuthMiddleware) Handle(req *http.Request, next func(req *http.Request) (rsp *http.Response, err error)) (rsp *http.Response, err error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	_ = req.ParseForm()

	allData := make(map[string]interface{})

	q := req.URL.Query()
	q.Set("timestamp", timestamp)
	q.Set("app_key", a.top.AppKey)

	for k, v := range a.top.PublicParams {
		q.Set(k, fmt.Sprintf("%v", v))
	}

	for k, v := range q {
		if len(v) > 0 {
			allData[k] = v[0]
		}
	}

	jsonData := make(map[string]interface{})
	b, _ := ioutil.ReadAll(req.Body)
	decoder := json.NewDecoder(bytes.NewBuffer(b))
	decoder.UseNumber()
	_ = decoder.Decode(&jsonData)

	for k, v := range jsonData {
		allData[k] = v
	}

	multipartFormData := make(map[string]interface{})
	if req.MultipartForm != nil {
		for k, v := range req.MultipartForm.Value {
			if len(v) > 0 {
				multipartFormData[k] = v[0]
				allData[k] = v[0]
			}
		}
	}

	formData := make(map[string]interface{})
	if len(req.PostForm) > 0 {
		for k, v := range req.PostForm {
			if len(v) > 0 {
				formData[k] = v[0]
				allData[k] = v[0]
			}
		}
	}

	keys := make([]string, 0, len(allData))
	for k := range allData {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	strKey := a.top.Secret
	for _, v := range keys {
		data := fmt.Sprintf("%v", allData[v])
		strKey += v + data
	}
	strKey += a.top.Secret

	h := md5.New()
	h.Write([]byte(strKey))
	sign := hex.EncodeToString(h.Sum(nil))
	sign = strings.ToUpper(sign)
	q.Set("sign", sign)
	req.URL.RawQuery = q.Encode()

	data := url.Values{}
	for k, v := range jsonData {
		data.Set(k, fmt.Sprintf("%v", v))
	}
	for k, v := range multipartFormData {
		data.Set(k, fmt.Sprintf("%v", v))
	}
	for k, v := range formData {
		data.Set(k, fmt.Sprintf("%v", v))
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))
	req.ContentLength = int64(len(data.Encode()))
	req.Body = ioutil.NopCloser(strings.NewReader(data.Encode()))
	rsp, err = next(req)
	return rsp, err
}
