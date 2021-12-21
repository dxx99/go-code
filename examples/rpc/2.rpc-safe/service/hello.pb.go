package service

type HelloServ struct {
	
}

func NewHelloServ() *HelloServ {
	return &HelloServ{}
}

func (h HelloServ) Hello(req string, resp *string) error {
	*resp = "hello " + req
	return nil
}

