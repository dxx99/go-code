package builder

import (
	"fmt"
	"testing"
)

func TestBuilder1_GetResult(t *testing.T) {
	builder := new(Builder1)
	director := NewDirector(builder)
	director.Construct()

	fmt.Println(builder.GetResult())	// output: string 123
}

func TestBuilder2_GetResult(t *testing.T) {
	builder := new(Builder2)
	director := NewDirector(builder)
	director.Construct()

	fmt.Println(builder.GetResult())
}



