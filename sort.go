package collection

import "reflect"

/**
f interface{} = func(v1, v2 <T>) bool
*/
func (c *collection) SortBy(f interface{}) *collection {
	if c == nil {
		return &collection{err: CollectionNilError}
	}
	if c.err != nil {
		return &collection{err: c.err}
	}
	sv, err := c.validateSortByInputSlice()
	if err != nil {
		return &collection{err: err}
	}
	ret := reflect.MakeSlice(sv.Type(), sv.Len(), sv.Cap())
	reflect.Copy(ret, sv)

	funcValue, _, err := c.validateSortByFunc(f)
	if err != nil {
		return &collection{err: err}
	}
	executeSort(funcValue, ret)
	return &collection{input: ret.Interface()}
}

type customFunc struct {
	length int
	less   func(i, j int) bool
	swap   func(i, j int)
}

func (cf *customFunc) Len() int           { return cf.length }
func (cf *customFunc) Less(i, j int) bool { return cf.less(i, j) }
func (cf *customFunc) Swap(i, j int)      { cf.swap(i, j) }
