package account

type Account[T Model] struct {
	Id uint64

	Status Status
}

func NewAccount[T Model]() *Account[T] {
	return &Account[T]{}
}
