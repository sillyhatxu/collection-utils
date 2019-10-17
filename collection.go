package collection

import (
	"fmt"
	"reflect"
)

type collection struct {
	input interface{}
	err   error
}

func Stream(input interface{}) *collection {
	return &collection{input: input}
}

func (c *collection) Result() (interface{}, error) {
	return c.input, c.err
}

func (c *collection) Count() (int, error) {
	if c.err != nil {
		return 0, c.err
	}
	sv, err := c.validateCountInputSlice()
	if err != nil {
		return 0, err
	}
	return sv.Len(), nil
}

func (c *collection) Sum(f interface{}) (int64, error) {
	if c.err != nil {
		return 0, c.err
	}
	result, err := c.sum(f).Result()
	if err != nil {
		return 0, err
	}
	sv := reflect.ValueOf(result)
	if sv.Kind() != reflect.Int64 {
		return 0, fmt.Errorf("collection.Sum called with non-input value of type %T", result)
	}
	return result.(int64), nil
}
