package lib

type Req struct {
	A int
	B int
}

type Resp struct {
	Total int
}

type Cal struct {

}

func (c *Cal) Add(req Req, resp *Resp) error {
	resp.Total = req.A + req.B
	return nil
}

func (c *Cal) Mul(req Req, resp *Resp) error {
	resp.Total = req.A	* req.B
	return nil
}