package xdi

import (
	"reflect"
	"sync"
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
	if _, exists := c.(*container).providers[typ]; !exists {
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

func TestConcurrentProvideAndResolve(t *testing.T) {
	c := NewContainer()
	var wg sync.WaitGroup

	// Concurrently provide values
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := c.Provide(func() string {
			return "Concurrent Test"
		})
		if err != nil {
			t.Errorf("Failed to provide value: %v", err)
		}
	}()

	// Concurrently resolve values
	wg.Add(1)
	go func() {
		defer wg.Done()
		var resolvedValue string
		err := c.Resolve(&resolvedValue)
		if err != nil && err.Error() != "no provider found for type string" {
			t.Errorf("Unexpected error during resolve: %v", err)
		}
	}()

	wg.Wait()

	// Validate resolved value
	var resolvedValue string
	err := c.Resolve(&resolvedValue)
	if err != nil {
		t.Fatalf("Failed to resolve value after concurrent operations: %v", err)
	}
	if resolvedValue != "Concurrent Test" {
		t.Fatalf("Expected resolved value to be 'Concurrent Test', but got '%s'", resolvedValue)
	}
}

func TestConcurrentAccessToProviders(t *testing.T) {
	c := NewContainer()
	var wg sync.WaitGroup

	// Provide a value
	err := c.Provide(func() string {
		return "Test Value"
	})
	if err != nil {
		t.Fatalf("Failed to provide value: %v", err)
	}

	// Concurrently resolve values
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var resolvedValue string
			err := c.Resolve(&resolvedValue)
			if err != nil {
				t.Errorf("Failed to resolve value: %v", err)
			}
			if resolvedValue != "Test Value" {
				t.Errorf("Expected resolved value to be 'Test Value', but got '%s'", resolvedValue)
			}
		}()
	}

	wg.Wait()
}
