//https://open.taobao.com/api.htm?docId=31127&docType=2
package api

import "github.com/truanter/top-go/pkg/model"

type TbkTpwdCreateService service

func (s *TbkTpwdCreateService) Methods() string {
	return "taobao.tbk.tpwd.create"
}

func (s *TbkTpwdCreateService) Do(data model.TbkTpwdCreateRequest) (model.TbkTpwdCreateResponse, error) {
	var resp model.TbkTpwdCreateResponse
	err := s.client.do(s, data, &resp)
	return resp, err
}
