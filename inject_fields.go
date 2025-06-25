package xdi

import (
	"errors"
	"reflect"
)

// InjectFields implements Container.
func (c container) InjectFields(target any) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
	}

	dataType := targetValue.Type().Elem()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		_, ok := field.Tag.Lookup("inject")
		if !ok {
			continue
		}

		fieldValue, err := c.getProviderValueByType(field.Type)
		if err != nil {
			return err
		}

		targetValue.Elem().FieldByName(field.Name).Set(fieldValue)
	}

	return nil
}
