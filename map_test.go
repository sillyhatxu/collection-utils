package collection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollection_MapIntToString(t *testing.T) {
	array := []int{9, 8, 7, 6}
	expect := []string{"90", "80", "70", "60"}
	result, err := Stream(array).Map(func(i int) string {
		return fmt.Sprintf("%d", i*10)
	}).Result()
	resultArray := result.([]string)
	assert.Nil(t, err)
	assert.EqualValues(t, len(resultArray), 4)
	assert.Equal(t, expect, result)
}

func TestCollection_FilterAndMapUserToDTO(t *testing.T) {
	type User struct {
		Id     int64
		Name   string
		Status bool
	}
	type UserDTO struct {
		Id   int64
		Name string
	}
	array := []User{
		{Id: 1, Name: "test-1", Status: true},
		{Id: 2, Name: "test-2", Status: false},
		{Id: 3, Name: "test-3", Status: true},
		{Id: 4, Name: "test-4", Status: false},
		{Id: 5, Name: "test-5", Status: true},
	}
	expect := []UserDTO{
		{Id: 1, Name: "test-1"},
		{Id: 3, Name: "test-3"},
		{Id: 5, Name: "test-5"},
	}
	result, err := Stream(array).Filter(func(user User) bool {

		return user.Status == true
	}).Map(func(user User) UserDTO {
		return UserDTO{Id: user.Id, Name: user.Name}
	}).Result()
	resultArray := result.([]UserDTO)
	assert.Nil(t, err)
	assert.EqualValues(t, len(resultArray), 3)
	assert.Equal(t, expect, result)
}
