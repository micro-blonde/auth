package profile

type Base struct {
	Id uint64
}

type Profile[T Model] struct {
	Base `gorm:"embedded" json:",inline"`

	T T `gorm:"embedded" json:",inline"`
}

func NewProfile[T Model]() *Profile[T] {
	return &Profile[T]{}
}
