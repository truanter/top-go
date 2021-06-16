package errors

type ErrorResponse struct {
	Code      int    `json:"code" xml:"code"`
	Msg       string `json:"msg" xml:"msg"`
	SubCode   string `json:"sub_code" xml:"sub_code"`
	SubMsg    string `json:"sub_msg" xml:"sub_msg"`
	RequestID string `json:"request_id" xml:"request_id"`
}
