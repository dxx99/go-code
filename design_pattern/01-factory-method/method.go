package factory_method

// Operator 定义一个接口去运算两个int值
type Operator interface {
	// Cal 这个操作，就是对应的工厂方法，每一组结构体都实现了这个方法
	Cal(int, int) int
}

type ComputerFactory interface {
	Create() Operator
}

// PlusFactory 实现加法工厂类
type PlusFactory struct {

}

func NewPlusFactory() *PlusFactory {
	return &PlusFactory{}
}

func (p *PlusFactory) Create() Operator {
	return &PlusOperator{}
}

// PlusOperator 加法的执行体
type PlusOperator struct {

}

func (p *PlusOperator) Cal(x int, y int) int {
	return x + y
}



// MulFactory 实现
type MulFactory struct {

}

func NewMulFactory() *MulFactory {
	return &MulFactory{}
}

func (m MulFactory) Create() Operator {
	return &MulOperator{}
}

type MulOperator struct {

}

func (m *MulOperator) Cal(x int, y int) int {
	return x * y
}


// todo 可以继续实现(减法，除法，取模)








