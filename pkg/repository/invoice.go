package repository

type InvoiceRepository interface {
	GetInvoice(int)
	GetInvoices()
	CreateInvoice()
	UpdateInvoice(int)
	DeleteInvoice(int)
}