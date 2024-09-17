package account

type Base struct {
	Id uint64

	InternalStatus InternalStatus
	Status         Status
}

type Account[T Model] struct {
	Base `gorm:"embedded" json:",inline"`

	T T `gorm:"embedded" json:",inline"`
}

func NewAccount[T Model]() *Account[T] {
	return &Account[T]{}
}
