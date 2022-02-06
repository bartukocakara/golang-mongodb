package repository

type OrderItemRepository interface {
	GetOrderItem(int)
	GetOrderItems()
	CreateOrderItem()
	UpdateOrderItem(int)
	DeleteOrderItem(int)
}