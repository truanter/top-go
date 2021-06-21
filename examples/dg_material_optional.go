package examples

import (
	"fmt"
	"github.com/truanter/top-go/pkg/config"
	"github.com/truanter/top-go/pkg/model"
	"github.com/truanter/top-go/pkg/sdk"
)

func DgMaterialOptional() {
	c := sdk.NewSdkClient(&config.SDKConfig{})
	c.SetKeyAndSecret(appKey, secret)
	c.UseHttps()
	c.SetDebug(true)
	api := c.TbkDgMaterialOptional()
	resp, err := api.Do(model.TbkDgMaterialOptionalRequest{
		PageSize:  2,
		AdzoneID:  adzoneId,
		Q:         "A",
		HasCoupon: true,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("request error: %s", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("tpwd create resp: %v", resp))
	}
}
