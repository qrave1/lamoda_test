package model

type Product struct {
	ID       uint
	Name     string
	Code     uint
	Quantity uint
	Size     uint
}

func NewProduct(ID uint, name string, code uint, quantity uint, size uint, warehouseID uint) *Product {
	return &Product{
		ID:       ID,
		Name:     name,
		Code:     code,
		Quantity: quantity,
		Size:     size,
	}
}
