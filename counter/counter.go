package counter

func CharsCount(fileData []byte) (map[string]int, error) {
	res := map[string]int{}

	for _, ch := range string(fileData) {
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
