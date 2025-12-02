package utils

import (
	"fmt"
	"time"
)

func Run(name string, function func() int) {
	start := time.Now()
	result := function()
	fmt.Printf("%s: %v [%v]\n", name, result, time.Since(start))
}
