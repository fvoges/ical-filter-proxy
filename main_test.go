package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadSecretFile(t *testing.T) {
	// Create a temporary file with test content
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "secret.txt")
	
	testContent := "my-secret-token"
	err := os.WriteFile(tmpFile, []byte(testContent), 0600)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	// Test reading the file
	result, err := readSecretFile(tmpFile)
	if err != nil {
		t.Errorf("readSecretFile() error = %v", err)
	}
	
	if result != testContent {
		t.Errorf("readSecretFile() = %v, expected %v", result, testContent)
	}
}

func TestReadSecretFile_WithWhitespace(t *testing.T) {
	// Create a temporary file with whitespace
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "secret.txt")
	
	testContent := "  my-secret-token  \n"
	expectedContent := "my-secret-token"
	
	err := os.WriteFile(tmpFile, []byte(testContent), 0600)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	// Test reading the file - should trim whitespace
	result, err := readSecretFile(tmpFile)
	if err != nil {
		t.Errorf("readSecretFile() error = %v", err)
	}
	
	if result != expectedContent {
		t.Errorf("readSecretFile() = '%v', expected '%v'", result, expectedContent)
	}
}

func TestReadSecretFile_NonExistent(t *testing.T) {
	// Test reading a non-existent file
	result, err := readSecretFile("/nonexistent/file.txt")
	if err == nil {
		t.Error("readSecretFile() expected error for non-existent file, got nil")
	}
	
	if result != "" {
		t.Errorf("readSecretFile() = %v, expected empty string", result)
	}
}

func TestConfigLoadConfig_Valid(t *testing.T) {
	// Create a temporary valid config file
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "config.yaml")
	
	validConfig := `
calendars:
  - name: test-calendar
    public: true
    feed_url: https://example.com/calendar.ics
    filters:
      - description: Test filter
        match:
          summary:
            contains: meeting
`
	
	err := os.WriteFile(configFile, []byte(validConfig), 0600)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	
	// Test loading the config
	var config Config
	result := config.LoadConfig(configFile)
	
	if !result {
		t.Error("LoadConfig() = false, expected true for valid config")
	}
	
	if len(config.Calendars) != 1 {
		t.Errorf("LoadConfig() loaded %d calendars, expected 1", len(config.Calendars))
	}
	
	if config.Calendars[0].Name != "test-calendar" {
		t.Errorf("LoadConfig() calendar name = %v, expected 'test-calendar'", config.Calendars[0].Name)
	}
}

func TestConfigLoadConfig_NoCalendars(t *testing.T) {
	// Create a config with no calendars
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "config.yaml")
	
	invalidConfig := `
calendars: []
`
	
	err := os.WriteFile(configFile, []byte(invalidConfig), 0600)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	
	// Test loading the config - should fail
	var config Config
	result := config.LoadConfig(configFile)
	
	if result {
		t.Error("LoadConfig() = true, expected false for config with no calendars")
	}
}

func TestConfigLoadConfig_InvalidURL(t *testing.T) {
	// Create a config with invalid URL
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "config.yaml")
	
	invalidConfig := `
calendars:
  - name: test-calendar
    public: true
    feed_url: ftp://example.com/calendar.ics
`
	
	err := os.WriteFile(configFile, []byte(invalidConfig), 0600)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	
	// Test loading the config - should fail due to invalid URL scheme
	var config Config
	result := config.LoadConfig(configFile)
	
	if result {
		t.Error("LoadConfig() = true, expected false for invalid URL scheme")
	}
}

func TestConfigLoadConfig_NonPublicWithoutToken(t *testing.T) {
	// Create a config with public=false but no token
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "config.yaml")
	
	invalidConfig := `
calendars:
  - name: test-calendar
    public: false
    feed_url: https://example.com/calendar.ics
`
	
	err := os.WriteFile(configFile, []byte(invalidConfig), 0600)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	
	// Test loading the config - should fail
	var config Config
	result := config.LoadConfig(configFile)
	
	if result {
		t.Error("LoadConfig() = true, expected false for non-public calendar without token")
	}
}

func TestConfigLoadConfig_FreeBusyMode(t *testing.T) {
	// Create a config with freebusy_mode enabled
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "config.yaml")
	
	validConfig := `
calendars:
  - name: freebusy-calendar
    public: true
    feed_url: https://example.com/calendar.ics
    freebusy_mode: true
`
	
	err := os.WriteFile(configFile, []byte(validConfig), 0600)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	
	// Test loading the config
	var config Config
	result := config.LoadConfig(configFile)
	
	if !result {
		t.Error("LoadConfig() = false, expected true for valid config")
	}
	
	if !config.Calendars[0].FreeBusyMode {
		t.Error("LoadConfig() freebusy_mode = false, expected true")
	}
}
