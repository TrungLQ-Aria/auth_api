package usecase

type Admin interface {
}

type admin struct {
}

func NewAdmin() Admin {
	return &admin{}
}
