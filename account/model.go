package account

import "github.com/ginger-core/query"

type Model interface {
	GetDeliveryResult() any
	ProcessUpdates(update query.Update) (changed bool)
}

type StructValueGetter interface {
	GetValues() map[string]any
}
