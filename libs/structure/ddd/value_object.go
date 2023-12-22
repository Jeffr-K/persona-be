package ddd

type ValueObject interface {
	Create()
	Delete()
}

type ValueObjectImpl struct {
}

func (vo ValueObjectImpl) Create() {}

func (vo ValueObjectImpl) Delete() {}
