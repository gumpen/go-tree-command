package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPattern1(t *testing.T) {

	expected := []byte(
		`./example
├── dir1
│   ├── dir11
│   │   └── file3.go
│   ├── dir12
│   │   └── file4.go
│   └── file2.go
└── file1.go
`)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "./example"}
	out := captureStdout()
	if out != string(expected) {
		t.Errorf("Unexpected out")
	}
}

func TestPattern2(t *testing.T) {

	expected := []byte(
		`./example3
└── example
    ├── dir1
    │   ├── dir11
    │   │   └── file3.go
    │   ├── dir12
    │   │   └── file4.go
    │   └── file2.go
    └── dir2
        └── file1.go
`)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "./example3"}
	out := captureStdout()
	fmt.Printf("%#v\n", out)
	fmt.Printf("%#v\n", string(expected))
	if out != string(expected) {
		t.Errorf("Unexpected out")
	}
}

func TestPattern3(t *testing.T) {

	expected := []byte(
		`./example3
├── .gorc
└── example
    ├── dir1
    │   ├── dir11
    │   │   └── file3.go
    │   ├── dir12
    │   │   └── file4.go
    │   └── file2.go
    └── dir2
        └── file1.go
`)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-a", "./example3"}
	out := captureStdout()
	if out != string(expected) {
		t.Errorf("Unexpected out")
	}
}

func captureStdout() string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	main()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
