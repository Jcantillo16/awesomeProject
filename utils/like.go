package utils

import "fmt"

func Like(param string) string {
	return fmt.Sprintf(`{ "$regex": ".*%s.*" }`, param)
}
