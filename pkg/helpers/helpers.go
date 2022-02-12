package helpers

import (
	"fmt"
	"os"
)

func DD(data ...interface{}) {
	fmt.Println(data...)
	os.Exit(1)
}
