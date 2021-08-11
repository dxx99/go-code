package factory_method

import (
	"fmt"
	"testing"
)

// 测试加法操作
func TestPlusOperator_Cal(t *testing.T) {
	pf := NewPlusFactory()
	data := pf.Create().Cal(1, 2)
	t.Log(data)
}

// 测试乘法操作
func TestMulOperator_Cal(t *testing.T) {
	pf := NewMulFactory()
	data := pf.Create().Cal(1, 2)
	t.Log(data)
}

func TestBitOperate(t *testing.T)  {
	a := 1<<0
	b := 1<<1
	c := 1<<2
	d := 1<<3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(9&b)
}