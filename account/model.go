package account

type Model interface {
	comparable
	GetDeliveryResult() any
}

type StructValueGetter interface {
	GetValues() map[string]any
}
