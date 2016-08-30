# spold2
spold2 is a small [Go](https://golang.org/) package for processing
[EcoSpold 2](http://www.ecoinvent.org/data-provider/data-provider-toolkit/ecospold2/ecospold2.html)
files.

## Getting Started

### Installing
Install Go and run `go get`:

```
go get github.com/msrocka/spold2
```

### Reading an EcoSpold 2 file

```go
package main

import (
	"fmt"
	"github.com/msrocka/spold2"
)

func main() {
	spold, err := spold2.ReadFile("path/to/file")
	if err != nil {
		// handle error
		return
	}
	fmt.Println(spold.DataSet.Description.Name)
}
```