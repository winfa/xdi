package xdi

import (
	"errors"
	"reflect"
)

func (c *container) Provide(provider any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	providerValue := reflect.ValueOf(provider)
	providerType := reflect.TypeOf(provider)

	if providerType.Kind() != reflect.Func {
		c.providers[providerType] = providerValue
		c.instances[providerType] = providerValue
		return nil
	}

	providerFuncType := providerType
	if providerFuncType.NumOut() == 0 {
		return errors.New("provider function must return at least one value")
	}

	outType := providerFuncType.Out(0)
	c.providers[outType] = providerValue
	return nil
}
