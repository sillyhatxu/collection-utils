package collection

import (
	"fmt"
	"reflect"
)

func (c *collection) validateCountInputSlice() (reflect.Value, error) {
	return c.validateSlice("Count")
}

func (c *collection) validateSumInputSlice() (reflect.Value, error) {
	return c.validateSlice("Sum")
}

func (c *collection) validateFilterInputSlice() (reflect.Value, error) {
	return c.validateSlice("Filter")
}

func (c *collection) validateMapInputSlice() (reflect.Value, error) {
	return c.validateSlice("Map")
}
func (c *collection) validateSortByInputSlice() (reflect.Value, error) {
	return c.validateSlice("SortBy")
}

func (c *collection) validateSlice(funcName string) (reflect.Value, error) {
	sv := reflect.ValueOf(c.input)
	if sv.Kind() != reflect.Slice {
		return reflect.Value{}, fmt.Errorf("collection.%s called with non-input value of type %T", funcName, c.input)
	}
	return sv, nil
}

func (c *collection) validateFindFirstFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	return c.validateFilterFunc(f)
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

func (c *collection) validateSumFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 1 ||
		funcType.NumOut() != 1 ||
		funcType.Out(0).Kind() != reflect.Int64 {
		return reflect.Value{}, nil, fmt.Errorf("collection.Filter called with invalid func. required func(in <T>) bool but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (c *collection) validateMapFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 1 ||
		funcType.NumOut() != 1 {
		return reflect.Value{}, nil, fmt.Errorf("collection.Map called with invalid func. required func(in <T>) out <T> but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}

func (c *collection) validateSortByFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func ||
		funcType.NumIn() != 2 ||
		funcType.NumOut() != 1 ||
		funcType.Out(0).Kind() != reflect.Bool {
		return reflect.Value{}, nil, fmt.Errorf("collection.SortBy called with invalid func. required func(in1, in2 <T>) bool but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}
