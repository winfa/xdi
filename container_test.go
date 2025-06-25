package xdi

import (
	"reflect"
	"testing"
)

func TestNewContainer(t *testing.T) {
	c := NewContainer()
	if c == nil {
		t.Fatal("Expected non-nil container")
	}
}

func TestProvide(t *testing.T) {
	c := NewContainer()

	provider := func() string {
		return "test"
	}

	err := c.Provide(provider)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Ensure provider is stored
	typ := reflect.TypeOf(provider).Out(0)
	if _, exists := c.(container).providers[typ]; !exists {
		t.Fatalf("Expected provider of type %v to be stored", typ)
	}
}

func TestInvoke(t *testing.T) {
	c := NewContainer()

	provider := func() string {
		return "test"
	}
	_ = c.Provide(provider)

	var invokedValue string
	invoker := func(val string) {
		invokedValue = val
	}

	err := c.Invoke(invoker)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if invokedValue != "test" {
		t.Fatalf("Expected invoked value to be 'test', got %v", invokedValue)
	}
}

func TestResolve(t *testing.T) {
	c := NewContainer()

	provider := func() string {
		return "test"
	}
	_ = c.Provide(provider)

	var resolvedValue string
	err := c.Resolve(&resolvedValue)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resolvedValue != "test" {
		t.Fatalf("Expected resolved value to be 'test', got %v", resolvedValue)
	}
}

func TestProvideError(t *testing.T) {
	c := NewContainer()

	provider := func() {}
	err := c.Provide(provider)
	if err == nil {
		t.Fatal("Expected error for invalid provider, got nil")
	}
}

func TestInvokeError(t *testing.T) {
	c := NewContainer()

	invoker := "invalid invoker"
	err := c.Invoke(invoker)
	if err == nil {
		t.Fatal("Expected error for invalid invoker, got nil")
	}
}

func TestResolveError(t *testing.T) {
	c := NewContainer()

	var resolvedValue string
	err := c.Resolve(resolvedValue) // Passing non-pointer
	if err == nil {
		t.Fatal("Expected error for non-pointer resolve, got nil")
	}

	err = c.Resolve(&resolvedValue) // No provider available
	if err == nil {
		t.Fatal("Expected error for unresolved dependency, got nil")
	}
}
