package inner

import (
	_ "go.etcd.io/bbolt"
)

func Version() int {
	return 2
}
