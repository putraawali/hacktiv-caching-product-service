package product

type Repository interface {
	Create(product *Product) (err error)
	FindAll() (products []Product, err error)
}
