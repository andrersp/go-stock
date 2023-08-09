package product

type ProductRepository interface {
	CreateProduct(*Product) (*Product, error)
	GetProductByID(int) (*Product, error)
	ListProductsByCategory(string) ([]Product, error)
	ListProducts() ([]Product, error)
}
