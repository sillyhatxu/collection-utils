package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollection_SortBy(t *testing.T) {
	array := []int{1, 4, 7, 8, 5, 2, 9, 3, 6}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result, err := Stream(array).SortBy(func(i1, i2 int) bool {
		return i1 < i2
	}).Result()

	resultArray := result.([]int)
	assert.Nil(t, err)
	assert.EqualValues(t, len(resultArray), 9)
	assert.Equal(t, expect, result)
}

func TestCollection_SortByUser(t *testing.T) {
	type User struct {
		Id     int64
		Name   string
		Status bool
	}
	array := []User{
		{Id: 2, Name: "test-2", Status: false},
		{Id: 1, Name: "test-1", Status: true},
		{Id: 5, Name: "test-5", Status: true},
		{Id: 4, Name: "test-4", Status: false},
		{Id: 3, Name: "test-3", Status: true},
	}
	expect := []User{
		{Id: 1, Name: "test-1", Status: true},
		{Id: 2, Name: "test-2", Status: false},
		{Id: 3, Name: "test-3", Status: true},
		{Id: 4, Name: "test-4", Status: false},
		{Id: 5, Name: "test-5", Status: true},
	}
	result, err := Stream(array).SortBy(func(u1, u2 User) bool {
		return u1.Id < u2.Id
	}).Result()

	resultArray := result.([]User)
	assert.Nil(t, err)
	assert.EqualValues(t, len(resultArray), 5)
	assert.Equal(t, expect, result)
}
