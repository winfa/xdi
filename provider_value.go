package xdi

import (
	"errors"
	"reflect"
)

func (c container) getProviderValueByType(targetType reflect.Type) (reflect.Value, error) {
	provider, exists := c.providers[targetType]
	if !exists {
		if isAutoConstructStructDataType(targetType) {
			structValuePointer := reflect.New(targetType)
			c.InjectFields(structValuePointer)

			c.instances[targetType] = structValuePointer.Elem()
			c.providers[targetType] = structValuePointer.Elem()

			return structValuePointer.Elem(), nil
		}
		return reflect.Value{}, errors.New("no provider found for type " + targetType.String())
	}

	if provider.Kind() != reflect.Func {
		return provider, nil
	}

	tryProviderValue, exists := c.instances[targetType]
	if exists {
		return tryProviderValue, nil
	}

	result, err := c.invokeFunction(provider)
	if err != nil {
		return reflect.Value{}, err
	}
	if len(result) == 0 {
		return reflect.Value{}, errors.New("provider function must return at least one value")
	}

	providerValue := result[0]
	c.instances[targetType] = providerValue

	return providerValue, nil
}
