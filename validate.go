package collection

import (
	"fmt"
	"reflect"
)

func (c *collection) validateFilterInputSlice() (reflect.Value, error) {
	return c.validateSlice("Filter")
}

func (c *collection) validateSlice(funcName string) (reflect.Value, error) {
	sv := reflect.ValueOf(c.slice)
	if sv.Kind() != reflect.Slice {
		return reflect.Value{}, fmt.Errorf("collection.%s called with non-slice value of type %T", funcName, c.slice)
	}
	return sv, nil
}

func (c *collection) validateFilterFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 1 ||
		funcType.NumOut() != 1 ||
		funcType.Out(0).Kind() != reflect.Bool {
		return reflect.Value{}, nil, fmt.Errorf("collection.Filter called with invalid func. required func(in <T>) bool but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}
