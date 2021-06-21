package examples

import (
	"fmt"
	"github.com/truanter/top-go/pkg/config"
	"github.com/truanter/top-go/pkg/model"
	"github.com/truanter/top-go/pkg/sdk"
)

func TpwdCreate() {
	c := sdk.NewSdkClient(&config.SDKConfig{})
	c.SetKeyAndSecret(appKey, secret)
	c.UseHttps()
	c.SetDebug(true)
	c.SetSimplify(true)
	api := c.TbkTpwdCreate()
	data := model.TbkTpwdCreateRequest{
		Url: "http://s.click.taobao.com/t?e=m%3D2%26s%3D70BphK96UPUcQipKwQzePOeEDrYVVa64r4ll3HtqqoxyINtkUhsv0P9jwbiZeMGaWCxUz27ycMhuV8aSfTp%2BXkFDp%2Fi5lyD6nqClqLWUZzvEc5V8ARk%2BqSetxX4B1T06xl%2BhMPmMA%2BNth2X5ZtNqUe3q8wtoQ2jlxReAuuwT2KURjaVwnYTDwtv9UaFZkN4t5lrVhFFjHe0hhQs2DjqgEA%3D%3D&scm=null&pvid=100_11.178.156.30_10676_8461624255800761794&app_pvid=59590_33.39.254.6_748_1624255800757&ptl=floorId:2836;originalFloorId:2836;pvid:100_11.178.156.30_10676_8461624255800761794;app_pvid:59590_33.39.254.6_748_1624255800757&xId=4vElNPj10GWNqNMaPYQvYOimFuFOvWw6pfrUQlqQxmZashaWS6s7jzrPgU9cPriUzr6RvGn1feGtOgen0TOuRvfqNQF0hGF0nD9ux0Z4RZlM&union_lens=lensId%3AMAPI%401624255800%402127fe06_081c_17a2d30a5ec_3c07%4001",
	}
	resp, err := api.Do(data)

	if err != nil {
		fmt.Println(fmt.Sprintf("request error: %s", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("tpwd create resp: %v", resp))
	}
}

func TpwdCreateByGeneralApi() {
	c := sdk.NewSdkClient(&config.SDKConfig{})
	c.SetKeyAndSecret(appKey, secret)
	c.UseHttps()
	c.SetDebug(true)
	c.SetSimplify(false)
	api := c.GeneralAPI()
	api.SetMethods("taobao.tbk.tpwd.create")
	data := map[string]interface{}{
		"url": "http://s.click.taobao.com/t?e=m%3D2%26s%3D70BphK96UPUcQipKwQzePOeEDrYVVa64r4ll3HtqqoxyINtkUhsv0P9jwbiZeMGaWCxUz27ycMhuV8aSfTp%2BXkFDp%2Fi5lyD6nqClqLWUZzvEc5V8ARk%2BqSetxX4B1T06xl%2BhMPmMA%2BNth2X5ZtNqUe3q8wtoQ2jlxReAuuwT2KURjaVwnYTDwtv9UaFZkN4t5lrVhFFjHe0hhQs2DjqgEA%3D%3D&scm=null&pvid=100_11.178.156.30_10676_8461624255800761794&app_pvid=59590_33.39.254.6_748_1624255800757&ptl=floorId:2836;originalFloorId:2836;pvid:100_11.178.156.30_10676_8461624255800761794;app_pvid:59590_33.39.254.6_748_1624255800757&xId=4vElNPj10GWNqNMaPYQvYOimFuFOvWw6pfrUQlqQxmZashaWS6s7jzrPgU9cPriUzr6RvGn1feGtOgen0TOuRvfqNQF0hGF0nD9ux0Z4RZlM&union_lens=lensId%3AMAPI%401624255800%402127fe06_081c_17a2d30a5ec_3c07%4001",
	}
	resp, err := api.Do(data)
	if err != nil {
		fmt.Println(fmt.Sprintf("request error: %s", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("tpwd create resp: %v", resp))
	}
}
