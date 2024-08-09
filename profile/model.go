package profile

type Model interface {
	comparable
	GetDeliveryResult() any
	GetUpdates() map[string]any
}

type StructValueGetter interface {
	GetValues() map[string]any
}
