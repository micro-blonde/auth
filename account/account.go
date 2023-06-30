package account

type Base struct {
	Id uint64

	Status Status
}

type Account[T Model] struct {
	Base

	T T `gorm:"embedded" json:",inline"`
}

func NewAccount[T Model]() *Account[T] {
	return &Account[T]{}
}
