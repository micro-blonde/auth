package profile

type Model interface {
	GetDeliveryResult() any
	GetUpdates() map[string]any
}

type StructValueGetter interface {
	GetValues() map[string]any
}
