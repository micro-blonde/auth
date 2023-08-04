package account

type InternalStatus uint64

const (
	InternalStatusRegistered InternalStatus = 0b1 << iota
	InternalStatusVerified
	InternalStatusRejected
)

func (s InternalStatus) Uint64() uint64 {
	return uint64(s)
}

func (s *InternalStatus) Add(status InternalStatus) InternalStatus {
	*s |= status
	return *s
}

func (s InternalStatus) Is(status InternalStatus) bool {
	return s&status == status
}

func (s InternalStatus) Has(status InternalStatus) bool {
	return s&status > 0
}
