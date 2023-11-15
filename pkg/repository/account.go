package repository

type Account interface {
}

type account struct {
}

func NewAccount() Account {
	return &account{}
}
