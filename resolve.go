package xdi

import (
	"errors"
	"reflect"
)

func (c *container) Resolve(target any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
	}
	targetType := targetValue.Type().Elem()

	providerValue, err := c.getProviderValueByType(targetType)
	if err != nil {
		return err
	}

	targetValue.Elem().Set(providerValue)
	return nil
}
