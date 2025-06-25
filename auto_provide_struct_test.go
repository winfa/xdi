package xdi

import (
	"testing"
)

type AutoProvidedStruct struct {
	Name  string `inject:""`
	Age   int    `inject:""`
	Email string `inject:""`
}

type NestedAutoProvidedStruct struct {
	Injected AutoProvidedStruct `inject:""`
	Address  string             `inject:""`
}

type AutoProvidedStructWithNonInjectFields struct {
	Name  string `inject:""`
	Age   int    `inject:""`
	Email string `inject:""`
	Other string // Field without inject tag
}

func TestAutoProvideStruct(t *testing.T) {
	c := NewContainer()

	// Provide values for some fields
	_ = c.Provide(func() string {
		return "Auto Injected Name"
	})
	_ = c.Provide(func() int {
		return 50
	})
	_ = c.Provide(func() AutoProvidedStruct {
		return AutoProvidedStruct{
			Name:  "Nested Name",
			Age:   40,
			Email: "nested@example.com",
		}
	})

	// Resolve a struct with inject tags
	var result NestedAutoProvidedStruct
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

func TestAutoProvideStructWithoutProvider(t *testing.T) {
	c := NewContainer()

	// Resolve a struct without providing any values
	var result AutoProvidedStruct
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

func TestAutoProvideStructWithNonInjectFields(t *testing.T) {
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
	var result AutoProvidedStructWithNonInjectFields
	err := c.Resolve(&result)

	expectedError := "no provider found for type xdi.AutoProvidedStructWithNonInjectFields"
	if err.Error() != expectedError {
		t.Fatalf("Expected error '%s', but got '%s'", expectedError, err.Error())
	}
}
