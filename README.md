# Gachinate

Package gachinator provides utils to gachinate input text.

## Usage

### CLI

TODO: pre-build or how to build

### As package

```go
// ...
import "stefanitsky.me/gachinator"
// ...

func main() {
    // ...
    input := "some input message text"
    result := gachinator.Gachinate([]byte(input))
    fmt.Println(string(result))
    // ...
}
```

## Examples

### EN (no yet implemented)

* `message` -> `mASSage`
* `cool` -> `c♂♂l`

### RU

* `фактор` -> `FUCKт♂р`
* `тёмный` -> `DARK`

## TODO

* Multilingual regexp (EN/RU etc.)
* CLI auto ci build
