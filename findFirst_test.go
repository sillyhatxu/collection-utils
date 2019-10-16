package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollection_FindNumberFirst(t *testing.T) {
	array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	result, err := Stream(array).FindFirst(func(i int) bool {
		return i == 5
	}).Result()
	expect := 5
	resultInt := result.(int)
	assert.Nil(t, err)
	assert.EqualValues(t, resultInt, expect)
}

func TestCollection_FindUserFirst(t *testing.T) {
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
	expect := User{Id: 3, Name: "test-3", Status: true}
	result, err := Stream(array).FindFirst(func(user User) bool {
		return user.Id == 3
	}).Result()
	resultUser := result.(User)
	assert.Nil(t, err)
	assert.Equal(t, expect, resultUser)
}
