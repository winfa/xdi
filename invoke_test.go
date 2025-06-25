package xdi

import (
	"strconv"
	"testing"
)

func TestInvokeSingleParameter(t *testing.T) {
	c := NewContainer()

	provider := func() string {
		return "test"
	}
	_ = c.Provide(provider)

	invoker := func(val string) string {
		return val + " invoked"
	}

	var result string
	err := c.Invoke(func(val string) {
		result = invoker(val)
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "test invoked"
	if result != expected {
		t.Fatalf("Expected result to be '%s', got '%s'", expected, result)
	}
}

type CombinedData struct {
	Num     int
	Content string
}

func TestInvokeMultipleParameters(t *testing.T) {
	c := NewContainer()

	provider1 := func() string {
		return "param1"
	}
	provider2 := func() int {
		return 42
	}
	provider3 := func(input1 string, input2 int) CombinedData {
		return CombinedData{
			Num:     input2,
			Content: input1,
		}
	}
	_ = c.Provide(provider1)
	_ = c.Provide(provider2)
	_ = c.Provide(provider3)

	var result string
	err := c.Invoke(func(val CombinedData) {
		result = val.Content + " and " + strconv.Itoa(val.Num)
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "param1 and 42"
	if result != expected {
		t.Fatalf("Expected result to be '%s', got '%s'", expected, result)
	}
}

func TestInvokeNoParameters(t *testing.T) {
	c := NewContainer()

	invoker := func() string {
		return "no params invoked"
	}

	var result string
	err := c.Invoke(func() {
		result = invoker()
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "no params invoked"
	if result != expected {
		t.Fatalf("Expected result to be '%s', got '%s'", expected, result)
	}
}

func TestInvokeErrorInvalidInvoker(t *testing.T) {
	c := NewContainer()

	invoker := "invalid invoker"
	err := c.Invoke(invoker)
	if err == nil {
		t.Fatal("Expected error for invalid invoker, got nil")
	}
}

func TestInvokeErrorMissingProvider(t *testing.T) {
	c := NewContainer()

	invoker := func(val string) {}
	err := c.Invoke(invoker)
	if err == nil {
		t.Fatal("Expected error for missing provider, got nil")
	}
}

func TestInvokeWithAutoProvideStruct(t *testing.T) {
	c := NewContainer()

	// Provide values for fields with inject tags
	_ = c.Provide(func() string {
		return "Auto Injected Name"
	})
	_ = c.Provide(func() int {
		return 50
	})

	// Invoke a function with AutoProvidedStruct as parameter
	err := c.Invoke(func(val AutoProvidedStruct) {
		// Validate the resolved struct
		if val.Name != "Auto Injected Name" || val.Age != 50 {
			t.Fatalf("Resolved struct does not match expected values: %+v", val)
		}
	})
	if err != nil {
		t.Fatalf("Failed to invoke function with auto-provided struct: %v", err)
	}
}

func TestInvokeWithMissingAutoProvideStructFields(t *testing.T) {
	c := NewContainer()

	// Invoke a function with AutoProvidedStruct as parameter, but missing providers
	err := c.Invoke(func(val AutoProvidedStruct) {
		// This block should not be executed due to missing providers
		t.Fatal("Expected error for missing providers, but function was invoked")
	})
	if err == nil {
		t.Fatal("Expected error for missing providers, got nil")
	}
}
