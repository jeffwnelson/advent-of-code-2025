package utils

import (
	"fmt"
	"time"
)

func Run(name string, fn func() any) {
	start := time.Now()
	result := fn()
	fmt.Printf("%s: %v [%v]\n", name, result, time.Since(start))
}
