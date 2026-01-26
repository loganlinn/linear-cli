// Package taskwriter handles writing Claude Code task JSON files
package taskwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ClaudeTask represents a Claude Code task in JSON format
// This schema matches Claude Code's task management system
type ClaudeTask struct {
	ID          string   `json:"id"`
	Subject     string   `json:"subject"`
	Description string   `json:"description"`
	ActiveForm  string   `json:"activeForm"`
	Status      string   `json:"status"`
	Blocks      []string `json:"blocks"`
	BlockedBy   []string `json:"blockedBy"`
}

// Writer handles writing tasks to the filesystem
type Writer struct{}

// NewWriter creates a new Writer instance
func NewWriter() *Writer {
	return &Writer{}
}

// WriteTasks writes a collection of tasks to the specified folder
// Each task is written as a separate JSON file named {task.ID}.json
func (w *Writer) WriteTasks(folder string, tasks []ClaudeTask) error {
	if folder == "" {
		return fmt.Errorf("output folder cannot be empty")
	}

	// Create folder if it doesn't exist
	if err := os.MkdirAll(folder, 0755); err != nil {
		return fmt.Errorf("failed to create folder %s: %w", folder, err)
	}

	// Write each task as a separate JSON file
	for _, task := range tasks {
		filename := filepath.Join(folder, fmt.Sprintf("%s.json", task.ID))

		// Marshal task to JSON with indentation
		data, err := json.MarshalIndent(task, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal task %s: %w", task.ID, err)
		}

		// Write to file
		if err := os.WriteFile(filename, data, 0644); err != nil {
			return fmt.Errorf("failed to write task %s to %s: %w", task.ID, filename, err)
		}
	}

	return nil
}
