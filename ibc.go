// Program ibc operates like the program bc but also copies standard input to
// standard output.  This is useful in vi by using "!}ibc" to process a
// paragraph of input without losing the input.
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	os.Stdout.Write(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cmd := exec.Command("bc", os.Args[1:]...)
	cmd.Stdin = bytes.NewReader(data)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout
	err = cmd.Run()
	if len(stdout.Bytes()) + len(stderr.Bytes()) > 0 {
		fmt.Println()
	}
	os.Stdout.Write(stdout.Bytes())
	os.Stderr.Write(stderr.Bytes())
	switch err := err.(type) {
	case nil:
	case *exec.ExitError:
		os.Exit(err.ProcessState.ExitCode())
	default:
		fmt.Fprintln(os.Stderr, err)
	}
}
