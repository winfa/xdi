package xdi

import (
	"reflect"
	"sync"
)

type Container interface {
	Provide(provider any) error
	Invoke(fn any) error
	Resolve(target any) error
	InjectFields(target any) error
}

type container struct {
	providers map[reflect.Type]reflect.Value
	instances map[reflect.Type]reflect.Value
	mu        sync.Mutex
}

func NewContainer() Container {
	return &container{
		providers: make(map[reflect.Type]reflect.Value),
		instances: make(map[reflect.Type]reflect.Value),
	}
}
