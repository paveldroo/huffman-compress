package counter

import (
	"fmt"
	"os"
)

func CharsCount(fname string) (map[string]int, error) {
	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %w", err)
	}

	res := map[string]int{}

	for _, ch := range string(data) {
		str := string(ch)
		_, ok := res[str]
		if !ok {
			res[str] = 1
			continue
		}
		res[str]++
	}

	return res, nil
}
