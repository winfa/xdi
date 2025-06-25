package xdi

import (
	"reflect"
	"strconv"
	"testing"
)

func TestProvideValidFunctionNoParams(t *testing.T) {
	c := NewContainer()

	provider := func() string {
		return "test"
	}

	err := c.Provide(provider)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	typ := reflect.TypeOf(provider).Out(0)
	if _, exists := c.(container).providers[typ]; !exists {
		t.Fatalf("Expected provider of type %v to be stored", typ)
	}
}

func TestProvideValidFunctionMultipleParams(t *testing.T) {
	c := NewContainer()

	// Define providers
	provider1 := func() string {
		return "param1"
	}
	provider2 := func() int {
		return 42
	}

	// Provide the providers to the container
	_ = c.Provide(provider1)
	_ = c.Provide(provider2)

	// Define an invoker function that uses the provided values
	invoker := func(val1 string, val2 int) string {
		return val1 + " and " + strconv.Itoa(val2)
	}

	// Invoke the function using the container
	var result string
	err := c.Invoke(func(val1 string, val2 int) {
		result = invoker(val1, val2)
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate the result
	expected := "param1 and 42"
	if result != expected {
		t.Fatalf("Expected result to be '%s', got '%s'", expected, result)
	}
}
