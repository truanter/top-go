// https://open.taobao.com/api.htm?docId=35896&docType=2
package api

import "github.com/truanter/top-go/pkg/model"

type TbkDgMaterialOptionalService service

func (s *TbkDgMaterialOptionalService) Methods() string {
	return "taobao.tbk.dg.material.optional"
}

func (s *TbkDgMaterialOptionalService) Do(data model.TbkDgMaterialOptionalRequest) (model.TbkDgMaterialOptionalResponse, error) {
	var resp model.TbkDgMaterialOptionalResponse
	err := s.client.do(s, data, &resp)
	return resp, err
}
