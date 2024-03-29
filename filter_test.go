package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollection_ArraySizeZeroFilter(t *testing.T) {
	array := []int{}
	expect := []int{}
	result, err := Stream(array).Filter(func(i int) bool {
		return i >= 3 && i <= 6
	}).Result()
	resultArray := result.([]int)
	assert.Nil(t, err)
	assert.EqualValues(t, len(resultArray), 0)
	assert.Equal(t, expect, result)
}

func TestCollection_NumberFilter(t *testing.T) {
	array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	expect := []int{6, 5, 4, 3}
	result, err := Stream(array).Filter(func(i int) bool {
		return i >= 3 && i <= 6
	}).Result()
	resultArray := result.([]int)
	assert.Nil(t, err)
	assert.EqualValues(t, len(resultArray), 4)
	assert.Equal(t, expect, result)
}

func TestCollection_Filter(t *testing.T) {
	type User struct {
		Id     int64
		Name   string
		Status bool
	}

	array := []User{
		{Id: 1, Name: "test-1", Status: true},
		{Id: 2, Name: "test-2", Status: false},
		{Id: 3, Name: "test-3", Status: true},
		{Id: 4, Name: "test-4", Status: false},
		{Id: 5, Name: "test-5", Status: true},
	}
	expect := []User{
		{Id: 1, Name: "test-1", Status: true},
		{Id: 3, Name: "test-3", Status: true},
		{Id: 5, Name: "test-5", Status: true},
	}
	result, err := Stream(array).Filter(func(user User) bool {
		return user.Status == true
	}).Result()
	userArray := result.([]User)
	assert.Nil(t, err)
	assert.EqualValues(t, len(userArray), 3)
	assert.Equal(t, expect, result)

	count, err := Stream(array).Filter(func(user User) bool {
		return user.Status == true
	}).Count()
	assert.Nil(t, err)
	assert.EqualValues(t, count, 3)
}

func TestCollection_PointerFilter(t *testing.T) {
	type User struct {
		Id     int64
		Name   string
		Status bool
	}

	array := []*User{
		{Id: 1, Name: "test-1", Status: true},
		{Id: 2, Name: "test-2", Status: false},
		{Id: 3, Name: "test-3", Status: true},
		{Id: 4, Name: "test-4", Status: false},
		{Id: 5, Name: "test-5", Status: true},
	}
	expect := []*User{
		{Id: 1, Name: "test-1", Status: true},
		{Id: 3, Name: "test-3", Status: true},
		{Id: 5, Name: "test-5", Status: true},
	}
	result, err := Stream(array).Filter(func(user *User) bool {
		return user.Status == true
	}).Result()
	userArray := result.([]*User)
	assert.Nil(t, err)
	assert.EqualValues(t, len(userArray), 3)
	assert.Equal(t, expect, result)
}
