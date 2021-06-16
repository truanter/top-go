package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"os"
)

// LogMiddleware ...
type LogMiddleware struct {
	top *SDKClient
}

// Handle ...
func (l *LogMiddleware) Handle(
	req *http.Request,
	next func(req *http.Request) (rsp *http.Response, err error),
) (rsp *http.Response, err error) {
	var debugFile *os.File
	top := l.top
	if top.IsDebug() {
		if top.GetDebugFile() == "" {
			debugFile = os.Stdout
		} else {
			debugFile, err = os.OpenFile(top.GetDebugFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err == nil {
				defer debugFile.Close()
			} else {
				return nil, err
			}
		}
		if request, err := httputil.DumpRequestOut(req, true); err == nil {
			if debugFile != nil {
				_, err = debugFile.WriteString("Request Host:" + req.URL.Host + "\n")
				_, err = debugFile.Write(request)
				_, err = debugFile.WriteString("\n")
			}
		}
	}
	rsp, err = next(req)
	if top.IsDebug() {
		if rsp != nil {
			if response, err := httputil.DumpResponse(rsp, true); err == nil {
				if debugFile != nil {
					_, err = debugFile.Write(response)
					_, err = debugFile.WriteString("\n")
				}
			}
		} else {
			errStr, _ := json.Marshal(err)
			_, err = debugFile.Write(errStr)
			_, err = debugFile.WriteString("\n")
		}
	}

	return rsp, err
}
