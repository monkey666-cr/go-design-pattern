package abstract_factory

import "fmt"

type MySQLMainDAO struct {
}

func (m *MySQLMainDAO) SaveOrderMain() {
	fmt.Println("MySQLMainDAO SaveOrderMain")
}

func (m *MySQLMainDAO) SearchOrderMain() {
	fmt.Println("MySQLMainDAO SearchOrderMain")
}

func (m *MySQLMainDAO) UpdateOrderMain() {
	fmt.Println("MySQLMainDAO UpdateOrderMain")
}

type MySQLDetailDAO struct {
}

func (m *MySQLDetailDAO) SaveOrderDetail() {
	fmt.Println("MySQLDetailDAO SaveOrderDetail")
}

type MySQLDAOFactory struct {
}

func (m *MySQLDAOFactory) CreateOrderMainDAO() OrderMainDao {
	return &MySQLMainDAO{}
}

func (m *MySQLDAOFactory) CreateOrderDetailDAO() OrderDetailDao {
	return &MySQLDetailDAO{}
}
