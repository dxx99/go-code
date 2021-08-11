package abstract_factory

import "fmt"

//抽象工厂

type OrderMainDao interface {
	SaveOrderMain()
}

type OrderDetailDao interface {
	SaveOrderDetail()
}

type DaoFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

type RdbMainDao struct {

}
func (*RdbMainDao) SaveOrderMain() {
	fmt.Println("rdb main save")
}

type RdbDetailDao struct {

}
func (*RdbDetailDao) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

type RdbDaoFactory struct {

}

// CreateOrderMainDao
func (*RdbDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &RdbMainDao{}
}

//CreateOrderDetailDao
func (*RdbDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RdbDetailDao{}
}

type XmlMainDao struct {

}
func (*XmlMainDao) SaveOrderMain() {
	fmt.Println("xml main save")
}

type XmlDetailDao struct {

}
func (*XmlDetailDao) SaveOrderDetail() {
	fmt.Println("xml detail save")
}

type XmlDaoFactory struct {

}

func (*XmlDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &XmlMainDao{}
}

func (*XmlDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XmlDetailDao{}
}





