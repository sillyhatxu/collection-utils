package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollection_Sum(t *testing.T) {
	array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	var expect int64 = 45
	result, err := Stream(array).Sum(func(i int) int64 {
		return int64(i)
	})
	assert.Nil(t, err)
	assert.Equal(t, expect, result)
}

func TestCollection_User(t *testing.T) {
	type Product struct {
		Id       int64
		Name     string
		Status   bool
		Quantity int
		Price    int64
	}

	type ProductGroup struct {
		Id       int64
		Name     string
		Products []Product
	}
	var groups []ProductGroup
	groups = append(groups, ProductGroup{Id: 1, Name: "group-1", Products: []Product{
		{Id: 11, Name: "p11", Status: true, Quantity: 3, Price: 8500},
		{Id: 12, Name: "p12", Status: false, Quantity: 5, Price: 1500},
		{Id: 13, Name: "p13", Status: true, Quantity: 8, Price: 3800},
	}})
	groups = append(groups, ProductGroup{Id: 2, Name: "group-2", Products: []Product{
		{Id: 21, Name: "p21", Status: true, Quantity: 8, Price: 2800},
		{Id: 22, Name: "p22", Status: false, Quantity: 3, Price: 4500},
		{Id: 23, Name: "p23", Status: false, Quantity: 2, Price: 3500},
	}})
	groups = append(groups, ProductGroup{Id: 3, Name: "group-3", Products: []Product{
		{Id: 31, Name: "p31", Status: true, Quantity: 1, Price: 7777},
		{Id: 32, Name: "p32", Status: true, Quantity: 2, Price: 8888},
		{Id: 33, Name: "p33", Status: true, Quantity: 3, Price: 9999},
	}})

	var expect int64 = 161850
	var expectFilter int64 = 133850
	result, err := Stream(groups).Sum(func(pg ProductGroup) int64 {
		resultTotal, err := Stream(pg.Products).Sum(func(product Product) int64 {
			return product.Price * int64(product.Quantity)
		})
		assert.Nil(t, err)
		return resultTotal
	})
	assert.Nil(t, err)
	assert.EqualValues(t, expect, result)

	result, err = Stream(groups).Sum(func(pg ProductGroup) int64 {
		resultTotal, err := Stream(pg.Products).Filter(func(product Product) bool {
			return product.Status == true
		}).Sum(func(product Product) int64 {
			return product.Price * int64(product.Quantity)
		})
		assert.Nil(t, err)
		return resultTotal
	})
	assert.Nil(t, err)
	assert.EqualValues(t, result, expectFilter)

}
