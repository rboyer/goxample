package main

import (
	"fmt"

	inner_v1 "github.com/rboyer/goxample/inner"
	"github.com/rboyer/goxample/v2/inner"
	_ "go.etcd.io/bbolt"
)

func main() {
	fmt.Printf("HI FROM VEE %d\n", inner.Version())
	fmt.Printf("BUT ALSO HI FROM VEE %d\n", inner_v1.Version())
}
