package collection

type collection struct {
	slice interface{}
	err   error
}

func New(slice interface{}) *collection {
	return &collection{slice: slice}
}

func (c *collection) Result() (interface{}, error) {
	return c.slice, c.err
}
