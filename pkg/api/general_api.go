package api

type GeneralAPI service

func (g *GeneralAPI) Methods() string {
	return g.Method
}

func (g *GeneralAPI) SetMethods(method string) {
	g.Method = method
}

func (g *GeneralAPI) Do(data interface{}) (interface{}, error) {
	var resp interface{}
	err := g.client.do(g, data, &resp)
	return resp, err
}
