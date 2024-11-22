package commands

import (
	"encoding/json"
	"fmt"
	"os"
)

// MockCommandExecutor implements CommandExecutor for testing
type MockCommandExecutor struct {
	OutputFile string
}

// NewMockCommandExecutor creates a new MockCommandExecutor
func NewMockCommandExecutor(outputFile string) *MockCommandExecutor {
	return &MockCommandExecutor{
		OutputFile: outputFile,
	}
}

// Output writes mock data to the specified file
func (e *MockCommandExecutor) Output() ([]byte, error) {
	// Mock discovered repositories data
	mockData := []string{
		"github.com/owner1/repo1",
		"github.com/owner2/repo2",
		"github.com/owner3/repo3",
	}

	// Convert mock data to JSON
	jsonData, err := json.Marshal(mockData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	// Write JSON data to the output file
	err = os.WriteFile(e.OutputFile, jsonData, 0644)
	if err != nil {
		return nil, fmt.Errorf("error writing JSON to file: %v", err)
	}

	return jsonData, nil // Return the mock data as output
}

// MockShellDiscoverCommandFunc creates a MockCommandExecutor for testing
func MockShellDiscoverCommandFunc(file string) Executor {
	return NewMockCommandExecutor(file)
}
