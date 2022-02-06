package repository

type CustomerAddressRepository interface {
	GetCustomerAddress(int)
	GetCustomerAddresses()
	CreateCustomerAddress()
	UpdateCustomerAddress(int)
	DeleteCustomerAddress(int)
}