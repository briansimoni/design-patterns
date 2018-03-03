// With the adapter design pattern, we've learned a quick way
// to achieve the open/close principle in applications. Instead
// of modifying your old source code (something which could not)
// be possible in some situations), you have created a way to use
// the old functionality with a new signature

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	pipeReader, pipeWriter := io.Pipe()
	defer pipeReader.Close()
	defer pipeWriter.Close()
	c := Counter{Writer: pipeWriter}

	f, err := os.Create("temp.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// defer os.Remove("temp.txt")
	tee := io.TeeReader(pipeReader, f)
	go func() {
		io.Copy(os.Stdout, tee)
	}()
	c.Count(100)
}

type Counter struct {
	Writer io.Writer
}

func (f *Counter) Count(n uint64) uint64 {
	if n == 0 {
		fmt.Println(0)
		return 0
	}
	cur := n
	out := strconv.FormatUint(cur, 10)
	out += "\n"
	f.Writer.Write([]byte(out))
	return f.Count(n - 1)
}
