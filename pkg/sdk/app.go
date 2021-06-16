package sdk

import "github.com/truanter/top-go/pkg/api"

func (c *SDKClient) TbkDgOptimusMaterial() *api.TbkDgOptimusMaterialService {
	return c.API.TbkDgOptimusMaterial
}

func (c *SDKClient) TbkDgMaterialOptional() *api.TbkDgMaterialOptionalService {
	return c.API.TbkDgMaterialOptional
}
