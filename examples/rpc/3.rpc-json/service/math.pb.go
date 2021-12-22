package service

type MathServ struct {
	
}

func (m *MathServ) Add(req ReqNumbers, resp *ResTotal) error {
	resp.Total = req.A + req.B
	return nil
}

func (m *MathServ) Sub(req ReqNumbers, resp *ResTotal) error {
	resp.Total = req.A - req.B
	return nil
}

func NewMathServ() *MathServ {
	return &MathServ{}
}


