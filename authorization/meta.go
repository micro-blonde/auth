package authorization

type Meta map[string]any

func (m Meta) Add(key string, value any) {
	m[key] = value
}

func (m Meta) Get(key string) any {
	return m[key]
}
