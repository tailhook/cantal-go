package cantal

import "unsafe"
import "fmt"

type Value interface {
    fmt.Formatter
    GetSize() int
    GetType() string
    GetName() *map[string]string
    set_pointer(unsafe.Pointer)
}
