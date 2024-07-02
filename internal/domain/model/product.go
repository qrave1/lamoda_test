package model

type Product struct {
	ID   uint
	Name string
	Code uint
	Size uint
}

func NewProduct(ID uint, name string, code, size uint) *Product {
	return &Product{
		ID:   ID,
		Name: name,
		Code: code,
		Size: size,
	}
}
