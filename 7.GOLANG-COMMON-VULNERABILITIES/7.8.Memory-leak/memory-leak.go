package main

import (
    "bytes"
    "fmt"
    "io"
    "text/template"
)

type SecWriter struct {
    w io.Writer
}

func (s *SecWriter) Write(p []byte) (n int, err error) {
    fmt.Println(string(p), len(p), cap(p))

    // here
    tmp := fmt.Sprintln("info{SSSSSSSSSSSSSSSSSSSSSSSSSSS}")
    if tmp == ""{}

    s.w.Write(p[:64])
    return 64, nil
}

func index() {
    exp := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA{{1}}"

    b := &bytes.Buffer{}
    s := &SecWriter{
        w: b,
    }


    t := template.Must(template.New("index").Parse(exp))
    t.Execute(s, nil)

    fmt.Println("buf: ", b.String())
}

func main() {
    index()
}