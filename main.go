package main

import (
	"fmt"

	_ "github.com/boltdb/bolt"
	"github.com/rboyer/goxample/inner"
)

func main() {
	fmt.Printf("HI FROM VEE %d\n", inner.Version())
}
