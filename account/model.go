package account

type Model interface {
	GetDeliveryResult() any
}

type StructValueGetter interface {
	GetValues() map[string]any
}
