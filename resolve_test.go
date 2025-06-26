package xdi

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Name  string
	Age   int
	Email string
}

func TestResolveWithProvider(t *testing.T) {
	c := NewContainer()

	// Define a provider for TestStruct
	provider := func() TestStruct {
		return TestStruct{
			Name:  "John Doe",
			Age:   30,
			Email: "john.doe@example.com",
		}
	}

	// Register the provider
	err := c.Provide(provider)
	if err != nil {
		t.Fatalf("Failed to provide struct: %v", err)
	}

	// Resolve the struct
	var result TestStruct
	err = c.Resolve(&result)
	if err != nil {
		t.Fatalf("Failed to resolve struct: %v", err)
	}

	// Validate the resolved struct
	if result.Name != "John Doe" || result.Age != 30 || result.Email != "john.doe@example.com" {
		t.Fatalf("Resolved struct does not match expected values: %+v", result)
	}
}

func TestResolveWithoutProvider(t *testing.T) {
	c := NewContainer()

	// Attempt to resolve a struct without a provider
	var result TestStruct
	err := c.Resolve(&result)

	// Validate the error
	if err == nil || err.Error() != "no provider found for type xdi.TestStruct" {
		t.Fatalf("Expected error for missing provider, got: %v", err)
	}
}

func TestResolveNonPointerTarget(t *testing.T) {
	c := NewContainer()

	// Attempt to resolve a non-pointer target
	var result TestStruct
	err := c.Resolve(result)

	// Validate the error
	if err == nil || err.Error() != "target must be a pointer" {
		t.Fatalf("Expected error for non-pointer target, got: %v", err)
	}
}

func TestResolveWithNonFunctionProvider(t *testing.T) {
	c := NewContainer()

	// Register a non-function provider (direct value)
	provider := TestStruct{
		Name:  "Jane Doe",
		Age:   25,
		Email: "jane.doe@example.com",
	}
	c.(*container).providers[reflect.TypeOf(TestStruct{})] = reflect.ValueOf(provider)

	// Resolve the struct
	var result TestStruct
	err := c.Resolve(&result)
	if err != nil {
		t.Fatalf("Failed to resolve struct: %v", err)
	}

	// Validate the resolved struct
	if result.Name != "Jane Doe" || result.Age != 25 || result.Email != "jane.doe@example.com" {
		t.Fatalf("Resolved struct does not match expected values: %+v", result)
	}
}
