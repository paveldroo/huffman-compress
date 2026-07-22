package reader

import (
	"fmt"
	"os"
)

func FileData(fname string) ([]byte, error) {
	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}

	return data, nil
}
