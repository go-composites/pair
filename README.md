<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/pair" width="720"></p>

# pair

[![ci](https://github.com/go-composites/pair/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/pair/actions/workflows/ci.yml)

The two-element composite of [go-composites](https://github.com/go-composites).
A `Pair` is a fixed, heterogeneous grouping of two values — the natural
key/value entry. It is interface-first and never goes `nil`: the Null-Object
`Pair` honours the whole `Interface`.

## Install

```sh
go get github.com/go-composites/pair
```

## Usage

```golang
package main

import (
    "fmt"

    Pair "github.com/go-composites/pair/src"
)

func main() {
    entry := Pair.New("answer", 42)

    fmt.Println(entry.ToGoString())                    // (answer, 42)
    fmt.Println(entry.First())                         // answer
    fmt.Println(entry.Second())                        // 42
    fmt.Println(entry.Equal(Pair.New("answer", 42)))   // true
    fmt.Println(entry.ToArray().Len())                 // 2
}
```

## API

### Constructors

- `New(first, second interface{}) Interface` — a Pair holding the two values.
- `Null() Interface` — the Null-Object `Pair`.

### Methods

| method | returns | notes |
| --- | --- | --- |
| `First()` | `interface{}` | the first slot |
| `Second()` | `interface{}` | the second slot |
| `Equal(other)` | Go `bool` | both slots `reflect.DeepEqual` (false against a null Pair) |
| `ToArray()` | `Array.Interface` | a two-element `[first, second]` Array |
| `ToGoString()` | Go `string` | `"(first, second)"` (`"(null)"` for the Null-Object) |
| `IsNull()` | Go `bool` | `true` only for the Null-Object |

## Null-Object

`Null()` returns the never-nil Null-Object `Pair`: its slots are `nil`, it
`Equal`s nothing (not even another null `Pair`), `ToArray()` is the empty
`Array`, and `ToGoString()` renders `"(null)"`. `IsNull()` reports `true` for
it.

## License

BSD-3-Clause © the go-composites/pair authors.
