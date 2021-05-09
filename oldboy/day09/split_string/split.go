package split_string

import (
	"strings"
)

func Split(str string, sep string) []string {
	var ret []string
	i := strings.Index(str, sep)
	for i > 0 {
		ret = append(ret, str[:i])
		str = str[i+1:]
		i = strings.Index(str, sep)
	}

	ret = append(ret, str)
	return ret
}
