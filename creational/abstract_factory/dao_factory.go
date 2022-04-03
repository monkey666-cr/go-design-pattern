package abstract_factory

// sqlserver
// mysql
// oracle

// OrderMainDao 订单记录
type OrderMainDao interface {
	// SaveOrderMain 保存
	SaveOrderMain()
	// SearchOrderMain 搜索
	SearchOrderMain()
	// UpdateOrderMain 更新
	UpdateOrderMain()
}

// OrderDetailDao 订单详情
type OrderDetailDao interface {
	// SaveOrderDetail 保存
	SaveOrderDetail()
}

// DAOFactory 抽象工厂函数
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDao
	CreateOrderDetailDAO() OrderDetailDao
}
