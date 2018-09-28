# leetspeak-go

Simple English to LeetSpeak converter.

## Installation
Assuming you have a working Go environment and `GOPATH/bin` is in your
`PATH`

```shell
go get -u github.com/iSuperMostafa/leetspeak-go
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/iSuperMostafa/leetspeak-go"
)

func main() {
    leet := leetspeak.ConvertToLeet("hello")
	fmt.Println(leet)
}
```