# Friendly Words
========================

This is a port of [glitchdotcom/friendly-words](https://github.com/glitchdotcom/friendly-words/tree/main/words) for golang.
The list of words are embedded in `generator/word/*.go`.

## Usage
```go
package main

import (
    "fmt"

    fw "github.com/danruto/friendly-words-go/generator"
)

func main() {
    s := fw.GenerateUniqueWordListID(5)

    // Generates something like `chime-worried-potpourri-phalanx-assembly`
    fmt.Println(s)
}

```
