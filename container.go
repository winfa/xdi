package xdi

import "reflect"

type Container interface {
	Provide(provider any) error
	Invoke(fn any) error
	Resolve(target any) error
	InjectFields(target any) error
}

type container struct {
	providers map[reflect.Type]reflect.Value
	instances map[reflect.Type]reflect.Value
}

func NewContainer() Container {
	return container{
		providers: make(map[reflect.Type]reflect.Value),
		instances: make(map[reflect.Type]reflect.Value),
	}
}
