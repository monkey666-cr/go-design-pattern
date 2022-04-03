package abstract_factory

import "testing"

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestAbstractFactory_MySQL_Create(t *testing.T) {
	var factory DAOFactory
	factory = &MySQLDAOFactory{}
	getMainAndDetail(factory)
}
