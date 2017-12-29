# epub
A pure go implementation of epub file format.

 [ ![Codeship Status for n3integration/epub](https://app.codeship.com/projects/1a12d410-d139-0135-6c51-26e28af241d2/status?branch=master)](https://app.codeship.com/projects/262382)
 [![Go Report Card](https://goreportcard.com/badge/github.com/n3integration/epub)](https://goreportcard.com/report/github.com/n3integration/epub)
 [![Documentation](https://godoc.org/github.com/n3integration/epub?status.svg)](http://godoc.org/github.com/n3integration/epub)

## Usage

```diff
import (
  "bytes"
  "fmt"
  "io"

  "github.com/n3integration/epub"
)

func main() {
+ // 1. Open the epub by passing its file path
  f := "some.epub"
  book, err := epub.Open(f)
  defer book.Close()

+ // 2. Iterate over book sections
  book.Each(func(title string, xhtml io.ReadCloser) {
+   // 3. Read and process section contents
    buf := new(bytes.Buffer)
    buf.ReadFrom(xhtml)

    fmt.Println("==========================================================")
    fmt.Println(title)
    fmt.Println("==========================================================")
    fmt.Println(buf.String())
  })
}
```

## Reference
- [Specification](http://idpf.org/epub)
