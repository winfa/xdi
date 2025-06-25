package xdi

import (
	"testing"
)

func TestInjectFields(t *testing.T) {
	c := NewContainer()

	// Provide values for fields with inject tags
	_ = c.Provide(func() string {
		return "Injected Name"
	})
	_ = c.Provide(func() int {
		return 30
	})

	// Create a struct to inject values into
	target := &AutoConstructdStruct{}

	// Inject fields
	err := c.InjectFields(target)
	if err != nil {
		t.Fatalf("Failed to inject fields: %v", err)
	}

	// Validate injected values
	if target.Name != "Injected Name" {
		t.Fatalf("Expected Name to be 'Injected Name', but got '%s'", target.Name)
	}
	if target.Age != 30 {
		t.Fatalf("Expected Age to be 30, but got %d", target.Age)
	}
}

func TestInjectFieldsWithMissingProvider(t *testing.T) {
	c := NewContainer()

	// Create a struct to inject values into
	target := &AutoConstructdStruct{}

	// Inject fields without providing values
	err := c.InjectFields(target)

	// Validate the error
	if err == nil {
		t.Fatalf("Expected error for missing provider, but got nil")
	}

	expectedError := "no provider found for type string"
	if err.Error() != expectedError {
		t.Fatalf("Expected error '%s', but got '%s'", expectedError, err.Error())
	}
}

func TestInjectFieldsWithNonPointerTarget(t *testing.T) {
	c := NewContainer()

	// Create a non-pointer struct
	var target TestStruct

	// Attempt to inject fields
	err := c.InjectFields(target)

	// Validate the error
	if err == nil {
		t.Fatalf("Expected error for non-pointer target, but got nil")
	}

	expectedError := "target must be a pointer"
	if err.Error() != expectedError {
		t.Fatalf("Expected error '%s', but got '%s'", expectedError, err.Error())
	}
}
