package xdi

import (
	"errors"
	"reflect"
)

type FunctionCallResult = []reflect.Value

func (c container) Invoke(invoker any) error {
	invokerValue := reflect.ValueOf(invoker)

	if invokerValue.Kind() != reflect.Func {
		return errors.New("invoker is not a function")
	}

	_, err := c.invokeFunction(invokerValue)
	if err != nil {
		return err
	}

	return nil
}

func (c container) invokeFunction(invoker reflect.Value) (FunctionCallResult, error) {
	if invoker.Kind() != reflect.Func {
		return nil, errors.New("invoker is not a function")
	}

	invokerType := invoker.Type()
	if invokerType.NumIn() == 0 {
		return invoker.Call(nil), nil
	}

	paramsList := []reflect.Value{}
	for i := 0; i < invokerType.NumIn(); i++ {
		paramType := invokerType.In(i)
		paramValue, err := c.getProviderValueByType(paramType)
		if err != nil {
			return nil, err
		}
		paramsList = append(paramsList, paramValue)
	}

	return invoker.Call(paramsList), nil
}
