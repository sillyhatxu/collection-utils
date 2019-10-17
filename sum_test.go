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
	type User struct {
		Id     int64
		Name   string
		Status bool
	}
}
