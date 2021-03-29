package main

import (
	"fmt"

	"github.com/rboyer/goxample/v2/inner"
	_ "go.etcd.io/bbolt"
)

func main() {
	fmt.Printf("HI FROM VEE %d\n", inner.Version())
}
