package test

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestExample(t *testing.T) {
	t.Log("TestExample")

	p := Person{
		Name: "test",
		Age:  10,
	}
	p.Name = "test2"

	fmt.Println(p)
}
