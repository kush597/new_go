package main

import (
    "fmt"
    "unsafe"

    "log"
)

func seeAnotherDay() {
    defer func() {
        if p := recover(); p != nil {
            err := fmt.Errorf("recover panic: panic call")
            log.Println(err)
            return
        }
    }()
    panic("oops")
}

func notSoMuch() {
    defer func() {
        if p := recover(); p != nil {
            err := fmt.Errorf("recover panic: sigseg")
            log.Println(err)
            return
        }
    }()
    b := make([]byte, 1)
    log.Println("access some memory")
    foo := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(9999999999999999)))
    fmt.Print(*foo + 1)
}

func main() {
    seeAnotherDay()
    notSoMuch()
}