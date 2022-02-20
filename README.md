# Gachinate

Package gachinator provides utils to gachinate input text.

## Usage

### CLI

TODO: pre-build or how to build

### As package

```go
// ...
import "github.com/stefanitsky/gachinator"
// ...

func main() {
    // ...
    input := "some input message text"
    result := gachinator.GachinateEN([]byte(input))
    fmt.Println(string(result))
    // ...
}
```

## Examples

### EN

* `manufacturable` -> `manuFUCKturable`
* `cool` -> `c♂♂l`

### RU

* `фактор` -> `FUCKт♂р`
* `тёмный` -> `DARK`

## TODO

* CLI auto ci build
