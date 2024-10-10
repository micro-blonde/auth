package profile

import "github.com/ginger-core/query"

type Model interface {
	comparable
	GetDeliveryResult() any
	ProcessUpdates(update query.Update) (changed bool)
}

type StructValueGetter interface {
	GetValues() map[string]any
}
