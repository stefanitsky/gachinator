# Gachinate

![workflow](https://github.com/stefanitsky/gachinator/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/stefanitsky/gachinator)](https://goreportcard.com/report/github.com/stefanitsky/gachinator)
[![codecov](https://codecov.io/gh/stefanitsky/gachinator/branch/master/graph/badge.svg)](https://codecov.io/gh/stefanitsky/gachinator)

Package gachinator provides utils to gachinate input text.

## Usage

### CLI

[Pre-built binaries](https://github.com/stefanitsky/gachinator/releases/)

#### CLI Usage

```sh
$ cat testdata/benchmark.txt | gachinator --lang=ru
П♂вседневная практика п♂казывает, чт♂ дальнейшее развитие...

$ echo 'Hello' | gachinator --lang=en
Hell♂
```

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

[More examples in tests](https://github.com/stefanitsky/gachinator/blob/master/gachinate_test.go)
