package collection

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Id     int64
	Name   string
	Status bool
}

func TestCollection_Filter(t *testing.T) {
	array := []*User{
		{Id: 1, Name: "test-1", Status: true},
		{Id: 2, Name: "test-2", Status: false},
		{Id: 3, Name: "test-3", Status: true},
		{Id: 4, Name: "test-4", Status: false},
		{Id: 5, Name: "test-5", Status: true},
	}
	result, err := New(array).Filter(func(i interface{}) bool {
		fmt.Println(i)
		return i.(*User).Status == true
	}).Result()
	userArray := result.([]*User)
	assert.Nil(t, err)
	assert.EqualValues(t, len(userArray), 3)
	fmt.Println(fmt.Sprintf("%#v", userArray))
}
