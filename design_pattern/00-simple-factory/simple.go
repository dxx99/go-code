package simple_factory

// Animals 人类
type Animals interface {
	Say() string
}

type Dog struct {
	
}

func (d *Dog) Say() string {
	return "I am is a dog"
}

type Cat struct {

}

func (c *Cat) Say() string {
	return "I am is a cat"
}

func NewAnimal(t int) Animals {
	switch t {
	case 1:
		return &Cat{}
	case 2:
		return &Dog{}
	default:
		return nil
	}
}









