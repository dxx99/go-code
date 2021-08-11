package abstract_factory

import "testing"

func Test_ExampleRdbFactory(t *testing.T)  {
	var factory DaoFactory
	factory = new(XmlDaoFactory)
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()

}

func saveMainAndDetail(factory DaoFactory)  {
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}