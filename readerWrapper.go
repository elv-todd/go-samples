package main

import (
    "fmt"
    "io"
    "os"
    "strings"
)

type IPAddr [4]byte

func (ipa IPAddr) String() string {
  return fmt.Sprintf("%d.%d.%d.%d", ipa[0], ipa[1], ipa[2], ipa[3])
}

type rot13Reader struct {
    r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
    n, err := r13.r.Read(b)
    for i := 0; i <= n; i++ {
        b[i] = b[i] + 13 // not really
    }
    return n, err
}

func main() {
    r := strings.NewReader("Hello, Eluvio!")

    b := make([]byte, 8)
    for {
        n, err := r.Read(b)
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n])
        if err == io.EOF {
            break
        }
    }

    r2 := strings.NewReader("Hello, Eluvio!")
    r13 := rot13Reader{r2}
    io.Copy(os.Stdout, &r13)
}