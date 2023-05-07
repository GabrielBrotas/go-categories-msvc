package utils

import (
	"strconv"
)

func StringToUint(str string) (i uint, err error) {
	u64, err := strconv.ParseUint(str, 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(u64), nil
}
