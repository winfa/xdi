package xdi

import (
	"testing"
)

type AutoConstructdStruct struct {
	Name  string `inject:""`
	Age   int    `inject:""`
	Email string `inject:""`
}

type NestedAutoConstructdStruct struct {
	Injected AutoConstructdStruct `inject:""`
	Address  string               `inject:""`
}

type AutoConstructdStructWithNonInjectFields struct {
	Name  string `inject:""`
	Age   int    `inject:""`
	Email string `inject:""`
	Other string // Field without inject tag
}

func TestAutoConstructStruct(t *testing.T) {
	c := NewContainer()

	// Provide values for some fields
	_ = c.Provide(func() string {
		return "Auto Injected Name"
	})
	_ = c.Provide(func() int {
		return 50
	})
	_ = c.Provide(func() AutoConstructdStruct {
		return AutoConstructdStruct{
			Name:  "Nested Name",
			Age:   40,
			Email: "nested@example.com",
		}
	})

	// Resolve a struct with inject tags
	var result NestedAutoConstructdStruct
	err := c.Resolve(&result)
	if err != nil {
		t.Fatalf("Failed to resolve struct: %v", err)
	}

	// Validate the resolved struct
	if result.Injected.Name != "Nested Name" || result.Injected.Age != 40 || result.Injected.Email != "nested@example.com" {
		t.Fatalf("Injected struct does not match expected values: %+v", result.Injected)
	}
	if result.Address != "Auto Injected Name" {
		t.Fatalf("Address field does not match expected value: %s", result.Address)
	}
}

func TestAutoConstructStructWithoutProvider(t *testing.T) {
	c := NewContainer()

	// Resolve a struct without providing any values
	var result AutoConstructdStruct
	err := c.Resolve(&result)

	// Validate the error
	if err == nil {
		t.Fatalf("Expected error for missing provider, but got nil")
	}

	expectedError := "no provider found for type string"
	if err.Error() != expectedError {
		t.Fatalf("Expected error '%s', but got '%s'", expectedError, err.Error())
	}
}

func TestAutoConstructStructWithNonInjectFields(t *testing.T) {
	c := NewContainer()

	// Provide values for fields with inject tags
	_ = c.Provide(func() string {
		return "Auto Injected Name"
	})
	_ = c.Provide(func() int {
		return 50
	})
	_ = c.Provide(func() string {
		return "auto.injected@example.com"
	})

	// Resolve a struct with both inject and non-inject fields
	var result AutoConstructdStructWithNonInjectFields
	err := c.Resolve(&result)

	expectedError := "no provider found for type xdi.AutoConstructdStructWithNonInjectFields"
	if err.Error() != expectedError {
		t.Fatalf("Expected error '%s', but got '%s'", expectedError, err.Error())
	}
}
