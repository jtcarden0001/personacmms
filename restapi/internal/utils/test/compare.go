package test

import (
	"reflect"
	"testing"
	"time"
)

// TODO: some optimization and code reduction to be had here in these comparison functions
func CompEntities(t *testing.T, expected interface{}, actual interface{}) {
	// exclude no fields
	CompEntitiesExcludeFields(t, expected, actual, make(map[string]struct{}))
}

func CompEntitiesExcludeFields(t *testing.T, expected interface{}, actual interface{}, fields map[string]struct{}) {
	// Compare all properties except for the specified fields
	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)

	for i := 0; i < expectedValue.NumField(); i++ {
		field := expectedValue.Type().Field(i)
		if _, ok := fields[field.Name]; ok {
			continue
		}

		expectedField := expectedValue.Field(i).Interface()
		actualField := actualValue.Field(i).Interface()

		isPtr := reflect.TypeOf(expectedField).Kind() == reflect.Ptr
		if isPtr {
			comparePointers(t, expectedField, actualField, field)
		} else {
			compareValues(t, expectedField, actualField, field)
		}
	}
}

func comparePointers(t *testing.T, expectedField interface{}, actualField interface{}, field reflect.StructField) {
	if expectedField == nil && actualField != nil {
		t.Errorf("Compare failed: expected %v for field %s, got %v", expectedField, field.Name, actualField)
	} else if expectedField != nil && actualField == nil {
		t.Errorf("Compare failed: expected %v for field %s, got %v", expectedField, field.Name, actualField)
	} else if expectedField != nil && actualField != nil {
		if !reflect.DeepEqual(expectedField, actualField) {
			t.Errorf("Compare failed: expected %v for field %s, got %v", expectedField, field.Name, actualField)
		}
	}
}

func compareValues(t *testing.T, expectedField interface{}, actualField interface{}, field reflect.StructField) {
	if field.Type == reflect.TypeOf(time.Time{}) {
		expectedTime := expectedField.(time.Time)
		actualTime := actualField.(time.Time)
		if !expectedTime.Truncate(time.Second).Equal(actualTime.Truncate(time.Second)) {
			t.Errorf("Compare failed: expected %v for field %s, got %v", expectedTime, field.Name, actualTime)
		}
	} else if !reflect.DeepEqual(expectedField, actualField) {
		t.Errorf("Compare failed: expected %v for field %s, got %v", expectedField, field.Name, actualField)
	}
}

func CompEntitiesFieldsShouldBeDifferent(t *testing.T, inital interface{}, updated interface{}, fields map[string]struct{}) {
	// Compare all properties make sure fields are different
	initalValue := reflect.ValueOf(inital)
	updatedValue := reflect.ValueOf(updated)

	for i := 0; i < initalValue.NumField(); i++ {
		field := initalValue.Type().Field(i)
		different := false
		if _, ok := fields[field.Name]; ok {
			different = true
		}

		initalField := initalValue.Field(i).Interface()
		updatedField := updatedValue.Field(i).Interface()
		// Time based comparisons have drift that causes DeepEqual to fail for some reason
		if field.Name == "Date" {
			if !initalField.(time.Time).Equal(updatedField.(time.Time)) {
				t.Errorf("Compare failed: expected %v for field %s to be the same, got %v", initalField, field.Name, updatedField)
			}
		} else {
			compResult := reflect.DeepEqual(initalField, updatedField)
			if different && compResult {
				t.Errorf("Compare failed: expected %v for field %s to be different, got %v", initalField, field.Name, updatedField)
			} else if !different && !compResult {
				t.Errorf("Compare failed: expected %v for field %s to be the same, got %v", initalField, field.Name, updatedField)
			}
		}
	}
}
