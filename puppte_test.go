package puppet

import (
	"fmt"
	"testing"
)

func TestWork(t *testing.T) {
	pool := NewPool(5)
	fmt.Println(pool)
}
