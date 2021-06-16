package api

import (
	. "github.com/truanter/top-go/pkg/model"
)

type TbkDgOptimusMaterialService service

func (s *TbkDgOptimusMaterialService) Methods() string {
	return "taobao.tbk.dg.optimus.material"
}

func (s *TbkDgOptimusMaterialService) Do(data TbkDgOptimusMaterialRequest) (TbkDgOptimusMaterialResponse, error) {
	var resp TbkDgOptimusMaterialResponse
	err := s.client.do(s, data, &resp)
	return resp, err
}
