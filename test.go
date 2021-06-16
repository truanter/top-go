package main

import (
	"fmt"
	"github.com/truanter/top-go/pkg/config"
	"github.com/truanter/top-go/pkg/model"
	"github.com/truanter/top-go/pkg/sdk"
)

func main() {
	c := sdk.NewSdkClient(&config.SDKConfig{})
	c.SetKeyAndSecret("your app key", "your secret")
	c.UseHttps()
	c.SetDebug(true)
	tbkDgOptimusMaterialResp, err := c.TbkDgOptimusMaterial().Do(model.TbkDgOptimusMaterialRequest{
		PageSize:   10,
		AdzoneID:   0000,
		PageNO:     1,
		MaterialID: 6708,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("tbkDgOptimusMaterialResp: %v, err: %s", tbkDgOptimusMaterialResp, err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("tbkDgOptimusMaterialResp: %v, err: %v", tbkDgOptimusMaterialResp, nil))
	}

	tbkDgMaterialOptionalResp, err := c.TbkDgMaterialOptional().Do(model.TbkDgMaterialOptionalRequest{
		AdzoneID: 00000000,
		Q:        "test",
		PageSize: 1,
	})

	if err != nil {
		fmt.Println(fmt.Sprintf("tbkDgMaterialOptionalResp: %v, err: %s", tbkDgMaterialOptionalResp, err.Error()))
	} else {
		if tbkDgMaterialOptionalResp.ErrorResponse.Code != 0 {
			fmt.Printf("error_code: %d, error_msg: %s, request_id: %s", tbkDgMaterialOptionalResp.ErrorResponse.Code, tbkDgMaterialOptionalResp.ErrorResponse.Msg, tbkDgMaterialOptionalResp.ErrorResponse.RequestID)
		} else {
			fmt.Println(fmt.Sprintf("tbkDgMaterialOptionalResp: %v, err: %v", tbkDgMaterialOptionalResp, nil))
		}
	}
}
