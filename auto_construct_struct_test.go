package xdi

import (
	"reflect"
	"testing"
)

type StructWithInjectTags struct {
	Field1 string `inject:""`
	Field2 int    `inject:""`
}

type StructWithoutInjectTags struct {
	Field1 string
	Field2 int
}

type StructWithMixedTags struct {
	Field1 string `inject:""`
	Field2 int
}

func TestIsAutoConstructStructDataType(t *testing.T) {
	tests := []struct {
		name     string
		dataType reflect.Type
		expected bool
	}{
		{
			name:     "Struct with inject tags",
			dataType: reflect.TypeOf(StructWithInjectTags{}),
			expected: true,
		},
		{
			name:     "Struct without inject tags",
			dataType: reflect.TypeOf(StructWithoutInjectTags{}),
			expected: false,
		},
		{
			name:     "Struct with mixed tags",
			dataType: reflect.TypeOf(StructWithMixedTags{}),
			expected: false,
		},
		{
			name:     "Pointer to struct with inject tags",
			dataType: reflect.TypeOf(&StructWithInjectTags{}),
			expected: true,
		},
		{
			name:     "Pointer to struct without inject tags",
			dataType: reflect.TypeOf(&StructWithoutInjectTags{}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAutoConstructStructDataType(tt.dataType)
			if result != tt.expected {
				t.Fatalf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestIsAutoConstructStruct(t *testing.T) {
	tests := []struct {
		name     string
		dataType reflect.Type
		expected bool
	}{
		{
			name:     "Struct with inject tags",
			dataType: reflect.TypeOf(StructWithInjectTags{}),
			expected: true,
		},
		{
			name:     "Struct without inject tags",
			dataType: reflect.TypeOf(StructWithoutInjectTags{}),
			expected: false,
		},
		{
			name:     "Struct with mixed tags",
			dataType: reflect.TypeOf(StructWithMixedTags{}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAutoConstructStruct(tt.dataType)
			if result != tt.expected {
				t.Fatalf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
