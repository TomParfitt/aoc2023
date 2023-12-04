package file

import (
	"fmt"
	"os"
)

func MustReadToString(file string) string {
	b, err := os.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
