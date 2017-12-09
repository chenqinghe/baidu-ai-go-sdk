package gosdk

import (
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestStructToMap(t *testing.T) {
	user := &User{
		Name: "bob",
		Age:  18,
	}
	m := StructToMap(*user)
	name, ok := m["Name"]
	if !ok || name != "bob" {
		t.Fail()
	}
	age, ok := m["Age"]
	if !ok || age != 18 {
		t.Fail()
	}
}
