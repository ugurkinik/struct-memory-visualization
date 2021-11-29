package main

import (
	"reflect"

	smv "github.com/kinix/struct-memory-visualization"
)

type struct1 struct {
	bool1 bool
	bool2 bool
	int1  int32
	bool3 bool
	bool4 bool
	bool5 bool
	int2  int64
}

func main() {
	smv.DrawMemory(reflect.TypeOf(struct1{}), "image.png")
}
