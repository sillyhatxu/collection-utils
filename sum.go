package collection

/**
f interface{} = func(v <T>) int64
*/
func (c *collection) sum(f interface{}) *collection {
	if c == nil {
		return &collection{err: CollectionNilError}
	}
	if c.err != nil {
		return &collection{err: c.err}
	}
	sv, err := c.validateFilterInputSlice()
	if err != nil {
		return &collection{err: err}
	}
	funcValue, _, err := c.validateSumFunc(f)
	if err != nil {
		return &collection{err: err}
	}
	var total int64 = 0
	if sv.Len() == 0 {
		return &collection{input: total}
	}
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		total += executeSumFunc(funcValue, v)
	}
	return &collection{input: total}
}
