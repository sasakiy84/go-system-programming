package main

import (
	"crypto/rand"
	"io"
	"os"
)

func myCopyN(dst io.Writer, src io.Reader, n int) {
	lReader := io.LimitReader(src, int64(n))
	io.Copy(dst, lReader)
}

func main() {
	r := rand.Reader

	myCopyN(os.Stdout, r, 1024)
}
