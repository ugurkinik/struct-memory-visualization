# Struct Memory Visualization Tool

It is a library to create images about how a struct is stored in the memory.

## Example

You can see how to use that library in the folder `example`. There is a small go app to visualize a struct.

```go
package main

import (
	"reflect"

	smv "github.com/kinix/struct-memory-visualization"
)

type exampleStruct struct {
	bool1 bool
	num1  int64
	bool2 bool
	bool4 bool
	num2  int32
	bool5 bool
}

func main() {
	smv.DrawMemory(reflect.TypeOf(exampleStruct{}), "image.png")
}

```

DrawMemory is the function that creates the image. The first parameter is the type of your struct and the second parameter is the name of the image file. The output of that code is below:

![example image](/example/image.png)

You can see the position of each field. Each line represents a word. The red blocks are paddings. 

You can check my article for more info: [medium.com](https://medium.com/@ugurkinik/optimizing-memory-by-changing-the-order-of-struct-field-485106504087)
