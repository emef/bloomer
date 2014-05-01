bloomer
=======

Simple bloom filter in golang.


**usage**

```go

import (
  "fmt"
  "github.com/emef/bloomer"
)

// create bloom filter with estimated 1 million elements
// and a false positive probability of 0.001
set := bloomer.NewSuggested(1e6, 1e-3)

// add some items
set.Add([]byte("hi, my name is"))
set.Add([]byte("..."))

// check membership:
fmt.Println("expect true:", set.Get([]byte("hi, my name is")))
fmt.Println("expect false:", set.Get([]byte("missing")))

```

