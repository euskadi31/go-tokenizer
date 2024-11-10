# Text Tokenizer for Golang ![Last release](https://img.shields.io/github/release/euskadi31/go-tokenizer.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/euskadi31/go-tokenizer)](https://goreportcard.com/report/github.com/euskadi31/go-tokenizer)

| Branch | Status                                                                                                                                                    | Coverage                                                                                                                                             |
| ------ | --------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| master | [![Go](https://github.com/euskadi31/go-tokenizer/actions/workflows/go.yml/badge.svg)](https://github.com/euskadi31/go-tokenizer/actions/workflows/go.yml) | [![Coveralls](https://img.shields.io/coveralls/euskadi31/go-tokenizer/master.svg)](https://coveralls.io/github/euskadi31/go-tokenizer?branch=master) |

```sh
go get -u github.com/euskadi31/go-tokenizer/v3
```

```go
import (
    "fmt"

    "github.com/euskadi31/go-tokenizer/v3"
)

func main() {
    t := tokenizer.New()

    tokens := t.Tokenize("I believe life is an intelligent thing: that things aren't random.")

    fmt.Print(tokens) // []string{"I", "believe", "life", "is", "an", "intelligent", "thing", "that", "things", "aren't", "random"}
}

```

## License

go-tokenizer is licensed under [the MIT license](LICENSE.md).
