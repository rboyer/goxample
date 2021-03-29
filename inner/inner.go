package inner

import (
	_ "github.com/boltdb/bolt"
)

func Version() int {
	return 1
}
