package model

type TbkTpwdCreateRequest struct {
	Url string `json:"url,omitempty"`
}

type TbkTpwdCreateResponse struct {
	BaseResponse
	TbkTpwdCreateResponse map[string]interface{} `json:"tbk_tpwd_create_response,omitempty" xml:"tbk_tpwd_create_response"`
}
