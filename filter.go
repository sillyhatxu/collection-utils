package collection

import (
	"fmt"
	"reflect"
)

func (c *collection) Filter(f func(i interface{}) bool) *collection {
	if c == nil {
		return &collection{err: CollectionNilError}
	}
	inputValue := reflect.ValueOf(c.slice)
	if inputValue.Kind() != reflect.Slice {
		return &collection{err: fmt.Errorf("collection.Filter called with non-slice value of type %T", c.slice)}
	}
	if inputValue.Len() == 0 {
		return &collection{slice: inputValue}
	}
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()

	resultSliceType := reflect.SliceOf(funcType.In(0))
	fmt.Println(resultSliceType)
	ret := reflect.MakeSlice(resultSliceType, 0, inputValue.Len()) //make slice
	for i := 0; i < inputValue.Len(); i++ {
		if v := inputValue.Index(i); f(v.Interface()) {
			ret = reflect.Append(ret, v)
		}
	}
	return &collection{
		slice: ret.Interface(),
	}
}

func (c *collection) validateFilterFunc(f interface{}) (reflect.Value, reflect.Type, error) {
	funcValue := reflect.ValueOf(f)
	funcType := funcValue.Type()
	if funcType.Kind() != reflect.Func {
		return reflect.Value{}, nil, fmt.Errorf("gollection.Filter called with invalid func. required func(in <T>) bool but supplied %v", funcType)
	}
	return funcValue, funcType, nil
}
