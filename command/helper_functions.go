package command

import "strconv"

func intToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
