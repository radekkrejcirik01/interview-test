package helpers

import "strconv"

// StringToInt64 convert string value to int64
func StringToInt64(s string) (int64, error) {
	intId, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return int64(intId), nil
}
