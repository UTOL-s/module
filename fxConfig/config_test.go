package fxconfig

import (
	"os"
	"testing"
)

func TestConfigAccessor(t *testing.T) {
	accessor := NewConfigAccessor()

	if accessor == nil {
		t.Error("ConfigAccessor should not be nil")
	}
}

func TestGetEnv(t *testing.T) {
	// Set a test environment variable
	testKey := "TEST_CONFIG_KEY"
	testValue := "test_value"
	os.Setenv(testKey, testValue)
	defer os.Unsetenv(testKey)

	result := GetEnv(testKey)
	if result != testValue {
		t.Errorf("Expected %s, got %s", testValue, result)
	}
}

func TestAccessorMethods(t *testing.T) {
	accessor := &Accessor{}

	// Test that methods don't panic
	_ = accessor.String("test")
	_ = accessor.Int("test")
	_ = accessor.Bool("test")
	_ = accessor.Float64("test")
	_ = accessor.AllSettings()
}

func TestConfigAccessorSingleton(t *testing.T) {
	accessor1 := ConfigAccessor()
	accessor2 := ConfigAccessor()

	if accessor1 != accessor2 {
		t.Error("ConfigAccessor should return the same instance")
	}
}
