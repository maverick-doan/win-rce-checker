package fileOps

import (
	"fmt"
	"os"
)

func WriteIndicatorFile(path string, content string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create indicator file: %w", err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		return nil, fmt.Errorf("failed to write to file: %w", err)
	}
	return file, err
}
